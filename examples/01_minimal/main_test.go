//go:build example

package main

import (
	"testing"

	"github.com/ozontech/testo"
	allure "github.com/ozontech/testo-allure"
)

func Test(t *testing.T) {
	testo.RunSuite(t, new(Suite))
}

type T struct {
	*testo.T
	*allure.PluginAllure
}

type Suite struct{ testo.Suite[T] }

func (Suite) TestFoo(t T) {
	t.Title("Example Test")
	t.Description("My first _Allure_ test with **testo**")
	t.Tags("testo", "example", "allure")

	allure.Step(t, "first step", func(t T) {
		t.Log("it works")
	})
}
