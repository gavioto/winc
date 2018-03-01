// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/winc/container/state"
	"code.cloudfoundry.org/winc/hcs"
	"github.com/Microsoft/hcsshim"
)

type HCSClient struct {
	GetContainerPropertiesStub        func(string) (hcsshim.ContainerProperties, error)
	getContainerPropertiesMutex       sync.RWMutex
	getContainerPropertiesArgsForCall []struct {
		arg1 string
	}
	getContainerPropertiesReturns struct {
		result1 hcsshim.ContainerProperties
		result2 error
	}
	getContainerPropertiesReturnsOnCall map[int]struct {
		result1 hcsshim.ContainerProperties
		result2 error
	}
	OpenContainerStub        func(string) (hcs.Container, error)
	openContainerMutex       sync.RWMutex
	openContainerArgsForCall []struct {
		arg1 string
	}
	openContainerReturns struct {
		result1 hcs.Container
		result2 error
	}
	openContainerReturnsOnCall map[int]struct {
		result1 hcs.Container
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *HCSClient) GetContainerProperties(arg1 string) (hcsshim.ContainerProperties, error) {
	fake.getContainerPropertiesMutex.Lock()
	ret, specificReturn := fake.getContainerPropertiesReturnsOnCall[len(fake.getContainerPropertiesArgsForCall)]
	fake.getContainerPropertiesArgsForCall = append(fake.getContainerPropertiesArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetContainerProperties", []interface{}{arg1})
	fake.getContainerPropertiesMutex.Unlock()
	if fake.GetContainerPropertiesStub != nil {
		return fake.GetContainerPropertiesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getContainerPropertiesReturns.result1, fake.getContainerPropertiesReturns.result2
}

func (fake *HCSClient) GetContainerPropertiesCallCount() int {
	fake.getContainerPropertiesMutex.RLock()
	defer fake.getContainerPropertiesMutex.RUnlock()
	return len(fake.getContainerPropertiesArgsForCall)
}

func (fake *HCSClient) GetContainerPropertiesArgsForCall(i int) string {
	fake.getContainerPropertiesMutex.RLock()
	defer fake.getContainerPropertiesMutex.RUnlock()
	return fake.getContainerPropertiesArgsForCall[i].arg1
}

func (fake *HCSClient) GetContainerPropertiesReturns(result1 hcsshim.ContainerProperties, result2 error) {
	fake.GetContainerPropertiesStub = nil
	fake.getContainerPropertiesReturns = struct {
		result1 hcsshim.ContainerProperties
		result2 error
	}{result1, result2}
}

func (fake *HCSClient) GetContainerPropertiesReturnsOnCall(i int, result1 hcsshim.ContainerProperties, result2 error) {
	fake.GetContainerPropertiesStub = nil
	if fake.getContainerPropertiesReturnsOnCall == nil {
		fake.getContainerPropertiesReturnsOnCall = make(map[int]struct {
			result1 hcsshim.ContainerProperties
			result2 error
		})
	}
	fake.getContainerPropertiesReturnsOnCall[i] = struct {
		result1 hcsshim.ContainerProperties
		result2 error
	}{result1, result2}
}

func (fake *HCSClient) OpenContainer(arg1 string) (hcs.Container, error) {
	fake.openContainerMutex.Lock()
	ret, specificReturn := fake.openContainerReturnsOnCall[len(fake.openContainerArgsForCall)]
	fake.openContainerArgsForCall = append(fake.openContainerArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("OpenContainer", []interface{}{arg1})
	fake.openContainerMutex.Unlock()
	if fake.OpenContainerStub != nil {
		return fake.OpenContainerStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.openContainerReturns.result1, fake.openContainerReturns.result2
}

func (fake *HCSClient) OpenContainerCallCount() int {
	fake.openContainerMutex.RLock()
	defer fake.openContainerMutex.RUnlock()
	return len(fake.openContainerArgsForCall)
}

func (fake *HCSClient) OpenContainerArgsForCall(i int) string {
	fake.openContainerMutex.RLock()
	defer fake.openContainerMutex.RUnlock()
	return fake.openContainerArgsForCall[i].arg1
}

func (fake *HCSClient) OpenContainerReturns(result1 hcs.Container, result2 error) {
	fake.OpenContainerStub = nil
	fake.openContainerReturns = struct {
		result1 hcs.Container
		result2 error
	}{result1, result2}
}

func (fake *HCSClient) OpenContainerReturnsOnCall(i int, result1 hcs.Container, result2 error) {
	fake.OpenContainerStub = nil
	if fake.openContainerReturnsOnCall == nil {
		fake.openContainerReturnsOnCall = make(map[int]struct {
			result1 hcs.Container
			result2 error
		})
	}
	fake.openContainerReturnsOnCall[i] = struct {
		result1 hcs.Container
		result2 error
	}{result1, result2}
}

func (fake *HCSClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getContainerPropertiesMutex.RLock()
	defer fake.getContainerPropertiesMutex.RUnlock()
	fake.openContainerMutex.RLock()
	defer fake.openContainerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *HCSClient) recordInvocation(key string, args []interface{}) {
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

var _ state.HCSClient = new(HCSClient)
