//go:build example

package main

import (
	"testing"
	"time"

	"github.com/ozontech/testo"
	// using a dot import like that allows us to use
	// allure functions with less verbosity
	. "github.com/ozontech/testo-allure"
)

func Test(t *testing.T) {
	testo.RunSuite(t, new(Suite))
}

type T struct {
	*testo.T
	*PluginAllure
}

type Suite struct{ testo.Suite[T] }

func (Suite) BeforeEach(t T) {
	Setup(t, "fill data", func(t T) {
		Step(t, "open connections", func(t T) { /* ... */ })
		Step(t, "fill db", func(t T) { /* ... */ })
	})
}

func (Suite) AfterEach(t T) {
	TearDown(t, "cleanup data", func(t T) {
		Step(t, "erase db", func(t T) { /* ... */ })
		Step(t, "close connections", func(t T) { /* ... */ })
	})
}

func (Suite) TestFoo(t T) {
	t.Title("Example test")

	t.Parameters(NewParameter("time", time.Now()), NewParameter("n", 5))
	t.Links(Issue("TESTO-1234"), NewLink("https://example.com"))
	t.Severity(SeverityTrivial)

	Step(t, "first step", func(t T) { t.Log("Hi") })
	Step(t, "second step", func(t T) { t.Log("Hi again") })
}
