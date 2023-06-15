// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	pegomock "github.com/petergtz/pegomock/v4"
	"reflect"

	command "github.com/runatlantis/atlantis/server/events/command"
)

func AnyPtrToCommandLock() *command.Lock {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*command.Lock))(nil)).Elem()))
	var nullValue *command.Lock
	return nullValue
}

func EqPtrToCommandLock(value *command.Lock) *command.Lock {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *command.Lock
	return nullValue
}

func NotEqPtrToCommandLock(value *command.Lock) *command.Lock {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue *command.Lock
	return nullValue
}

func PtrToCommandLockThat(matcher pegomock.ArgumentMatcher) *command.Lock {
	pegomock.RegisterMatcher(matcher)
	var nullValue *command.Lock
	return nullValue
}
