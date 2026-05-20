package allure

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ozontech/testo"
	"github.com/stretchr/testify/require"
)

type T struct {
	*testo.T
	*PluginAllure
}

type Suite struct{ testo.Suite[T] }

func TestAllure(t *testing.T) {
	// TODO(metafates): use t.Chdir() when go is updated
	wd, err := os.Getwd()
	require.NoError(t, err)

	dir := t.TempDir()
	require.NoError(t, os.Chdir(dir))
	t.Cleanup(func() { _ = os.Chdir(wd) })

	testo.RunSuite(t, new(Suite))

	require.DirExists(t, filepath.Join(dir, "allure-results"), "output dir does not exist")

	// TODO(metafates): other assertions
}

func (Suite) BeforeEach(t T) {
	Setup(t, "init", func(t T) {
		testo.Run(t, "nested", func(t T) {})
	})

	Setup(t, "extra init", func(t T) {})
}

func (Suite) AfterEach(t T) {
	Setup(t, "destroy", func(t T) {
		testo.Run(t, "nested", func(t T) {})
	})

	Setup(t, "extra destroy", func(t T) {})
}

func (Suite) TestFoo(t T) {
	t.Flaky()

	testo.Run(t, "subtest", func(t T) {
		t.Known()

		testo.Run(t, "nested", func(t T) {})
	})
}

func (Suite) CasesX() []int {
	return []int{1}
}

func (Suite) TestBar(t T, params struct{ X int }) {
}

func TestBaseName(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name string
		want string
	}{
		{name: "Test/some/long/name", want: "name"},
		{name: "foo", want: "foo"},
		{name: "", want: ""},
		{name: "Duplicate/test#01", want: "test#01"},
		{name: "Some/malformed/", want: ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := testBaseName(tc.name)

			require.Equal(t, tc.want, got)
		})
	}
}

func TestTrimLines(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name string
		text string
		want string
	}{
		{
			name: "empty string",
			text: "",
			want: "",
		},
		{
			name: "single line",
			text: "   foo bar ",
			want: "foo bar",
		},
		{
			name: "multiline",
			text: " foo bar \n\n foo  bar \n",
			want: "foo bar\n\nfoo  bar",
		},
		{
			name: "empty multiline",
			text: "\n",
			want: "",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := trimLines(tc.text)

			require.Equal(t, tc.want, got)
		})
	}
}

func TestTrimTestifyErrorTrace(t *testing.T) {
	for _, tc := range []struct {
		name string
		text string
		want string
	}{
		{
			name: "empty string",
			text: "",
			want: "",
		},
		{
			name: "with one-line error trace",
			text: `
Error Trace:	/Users/user/testo/pkg/plugins/allure/assert.gen.go:426
/Users/user/testo/testo.go:277
Error:      	An error is expected but got nil.
Test:       	Test/Suite/TestExample_case_3212728f34741cb7/send_request/assert:_successful_round_trip
Messages:   	successful round trip
Caller: /Users/user/testo/pkg/plugins/allure/examples/02_advanced/main_test.go:106`,
			want: `
Error:      	An error is expected but got nil.
Test:       	Test/Suite/TestExample_case_3212728f34741cb7/send_request/assert:_successful_round_trip
Messages:   	successful round trip
Caller: /Users/user/testo/pkg/plugins/allure/examples/02_advanced/main_test.go:106`,
		},
		{
			name: "without closing tag",
			text: `
Error Trace:	/Users/user/testo/pkg/plugins/allure/assert.gen.go:426
/Users/user/testo/testo.go:277
Test:       	Test/Suite/TestExample_case_3212728f34741cb7/send_request/assert:_successful_round_trip
Messages:   	successful round trip
Caller: /Users/user/testo/pkg/plugins/allure/examples/02_advanced/main_test.go:106`,
			want: `
Error Trace:	/Users/user/testo/pkg/plugins/allure/assert.gen.go:426
/Users/user/testo/testo.go:277
Test:       	Test/Suite/TestExample_case_3212728f34741cb7/send_request/assert:_successful_round_trip
Messages:   	successful round trip
Caller: /Users/user/testo/pkg/plugins/allure/examples/02_advanced/main_test.go:106`,
		},
		{
			name: "without closing tag",
			text: `
Error Trace:	/Users/user/testo/pkg/plugins/allure/assert.gen.go:426
/Users/user/testo/testo.go:277
Test:       	Test/Suite/TestExample_case_3212728f34741cb7/send_request/assert:_successful_round_trip
Messages:   	successful round trip
Caller: /Users/user/testo/pkg/plugins/allure/examples/02_advanced/main_test.go:106`,
			want: `
Error Trace:	/Users/user/testo/pkg/plugins/allure/assert.gen.go:426
/Users/user/testo/testo.go:277
Test:       	Test/Suite/TestExample_case_3212728f34741cb7/send_request/assert:_successful_round_trip
Messages:   	successful round trip
Caller: /Users/user/testo/pkg/plugins/allure/examples/02_advanced/main_test.go:106`,
		},
		{
			name: "with multiline error trace",
			text: `
Error Trace:	/Users/user/testo/pkg/plugins/allure/assert.gen.go:426
/Users/user/testo/testo.go:276
/Users/user/testo/testo.go:277
Error:      	An error is expected but got nil.
Test:       	Test/Suite/TestExample_case_3212728f34741cb7/send_request/assert:_successful_round_trip
Messages:   	successful round trip
Caller: /Users/user/testo/pkg/plugins/allure/examples/02_advanced/main_test.go:106`,
			want: `
Error:      	An error is expected but got nil.
Test:       	Test/Suite/TestExample_case_3212728f34741cb7/send_request/assert:_successful_round_trip
Messages:   	successful round trip
Caller: /Users/user/testo/pkg/plugins/allure/examples/02_advanced/main_test.go:106`,
		},
		{
			name: "without error trace",
			text: `
Error:      	An error is expected but got nil.
Test:       	Test/Suite/TestExample_case_3212728f34741cb7/send_request/assert:_successful_round_trip
Messages:   	successful round trip
Caller: /Users/user/testo/pkg/plugins/allure/examples/02_advanced/main_test.go:106`,
			want: `
Error:      	An error is expected but got nil.
Test:       	Test/Suite/TestExample_case_3212728f34741cb7/send_request/assert:_successful_round_trip
Messages:   	successful round trip
Caller: /Users/user/testo/pkg/plugins/allure/examples/02_advanced/main_test.go:106`,
		},
		{
			name: "multiple error traces",
			text: `
Error Trace:	/Users/user/testo/pkg/plugins/allure/assert.gen.go:426
/Users/user/testo/testo.go:276
/Users/user/testo/testo.go:277
Error:      	An error is expected but got nil.
Test:       	Test/Suite/TestExample_case_3212728f34741cb7/send_request/assert:_successful_round_trip
Messages:   	successful round trip
Caller: /Users/user/testo/pkg/plugins/allure/examples/02_advanced/main_test.go:106
Error Trace:	/Users/user/testo/pkg/plugins/allure/assert.gen.go:426
/Users/user/testo/testo.go:276
/Users/user/testo/testo.go:277
Error:      	An error is expected but got nil.
Test:       	Test/Suite/TestExample_case_3212728f34741cb7/send_request/assert:_successful_round_trip
Messages:   	successful round trip
Caller: /Users/user/testo/pkg/plugins/allure/examples/02_advanced/main_test.go:106`,
			want: `
Error:      	An error is expected but got nil.
Test:       	Test/Suite/TestExample_case_3212728f34741cb7/send_request/assert:_successful_round_trip
Messages:   	successful round trip
Caller: /Users/user/testo/pkg/plugins/allure/examples/02_advanced/main_test.go:106
Error:      	An error is expected but got nil.
Test:       	Test/Suite/TestExample_case_3212728f34741cb7/send_request/assert:_successful_round_trip
Messages:   	successful round trip
Caller: /Users/user/testo/pkg/plugins/allure/examples/02_advanced/main_test.go:106`,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := trimTestifyErrorTrace(tc.text)

			require.Equal(t, tc.want, got)
		})
	}
}

func Test_fullName(t *testing.T) {
	t.Parallel()

	tests := []struct {
		Name  string
		Input string
		Want  string
	}{
		{
			Name:  "regular name remains unchanged",
			Input: "Test/Suite/TestFoo",
			Want:  "Test/Suite/TestFoo",
		},
		{
			Name:  "name with hash in the middle remains unchanged",
			Input: "TestFoo/#02/TestFoo",
			Want:  "TestFoo/#02/TestFoo",
		},
		{
			Name:  "name with hash in the end is trimmed",
			Input: "Test/Suite/TestFoo#01",
			Want:  "Test/Suite/TestFoo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			got := fullName(tt.Input)

			require.Equal(t, tt.Want, got)
		})
	}
}
