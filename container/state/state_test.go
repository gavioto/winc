package state_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"

	"code.cloudfoundry.org/winc/container/state"
	"code.cloudfoundry.org/winc/container/state/fakes"
	hcsfakes "code.cloudfoundry.org/winc/hcs/fakes"
	"github.com/Microsoft/hcsshim"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StateManager", func() {
	const (
		containerId = "some-container-id"
		bundlePath  = "some-bundle-path"
	)

	var (
		hcsClient      *fakes.HCSClient
		rootDir        string
		sm             *state.Manager
		processManager *fakes.ProcessManager
		container      *hcsfakes.Container
	)

	BeforeEach(func() {
		var err error

		rootDir, err = ioutil.TempDir("", "winc.container.state.test")
		Expect(err).NotTo(HaveOccurred())

		hcsClient = &fakes.HCSClient{}
		processManager = &fakes.ProcessManager{}
		sm = state.NewManager(hcsClient, containerId, rootDir, processManager)
	})

	AfterEach(func() {
		Expect(os.RemoveAll(rootDir)).To(Succeed())
	})

	Context("Get", func() {
		const (
			containerPid = 99
		)

		Context("before the manager has been initialized", func() {
			It("doing anything else returns an error", func() {
				_, _, err := sm.Get()
				Expect(err).To(MatchError(&state.FileNotFoundError{Id: containerId}))
				err = sm.SetRunning(0)
				Expect(err).To(MatchError(&state.FileNotFoundError{Id: containerId}))
				err = sm.SetExecFailed()
				Expect(err).To(MatchError(&state.FileNotFoundError{Id: containerId}))
			})
		})

		Context("initializing the manager", func() {
			It("writes the bundle path to state.json in <rootDir>/<containerId>/", func() {
				err := sm.Initialize(bundlePath)
				Expect(err).To(Succeed())

				var state state.ContainerState
				contents, err := ioutil.ReadFile(filepath.Join(rootDir, containerId, "state.json"))
				Expect(err).NotTo(HaveOccurred())
				Expect(json.Unmarshal(contents, &state)).To(Succeed())

				Expect(state.Bundle).To(Equal(bundlePath))
			})
		})

		Context("after the manager has been initialized", func() {
			BeforeEach(func() {
				Expect(sm.Initialize(bundlePath)).To(Succeed())

				container = &hcsfakes.Container{}
			})

			Context("when the container pid is found", func() {
				BeforeEach(func() {
					processManager.ContainerPidReturnsOnCall(0, containerPid, nil)
				})

				It("returns the status as 'created' along with the other expected state fields", func() {
					status, resultBundlePath, err := sm.Get()

					Expect(err).NotTo(HaveOccurred())
					Expect(status).To(Equal("created"))
					Expect(resultBundlePath).To(Equal(bundlePath))

					Expect(hcsClient.GetContainerPropertiesCallCount()).To(Equal(1))
					Expect(hcsClient.GetContainerPropertiesArgsForCall(0)).To(Equal(containerId))
				})

				Context("after the container has been stopped", func() {
					BeforeEach(func() {
						hcsClient.GetContainerPropertiesReturnsOnCall(0, hcsshim.ContainerProperties{Stopped: true}, nil)
					})

					It("returns the status as 'stopped' along with the other expected state fields", func() {
						status, resultBundlePath, err := sm.Get()

						Expect(err).NotTo(HaveOccurred())
						Expect(status).To(Equal("stopped"))
						Expect(resultBundlePath).To(Equal(bundlePath))

						Expect(hcsClient.GetContainerPropertiesCallCount()).To(Equal(1))
						Expect(hcsClient.GetContainerPropertiesArgsForCall(0)).To(Equal(containerId))
					})
				})

				Context("after the init process has been successfully started", func() {
					var initProcessPid int
					BeforeEach(func() {
						initProcessPid = 89
						processManager.ProcessStartTimeReturns(syscall.Filetime{LowDateTime: 10, HighDateTime: 100}, nil)
						Expect(sm.SetRunning(initProcessPid)).To(Succeed())

						hcsClient.OpenContainerReturnsOnCall(0, container, nil)
						hcsClient.OpenContainerReturnsOnCall(1, container, nil)
					})

					Context("and the init process is still running", func() {
						BeforeEach(func() {
							processList := []hcsshim.ProcessListItem{
								hcsshim.ProcessListItem{
									ProcessId: uint32(containerPid),
									ImageName: "wininit.exe",
								},
								hcsshim.ProcessListItem{
									ProcessId: uint32(initProcessPid),
									ImageName: "init-process.exe",
								},
							}
							container.ProcessListReturnsOnCall(0, processList, nil)
							container.ProcessListReturnsOnCall(1, processList, nil)
						})

						It("returns the status as 'running' along with the other expected state fields", func() {
							status, resultBundlePath, err := sm.Get()

							Expect(err).NotTo(HaveOccurred())
							Expect(status).To(Equal("running"))
							Expect(resultBundlePath).To(Equal(bundlePath))

							Expect(hcsClient.OpenContainerCallCount()).To(Equal(1))
							Expect(hcsClient.OpenContainerArgsForCall(0)).To(Equal(containerId))

							Expect(container.ProcessListCallCount()).To(Equal(1))
							Expect(container.CloseCallCount()).To(Equal(1))

							Expect(processManager.ProcessStartTimeCallCount()).To(Equal(2))
						})
					})

					Context("and the init process has returned", func() {
						BeforeEach(func() {

							processList := []hcsshim.ProcessListItem{
								hcsshim.ProcessListItem{
									ProcessId: uint32(containerPid),
									ImageName: "wininit.exe",
								},
							}
							container.ProcessListReturnsOnCall(0, processList, nil)
						})

						It("returns the status as 'exited' along with the other expected state fields", func() {
							status, resultBundlePath, err := sm.Get()

							Expect(err).NotTo(HaveOccurred())
							Expect(status).To(Equal("exited"))
							Expect(resultBundlePath).To(Equal(bundlePath))

							Expect(hcsClient.OpenContainerCallCount()).To(Equal(1))
							Expect(hcsClient.OpenContainerArgsForCall(0)).To(Equal(containerId))

							Expect(container.ProcessListCallCount()).To(Equal(1))
							Expect(container.CloseCallCount()).To(Equal(1))
						})
					})
				})
			})

			Context("FAILURE", func() {
				Context("after the init process has failed to start", func() {
					BeforeEach(func() {
						Expect(sm.SetExecFailed()).To(Succeed())
					})

					It("returns the status as 'exited' along with the other expected state fields", func() {
						status, resultBundlePath, err := sm.Get()

						Expect(err).NotTo(HaveOccurred())
						Expect(status).To(Equal("exited"))
						Expect(resultBundlePath).To(Equal(bundlePath))
					})
				})

				Context("when the specified container does not exist", func() {
					BeforeEach(func() {
						hcsClient.GetContainerPropertiesReturns(hcsshim.ContainerProperties{}, errors.New("container does not exist"))
					})

					It("errors", func() {
						_, _, err := sm.Get()
						Expect(err).To(Equal(&state.ContainerNotFoundError{Id: containerId}))
					})
				})
			})
		})
	})
	Context("modifying the state file", func() {
		Context("after the manager has been initialized", func() {
			BeforeEach(func() {
				Expect(sm.Initialize(bundlePath)).To(Succeed())

				//container = &hcsfakes.Container{}
			})

			Context("SetExecFailed", func() {
				It("adds UserProgramExecFailed to the state file", func() {
					var beforeState state.ContainerState
					contents, err := ioutil.ReadFile(filepath.Join(rootDir, containerId, "state.json"))
					Expect(err).NotTo(HaveOccurred())
					Expect(json.Unmarshal(contents, &beforeState)).To(Succeed())

					Expect(beforeState.UserProgramExecFailed).To(BeFalse())

					sm.SetExecFailed()

					var afterState state.ContainerState
					contents, err = ioutil.ReadFile(filepath.Join(rootDir, containerId, "state.json"))
					Expect(err).NotTo(HaveOccurred())
					Expect(json.Unmarshal(contents, &afterState)).To(Succeed())

					Expect(afterState.UserProgramExecFailed).To(BeTrue())
				})
			})

			Context("SetRunning", func() {
				It("adds UserProgramPID and UserProgramStartTime to the state file", func() {
					var beforeState state.ContainerState
					contents, err := ioutil.ReadFile(filepath.Join(rootDir, containerId, "state.json"))
					Expect(err).NotTo(HaveOccurred())
					Expect(json.Unmarshal(contents, &beforeState)).To(Succeed())

					Expect(beforeState.UserProgramPID).To(Equal(0))
					Expect(beforeState.UserProgramStartTime).To(Equal(syscall.Filetime{LowDateTime: 0, HighDateTime: 0}))

					expectedStartTime := syscall.Filetime{LowDateTime: 1, HighDateTime: 2}
					processManager.ProcessStartTimeReturnsOnCall(0, expectedStartTime, nil)

					err = sm.SetRunning(1)
					Expect(err).NotTo(HaveOccurred())

					var afterState state.ContainerState
					contents, err = ioutil.ReadFile(filepath.Join(rootDir, containerId, "state.json"))
					Expect(err).NotTo(HaveOccurred())
					Expect(json.Unmarshal(contents, &afterState)).To(Succeed())

					Expect(afterState.UserProgramPID).To(Equal(1))
					Expect(afterState.UserProgramStartTime).To(Equal(expectedStartTime))
				})
			})
		})
	})
})