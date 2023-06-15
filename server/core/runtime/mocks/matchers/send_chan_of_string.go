// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	pegomock "github.com/petergtz/pegomock/v4"
	"reflect"
)

func AnySendChanOfString() chan<- string {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(chan<- string))(nil)).Elem()))
	var nullValue chan<- string
	return nullValue
}

func EqSendChanOfString(value chan<- string) chan<- string {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue chan<- string
	return nullValue
}

func NotEqSendChanOfString(value chan<- string) chan<- string {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue chan<- string
	return nullValue
}

func SendChanOfStringThat(matcher pegomock.ArgumentMatcher) chan<- string {
	pegomock.RegisterMatcher(matcher)
	var nullValue chan<- string
	return nullValue
}
