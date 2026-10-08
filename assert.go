package allure

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
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

var testifyKV = regexp.MustCompile(`^\s*(?P<key>.+?):\s(?P<value>.*)`)

func testifySplitKeyValue(s string) (key, value string, found bool) {
	matches := testifyKV.FindStringSubmatch(s)
	if len(matches) == 0 {
		return "", "", false
	}

	groups := testifyKV.SubexpNames()

	res := make(map[string]string, len(groups))

	for i, name := range groups {
		if i == 0 {
			continue
		}

		res[name] = matches[i]
	}

	return res["key"], strings.TrimSpace(res["value"]), true
}

func transformTestifyErrorMsg(s string) string {
	var (
		prevKey    string
		prevValues []string
		keys       []string
	)

	byKey := make(map[string][]string)

	for line := range strings.Lines(s) {
		const lenLimit = 2000

		if len(line) > lenLimit {
			line = line[:lenLimit] + "..."
		}

		key, value, ok := testifySplitKeyValue(line)
		if !ok {
			prevValues = append(prevValues, strings.TrimSpace(line))

			continue
		}

		if prevKey != "" {
			keys = append(keys, prevKey)
			byKey[prevKey] = prevValues
		}

		prevKey = key
		prevValues = []string{value}
	}

	keys = append(keys, prevKey)

	byKey[prevKey] = prevValues

	adjustTestifyLines(byKey)

	lines := make([]string, 1, len(keys))

	for _, k := range keys {
		v, ok := byKey[k]
		if !ok || k == "" {
			continue
		}

		prefix := k + ": "

		if len(v) == 0 {
			lines = append(lines, prefix)
		} else {
			lines = append(lines, prefix+v[0])
			lines = append(lines, v[1:]...)
		}
	}

	return strings.Join(lines, "\n")
}

// adjustTestifyLines accepts a testify assertion log lines
// groupped by their keys. For example:
//
//	Error: foobar
//	Error Trace: ...
//	Messages: one
//	          two
//	          three
func adjustTestifyLines(byKey map[string][]string) {
	const (
		keyMessages   = "Messages"
		keyError      = "Error"
		keyErrorTrace = "Error Trace"
	)

	{
		_, hasMessages := byKey[keyMessages]
		_, hasError := byKey[keyError]

		if hasMessages && hasError {
			byKey[keyError], byKey[keyMessages] = byKey[keyMessages], byKey[keyError]
		}
	}

	delete(byKey, keyErrorTrace)
}

func trimCallerLine(s string) string {
	lines := strings.Split(s, "\n")

	lines = slices.DeleteFunc(lines, func(l string) bool {
		return l == "" || strings.HasPrefix(l, "Caller: ")
	})

	return strings.Join(lines, "\n")
}
