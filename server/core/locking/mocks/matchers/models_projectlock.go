// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	pegomock "github.com/petergtz/pegomock/v4"
	"reflect"

	models "github.com/runatlantis/atlantis/server/events/models"
)

func AnyModelsProjectLock() models.ProjectLock {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(models.ProjectLock))(nil)).Elem()))
	var nullValue models.ProjectLock
	return nullValue
}

func EqModelsProjectLock(value models.ProjectLock) models.ProjectLock {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue models.ProjectLock
	return nullValue
}

func NotEqModelsProjectLock(value models.ProjectLock) models.ProjectLock {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue models.ProjectLock
	return nullValue
}

func ModelsProjectLockThat(matcher pegomock.ArgumentMatcher) models.ProjectLock {
	pegomock.RegisterMatcher(matcher)
	var nullValue models.ProjectLock
	return nullValue
}
