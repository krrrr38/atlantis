// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/core/runtime (interfaces: VersionedExecutorWorkflow)

package mocks

import (
	go_version "github.com/hashicorp/go-version"
	pegomock "github.com/petergtz/pegomock/v4"
	command "github.com/runatlantis/atlantis/server/events/command"
	logging "github.com/runatlantis/atlantis/server/logging"
	"reflect"
	"time"
)

type MockVersionedExecutorWorkflow struct {
	fail func(message string, callerSkip ...int)
}

func NewMockVersionedExecutorWorkflow(options ...pegomock.Option) *MockVersionedExecutorWorkflow {
	mock := &MockVersionedExecutorWorkflow{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockVersionedExecutorWorkflow) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockVersionedExecutorWorkflow) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockVersionedExecutorWorkflow) EnsureExecutorVersion(_param0 logging.SimpleLogging, _param1 *go_version.Version) (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockVersionedExecutorWorkflow().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("EnsureExecutorVersion", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockVersionedExecutorWorkflow) Run(_param0 command.ProjectContext, _param1 string, _param2 map[string]string, _param3 string, _param4 []string) (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockVersionedExecutorWorkflow().")
	}
	params := []pegomock.Param{_param0, _param1, _param2, _param3, _param4}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Run", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockVersionedExecutorWorkflow) VerifyWasCalledOnce() *VerifierMockVersionedExecutorWorkflow {
	return &VerifierMockVersionedExecutorWorkflow{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockVersionedExecutorWorkflow) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockVersionedExecutorWorkflow {
	return &VerifierMockVersionedExecutorWorkflow{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockVersionedExecutorWorkflow) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockVersionedExecutorWorkflow {
	return &VerifierMockVersionedExecutorWorkflow{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockVersionedExecutorWorkflow) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockVersionedExecutorWorkflow {
	return &VerifierMockVersionedExecutorWorkflow{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockVersionedExecutorWorkflow struct {
	mock                   *MockVersionedExecutorWorkflow
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockVersionedExecutorWorkflow) EnsureExecutorVersion(_param0 logging.SimpleLogging, _param1 *go_version.Version) *MockVersionedExecutorWorkflow_EnsureExecutorVersion_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "EnsureExecutorVersion", params, verifier.timeout)
	return &MockVersionedExecutorWorkflow_EnsureExecutorVersion_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockVersionedExecutorWorkflow_EnsureExecutorVersion_OngoingVerification struct {
	mock              *MockVersionedExecutorWorkflow
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockVersionedExecutorWorkflow_EnsureExecutorVersion_OngoingVerification) GetCapturedArguments() (logging.SimpleLogging, *go_version.Version) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockVersionedExecutorWorkflow_EnsureExecutorVersion_OngoingVerification) GetAllCapturedArguments() (_param0 []logging.SimpleLogging, _param1 []*go_version.Version) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]logging.SimpleLogging, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(logging.SimpleLogging)
		}
		_param1 = make([]*go_version.Version, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(*go_version.Version)
		}
	}
	return
}

func (verifier *VerifierMockVersionedExecutorWorkflow) Run(_param0 command.ProjectContext, _param1 string, _param2 map[string]string, _param3 string, _param4 []string) *MockVersionedExecutorWorkflow_Run_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2, _param3, _param4}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Run", params, verifier.timeout)
	return &MockVersionedExecutorWorkflow_Run_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockVersionedExecutorWorkflow_Run_OngoingVerification struct {
	mock              *MockVersionedExecutorWorkflow
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockVersionedExecutorWorkflow_Run_OngoingVerification) GetCapturedArguments() (command.ProjectContext, string, map[string]string, string, []string) {
	_param0, _param1, _param2, _param3, _param4 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1], _param3[len(_param3)-1], _param4[len(_param4)-1]
}

func (c *MockVersionedExecutorWorkflow_Run_OngoingVerification) GetAllCapturedArguments() (_param0 []command.ProjectContext, _param1 []string, _param2 []map[string]string, _param3 []string, _param4 [][]string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]command.ProjectContext, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(command.ProjectContext)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([]map[string]string, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(map[string]string)
		}
		_param3 = make([]string, len(c.methodInvocations))
		for u, param := range params[3] {
			_param3[u] = param.(string)
		}
		_param4 = make([][]string, len(c.methodInvocations))
		for u, param := range params[4] {
			_param4[u] = param.([]string)
		}
	}
	return
}
