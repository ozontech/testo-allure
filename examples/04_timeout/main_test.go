//go:build example

package main

import (
	"testing"
	"time"

	"github.com/ozontech/testo"
	allure "github.com/ozontech/testo-allure"
)

type T struct {
	*testo.T
	*allure.PluginAllure
}

type Suite struct{ testo.Suite[T] }

func (*Suite) TestFast(t T) {
	t.Parallel()

	allure.Step(t, "just a moment", func(t T) {
		t.Log("Inside step")
	})
}

func (*Suite) TestSlow(t T) {
	t.Parallel()

	allure.Step(t, "wait...", func(t T) {
		time.Sleep(5 * time.Second)
	})
}

func (*Suite) AfterAll(t T) {
	t.Log("inside after all")
}

func Test(t *testing.T) {
	testo.RunSuite(t, new(Suite))
}
