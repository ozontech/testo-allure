package allure

import (
	"github.com/ozontech/testo"
	"github.com/ozontech/testo/testoplugin"
)

// CommonT is the common interface
// for all T's with Allure plugin.
type CommonT interface {
	testo.CommonT

	Interface
}

// Step wraps [testo.Run] with some differences:
//   - if the step fails (before call to t.Parallel) with fatal error, outer test execution will stop.
//   - context returned by the t.Context will cancel when parent context does.
//
// NOTE: [testo.Run] is also treated as Allure step in the report.
//
// WARN: Running this function during t.Cleanup panics.
//
// See also [Setup] and [TearDown].
func Step[T CommonT](
	t T,
	name string,
	f func(t T),
	options ...testoplugin.Option,
) {
	t.Helper()

	options = append(options, asStep())

	if !testo.Run(t, name, f, options...) {
		// propagate fatal error
		if testo.Reflect(t).HasFatalSubtest {
			u := t.unwrapAllure()

			// to skip trace capturing
			u.inAssertion.Store(true)
			defer u.inAssertion.Store(false)

			t.FailNow()
		}
	}
}

// Setup runs a [Step] marked as Setup in Allure report.
//
// See also [TearDown].
func Setup[T CommonT](
	t T,
	name string,
	f func(t T),
	options ...testoplugin.Option,
) {
	t.Helper()

	options = append(options, asSetup())

	Step(t, name, f, options...)
}

// TearDown runs a [Step] marked as TearDown in Allure report.
//
// See also [Setup].
func TearDown[T CommonT](
	t T,
	name string,
	f func(t T),
	options ...testoplugin.Option,
) {
	t.Helper()

	options = append(options, asTearDown())

	Step(t, name, f, options...)
}
