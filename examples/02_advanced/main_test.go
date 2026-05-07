//go:build example

package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/ozontech/testo"
	allure "github.com/ozontech/testo-allure"
	"github.com/ozontech/testo/testoplugin"
)

var categories = []allure.Category{
	{
		Name:            "Example category",
		MatchedStatuses: []allure.Status{allure.StatusBroken, allure.StatusFailed},
	},
	{
		Name:         "Another category",
		MessageRegex: ".*testo.*",
	},
}

func Test(t *testing.T) {
	options := []testoplugin.Option{
		allure.WithCategories(categories...),
		allure.WithDeduplicateAttachments(true),
		allure.WithLinkTransformer(func(link allure.Link) allure.Link {
			u, err := url.Parse(link.URL)
			if err == nil && u.Scheme != "" {
				return link
			}

			switch link.Type {
			case allure.LinkTypeIssue:
				link.URL = fmt.Sprintf("https://example.com/issues/%s", link.URL)

			case allure.LinkTypeTMS:
				link.URL = fmt.Sprintf("https://example.com/tms/%s", link.URL)
			}

			return link
		}),
	}

	testo.RunSuite(t, new(Suite), options...)
}

type T = struct {
	*testo.T
	*allure.PluginAllure
}

type Suite struct {
	testo.Suite[T]

	client *http.Client
}

// BeforeAll hook is executed before running any tests.
func (s *Suite) BeforeAll(t T) {
	allure.Step(t, "Create HTTP client", func(t T) {
		s.client = &http.Client{Timeout: 10 * time.Second}
	})
}

// CasesURL provides URLs.
func (s *Suite) CasesURL() []string {
	return []string{
		"https://example.com",
		"https://example.net",
		"https://example.org",
	}
}

// TestExample is executed for each URL from [Suite.CasesURL] output.
func (s *Suite) TestExample(t T, params struct{ URL string }) {
	t.Parallel()

	t.Titlef("Send request to %s", params.URL)
	t.Tags("http", "html", "example")
	t.Descriptionf(
		"This test sends a simple request to %s and inspects its response.",
		params.URL,
	)
	t.Severity(allure.SeverityCritical)
	t.Links(
		allure.Issue("EXAMPLE-2489"),
		allure.NewLink("https://example.com").Named("example link"),
	)

	var res *http.Response

	allure.Step(t, "send request", func(t T) {
		t.Parameters(
			allure.NewParameter("method", http.MethodGet),
			allure.NewParameter("token", "super secret here").Masked(),
		)

		var err error
		res, err = s.client.Get(params.URL)

		is := t.Require()
		is.NoError(err, "successful round trip")
		is.Equal(http.StatusOK, res.StatusCode, "status is OK")
	})

	allure.Step(t, "inspect response", func(t T) {
		allure.Step(t, "check headers", func(t T) {
			contentType := res.Header.Get("Content-Type")

			t.Assert().Equal("text/html", contentType, "Content-Type is HTML")
		})

		allure.Step(t, "read body", func(t T) {
			body, err := io.ReadAll(res.Body)

			t.Require().NoError(err)

			t.Attach("body", allure.Bytes(body).As(allure.TextHTML))

			t.Assert().Greater(len(body), 0, "not empty")
		})
	})
}
