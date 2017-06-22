// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/winc/lib/filelock"
)

type FileLocker struct {
	OpenStub        func() (filelock.LockedFile, error)
	openMutex       sync.RWMutex
	openArgsForCall []struct{}
	openReturns     struct {
		result1 filelock.LockedFile
		result2 error
	}
	openReturnsOnCall map[int]struct {
		result1 filelock.LockedFile
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FileLocker) Open() (filelock.LockedFile, error) {
	fake.openMutex.Lock()
	ret, specificReturn := fake.openReturnsOnCall[len(fake.openArgsForCall)]
	fake.openArgsForCall = append(fake.openArgsForCall, struct{}{})
	fake.recordInvocation("Open", []interface{}{})
	fake.openMutex.Unlock()
	if fake.OpenStub != nil {
		return fake.OpenStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.openReturns.result1, fake.openReturns.result2
}

func (fake *FileLocker) OpenCallCount() int {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	return len(fake.openArgsForCall)
}

func (fake *FileLocker) OpenReturns(result1 filelock.LockedFile, result2 error) {
	fake.OpenStub = nil
	fake.openReturns = struct {
		result1 filelock.LockedFile
		result2 error
	}{result1, result2}
}

func (fake *FileLocker) OpenReturnsOnCall(i int, result1 filelock.LockedFile, result2 error) {
	fake.OpenStub = nil
	if fake.openReturnsOnCall == nil {
		fake.openReturnsOnCall = make(map[int]struct {
			result1 filelock.LockedFile
			result2 error
		})
	}
	fake.openReturnsOnCall[i] = struct {
		result1 filelock.LockedFile
		result2 error
	}{result1, result2}
}

func (fake *FileLocker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FileLocker) recordInvocation(key string, args []interface{}) {
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

var _ filelock.FileLocker = new(FileLocker)
