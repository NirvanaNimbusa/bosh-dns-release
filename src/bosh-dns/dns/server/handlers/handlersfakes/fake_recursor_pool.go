// Code generated by counterfeiter. DO NOT EDIT.
package handlersfakes

import (
	"bosh-dns/dns/server/handlers"
	"sync"
)

type FakeRecursorPool struct {
	PerformStrategicallyStub        func(func(string) error) error
	performStrategicallyMutex       sync.RWMutex
	performStrategicallyArgsForCall []struct {
		arg1 func(string) error
	}
	performStrategicallyReturns struct {
		result1 error
	}
	performStrategicallyReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRecursorPool) PerformStrategically(arg1 func(string) error) error {
	fake.performStrategicallyMutex.Lock()
	ret, specificReturn := fake.performStrategicallyReturnsOnCall[len(fake.performStrategicallyArgsForCall)]
	fake.performStrategicallyArgsForCall = append(fake.performStrategicallyArgsForCall, struct {
		arg1 func(string) error
	}{arg1})
	fake.recordInvocation("PerformStrategically", []interface{}{arg1})
	fake.performStrategicallyMutex.Unlock()
	if fake.PerformStrategicallyStub != nil {
		return fake.PerformStrategicallyStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.performStrategicallyReturns.result1
}

func (fake *FakeRecursorPool) PerformStrategicallyCallCount() int {
	fake.performStrategicallyMutex.RLock()
	defer fake.performStrategicallyMutex.RUnlock()
	return len(fake.performStrategicallyArgsForCall)
}

func (fake *FakeRecursorPool) PerformStrategicallyArgsForCall(i int) func(string) error {
	fake.performStrategicallyMutex.RLock()
	defer fake.performStrategicallyMutex.RUnlock()
	return fake.performStrategicallyArgsForCall[i].arg1
}

func (fake *FakeRecursorPool) PerformStrategicallyReturns(result1 error) {
	fake.PerformStrategicallyStub = nil
	fake.performStrategicallyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRecursorPool) PerformStrategicallyReturnsOnCall(i int, result1 error) {
	fake.PerformStrategicallyStub = nil
	if fake.performStrategicallyReturnsOnCall == nil {
		fake.performStrategicallyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.performStrategicallyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRecursorPool) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.performStrategicallyMutex.RLock()
	defer fake.performStrategicallyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRecursorPool) recordInvocation(key string, args []interface{}) {
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

var _ handlers.RecursorPool = new(FakeRecursorPool)
