package tests

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/stretchr/testify/assert"
)

var ErrFieldDoesNotExist = errors.New("field does not exist")
var ErrFieldDoesNotMatch = errors.New("field does not match")

func CompareFields(a, b interface{}, fields []string) error {
	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)

	for _, field := range fields {
		fa := va.FieldByName(field)
		fb := vb.FieldByName(field)

		if !fa.IsValid() || !fb.IsValid() {
			return fmt.Errorf("%w: %s", ErrFieldDoesNotExist, field)
		}

		if fa.Interface() != fb.Interface() {
			return fmt.Errorf("%w: %s", ErrFieldDoesNotMatch, field)
		}
	}

	return nil
}

// assertExpectedAndActual is a helper function to allow the step function to call
// assertion functions where you want to compare an expected and an actual value.

func AssertExpectedAndActual(a expectedAndActualAssertion, expected, actual interface{}, msgAndArgs ...interface{}) error {
	var t asserter

	a(&t, expected, actual, msgAndArgs...)

	return t.err
}

type expectedAndActualAssertion func(t assert.TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool

// assertActual is a helper function to allow the step function to call
// assertion functions where you want to compare an actual value to a
// predined state like nil, empty or true/false.
func AssertActual(a actualAssertion, actual interface{}, msgAndArgs ...interface{}) error {
	var t asserter

	a(&t, actual, msgAndArgs...)

	return t.err
}

type actualAssertion func(t assert.TestingT, actual interface{}, msgAndArgs ...interface{}) bool

// asserter is used to be able to retrieve the error reported by the called assertion.
type asserter struct {
	err error
}

var ErrAssertionFailed = errors.New("assertion failed")

// Errorf is used by the called assertion to report an error.
func (a *asserter) Errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	a.err = fmt.Errorf("%w: %s", ErrAssertionFailed, msg)
}
