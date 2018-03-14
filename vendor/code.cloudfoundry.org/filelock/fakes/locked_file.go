// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/filelock"
)

type LockedFile struct {
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	ReadStub        func([]byte) (int, error)
	readMutex       sync.RWMutex
	readArgsForCall []struct {
		arg1 []byte
	}
	readReturns struct {
		result1 int
		result2 error
	}
	readReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	TruncateStub        func(int64) error
	truncateMutex       sync.RWMutex
	truncateArgsForCall []struct {
		arg1 int64
	}
	truncateReturns struct {
		result1 error
	}
	truncateReturnsOnCall map[int]struct {
		result1 error
	}
	WriteStub        func([]byte) (int, error)
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		arg1 []byte
	}
	writeReturns struct {
		result1 int
		result2 error
	}
	writeReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	SeekStub        func(int64, int) (int64, error)
	seekMutex       sync.RWMutex
	seekArgsForCall []struct {
		arg1 int64
		arg2 int
	}
	seekReturns struct {
		result1 int64
		result2 error
	}
	seekReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *LockedFile) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.closeReturns.result1
}

func (fake *LockedFile) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *LockedFile) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *LockedFile) CloseReturnsOnCall(i int, result1 error) {
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *LockedFile) Read(arg1 []byte) (int, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.readMutex.Lock()
	ret, specificReturn := fake.readReturnsOnCall[len(fake.readArgsForCall)]
	fake.readArgsForCall = append(fake.readArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	fake.recordInvocation("Read", []interface{}{arg1Copy})
	fake.readMutex.Unlock()
	if fake.ReadStub != nil {
		return fake.ReadStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.readReturns.result1, fake.readReturns.result2
}

func (fake *LockedFile) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *LockedFile) ReadArgsForCall(i int) []byte {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return fake.readArgsForCall[i].arg1
}

func (fake *LockedFile) ReadReturns(result1 int, result2 error) {
	fake.ReadStub = nil
	fake.readReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *LockedFile) ReadReturnsOnCall(i int, result1 int, result2 error) {
	fake.ReadStub = nil
	if fake.readReturnsOnCall == nil {
		fake.readReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.readReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *LockedFile) Truncate(arg1 int64) error {
	fake.truncateMutex.Lock()
	ret, specificReturn := fake.truncateReturnsOnCall[len(fake.truncateArgsForCall)]
	fake.truncateArgsForCall = append(fake.truncateArgsForCall, struct {
		arg1 int64
	}{arg1})
	fake.recordInvocation("Truncate", []interface{}{arg1})
	fake.truncateMutex.Unlock()
	if fake.TruncateStub != nil {
		return fake.TruncateStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.truncateReturns.result1
}

func (fake *LockedFile) TruncateCallCount() int {
	fake.truncateMutex.RLock()
	defer fake.truncateMutex.RUnlock()
	return len(fake.truncateArgsForCall)
}

func (fake *LockedFile) TruncateArgsForCall(i int) int64 {
	fake.truncateMutex.RLock()
	defer fake.truncateMutex.RUnlock()
	return fake.truncateArgsForCall[i].arg1
}

func (fake *LockedFile) TruncateReturns(result1 error) {
	fake.TruncateStub = nil
	fake.truncateReturns = struct {
		result1 error
	}{result1}
}

func (fake *LockedFile) TruncateReturnsOnCall(i int, result1 error) {
	fake.TruncateStub = nil
	if fake.truncateReturnsOnCall == nil {
		fake.truncateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.truncateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *LockedFile) Write(arg1 []byte) (int, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.writeMutex.Lock()
	ret, specificReturn := fake.writeReturnsOnCall[len(fake.writeArgsForCall)]
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	fake.recordInvocation("Write", []interface{}{arg1Copy})
	fake.writeMutex.Unlock()
	if fake.WriteStub != nil {
		return fake.WriteStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.writeReturns.result1, fake.writeReturns.result2
}

func (fake *LockedFile) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *LockedFile) WriteArgsForCall(i int) []byte {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return fake.writeArgsForCall[i].arg1
}

func (fake *LockedFile) WriteReturns(result1 int, result2 error) {
	fake.WriteStub = nil
	fake.writeReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *LockedFile) WriteReturnsOnCall(i int, result1 int, result2 error) {
	fake.WriteStub = nil
	if fake.writeReturnsOnCall == nil {
		fake.writeReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.writeReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *LockedFile) Seek(arg1 int64, arg2 int) (int64, error) {
	fake.seekMutex.Lock()
	ret, specificReturn := fake.seekReturnsOnCall[len(fake.seekArgsForCall)]
	fake.seekArgsForCall = append(fake.seekArgsForCall, struct {
		arg1 int64
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("Seek", []interface{}{arg1, arg2})
	fake.seekMutex.Unlock()
	if fake.SeekStub != nil {
		return fake.SeekStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.seekReturns.result1, fake.seekReturns.result2
}

func (fake *LockedFile) SeekCallCount() int {
	fake.seekMutex.RLock()
	defer fake.seekMutex.RUnlock()
	return len(fake.seekArgsForCall)
}

func (fake *LockedFile) SeekArgsForCall(i int) (int64, int) {
	fake.seekMutex.RLock()
	defer fake.seekMutex.RUnlock()
	return fake.seekArgsForCall[i].arg1, fake.seekArgsForCall[i].arg2
}

func (fake *LockedFile) SeekReturns(result1 int64, result2 error) {
	fake.SeekStub = nil
	fake.seekReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *LockedFile) SeekReturnsOnCall(i int, result1 int64, result2 error) {
	fake.SeekStub = nil
	if fake.seekReturnsOnCall == nil {
		fake.seekReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.seekReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *LockedFile) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	fake.truncateMutex.RLock()
	defer fake.truncateMutex.RUnlock()
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	fake.seekMutex.RLock()
	defer fake.seekMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *LockedFile) recordInvocation(key string, args []interface{}) {
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

var _ filelock.LockedFile = new(LockedFile)