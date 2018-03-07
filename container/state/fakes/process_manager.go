// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
	"syscall"

	"code.cloudfoundry.org/winc/container/state"
)

type ProcessManager struct {
	ContainerPidStub        func(string) (int, error)
	containerPidMutex       sync.RWMutex
	containerPidArgsForCall []struct {
		arg1 string
	}
	containerPidReturns struct {
		result1 int
		result2 error
	}
	containerPidReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	ProcessStartTimeStub        func(uint32) (syscall.Filetime, error)
	processStartTimeMutex       sync.RWMutex
	processStartTimeArgsForCall []struct {
		arg1 uint32
	}
	processStartTimeReturns struct {
		result1 syscall.Filetime
		result2 error
	}
	processStartTimeReturnsOnCall map[int]struct {
		result1 syscall.Filetime
		result2 error
	}
	IsProcessRunningStub        func(uint32, syscall.Filetime) (bool, error)
	isProcessRunningMutex       sync.RWMutex
	isProcessRunningArgsForCall []struct {
		arg1 uint32
		arg2 syscall.Filetime
	}
	isProcessRunningReturns struct {
		result1 bool
		result2 error
	}
	isProcessRunningReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ProcessManager) ContainerPid(arg1 string) (int, error) {
	fake.containerPidMutex.Lock()
	ret, specificReturn := fake.containerPidReturnsOnCall[len(fake.containerPidArgsForCall)]
	fake.containerPidArgsForCall = append(fake.containerPidArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ContainerPid", []interface{}{arg1})
	fake.containerPidMutex.Unlock()
	if fake.ContainerPidStub != nil {
		return fake.ContainerPidStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.containerPidReturns.result1, fake.containerPidReturns.result2
}

func (fake *ProcessManager) ContainerPidCallCount() int {
	fake.containerPidMutex.RLock()
	defer fake.containerPidMutex.RUnlock()
	return len(fake.containerPidArgsForCall)
}

func (fake *ProcessManager) ContainerPidArgsForCall(i int) string {
	fake.containerPidMutex.RLock()
	defer fake.containerPidMutex.RUnlock()
	return fake.containerPidArgsForCall[i].arg1
}

func (fake *ProcessManager) ContainerPidReturns(result1 int, result2 error) {
	fake.ContainerPidStub = nil
	fake.containerPidReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *ProcessManager) ContainerPidReturnsOnCall(i int, result1 int, result2 error) {
	fake.ContainerPidStub = nil
	if fake.containerPidReturnsOnCall == nil {
		fake.containerPidReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.containerPidReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *ProcessManager) ProcessStartTime(arg1 uint32) (syscall.Filetime, error) {
	fake.processStartTimeMutex.Lock()
	ret, specificReturn := fake.processStartTimeReturnsOnCall[len(fake.processStartTimeArgsForCall)]
	fake.processStartTimeArgsForCall = append(fake.processStartTimeArgsForCall, struct {
		arg1 uint32
	}{arg1})
	fake.recordInvocation("ProcessStartTime", []interface{}{arg1})
	fake.processStartTimeMutex.Unlock()
	if fake.ProcessStartTimeStub != nil {
		return fake.ProcessStartTimeStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.processStartTimeReturns.result1, fake.processStartTimeReturns.result2
}

func (fake *ProcessManager) ProcessStartTimeCallCount() int {
	fake.processStartTimeMutex.RLock()
	defer fake.processStartTimeMutex.RUnlock()
	return len(fake.processStartTimeArgsForCall)
}

func (fake *ProcessManager) ProcessStartTimeArgsForCall(i int) uint32 {
	fake.processStartTimeMutex.RLock()
	defer fake.processStartTimeMutex.RUnlock()
	return fake.processStartTimeArgsForCall[i].arg1
}

func (fake *ProcessManager) ProcessStartTimeReturns(result1 syscall.Filetime, result2 error) {
	fake.ProcessStartTimeStub = nil
	fake.processStartTimeReturns = struct {
		result1 syscall.Filetime
		result2 error
	}{result1, result2}
}

func (fake *ProcessManager) ProcessStartTimeReturnsOnCall(i int, result1 syscall.Filetime, result2 error) {
	fake.ProcessStartTimeStub = nil
	if fake.processStartTimeReturnsOnCall == nil {
		fake.processStartTimeReturnsOnCall = make(map[int]struct {
			result1 syscall.Filetime
			result2 error
		})
	}
	fake.processStartTimeReturnsOnCall[i] = struct {
		result1 syscall.Filetime
		result2 error
	}{result1, result2}
}

func (fake *ProcessManager) IsProcessRunning(arg1 uint32, arg2 syscall.Filetime) (bool, error) {
	fake.isProcessRunningMutex.Lock()
	ret, specificReturn := fake.isProcessRunningReturnsOnCall[len(fake.isProcessRunningArgsForCall)]
	fake.isProcessRunningArgsForCall = append(fake.isProcessRunningArgsForCall, struct {
		arg1 uint32
		arg2 syscall.Filetime
	}{arg1, arg2})
	fake.recordInvocation("IsProcessRunning", []interface{}{arg1, arg2})
	fake.isProcessRunningMutex.Unlock()
	if fake.IsProcessRunningStub != nil {
		return fake.IsProcessRunningStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.isProcessRunningReturns.result1, fake.isProcessRunningReturns.result2
}

func (fake *ProcessManager) IsProcessRunningCallCount() int {
	fake.isProcessRunningMutex.RLock()
	defer fake.isProcessRunningMutex.RUnlock()
	return len(fake.isProcessRunningArgsForCall)
}

func (fake *ProcessManager) IsProcessRunningArgsForCall(i int) (uint32, syscall.Filetime) {
	fake.isProcessRunningMutex.RLock()
	defer fake.isProcessRunningMutex.RUnlock()
	return fake.isProcessRunningArgsForCall[i].arg1, fake.isProcessRunningArgsForCall[i].arg2
}

func (fake *ProcessManager) IsProcessRunningReturns(result1 bool, result2 error) {
	fake.IsProcessRunningStub = nil
	fake.isProcessRunningReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *ProcessManager) IsProcessRunningReturnsOnCall(i int, result1 bool, result2 error) {
	fake.IsProcessRunningStub = nil
	if fake.isProcessRunningReturnsOnCall == nil {
		fake.isProcessRunningReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.isProcessRunningReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *ProcessManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.containerPidMutex.RLock()
	defer fake.containerPidMutex.RUnlock()
	fake.processStartTimeMutex.RLock()
	defer fake.processStartTimeMutex.RUnlock()
	fake.isProcessRunningMutex.RLock()
	defer fake.isProcessRunningMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ProcessManager) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ state.ProcessManager = new(ProcessManager)
