// Code generated from https://github.com/stretchr/testify/archive/a53be35c3b0cfcd5189cffcfd75df60ea581104c.zip; DO NOT EDIT.

package allure

import (
	"cmp"
	"fmt"
	testo "github.com/ozontech/testo"
	stacktrace "github.com/ozontech/testo-allure/internal/stacktrace"
	assert "github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"runtime"
	"time"
)

// Condition uses a Comparison to assert a complex condition.
func (x Requirements) Condition(comp assert.Comparison, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "condition")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "condition"), NewParameter("comp", asShortString(comp)).withMode(x.mode))
		}()
		if assert.Condition(t, comp, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// Condition uses a Comparison to assert a complex condition.
func (x Assertions) Condition(comp assert.Comparison, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "condition")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "condition"), NewParameter("comp", asShortString(comp)).withMode(x.mode))
		}()
		if assert.Condition(t, comp, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// Contains asserts that the specified string, list(array, slice...) or map contains the
// specified substring or element.
//
// 	assert.Contains(t, "Hello World", "World")
// 	assert.Contains(t, ["Hello", "World"], "World")
// 	assert.Contains(t, {"Hello": "World"}, "Hello")
func (x Requirements) Contains(s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "contains")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "contains"), NewParameter("s", asShortString(s)).withMode(x.mode), NewParameter("contains", asShortString(contains)).withMode(x.mode))
		}()
		if assert.Contains(t, s, contains, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// Contains asserts that the specified string, list(array, slice...) or map contains the
// specified substring or element.
//
// 	assert.Contains(t, "Hello World", "World")
// 	assert.Contains(t, ["Hello", "World"], "World")
// 	assert.Contains(t, {"Hello": "World"}, "Hello")
func (x Assertions) Contains(s interface{}, contains interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "contains")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "contains"), NewParameter("s", asShortString(s)).withMode(x.mode), NewParameter("contains", asShortString(contains)).withMode(x.mode))
		}()
		if assert.Contains(t, s, contains, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// DirExists checks whether a directory exists in the given path. It also fails
// if the path is a file rather a directory or there is an error checking whether it exists.
func (x Requirements) DirExists(path string, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "dir exists")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "dir exists"), NewParameter("path", asShortString(path)).withMode(x.mode))
		}()
		if assert.DirExists(t, path, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// DirExists checks whether a directory exists in the given path. It also fails
// if the path is a file rather a directory or there is an error checking whether it exists.
func (x Assertions) DirExists(path string, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "dir exists")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "dir exists"), NewParameter("path", asShortString(path)).withMode(x.mode))
		}()
		if assert.DirExists(t, path, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// ElementsMatch asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
//
// assert.ElementsMatch(t, [1, 3, 2, 3], [1, 3, 3, 2])
func (x Requirements) ElementsMatch(listA interface{}, listB interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "elements match")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "elements match"), NewParameter("list A", asShortString(listA)).withMode(x.mode), NewParameter("list B", asShortString(listB)).withMode(x.mode))
		}()
		if assert.ElementsMatch(t, listA, listB, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// ElementsMatch asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
//
// assert.ElementsMatch(t, [1, 3, 2, 3], [1, 3, 3, 2])
func (x Assertions) ElementsMatch(listA interface{}, listB interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "elements match")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "elements match"), NewParameter("list A", asShortString(listA)).withMode(x.mode), NewParameter("list B", asShortString(listB)).withMode(x.mode))
		}()
		if assert.ElementsMatch(t, listA, listB, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// Empty asserts that the given value is "empty".
//
// [Zero values] are "empty".
//
// Arrays are "empty" if every element is the zero value of the type (stricter than "empty").
//
// Slices, maps and channels with zero length are "empty".
//
// Pointer values are "empty" if the pointer is nil or if the pointed value is "empty".
//
// 	assert.Empty(t, obj)
//
// [Zero values]: https://go.dev/ref/spec#The_zero_value
func (x Requirements) Empty(object interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "empty")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "empty"), NewParameter("object", asShortString(object)).withMode(x.mode))
		}()
		if assert.Empty(t, object, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// Empty asserts that the given value is "empty".
//
// [Zero values] are "empty".
//
// Arrays are "empty" if every element is the zero value of the type (stricter than "empty").
//
// Slices, maps and channels with zero length are "empty".
//
// Pointer values are "empty" if the pointer is nil or if the pointed value is "empty".
//
// 	assert.Empty(t, obj)
//
// [Zero values]: https://go.dev/ref/spec#The_zero_value
func (x Assertions) Empty(object interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "empty")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "empty"), NewParameter("object", asShortString(object)).withMode(x.mode))
		}()
		if assert.Empty(t, object, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// Equal asserts that two objects are equal.
//
// 	assert.Equal(t, 123, 123)
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses). Function equality
// cannot be determined and will always fail.
func (x Requirements) Equal(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "equal")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "equal"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode))
		}()
		if assert.Equal(t, expected, actual, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// Equal asserts that two objects are equal.
//
// 	assert.Equal(t, 123, 123)
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses). Function equality
// cannot be determined and will always fail.
func (x Assertions) Equal(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "equal")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "equal"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode))
		}()
		if assert.Equal(t, expected, actual, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// EqualError asserts that a function returned an error (i.e. not `nil`)
// and that it is equal to the provided error.
//
// 	actualObj, err := SomeFunction()
// 	assert.EqualError(t, err,  expectedErrorString)
func (x Requirements) EqualError(theError error, errString string, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "equal error")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "equal error"), NewParameter("the error", asShortString(theError)).withMode(x.mode), NewParameter("err string", asShortString(errString)).withMode(x.mode))
		}()
		if assert.EqualError(t, theError, errString, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// EqualError asserts that a function returned an error (i.e. not `nil`)
// and that it is equal to the provided error.
//
// 	actualObj, err := SomeFunction()
// 	assert.EqualError(t, err,  expectedErrorString)
func (x Assertions) EqualError(theError error, errString string, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "equal error")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "equal error"), NewParameter("the error", asShortString(theError)).withMode(x.mode), NewParameter("err string", asShortString(errString)).withMode(x.mode))
		}()
		if assert.EqualError(t, theError, errString, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// EqualExportedValues asserts that the types of two objects are equal and their public
// fields are also equal. This is useful for comparing structs that have private fields
// that could potentially differ.
//
// 	 type S struct {
// 		Exported     	int
// 		notExported   	int
// 	 }
// 	 assert.EqualExportedValues(t, S{1, 2}, S{1, 3}) => true
// 	 assert.EqualExportedValues(t, S{1, 2}, S{2, 3}) => false
func (x Requirements) EqualExportedValues(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "equal exported values")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "equal exported values"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode))
		}()
		if assert.EqualExportedValues(t, expected, actual, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// EqualExportedValues asserts that the types of two objects are equal and their public
// fields are also equal. This is useful for comparing structs that have private fields
// that could potentially differ.
//
// 	 type S struct {
// 		Exported     	int
// 		notExported   	int
// 	 }
// 	 assert.EqualExportedValues(t, S{1, 2}, S{1, 3}) => true
// 	 assert.EqualExportedValues(t, S{1, 2}, S{2, 3}) => false
func (x Assertions) EqualExportedValues(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "equal exported values")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "equal exported values"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode))
		}()
		if assert.EqualExportedValues(t, expected, actual, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// EqualValues asserts that two objects are equal or convertible to the larger
// type and equal.
//
// 	assert.EqualValues(t, uint32(123), int32(123))
func (x Requirements) EqualValues(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "equal values")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "equal values"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode))
		}()
		if assert.EqualValues(t, expected, actual, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// EqualValues asserts that two objects are equal or convertible to the larger
// type and equal.
//
// 	assert.EqualValues(t, uint32(123), int32(123))
func (x Assertions) EqualValues(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "equal values")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "equal values"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode))
		}()
		if assert.EqualValues(t, expected, actual, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// Error asserts that a function returned an error (i.e. not `nil`).
//
// 	actualObj, err := SomeFunction()
// 	assert.Error(t, err)
func (x Requirements) Error(err error, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "error")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "error"), NewParameter("err", asShortString(err)).withMode(x.mode))
		}()
		if assert.Error(t, err, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// Error asserts that a function returned an error (i.e. not `nil`).
//
// 	actualObj, err := SomeFunction()
// 	assert.Error(t, err)
func (x Assertions) Error(err error, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "error")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "error"), NewParameter("err", asShortString(err)).withMode(x.mode))
		}()
		if assert.Error(t, err, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// ErrorAs asserts that at least one of the errors in err's chain matches target, and if so, sets target to that error value.
// This is a wrapper for errors.As.
func (x Requirements) ErrorAs(err error, target interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "error as")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "error as"), NewParameter("err", asShortString(err)).withMode(x.mode), NewParameter("target", asShortString(target)).withMode(x.mode))
		}()
		if assert.ErrorAs(t, err, target, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// ErrorAs asserts that at least one of the errors in err's chain matches target, and if so, sets target to that error value.
// This is a wrapper for errors.As.
func (x Assertions) ErrorAs(err error, target interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "error as")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "error as"), NewParameter("err", asShortString(err)).withMode(x.mode), NewParameter("target", asShortString(target)).withMode(x.mode))
		}()
		if assert.ErrorAs(t, err, target, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// ErrorContains asserts that a function returned an error (i.e. not `nil`)
// and that the error contains the specified substring.
//
// 	actualObj, err := SomeFunction()
// 	assert.ErrorContains(t, err,  expectedErrorSubString)
func (x Requirements) ErrorContains(theError error, contains string, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "error contains")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "error contains"), NewParameter("the error", asShortString(theError)).withMode(x.mode), NewParameter("contains", asShortString(contains)).withMode(x.mode))
		}()
		if assert.ErrorContains(t, theError, contains, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// ErrorContains asserts that a function returned an error (i.e. not `nil`)
// and that the error contains the specified substring.
//
// 	actualObj, err := SomeFunction()
// 	assert.ErrorContains(t, err,  expectedErrorSubString)
func (x Assertions) ErrorContains(theError error, contains string, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "error contains")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "error contains"), NewParameter("the error", asShortString(theError)).withMode(x.mode), NewParameter("contains", asShortString(contains)).withMode(x.mode))
		}()
		if assert.ErrorContains(t, theError, contains, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// ErrorIs asserts that at least one of the errors in err's chain matches target.
// This is a wrapper for errors.Is.
func (x Requirements) ErrorIs(err error, target error, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "error is")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "error is"), NewParameter("err", asShortString(err)).withMode(x.mode), NewParameter("target", asShortString(target)).withMode(x.mode))
		}()
		if assert.ErrorIs(t, err, target, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// ErrorIs asserts that at least one of the errors in err's chain matches target.
// This is a wrapper for errors.Is.
func (x Assertions) ErrorIs(err error, target error, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "error is")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "error is"), NewParameter("err", asShortString(err)).withMode(x.mode), NewParameter("target", asShortString(target)).withMode(x.mode))
		}()
		if assert.ErrorIs(t, err, target, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// Eventually asserts that given condition will be met in waitFor time,
// periodically checking target function each tick.
//
// 	assert.Eventually(t, func() bool { return true; }, time.Second, 10*time.Millisecond)
func (x Requirements) Eventually(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "eventually")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "eventually"), NewParameter("condition", asShortString(condition)).withMode(x.mode), NewParameter("wait for", asShortString(waitFor)).withMode(x.mode), NewParameter("tick", asShortString(tick)).withMode(x.mode))
		}()
		if assert.Eventually(t, condition, waitFor, tick, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// Eventually asserts that given condition will be met in waitFor time,
// periodically checking target function each tick.
//
// 	assert.Eventually(t, func() bool { return true; }, time.Second, 10*time.Millisecond)
func (x Assertions) Eventually(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "eventually")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "eventually"), NewParameter("condition", asShortString(condition)).withMode(x.mode), NewParameter("wait for", asShortString(waitFor)).withMode(x.mode), NewParameter("tick", asShortString(tick)).withMode(x.mode))
		}()
		if assert.Eventually(t, condition, waitFor, tick, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// Exactly asserts that two objects are equal in value and type.
//
// 	assert.Exactly(t, int32(123), int64(123))
func (x Requirements) Exactly(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "exactly")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "exactly"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode))
		}()
		if assert.Exactly(t, expected, actual, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// Exactly asserts that two objects are equal in value and type.
//
// 	assert.Exactly(t, int32(123), int64(123))
func (x Assertions) Exactly(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "exactly")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "exactly"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode))
		}()
		if assert.Exactly(t, expected, actual, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// False asserts that the specified value is false.
//
// 	assert.False(t, myBool)
func (x Requirements) False(value bool, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "false")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "false"), NewParameter("value", asShortString(value)).withMode(x.mode))
		}()
		if assert.False(t, value, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// False asserts that the specified value is false.
//
// 	assert.False(t, myBool)
func (x Assertions) False(value bool, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "false")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "false"), NewParameter("value", asShortString(value)).withMode(x.mode))
		}()
		if assert.False(t, value, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// FileExists checks whether a file exists in the given path. It also fails if
// the path points to a directory or there is an error when trying to check the file.
func (x Requirements) FileExists(path string, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "file exists")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "file exists"), NewParameter("path", asShortString(path)).withMode(x.mode))
		}()
		if assert.FileExists(t, path, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// FileExists checks whether a file exists in the given path. It also fails if
// the path points to a directory or there is an error when trying to check the file.
func (x Assertions) FileExists(path string, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "file exists")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "file exists"), NewParameter("path", asShortString(path)).withMode(x.mode))
		}()
		if assert.FileExists(t, path, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// Greater asserts that the first element is greater than the second
//
// 	assert.Greater(t, 2, 1)
// 	assert.Greater(t, float64(2), float64(1))
// 	assert.Greater(t, "b", "a")
func (x Requirements) Greater(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "greater")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "greater"), NewParameter("e 1", asShortString(e1)).withMode(x.mode), NewParameter("e 2", asShortString(e2)).withMode(x.mode))
		}()
		if assert.Greater(t, e1, e2, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// Greater asserts that the first element is greater than the second
//
// 	assert.Greater(t, 2, 1)
// 	assert.Greater(t, float64(2), float64(1))
// 	assert.Greater(t, "b", "a")
func (x Assertions) Greater(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "greater")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "greater"), NewParameter("e 1", asShortString(e1)).withMode(x.mode), NewParameter("e 2", asShortString(e2)).withMode(x.mode))
		}()
		if assert.Greater(t, e1, e2, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// GreaterOrEqual asserts that the first element is greater than or equal to the second
//
// 	assert.GreaterOrEqual(t, 2, 1)
// 	assert.GreaterOrEqual(t, 2, 2)
// 	assert.GreaterOrEqual(t, "b", "a")
// 	assert.GreaterOrEqual(t, "b", "b")
func (x Requirements) GreaterOrEqual(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "greater or equal")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "greater or equal"), NewParameter("e 1", asShortString(e1)).withMode(x.mode), NewParameter("e 2", asShortString(e2)).withMode(x.mode))
		}()
		if assert.GreaterOrEqual(t, e1, e2, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// GreaterOrEqual asserts that the first element is greater than or equal to the second
//
// 	assert.GreaterOrEqual(t, 2, 1)
// 	assert.GreaterOrEqual(t, 2, 2)
// 	assert.GreaterOrEqual(t, "b", "a")
// 	assert.GreaterOrEqual(t, "b", "b")
func (x Assertions) GreaterOrEqual(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "greater or equal")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "greater or equal"), NewParameter("e 1", asShortString(e1)).withMode(x.mode), NewParameter("e 2", asShortString(e2)).withMode(x.mode))
		}()
		if assert.GreaterOrEqual(t, e1, e2, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// HTTPBodyContains asserts that a specified handler returns a
// body that contains a string.
//
// 	assert.HTTPBodyContains(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
//
// Returns whether the assertion was successful (true) or not (false).
func (x Requirements) HTTPBodyContains(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "HTTP body contains")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "HTTP body contains"), NewParameter("handler", asShortString(handler)).withMode(x.mode), NewParameter("method", asShortString(method)).withMode(x.mode), NewParameter("url", asShortString(url)).withMode(x.mode), NewParameter("values", asShortString(values)).withMode(x.mode), NewParameter("str", asShortString(str)).withMode(x.mode))
		}()
		if assert.HTTPBodyContains(t, handler, method, url, values, str, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// HTTPBodyContains asserts that a specified handler returns a
// body that contains a string.
//
// 	assert.HTTPBodyContains(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
//
// Returns whether the assertion was successful (true) or not (false).
func (x Assertions) HTTPBodyContains(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "HTTP body contains")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "HTTP body contains"), NewParameter("handler", asShortString(handler)).withMode(x.mode), NewParameter("method", asShortString(method)).withMode(x.mode), NewParameter("url", asShortString(url)).withMode(x.mode), NewParameter("values", asShortString(values)).withMode(x.mode), NewParameter("str", asShortString(str)).withMode(x.mode))
		}()
		if assert.HTTPBodyContains(t, handler, method, url, values, str, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// HTTPBodyNotContains asserts that a specified handler returns a
// body that does not contain a string.
//
// 	assert.HTTPBodyNotContains(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
//
// Returns whether the assertion was successful (true) or not (false).
func (x Requirements) HTTPBodyNotContains(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "HTTP body not contains")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "HTTP body not contains"), NewParameter("handler", asShortString(handler)).withMode(x.mode), NewParameter("method", asShortString(method)).withMode(x.mode), NewParameter("url", asShortString(url)).withMode(x.mode), NewParameter("values", asShortString(values)).withMode(x.mode), NewParameter("str", asShortString(str)).withMode(x.mode))
		}()
		if assert.HTTPBodyNotContains(t, handler, method, url, values, str, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// HTTPBodyNotContains asserts that a specified handler returns a
// body that does not contain a string.
//
// 	assert.HTTPBodyNotContains(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
//
// Returns whether the assertion was successful (true) or not (false).
func (x Assertions) HTTPBodyNotContains(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "HTTP body not contains")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "HTTP body not contains"), NewParameter("handler", asShortString(handler)).withMode(x.mode), NewParameter("method", asShortString(method)).withMode(x.mode), NewParameter("url", asShortString(url)).withMode(x.mode), NewParameter("values", asShortString(values)).withMode(x.mode), NewParameter("str", asShortString(str)).withMode(x.mode))
		}()
		if assert.HTTPBodyNotContains(t, handler, method, url, values, str, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// HTTPError asserts that a specified handler returns an error status code.
//
// 	assert.HTTPError(t, myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
func (x Requirements) HTTPError(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "HTTP error")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "HTTP error"), NewParameter("handler", asShortString(handler)).withMode(x.mode), NewParameter("method", asShortString(method)).withMode(x.mode), NewParameter("url", asShortString(url)).withMode(x.mode), NewParameter("values", asShortString(values)).withMode(x.mode))
		}()
		if assert.HTTPError(t, handler, method, url, values, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// HTTPError asserts that a specified handler returns an error status code.
//
// 	assert.HTTPError(t, myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
func (x Assertions) HTTPError(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "HTTP error")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "HTTP error"), NewParameter("handler", asShortString(handler)).withMode(x.mode), NewParameter("method", asShortString(method)).withMode(x.mode), NewParameter("url", asShortString(url)).withMode(x.mode), NewParameter("values", asShortString(values)).withMode(x.mode))
		}()
		if assert.HTTPError(t, handler, method, url, values, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// HTTPRedirect asserts that a specified handler returns a redirect status code.
//
// 	assert.HTTPRedirect(t, myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
func (x Requirements) HTTPRedirect(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "HTTP redirect")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "HTTP redirect"), NewParameter("handler", asShortString(handler)).withMode(x.mode), NewParameter("method", asShortString(method)).withMode(x.mode), NewParameter("url", asShortString(url)).withMode(x.mode), NewParameter("values", asShortString(values)).withMode(x.mode))
		}()
		if assert.HTTPRedirect(t, handler, method, url, values, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// HTTPRedirect asserts that a specified handler returns a redirect status code.
//
// 	assert.HTTPRedirect(t, myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
func (x Assertions) HTTPRedirect(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "HTTP redirect")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "HTTP redirect"), NewParameter("handler", asShortString(handler)).withMode(x.mode), NewParameter("method", asShortString(method)).withMode(x.mode), NewParameter("url", asShortString(url)).withMode(x.mode), NewParameter("values", asShortString(values)).withMode(x.mode))
		}()
		if assert.HTTPRedirect(t, handler, method, url, values, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// HTTPStatusCode asserts that a specified handler returns a specified status code.
//
// 	assert.HTTPStatusCode(t, myHandler, "GET", "/notImplemented", nil, 501)
//
// Returns whether the assertion was successful (true) or not (false).
func (x Requirements) HTTPStatusCode(handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "HTTP status code")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "HTTP status code"), NewParameter("handler", asShortString(handler)).withMode(x.mode), NewParameter("method", asShortString(method)).withMode(x.mode), NewParameter("url", asShortString(url)).withMode(x.mode), NewParameter("values", asShortString(values)).withMode(x.mode), NewParameter("statuscode", asShortString(statuscode)).withMode(x.mode))
		}()
		if assert.HTTPStatusCode(t, handler, method, url, values, statuscode, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// HTTPStatusCode asserts that a specified handler returns a specified status code.
//
// 	assert.HTTPStatusCode(t, myHandler, "GET", "/notImplemented", nil, 501)
//
// Returns whether the assertion was successful (true) or not (false).
func (x Assertions) HTTPStatusCode(handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "HTTP status code")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "HTTP status code"), NewParameter("handler", asShortString(handler)).withMode(x.mode), NewParameter("method", asShortString(method)).withMode(x.mode), NewParameter("url", asShortString(url)).withMode(x.mode), NewParameter("values", asShortString(values)).withMode(x.mode), NewParameter("statuscode", asShortString(statuscode)).withMode(x.mode))
		}()
		if assert.HTTPStatusCode(t, handler, method, url, values, statuscode, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// HTTPSuccess asserts that a specified handler returns a success status code.
//
// 	assert.HTTPSuccess(t, myHandler, "POST", "http://www.google.com", nil)
//
// Returns whether the assertion was successful (true) or not (false).
func (x Requirements) HTTPSuccess(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "HTTP success")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "HTTP success"), NewParameter("handler", asShortString(handler)).withMode(x.mode), NewParameter("method", asShortString(method)).withMode(x.mode), NewParameter("url", asShortString(url)).withMode(x.mode), NewParameter("values", asShortString(values)).withMode(x.mode))
		}()
		if assert.HTTPSuccess(t, handler, method, url, values, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// HTTPSuccess asserts that a specified handler returns a success status code.
//
// 	assert.HTTPSuccess(t, myHandler, "POST", "http://www.google.com", nil)
//
// Returns whether the assertion was successful (true) or not (false).
func (x Assertions) HTTPSuccess(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "HTTP success")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "HTTP success"), NewParameter("handler", asShortString(handler)).withMode(x.mode), NewParameter("method", asShortString(method)).withMode(x.mode), NewParameter("url", asShortString(url)).withMode(x.mode), NewParameter("values", asShortString(values)).withMode(x.mode))
		}()
		if assert.HTTPSuccess(t, handler, method, url, values, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// Implements asserts that an object is implemented by the specified interface.
//
// 	assert.Implements(t, (*MyInterface)(nil), new(MyObject))
func (x Requirements) Implements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "implements")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "implements"), NewParameter("interface object", asShortString(interfaceObject)).withMode(x.mode), NewParameter("object", asShortString(object)).withMode(x.mode))
		}()
		if assert.Implements(t, interfaceObject, object, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// Implements asserts that an object is implemented by the specified interface.
//
// 	assert.Implements(t, (*MyInterface)(nil), new(MyObject))
func (x Assertions) Implements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "implements")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "implements"), NewParameter("interface object", asShortString(interfaceObject)).withMode(x.mode), NewParameter("object", asShortString(object)).withMode(x.mode))
		}()
		if assert.Implements(t, interfaceObject, object, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// InDelta asserts that the two numerals are within delta of each other.
//
// 	assert.InDelta(t, math.Pi, 22/7.0, 0.01)
func (x Requirements) InDelta(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "in delta")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "in delta"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode), NewParameter("delta", asShortString(delta)).withMode(x.mode))
		}()
		if assert.InDelta(t, expected, actual, delta, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// InDelta asserts that the two numerals are within delta of each other.
//
// 	assert.InDelta(t, math.Pi, 22/7.0, 0.01)
func (x Assertions) InDelta(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "in delta")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "in delta"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode), NewParameter("delta", asShortString(delta)).withMode(x.mode))
		}()
		if assert.InDelta(t, expected, actual, delta, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// InDeltaMapValues is the same as InDelta, but it compares all values between two maps. Both maps must have exactly the same keys.
func (x Requirements) InDeltaMapValues(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "in delta map values")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "in delta map values"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode), NewParameter("delta", asShortString(delta)).withMode(x.mode))
		}()
		if assert.InDeltaMapValues(t, expected, actual, delta, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// InDeltaMapValues is the same as InDelta, but it compares all values between two maps. Both maps must have exactly the same keys.
func (x Assertions) InDeltaMapValues(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "in delta map values")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "in delta map values"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode), NewParameter("delta", asShortString(delta)).withMode(x.mode))
		}()
		if assert.InDeltaMapValues(t, expected, actual, delta, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// InDeltaSlice is the same as InDelta, except it compares two slices.
func (x Requirements) InDeltaSlice(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "in delta slice")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "in delta slice"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode), NewParameter("delta", asShortString(delta)).withMode(x.mode))
		}()
		if assert.InDeltaSlice(t, expected, actual, delta, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// InDeltaSlice is the same as InDelta, except it compares two slices.
func (x Assertions) InDeltaSlice(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "in delta slice")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "in delta slice"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode), NewParameter("delta", asShortString(delta)).withMode(x.mode))
		}()
		if assert.InDeltaSlice(t, expected, actual, delta, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// InEpsilon asserts that expected and actual have a relative error less than epsilon
func (x Requirements) InEpsilon(expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "in epsilon")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "in epsilon"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode), NewParameter("epsilon", asShortString(epsilon)).withMode(x.mode))
		}()
		if assert.InEpsilon(t, expected, actual, epsilon, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// InEpsilon asserts that expected and actual have a relative error less than epsilon
func (x Assertions) InEpsilon(expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "in epsilon")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "in epsilon"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode), NewParameter("epsilon", asShortString(epsilon)).withMode(x.mode))
		}()
		if assert.InEpsilon(t, expected, actual, epsilon, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// InEpsilonSlice is the same as InEpsilon, except it compares each value from two slices.
func (x Requirements) InEpsilonSlice(expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "in epsilon slice")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "in epsilon slice"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode), NewParameter("epsilon", asShortString(epsilon)).withMode(x.mode))
		}()
		if assert.InEpsilonSlice(t, expected, actual, epsilon, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// InEpsilonSlice is the same as InEpsilon, except it compares each value from two slices.
func (x Assertions) InEpsilonSlice(expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "in epsilon slice")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "in epsilon slice"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode), NewParameter("epsilon", asShortString(epsilon)).withMode(x.mode))
		}()
		if assert.InEpsilonSlice(t, expected, actual, epsilon, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// IsDecreasing asserts that the collection is decreasing
//
// 	assert.IsDecreasing(t, []int{2, 1, 0})
// 	assert.IsDecreasing(t, []float{2, 1})
// 	assert.IsDecreasing(t, []string{"b", "a"})
func (x Requirements) IsDecreasing(object interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "is decreasing")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "is decreasing"), NewParameter("object", asShortString(object)).withMode(x.mode))
		}()
		if assert.IsDecreasing(t, object, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// IsDecreasing asserts that the collection is decreasing
//
// 	assert.IsDecreasing(t, []int{2, 1, 0})
// 	assert.IsDecreasing(t, []float{2, 1})
// 	assert.IsDecreasing(t, []string{"b", "a"})
func (x Assertions) IsDecreasing(object interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "is decreasing")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "is decreasing"), NewParameter("object", asShortString(object)).withMode(x.mode))
		}()
		if assert.IsDecreasing(t, object, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// IsIncreasing asserts that the collection is increasing
//
// 	assert.IsIncreasing(t, []int{1, 2, 3})
// 	assert.IsIncreasing(t, []float{1, 2})
// 	assert.IsIncreasing(t, []string{"a", "b"})
func (x Requirements) IsIncreasing(object interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "is increasing")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "is increasing"), NewParameter("object", asShortString(object)).withMode(x.mode))
		}()
		if assert.IsIncreasing(t, object, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// IsIncreasing asserts that the collection is increasing
//
// 	assert.IsIncreasing(t, []int{1, 2, 3})
// 	assert.IsIncreasing(t, []float{1, 2})
// 	assert.IsIncreasing(t, []string{"a", "b"})
func (x Assertions) IsIncreasing(object interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "is increasing")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "is increasing"), NewParameter("object", asShortString(object)).withMode(x.mode))
		}()
		if assert.IsIncreasing(t, object, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// IsNonDecreasing asserts that the collection is not decreasing
//
// 	assert.IsNonDecreasing(t, []int{1, 1, 2})
// 	assert.IsNonDecreasing(t, []float{1, 2})
// 	assert.IsNonDecreasing(t, []string{"a", "b"})
func (x Requirements) IsNonDecreasing(object interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "is non decreasing")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "is non decreasing"), NewParameter("object", asShortString(object)).withMode(x.mode))
		}()
		if assert.IsNonDecreasing(t, object, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// IsNonDecreasing asserts that the collection is not decreasing
//
// 	assert.IsNonDecreasing(t, []int{1, 1, 2})
// 	assert.IsNonDecreasing(t, []float{1, 2})
// 	assert.IsNonDecreasing(t, []string{"a", "b"})
func (x Assertions) IsNonDecreasing(object interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "is non decreasing")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "is non decreasing"), NewParameter("object", asShortString(object)).withMode(x.mode))
		}()
		if assert.IsNonDecreasing(t, object, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// IsNonIncreasing asserts that the collection is not increasing
//
// 	assert.IsNonIncreasing(t, []int{2, 1, 1})
// 	assert.IsNonIncreasing(t, []float{2, 1})
// 	assert.IsNonIncreasing(t, []string{"b", "a"})
func (x Requirements) IsNonIncreasing(object interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "is non increasing")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "is non increasing"), NewParameter("object", asShortString(object)).withMode(x.mode))
		}()
		if assert.IsNonIncreasing(t, object, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// IsNonIncreasing asserts that the collection is not increasing
//
// 	assert.IsNonIncreasing(t, []int{2, 1, 1})
// 	assert.IsNonIncreasing(t, []float{2, 1})
// 	assert.IsNonIncreasing(t, []string{"b", "a"})
func (x Assertions) IsNonIncreasing(object interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "is non increasing")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "is non increasing"), NewParameter("object", asShortString(object)).withMode(x.mode))
		}()
		if assert.IsNonIncreasing(t, object, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// IsType asserts that the specified objects are of the same type.
//
// 	assert.IsType(t, &MyStruct{}, &MyStruct{})
func (x Requirements) IsType(expectedType interface{}, object interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "is type")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "is type"), NewParameter("expected type", asShortString(expectedType)).withMode(x.mode), NewParameter("object", asShortString(object)).withMode(x.mode))
		}()
		if assert.IsType(t, expectedType, object, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// IsType asserts that the specified objects are of the same type.
//
// 	assert.IsType(t, &MyStruct{}, &MyStruct{})
func (x Assertions) IsType(expectedType interface{}, object interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "is type")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "is type"), NewParameter("expected type", asShortString(expectedType)).withMode(x.mode), NewParameter("object", asShortString(object)).withMode(x.mode))
		}()
		if assert.IsType(t, expectedType, object, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// JSONEq asserts that two JSON strings are equivalent.
//
// 	assert.JSONEq(t, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
func (x Requirements) JSONEq(expected string, actual string, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "JSON eq")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "JSON eq"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode))
		}()
		if assert.JSONEq(t, expected, actual, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// JSONEq asserts that two JSON strings are equivalent.
//
// 	assert.JSONEq(t, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
func (x Assertions) JSONEq(expected string, actual string, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "JSON eq")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "JSON eq"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode))
		}()
		if assert.JSONEq(t, expected, actual, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// Len asserts that the specified object has specific length.
// Len also fails if the object has a type that len() not accept.
//
// 	assert.Len(t, mySlice, 3)
func (x Requirements) Len(object interface{}, length int, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "len")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "len"), NewParameter("object", asShortString(object)).withMode(x.mode), NewParameter("length", asShortString(length)).withMode(x.mode))
		}()
		if assert.Len(t, object, length, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// Len asserts that the specified object has specific length.
// Len also fails if the object has a type that len() not accept.
//
// 	assert.Len(t, mySlice, 3)
func (x Assertions) Len(object interface{}, length int, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "len")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "len"), NewParameter("object", asShortString(object)).withMode(x.mode), NewParameter("length", asShortString(length)).withMode(x.mode))
		}()
		if assert.Len(t, object, length, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// Less asserts that the first element is less than the second
//
// 	assert.Less(t, 1, 2)
// 	assert.Less(t, float64(1), float64(2))
// 	assert.Less(t, "a", "b")
func (x Requirements) Less(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "less")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "less"), NewParameter("e 1", asShortString(e1)).withMode(x.mode), NewParameter("e 2", asShortString(e2)).withMode(x.mode))
		}()
		if assert.Less(t, e1, e2, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// Less asserts that the first element is less than the second
//
// 	assert.Less(t, 1, 2)
// 	assert.Less(t, float64(1), float64(2))
// 	assert.Less(t, "a", "b")
func (x Assertions) Less(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "less")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "less"), NewParameter("e 1", asShortString(e1)).withMode(x.mode), NewParameter("e 2", asShortString(e2)).withMode(x.mode))
		}()
		if assert.Less(t, e1, e2, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// LessOrEqual asserts that the first element is less than or equal to the second
//
// 	assert.LessOrEqual(t, 1, 2)
// 	assert.LessOrEqual(t, 2, 2)
// 	assert.LessOrEqual(t, "a", "b")
// 	assert.LessOrEqual(t, "b", "b")
func (x Requirements) LessOrEqual(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "less or equal")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "less or equal"), NewParameter("e 1", asShortString(e1)).withMode(x.mode), NewParameter("e 2", asShortString(e2)).withMode(x.mode))
		}()
		if assert.LessOrEqual(t, e1, e2, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// LessOrEqual asserts that the first element is less than or equal to the second
//
// 	assert.LessOrEqual(t, 1, 2)
// 	assert.LessOrEqual(t, 2, 2)
// 	assert.LessOrEqual(t, "a", "b")
// 	assert.LessOrEqual(t, "b", "b")
func (x Assertions) LessOrEqual(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "less or equal")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "less or equal"), NewParameter("e 1", asShortString(e1)).withMode(x.mode), NewParameter("e 2", asShortString(e2)).withMode(x.mode))
		}()
		if assert.LessOrEqual(t, e1, e2, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// Negative asserts that the specified element is negative
//
// 	assert.Negative(t, -1)
// 	assert.Negative(t, -1.23)
func (x Requirements) Negative(e interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "negative")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "negative"), NewParameter("e", asShortString(e)).withMode(x.mode))
		}()
		if assert.Negative(t, e, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// Negative asserts that the specified element is negative
//
// 	assert.Negative(t, -1)
// 	assert.Negative(t, -1.23)
func (x Assertions) Negative(e interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "negative")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "negative"), NewParameter("e", asShortString(e)).withMode(x.mode))
		}()
		if assert.Negative(t, e, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// Never asserts that the given condition doesn't satisfy in waitFor time,
// periodically checking the target function each tick.
//
// 	assert.Never(t, func() bool { return false; }, time.Second, 10*time.Millisecond)
func (x Requirements) Never(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "never")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "never"), NewParameter("condition", asShortString(condition)).withMode(x.mode), NewParameter("wait for", asShortString(waitFor)).withMode(x.mode), NewParameter("tick", asShortString(tick)).withMode(x.mode))
		}()
		if assert.Never(t, condition, waitFor, tick, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// Never asserts that the given condition doesn't satisfy in waitFor time,
// periodically checking the target function each tick.
//
// 	assert.Never(t, func() bool { return false; }, time.Second, 10*time.Millisecond)
func (x Assertions) Never(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "never")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "never"), NewParameter("condition", asShortString(condition)).withMode(x.mode), NewParameter("wait for", asShortString(waitFor)).withMode(x.mode), NewParameter("tick", asShortString(tick)).withMode(x.mode))
		}()
		if assert.Never(t, condition, waitFor, tick, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// Nil asserts that the specified object is nil.
//
// 	assert.Nil(t, err)
func (x Requirements) Nil(object interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "nil")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "nil"), NewParameter("object", asShortString(object)).withMode(x.mode))
		}()
		if assert.Nil(t, object, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// Nil asserts that the specified object is nil.
//
// 	assert.Nil(t, err)
func (x Assertions) Nil(object interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "nil")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "nil"), NewParameter("object", asShortString(object)).withMode(x.mode))
		}()
		if assert.Nil(t, object, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// NoDirExists checks whether a directory does not exist in the given path.
// It fails if the path points to an existing _directory_ only.
func (x Requirements) NoDirExists(path string, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "no dir exists")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "no dir exists"), NewParameter("path", asShortString(path)).withMode(x.mode))
		}()
		if assert.NoDirExists(t, path, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// NoDirExists checks whether a directory does not exist in the given path.
// It fails if the path points to an existing _directory_ only.
func (x Assertions) NoDirExists(path string, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "no dir exists")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "no dir exists"), NewParameter("path", asShortString(path)).withMode(x.mode))
		}()
		if assert.NoDirExists(t, path, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// NoError asserts that a function returned no error (i.e. `nil`).
//
// 	  actualObj, err := SomeFunction()
// 	  if assert.NoError(t, err) {
// 		   assert.Equal(t, expectedObj, actualObj)
// 	  }
func (x Requirements) NoError(err error, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "no error")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "no error"), NewParameter("err", asShortString(err)).withMode(x.mode))
		}()
		if assert.NoError(t, err, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// NoError asserts that a function returned no error (i.e. `nil`).
//
// 	  actualObj, err := SomeFunction()
// 	  if assert.NoError(t, err) {
// 		   assert.Equal(t, expectedObj, actualObj)
// 	  }
func (x Assertions) NoError(err error, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "no error")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "no error"), NewParameter("err", asShortString(err)).withMode(x.mode))
		}()
		if assert.NoError(t, err, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// NoFileExists checks whether a file does not exist in a given path. It fails
// if the path points to an existing _file_ only.
func (x Requirements) NoFileExists(path string, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "no file exists")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "no file exists"), NewParameter("path", asShortString(path)).withMode(x.mode))
		}()
		if assert.NoFileExists(t, path, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// NoFileExists checks whether a file does not exist in a given path. It fails
// if the path points to an existing _file_ only.
func (x Assertions) NoFileExists(path string, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "no file exists")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "no file exists"), NewParameter("path", asShortString(path)).withMode(x.mode))
		}()
		if assert.NoFileExists(t, path, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// NotContains asserts that the specified string, list(array, slice...) or map does NOT contain the
// specified substring or element.
//
// 	assert.NotContains(t, "Hello World", "Earth")
// 	assert.NotContains(t, ["Hello", "World"], "Earth")
// 	assert.NotContains(t, {"Hello": "World"}, "Earth")
func (x Requirements) NotContains(s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "not contains")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "not contains"), NewParameter("s", asShortString(s)).withMode(x.mode), NewParameter("contains", asShortString(contains)).withMode(x.mode))
		}()
		if assert.NotContains(t, s, contains, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// NotContains asserts that the specified string, list(array, slice...) or map does NOT contain the
// specified substring or element.
//
// 	assert.NotContains(t, "Hello World", "Earth")
// 	assert.NotContains(t, ["Hello", "World"], "Earth")
// 	assert.NotContains(t, {"Hello": "World"}, "Earth")
func (x Assertions) NotContains(s interface{}, contains interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "not contains")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "not contains"), NewParameter("s", asShortString(s)).withMode(x.mode), NewParameter("contains", asShortString(contains)).withMode(x.mode))
		}()
		if assert.NotContains(t, s, contains, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// NotElementsMatch asserts that the specified listA(array, slice...) is NOT equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should not match.
// This is an inverse of ElementsMatch.
//
// assert.NotElementsMatch(t, [1, 1, 2, 3], [1, 1, 2, 3]) -> false
//
// assert.NotElementsMatch(t, [1, 1, 2, 3], [1, 2, 3]) -> true
//
// assert.NotElementsMatch(t, [1, 2, 3], [1, 2, 4]) -> true
func (x Requirements) NotElementsMatch(listA interface{}, listB interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "not elements match")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "not elements match"), NewParameter("list A", asShortString(listA)).withMode(x.mode), NewParameter("list B", asShortString(listB)).withMode(x.mode))
		}()
		if assert.NotElementsMatch(t, listA, listB, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// NotElementsMatch asserts that the specified listA(array, slice...) is NOT equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should not match.
// This is an inverse of ElementsMatch.
//
// assert.NotElementsMatch(t, [1, 1, 2, 3], [1, 1, 2, 3]) -> false
//
// assert.NotElementsMatch(t, [1, 1, 2, 3], [1, 2, 3]) -> true
//
// assert.NotElementsMatch(t, [1, 2, 3], [1, 2, 4]) -> true
func (x Assertions) NotElementsMatch(listA interface{}, listB interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "not elements match")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "not elements match"), NewParameter("list A", asShortString(listA)).withMode(x.mode), NewParameter("list B", asShortString(listB)).withMode(x.mode))
		}()
		if assert.NotElementsMatch(t, listA, listB, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// NotEmpty asserts that the specified object is NOT [Empty].
//
// 	if assert.NotEmpty(t, obj) {
// 	  assert.Equal(t, "two", obj[1])
// 	}
func (x Requirements) NotEmpty(object interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "not empty")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "not empty"), NewParameter("object", asShortString(object)).withMode(x.mode))
		}()
		if assert.NotEmpty(t, object, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// NotEmpty asserts that the specified object is NOT [Empty].
//
// 	if assert.NotEmpty(t, obj) {
// 	  assert.Equal(t, "two", obj[1])
// 	}
func (x Assertions) NotEmpty(object interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "not empty")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "not empty"), NewParameter("object", asShortString(object)).withMode(x.mode))
		}()
		if assert.NotEmpty(t, object, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// NotEqual asserts that the specified values are NOT equal.
//
// 	assert.NotEqual(t, obj1, obj2)
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses).
func (x Requirements) NotEqual(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "not equal")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "not equal"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode))
		}()
		if assert.NotEqual(t, expected, actual, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// NotEqual asserts that the specified values are NOT equal.
//
// 	assert.NotEqual(t, obj1, obj2)
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses).
func (x Assertions) NotEqual(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "not equal")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "not equal"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode))
		}()
		if assert.NotEqual(t, expected, actual, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// NotEqualValues asserts that two objects are not equal even when converted to the same type
//
// 	assert.NotEqualValues(t, obj1, obj2)
func (x Requirements) NotEqualValues(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "not equal values")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "not equal values"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode))
		}()
		if assert.NotEqualValues(t, expected, actual, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// NotEqualValues asserts that two objects are not equal even when converted to the same type
//
// 	assert.NotEqualValues(t, obj1, obj2)
func (x Assertions) NotEqualValues(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "not equal values")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "not equal values"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode))
		}()
		if assert.NotEqualValues(t, expected, actual, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// NotErrorAs asserts that none of the errors in err's chain matches target,
// but if so, sets target to that error value.
func (x Requirements) NotErrorAs(err error, target interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "not error as")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "not error as"), NewParameter("err", asShortString(err)).withMode(x.mode), NewParameter("target", asShortString(target)).withMode(x.mode))
		}()
		if assert.NotErrorAs(t, err, target, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// NotErrorAs asserts that none of the errors in err's chain matches target,
// but if so, sets target to that error value.
func (x Assertions) NotErrorAs(err error, target interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "not error as")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "not error as"), NewParameter("err", asShortString(err)).withMode(x.mode), NewParameter("target", asShortString(target)).withMode(x.mode))
		}()
		if assert.NotErrorAs(t, err, target, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// NotErrorIs asserts that none of the errors in err's chain matches target.
// This is a wrapper for errors.Is.
func (x Requirements) NotErrorIs(err error, target error, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "not error is")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "not error is"), NewParameter("err", asShortString(err)).withMode(x.mode), NewParameter("target", asShortString(target)).withMode(x.mode))
		}()
		if assert.NotErrorIs(t, err, target, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// NotErrorIs asserts that none of the errors in err's chain matches target.
// This is a wrapper for errors.Is.
func (x Assertions) NotErrorIs(err error, target error, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "not error is")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "not error is"), NewParameter("err", asShortString(err)).withMode(x.mode), NewParameter("target", asShortString(target)).withMode(x.mode))
		}()
		if assert.NotErrorIs(t, err, target, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// NotImplements asserts that an object does not implement the specified interface.
//
// 	assert.NotImplements(t, (*MyInterface)(nil), new(MyObject))
func (x Requirements) NotImplements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "not implements")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "not implements"), NewParameter("interface object", asShortString(interfaceObject)).withMode(x.mode), NewParameter("object", asShortString(object)).withMode(x.mode))
		}()
		if assert.NotImplements(t, interfaceObject, object, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// NotImplements asserts that an object does not implement the specified interface.
//
// 	assert.NotImplements(t, (*MyInterface)(nil), new(MyObject))
func (x Assertions) NotImplements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "not implements")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "not implements"), NewParameter("interface object", asShortString(interfaceObject)).withMode(x.mode), NewParameter("object", asShortString(object)).withMode(x.mode))
		}()
		if assert.NotImplements(t, interfaceObject, object, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// NotNil asserts that the specified object is not nil.
//
// 	assert.NotNil(t, err)
func (x Requirements) NotNil(object interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "not nil")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "not nil"), NewParameter("object", asShortString(object)).withMode(x.mode))
		}()
		if assert.NotNil(t, object, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// NotNil asserts that the specified object is not nil.
//
// 	assert.NotNil(t, err)
func (x Assertions) NotNil(object interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "not nil")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "not nil"), NewParameter("object", asShortString(object)).withMode(x.mode))
		}()
		if assert.NotNil(t, object, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// NotPanics asserts that the code inside the specified PanicTestFunc does NOT panic.
//
// 	assert.NotPanics(t, func(){ RemainCalm() })
func (x Requirements) NotPanics(f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "not panics")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "not panics"), NewParameter("f", asShortString(f)).withMode(x.mode))
		}()
		if assert.NotPanics(t, f, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// NotPanics asserts that the code inside the specified PanicTestFunc does NOT panic.
//
// 	assert.NotPanics(t, func(){ RemainCalm() })
func (x Assertions) NotPanics(f assert.PanicTestFunc, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "not panics")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "not panics"), NewParameter("f", asShortString(f)).withMode(x.mode))
		}()
		if assert.NotPanics(t, f, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// NotRegexp asserts that a specified regexp does not match a string.
//
// 	assert.NotRegexp(t, regexp.MustCompile("starts"), "it's starting")
// 	assert.NotRegexp(t, "^start", "it's not starting")
func (x Requirements) NotRegexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "not regexp")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "not regexp"), NewParameter("rx", asShortString(rx)).withMode(x.mode), NewParameter("str", asShortString(str)).withMode(x.mode))
		}()
		if assert.NotRegexp(t, rx, str, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// NotRegexp asserts that a specified regexp does not match a string.
//
// 	assert.NotRegexp(t, regexp.MustCompile("starts"), "it's starting")
// 	assert.NotRegexp(t, "^start", "it's not starting")
func (x Assertions) NotRegexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "not regexp")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "not regexp"), NewParameter("rx", asShortString(rx)).withMode(x.mode), NewParameter("str", asShortString(str)).withMode(x.mode))
		}()
		if assert.NotRegexp(t, rx, str, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// NotSame asserts that two pointers do not reference the same object.
//
// 	assert.NotSame(t, ptr1, ptr2)
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
func (x Requirements) NotSame(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "not same")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "not same"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode))
		}()
		if assert.NotSame(t, expected, actual, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// NotSame asserts that two pointers do not reference the same object.
//
// 	assert.NotSame(t, ptr1, ptr2)
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
func (x Assertions) NotSame(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "not same")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "not same"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode))
		}()
		if assert.NotSame(t, expected, actual, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// NotSubset asserts that the list (array, slice, or map) does NOT contain all
// elements given in the subset (array, slice, or map).
// Map elements are key-value pairs unless compared with an array or slice where
// only the map key is evaluated.
//
// 	assert.NotSubset(t, [1, 3, 4], [1, 2])
// 	assert.NotSubset(t, {"x": 1, "y": 2}, {"z": 3})
// 	assert.NotSubset(t, [1, 3, 4], {1: "one", 2: "two"})
// 	assert.NotSubset(t, {"x": 1, "y": 2}, ["z"])
func (x Requirements) NotSubset(list interface{}, subset interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "not subset")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "not subset"), NewParameter("list", asShortString(list)).withMode(x.mode), NewParameter("subset", asShortString(subset)).withMode(x.mode))
		}()
		if assert.NotSubset(t, list, subset, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// NotSubset asserts that the list (array, slice, or map) does NOT contain all
// elements given in the subset (array, slice, or map).
// Map elements are key-value pairs unless compared with an array or slice where
// only the map key is evaluated.
//
// 	assert.NotSubset(t, [1, 3, 4], [1, 2])
// 	assert.NotSubset(t, {"x": 1, "y": 2}, {"z": 3})
// 	assert.NotSubset(t, [1, 3, 4], {1: "one", 2: "two"})
// 	assert.NotSubset(t, {"x": 1, "y": 2}, ["z"])
func (x Assertions) NotSubset(list interface{}, subset interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "not subset")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "not subset"), NewParameter("list", asShortString(list)).withMode(x.mode), NewParameter("subset", asShortString(subset)).withMode(x.mode))
		}()
		if assert.NotSubset(t, list, subset, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// NotZero asserts that i is not the zero value for its type.
func (x Requirements) NotZero(i interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "not zero")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "not zero"), NewParameter("i", asShortString(i)).withMode(x.mode))
		}()
		if assert.NotZero(t, i, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// NotZero asserts that i is not the zero value for its type.
func (x Assertions) NotZero(i interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "not zero")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "not zero"), NewParameter("i", asShortString(i)).withMode(x.mode))
		}()
		if assert.NotZero(t, i, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// Panics asserts that the code inside the specified PanicTestFunc panics.
//
// 	assert.Panics(t, func(){ GoCrazy() })
func (x Requirements) Panics(f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "panics")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "panics"), NewParameter("f", asShortString(f)).withMode(x.mode))
		}()
		if assert.Panics(t, f, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// Panics asserts that the code inside the specified PanicTestFunc panics.
//
// 	assert.Panics(t, func(){ GoCrazy() })
func (x Assertions) Panics(f assert.PanicTestFunc, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "panics")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "panics"), NewParameter("f", asShortString(f)).withMode(x.mode))
		}()
		if assert.Panics(t, f, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// PanicsWithError asserts that the code inside the specified PanicTestFunc
// panics, and that the recovered panic value is an error that satisfies the
// EqualError comparison.
//
// 	assert.PanicsWithError(t, "crazy error", func(){ GoCrazy() })
func (x Requirements) PanicsWithError(errString string, f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "panics with error")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "panics with error"), NewParameter("err string", asShortString(errString)).withMode(x.mode), NewParameter("f", asShortString(f)).withMode(x.mode))
		}()
		if assert.PanicsWithError(t, errString, f, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// PanicsWithError asserts that the code inside the specified PanicTestFunc
// panics, and that the recovered panic value is an error that satisfies the
// EqualError comparison.
//
// 	assert.PanicsWithError(t, "crazy error", func(){ GoCrazy() })
func (x Assertions) PanicsWithError(errString string, f assert.PanicTestFunc, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "panics with error")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "panics with error"), NewParameter("err string", asShortString(errString)).withMode(x.mode), NewParameter("f", asShortString(f)).withMode(x.mode))
		}()
		if assert.PanicsWithError(t, errString, f, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// PanicsWithValue asserts that the code inside the specified PanicTestFunc panics, and that
// the recovered panic value equals the expected panic value.
//
// 	assert.PanicsWithValue(t, "crazy error", func(){ GoCrazy() })
func (x Requirements) PanicsWithValue(expected interface{}, f assert.PanicTestFunc, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "panics with value")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "panics with value"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("f", asShortString(f)).withMode(x.mode))
		}()
		if assert.PanicsWithValue(t, expected, f, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// PanicsWithValue asserts that the code inside the specified PanicTestFunc panics, and that
// the recovered panic value equals the expected panic value.
//
// 	assert.PanicsWithValue(t, "crazy error", func(){ GoCrazy() })
func (x Assertions) PanicsWithValue(expected interface{}, f assert.PanicTestFunc, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "panics with value")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "panics with value"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("f", asShortString(f)).withMode(x.mode))
		}()
		if assert.PanicsWithValue(t, expected, f, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// Positive asserts that the specified element is positive
//
// 	assert.Positive(t, 1)
// 	assert.Positive(t, 1.23)
func (x Requirements) Positive(e interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "positive")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "positive"), NewParameter("e", asShortString(e)).withMode(x.mode))
		}()
		if assert.Positive(t, e, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// Positive asserts that the specified element is positive
//
// 	assert.Positive(t, 1)
// 	assert.Positive(t, 1.23)
func (x Assertions) Positive(e interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "positive")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "positive"), NewParameter("e", asShortString(e)).withMode(x.mode))
		}()
		if assert.Positive(t, e, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// Regexp asserts that a specified regexp matches a string.
//
// 	assert.Regexp(t, regexp.MustCompile("start"), "it's starting")
// 	assert.Regexp(t, "start...$", "it's not starting")
func (x Requirements) Regexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "regexp")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "regexp"), NewParameter("rx", asShortString(rx)).withMode(x.mode), NewParameter("str", asShortString(str)).withMode(x.mode))
		}()
		if assert.Regexp(t, rx, str, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// Regexp asserts that a specified regexp matches a string.
//
// 	assert.Regexp(t, regexp.MustCompile("start"), "it's starting")
// 	assert.Regexp(t, "start...$", "it's not starting")
func (x Assertions) Regexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "regexp")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "regexp"), NewParameter("rx", asShortString(rx)).withMode(x.mode), NewParameter("str", asShortString(str)).withMode(x.mode))
		}()
		if assert.Regexp(t, rx, str, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// Same asserts that two pointers reference the same object.
//
// 	assert.Same(t, ptr1, ptr2)
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
func (x Requirements) Same(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "same")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "same"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode))
		}()
		if assert.Same(t, expected, actual, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// Same asserts that two pointers reference the same object.
//
// 	assert.Same(t, ptr1, ptr2)
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
func (x Assertions) Same(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "same")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "same"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode))
		}()
		if assert.Same(t, expected, actual, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// Subset asserts that the list (array, slice, or map) contains all elements
// given in the subset (array, slice, or map).
// Map elements are key-value pairs unless compared with an array or slice where
// only the map key is evaluated.
//
// 	assert.Subset(t, [1, 2, 3], [1, 2])
// 	assert.Subset(t, {"x": 1, "y": 2}, {"x": 1})
// 	assert.Subset(t, [1, 2, 3], {1: "one", 2: "two"})
// 	assert.Subset(t, {"x": 1, "y": 2}, ["x"])
func (x Requirements) Subset(list interface{}, subset interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "subset")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "subset"), NewParameter("list", asShortString(list)).withMode(x.mode), NewParameter("subset", asShortString(subset)).withMode(x.mode))
		}()
		if assert.Subset(t, list, subset, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// Subset asserts that the list (array, slice, or map) contains all elements
// given in the subset (array, slice, or map).
// Map elements are key-value pairs unless compared with an array or slice where
// only the map key is evaluated.
//
// 	assert.Subset(t, [1, 2, 3], [1, 2])
// 	assert.Subset(t, {"x": 1, "y": 2}, {"x": 1})
// 	assert.Subset(t, [1, 2, 3], {1: "one", 2: "two"})
// 	assert.Subset(t, {"x": 1, "y": 2}, ["x"])
func (x Assertions) Subset(list interface{}, subset interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "subset")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "subset"), NewParameter("list", asShortString(list)).withMode(x.mode), NewParameter("subset", asShortString(subset)).withMode(x.mode))
		}()
		if assert.Subset(t, list, subset, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// True asserts that the specified value is true.
//
// 	assert.True(t, myBool)
func (x Requirements) True(value bool, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "true")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "true"), NewParameter("value", asShortString(value)).withMode(x.mode))
		}()
		if assert.True(t, value, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// True asserts that the specified value is true.
//
// 	assert.True(t, myBool)
func (x Assertions) True(value bool, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "true")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "true"), NewParameter("value", asShortString(value)).withMode(x.mode))
		}()
		if assert.True(t, value, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// WithinDuration asserts that the two times are within duration delta of each other.
//
// 	assert.WithinDuration(t, time.Now(), time.Now(), 10*time.Second)
func (x Requirements) WithinDuration(expected time.Time, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "within duration")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "within duration"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode), NewParameter("delta", asShortString(delta)).withMode(x.mode))
		}()
		if assert.WithinDuration(t, expected, actual, delta, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// WithinDuration asserts that the two times are within duration delta of each other.
//
// 	assert.WithinDuration(t, time.Now(), time.Now(), 10*time.Second)
func (x Assertions) WithinDuration(expected time.Time, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "within duration")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "within duration"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode), NewParameter("delta", asShortString(delta)).withMode(x.mode))
		}()
		if assert.WithinDuration(t, expected, actual, delta, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// WithinRange asserts that a time is within a time range (inclusive).
//
// 	assert.WithinRange(t, time.Now(), time.Now().Add(-time.Second), time.Now().Add(time.Second))
func (x Requirements) WithinRange(actual time.Time, start time.Time, end time.Time, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "within range")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "within range"), NewParameter("actual", asShortString(actual)).withMode(x.mode), NewParameter("start", asShortString(start)).withMode(x.mode), NewParameter("end", asShortString(end)).withMode(x.mode))
		}()
		if assert.WithinRange(t, actual, start, end, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// WithinRange asserts that a time is within a time range (inclusive).
//
// 	assert.WithinRange(t, time.Now(), time.Now().Add(-time.Second), time.Now().Add(time.Second))
func (x Assertions) WithinRange(actual time.Time, start time.Time, end time.Time, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "within range")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "within range"), NewParameter("actual", asShortString(actual)).withMode(x.mode), NewParameter("start", asShortString(start)).withMode(x.mode), NewParameter("end", asShortString(end)).withMode(x.mode))
		}()
		if assert.WithinRange(t, actual, start, end, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// YAMLEq asserts that two YAML strings are equivalent.
func (x Requirements) YAMLEq(expected string, actual string, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "YAML eq")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "YAML eq"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode))
		}()
		if assert.YAMLEq(t, expected, actual, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// YAMLEq asserts that two YAML strings are equivalent.
func (x Assertions) YAMLEq(expected string, actual string, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "YAML eq")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "YAML eq"), NewParameter("expected", asShortString(expected)).withMode(x.mode), NewParameter("actual", asShortString(actual)).withMode(x.mode))
		}()
		if assert.YAMLEq(t, expected, actual, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}

// Zero asserts that i is the zero value for its type.
func (x Requirements) Zero(i interface{}, msgAndArgs ...interface{}) {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "zero")
	Step(x.t, "require: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("require", "zero"), NewParameter("i", asShortString(i)).withMode(x.mode))
		}()
		if assert.Zero(t, i, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
		t.FailNow()
	}, asAssertion())
}

// Zero asserts that i is the zero value for its type.
func (x Assertions) Zero(i interface{}, msgAndArgs ...interface{}) bool {
	x.t.Helper()
	callerTrace := stacktrace.Take(1)
	_, callerFile, callerLine, _ := runtime.Caller(1)
	name := cmp.Or(messageFromMsgAndArgs(msgAndArgs...), "zero")
	return testo.Run(x.t, "assert: "+name, func(t *PluginAllure) {
		t.Helper()
		defer func() {
			t.Parameters(NewParameter("assert", "zero"), NewParameter("i", asShortString(i)).withMode(x.mode))
		}()
		if assert.Zero(t, i, msgAndArgs...) {
			return
		}
		callerLog := fmt.Sprintf("Caller: %s:%d", callerFile, callerLine)
		t.Log(callerLog)
		t.addMessage(callerLog)
		t.addTrace(callerTrace)
	}, asAssertion())
}
