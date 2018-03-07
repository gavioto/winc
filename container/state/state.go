package state

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"

	"code.cloudfoundry.org/winc/hcs"
	"github.com/Microsoft/hcsshim"
)

const stateFile = "state.json"

//go:generate counterfeiter -o fakes/hcsclient.go --fake-name HCSClient . HCSClient
type HCSClient interface {
	GetContainerProperties(string) (hcsshim.ContainerProperties, error)
	OpenContainer(string) (hcs.Container, error)
}

//go:generate counterfeiter -o fakes/process_manager.go --fake-name ProcessManager . ProcessManager
type ProcessManager interface {
	ContainerPid(string) (int, error)
	ProcessStartTime(uint32) (syscall.Filetime, error)
	IsProcessRunning(uint32, syscall.Filetime) (bool, error)
}

type Manager struct {
	hcsClient      HCSClient
	id             string
	rootDir        string
	processManager ProcessManager
}

type ContainerState struct {
	Bundle                string           `json:"bundle"`
	UserProgramPID        int              `json:"user_program_pid"`
	UserProgramStartTime  syscall.Filetime `json:"user_program_start_time"`
	UserProgramExecFailed bool             `json:"user_program_exec_failed"`
}

type ContainerNotFoundError struct {
	Id string
}

func (e *ContainerNotFoundError) Error() string {
	return fmt.Sprintf("container does not exist: %s", e.Id)
}

type FileNotFoundError struct {
	Id string
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("unable to find state file for container: %s", e.Id)
}

func NewManager(hcsClient HCSClient, id, rootDir string, processManager ProcessManager) *Manager {
	return &Manager{
		hcsClient:      hcsClient,
		id:             id,
		rootDir:        rootDir,
		processManager: processManager,
	}
}

func (m *Manager) Initialize(bundlePath string) error {
	if err := os.MkdirAll(m.stateDir(), 0755); err != nil {
		return err
	}

	state := ContainerState{Bundle: bundlePath}
	return m.writeState(state)
}

func (m *Manager) Get() (string, string, error) {
	if !m.isInitialized() {
		return "", "", &FileNotFoundError{Id: m.id}
	}

	cp, err := m.hcsClient.GetContainerProperties(m.id)
	if err != nil {
		return "", "", &ContainerNotFoundError{Id: m.id}
	}

	state, err := m.readState()
	if err != nil {
		return "", "", err
	}

	var status string
	if cp.Stopped {
		status = "stopped"
	} else {
		status, err = m.userProgramStatus(state)
		if err != nil {
			panic(err)
		}
	}

	return status, state.Bundle, nil
}

func (m *Manager) SetRunning(pid int) error {
	if !m.isInitialized() {
		return &FileNotFoundError{Id: m.id}
	}

	state, err := m.readState()
	if err != nil {
		return err
	}

	state.UserProgramPID = pid
	state.UserProgramStartTime, err = m.processManager.ProcessStartTime(uint32(pid))
	if err != nil {
		return err
	}

	return m.writeState(state)
}

func (m *Manager) SetExecFailed() error {
	if !m.isInitialized() {
		return &FileNotFoundError{Id: m.id}
	}

	state, err := m.readState()
	if err != nil {
		return err
	}

	state.UserProgramExecFailed = true

	return m.writeState(state)
}

func (m *Manager) stateDir() string {
	return filepath.Join(m.rootDir, m.id)
}

func (m *Manager) isInitialized() bool {
	_, err := os.Stat(m.stateDir())
	return err == nil
}

func (m *Manager) readState() (ContainerState, error) {
	contents, err := ioutil.ReadFile(filepath.Join(m.stateDir(), stateFile))
	if err != nil {
		return ContainerState{}, err
	}

	var state ContainerState
	if err := json.Unmarshal(contents, &state); err != nil {
		return ContainerState{}, err
	}

	return state, nil
}

func (m *Manager) writeState(state ContainerState) error {
	contents, err := json.Marshal(state)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filepath.Join(m.stateDir(), stateFile), contents, 0644)
}

func (m *Manager) WriteContainerState(ContainerState ContainerState) error {
	contents, err := json.Marshal(ContainerState)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filepath.Join(m.stateDir(), stateFile), contents, 0644)
}

func (m *Manager) userProgramStatus(state ContainerState) (string, error) {
	if !stateValid(state) {
		panic("invalid state")
	}

	if state.UserProgramExecFailed {
		return "exited", nil
	}

	if (state.UserProgramPID == 0) && (state.UserProgramStartTime == syscall.Filetime{}) {
		return "created", nil
	}

	isRunning, err := m.processManager.IsProcessRunning(uint32(state.UserProgramPID), state.UserProgramStartTime)
	if err != nil {
		return "", err
	}
	if isRunning {
		return "running", nil
	}

	//start(containerId, state.UserProgramPID, state.UserProgramStartTime) (bool)
	//container, err := m.hcsClient.OpenContainer(m.id)
	//if err != nil {
	//	return "", err
	//}
	//defer container.Close()

	//pl, err := container.ProcessList()
	//if err != nil {
	//	return "", err
	//}

	//for _, v := range pl {
	//	if v.ProcessId == uint32(state.UserProgramPID) {
	//		s, err := m.processManager.ProcessStartTime(v.ProcessId)
	//		if err != nil {
	//			return "", err
	//		}

	//		if s == state.UserProgramStartTime {
	//			return "running", nil
	//		}
	//	}
	//}
	//end

	return "exited", nil
}

func stateValid(state ContainerState) bool {
	return (state.UserProgramPID == 0 && state.UserProgramStartTime == syscall.Filetime{}) ||
		(state.UserProgramPID != 0 && state.UserProgramStartTime != syscall.Filetime{})
}
