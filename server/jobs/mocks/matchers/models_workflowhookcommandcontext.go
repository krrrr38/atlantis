// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"github.com/petergtz/pegomock"
	"reflect"

	models "github.com/runatlantis/atlantis/server/events/models"
)

func AnyModelsWorkflowHookCommandContext() models.WorkflowHookCommandContext {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(models.WorkflowHookCommandContext))(nil)).Elem()))
	var nullValue models.WorkflowHookCommandContext
	return nullValue
}

func EqModelsWorkflowHookCommandContext(value models.WorkflowHookCommandContext) models.WorkflowHookCommandContext {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue models.WorkflowHookCommandContext
	return nullValue
}

func NotEqModelsWorkflowHookCommandContext(value models.WorkflowHookCommandContext) models.WorkflowHookCommandContext {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue models.WorkflowHookCommandContext
	return nullValue
}

func ModelsWorkflowHookCommandContextThat(matcher pegomock.ArgumentMatcher) models.WorkflowHookCommandContext {
	pegomock.RegisterMatcher(matcher)
	var nullValue models.WorkflowHookCommandContext
	return nullValue
}
