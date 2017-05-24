package main

import (
	"io"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"

	"code.cloudfoundry.org/winc/container"
	"code.cloudfoundry.org/winc/hcsclient"
	"code.cloudfoundry.org/winc/sandbox"
	"github.com/Sirupsen/logrus"
	specs "github.com/opencontainers/runtime-spec/specs-go"
	"github.com/urfave/cli"
)

var execCommand = cli.Command{
	Name:  "exec",
	Usage: "execute new process inside a container",
	ArgsUsage: `<container-id> <command> [command options]  || -p process.json <container-id>

Where "<container-id>" is the name for the instance of the container and
"<command>" is the command to be executed in the container.
"<command>" can't be empty unless a "-p" flag provided.

EXAMPLE:
For example, if the container is configured to run the Windows ps command the
following will output a list of processes running in the container:

       # runc exec <container-id> ps`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "pid-file",
			Value: "",
			Usage: "specify the file to write the process id to",
		},
		cli.BoolFlag{
			Name:  "detach,d",
			Usage: "detach from the container's process",
		},
		cli.StringFlag{
			Name:  "process, p",
			Value: "",
			Usage: "path to the process.json",
		},
		cli.StringFlag{
			Name:  "cwd",
			Value: "",
			Usage: "current working directory in the container",
		},
		cli.StringFlag{
			Name:  "user, u",
			Value: "",
			Usage: "Username to execute process as",
		},
		cli.StringSliceFlag{
			Name:  "env, e",
			Usage: "set environment variables",
		},
		// 	cli.BoolFlag{
		// 		Name:  "tty, t",
		// 		Usage: "allocate a pseudo-TTY",
		// 	},
	},
	Action: func(context *cli.Context) error {
		if err := checkArgs(context, 1, minArgs); err != nil {
			return err
		}

		containerId := context.Args().First()
		processConfig := context.String("process")
		command := context.Args()[1:]
		cwd := context.String("cwd")
		user := context.String("user")
		env := context.StringSlice("env")
		pidFile := context.String("pid-file")
		detach := context.Bool("detach")

		logger := logrus.WithField("containerId", containerId)

		spec, err := ValidateProcess(logger, processConfig, &specs.Process{
			Args: command,
			Cwd:  cwd,
			User: specs.User{
				Username: user,
			},
			Env: env,
		})
		if err != nil {
			return err
		}

		logger.WithFields(logrus.Fields{
			"processConfig": processConfig,
			"pidFile":       pidFile,
			"command":       spec.Args,
			"cwd":           spec.Cwd,
			"user":          spec.User.Username,
			"env":           env,
			"detach":        detach,
		}).Debug("executing process in container")

		client := hcsclient.HCSClient{}
		cp, err := client.GetContainerProperties(containerId)
		if err != nil {
			return err
		}
		sm := sandbox.NewManager(&client, cp.Name)
		cm := container.NewManager(&client, sm, containerId)

		process, err := cm.Exec(spec)
		if err != nil {
			return err
		}

		if pidFile != "" {
			if err := ioutil.WriteFile(pidFile, []byte(strconv.FormatInt(int64(process.Pid()), 10)), 0666); err != nil {
				return err
			}
		}

		if !detach {
			_, stdout, stderr, err := process.Stdio()
			if err != nil {
				return err
			}

			go func() {
				io.Copy(os.Stdout, stdout)
				stdout.Close()
			}()
			go func() {
				io.Copy(os.Stderr, stderr)
				stderr.Close()
			}()

			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt)
			go func() {
				<-c
				process.Kill()
			}()

			if err := process.Wait(); err != nil {
				return err
			}

			exitCode, err := process.ExitCode()
			if err != nil {
				return err
			}
			os.Exit(exitCode)
		}

		return nil
	},
	SkipArgReorder: true,
}
