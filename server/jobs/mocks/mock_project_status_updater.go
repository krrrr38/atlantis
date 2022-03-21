// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/jobs (interfaces: ProjectStatusUpdater)

package mocks

import (
	"reflect"
	"time"

	pegomock "github.com/petergtz/pegomock"
	"github.com/runatlantis/atlantis/server/events/command"
	models "github.com/runatlantis/atlantis/server/events/models"
)

type MockProjectStatusUpdater struct {
	fail func(message string, callerSkip ...int)
}

func NewMockProjectStatusUpdater(options ...pegomock.Option) *MockProjectStatusUpdater {
	mock := &MockProjectStatusUpdater{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockProjectStatusUpdater) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockProjectStatusUpdater) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockProjectStatusUpdater) UpdateProject(_param0 command.ProjectContext, _param1 command.Name, _param2 models.CommitStatus, _param3 string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectStatusUpdater().")
	}
	params := []pegomock.Param{_param0, _param1, _param2, _param3}
	result := pegomock.GetGenericMockFrom(mock).Invoke("UpdateProject", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockProjectStatusUpdater) VerifyWasCalledOnce() *VerifierMockProjectStatusUpdater {
	return &VerifierMockProjectStatusUpdater{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockProjectStatusUpdater) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockProjectStatusUpdater {
	return &VerifierMockProjectStatusUpdater{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockProjectStatusUpdater) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockProjectStatusUpdater {
	return &VerifierMockProjectStatusUpdater{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockProjectStatusUpdater) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockProjectStatusUpdater {
	return &VerifierMockProjectStatusUpdater{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockProjectStatusUpdater struct {
	mock                   *MockProjectStatusUpdater
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockProjectStatusUpdater) UpdateProject(_param0 command.ProjectContext, _param1 command.Name, _param2 models.CommitStatus, _param3 string) *MockProjectStatusUpdater_UpdateProject_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2, _param3}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "UpdateProject", params, verifier.timeout)
	return &MockProjectStatusUpdater_UpdateProject_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockProjectStatusUpdater_UpdateProject_OngoingVerification struct {
	mock              *MockProjectStatusUpdater
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockProjectStatusUpdater_UpdateProject_OngoingVerification) GetCapturedArguments() (command.ProjectContext, command.Name, models.CommitStatus, string) {
	_param0, _param1, _param2, _param3 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1], _param3[len(_param3)-1]
}

func (c *MockProjectStatusUpdater_UpdateProject_OngoingVerification) GetAllCapturedArguments() (_param0 []command.ProjectContext, _param1 []command.Name, _param2 []models.CommitStatus, _param3 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]command.ProjectContext, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(command.ProjectContext)
		}
		_param1 = make([]command.Name, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(command.Name)
		}
		_param2 = make([]models.CommitStatus, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(models.CommitStatus)
		}
		_param3 = make([]string, len(c.methodInvocations))
		for u, param := range params[3] {
			_param3[u] = param.(string)
		}
	}
	return
}
