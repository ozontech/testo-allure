package allure

import (
	"flag"

	"github.com/ozontech/testo/testoplugin"
)

var (
	flagDir = flag.String(
		"allure.dir",
		"allure-results",
		"path to the directory where Allure will save the test results. if the directory does not exist, it will be created.",
	)
	flagInvert = flag.Bool(
		"allure.invert",
		false,
		"only run the tests that do not match the conditions specified by the test plan.",
	)
)

type option func(*PluginAllure)

// LinkTransformerFunc takes a raw link provided by the user
// and transforms it before writing the report.
type LinkTransformerFunc func(link Link) Link

// WithLinkTransformer specifies a function for
// transforming links before writing the report.
//
// For example, may be useful to support short
// identifiers of issues and TMS links and use URL templates to generate full URLs.
func WithLinkTransformer(f LinkTransformerFunc) testoplugin.Option {
	return testoplugin.Option{
		Propagate: true,
		Value: option(func(a *PluginAllure) {
			a.linkTransformer = f
		}),
	}
}

// WithCategories adds [custom categories] to the report.
// This option should be passed to the top-level [testo.RunSuite] call.
//
// [custom categories]: https://allurereport.org/docs/categories/#custom-categories
func WithCategories(categories ...Category) testoplugin.Option {
	return testoplugin.Option{
		Value: option(func(a *PluginAllure) {
			a.categories = append(a.categories, categories...)
		}),
	}
}

// WithOutputDir sets output directory for test results.
//
// By default, it is "allure-results" in the current working directory.
func WithOutputDir(dir string) testoplugin.Option {
	return testoplugin.Option{
		Value: option(func(a *PluginAllure) {
			a.outputDir = dir
		}),
		Propagate: true,
	}
}

// WithExcluded will mark this test or step as excluded from the report.
func WithExcluded(excluded bool) testoplugin.Option {
	return testoplugin.Option{
		Value: option(func(a *PluginAllure) {
			a.excluded = excluded
		}),
		Propagate: true,
	}
}

// WithDeduplicateAttachments enables deduplication of attachment files.
//
// If enabled, plugin will keep track of all written attachments, so that each
// attachment is written at most once and multiple tests will reference the same attachment.
func WithDeduplicateAttachments(deduplicate bool) testoplugin.Option {
	return testoplugin.Option{
		Value: option(func(a *PluginAllure) {
			a.deduplicateAttachments = deduplicate
		}),
		Propagate: true,
	}
}

// WithID will set the given id as AllureID for this test.
//
// If Allure TestOps discovers ID in a test result, it ignores all
// the information related to testCaseId and links the test result to a particular test case.
//
// See [Cooking the AllureID] for more information.
//
// [Cooking the AllureID]: https://help.qameta.io/support/solutions/articles/101000480600-cooking-the-allureid
func WithID(id string) testoplugin.Option {
	return testoplugin.Option{
		Value: option(func(a *PluginAllure) {
			a.ID(id)
		}),
	}
}

// WithLabels adds given labels to this test and all of its subtests.
func WithLabels(labels ...Label) testoplugin.Option {
	return testoplugin.Option{
		Value: option(func(a *PluginAllure) {
			a.Labels(labels...)
		}),
		Propagate: true,
	}
}

// WithTags adds given tags to this test and all of its subtests.
func WithTags(tags ...string) testoplugin.Option {
	return testoplugin.Option{
		Value: option(func(a *PluginAllure) {
			a.Tags(tags...)
		}),
		Propagate: true,
	}
}

// WithOwner sets an owner for this test and all of its subtests.
func WithOwner(owner string) testoplugin.Option {
	return testoplugin.Option{
		Value: option(func(a *PluginAllure) {
			a.Owner(owner)
		}),
		Propagate: true,
	}
}

// WithLinks sets passed links for this test and all of its subtests.
func WithLinks(links ...Link) testoplugin.Option {
	return testoplugin.Option{
		Value: option(func(a *PluginAllure) {
			a.Links(links...)
		}),
		Propagate: true,
	}
}

func asStep() testoplugin.Option {
	return testoplugin.Option{
		Value: option(func(a *PluginAllure) {
			a.inStep.Store(true)
		}),
	}
}

func asSetup() testoplugin.Option {
	return testoplugin.Option{
		Value: option(func(a *PluginAllure) {
			a.stage = stageSetup
		}),
	}
}

func asTearDown() testoplugin.Option {
	return testoplugin.Option{
		Value: option(func(a *PluginAllure) {
			a.stage = stageTearDown
		}),
	}
}

func asAssertion() testoplugin.Option {
	return testoplugin.Option{
		Value: option(func(a *PluginAllure) {
			a.inAssertion.Store(true)
		}),
	}
}
