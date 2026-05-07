package allure

import (
	"fmt"
)

//go:generate sh -c "cd _codegen && go run . -pkg allure -path ../assert.gen.go"

// Assertions provides a set of helpers to perform assertions in tests.
//
// Each assertion is included in the Allure report
// as a step with passed parameters.
type Assertions struct {
	t    *PluginAllure
	mode ParameterMode
}

// Requirements implements the same assertions as [Assertions]
// but stops test execution when assertion fails.
type Requirements struct {
	t    *PluginAllure
	mode ParameterMode
}

// Masked returns a new requirements instance
// which will mask its parameters.
func (r Requirements) Masked() Requirements {
	r.mode = ParameterModeMasked

	return r
}

// Masked returns a new assertions instance
// which will mask its parameters.
func (a Assertions) Masked() Assertions {
	a.mode = ParameterModeMasked

	return a
}

// Hidden returns a new assertions instance
// which will hide its parameters.
func (a Assertions) Hidden() Assertions {
	a.mode = ParameterModeHidden

	return a
}

// Hidden returns a new requirements instance
// which will hide its parameters.
func (r Requirements) Hidden() Requirements {
	r.mode = ParameterModeHidden

	return r
}

func messageFromMsgAndArgs(msgAndArgs ...any) string {
	switch len(msgAndArgs) {
	case 0:
		return ""

	case 1:
		msg := msgAndArgs[0]

		if s, ok := msg.(string); ok {
			return s
		}

		return fmt.Sprintf("%+v", msg)

	default:
		format, ok := msgAndArgs[0].(string)
		if !ok {
			panic(fmt.Sprintf("format must be a string, got %T", msgAndArgs[0]))
		}

		return fmt.Sprintf(format, msgAndArgs[1:]...)
	}
}

func asShortString(v any) string {
	s := fmt.Sprintf("%+v", v)

	const limit = 200

	if len(s) > limit {
		s = s[:limit] + "..."
	}

	return s
}
