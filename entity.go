package allure

import (
	"bytes"
	"encoding"
	"fmt"
	"strconv"
	"sync/atomic"

	"github.com/ozontech/testo-allure/internal/syncutil"
)

// UUID is unique identifier.
type UUID = string

type stage int

const (
	stageTest stage = iota
	stageSetup
	stageTearDown
)

var _ encoding.TextMarshaler = (*Severity)(nil)

// Severity is a value indicating how important the test is.
//
// This may give the future reader an idea of how to prioritize
// the investigations of different test failures.
type Severity int

// Possible severity values.
const (
	// SeverityTrivial is the least severe level.
	//
	// Usually related to minor, non-critical issues that have
	// little to no impact on the overall user experience.
	SeverityTrivial Severity = -2

	// SeverityMinor indicates minor issues, such as cosmetic problems
	// or small functional discrepancies.
	SeverityMinor Severity = -1

	// WARN(metafates): ensure that [SeverityNormal] is always of value 0
	// since we want it to be the default value.

	// SeverityNormal is the default severity.
	//
	// Tests at this level indicate regular or expected issues
	// and do not pose a major threat to the application's stability.
	SeverityNormal Severity = 0

	// SeverityCritical is for tests that identify serious problems
	// with the system's functionality, though not necessarily a complete blocker.
	SeverityCritical Severity = 1

	// SeverityBlocker is the most severe level.
	//
	// Tests flagged as "blocker" are critical issues that
	// prevent core functionality and require immediate attention on failure.
	SeverityBlocker Severity = 2
)

func (s Severity) String() string {
	switch s {
	case SeverityTrivial:
		return "trivial"

	case SeverityMinor:
		return "minor"

	case SeverityNormal:
		return "normal"

	case SeverityCritical:
		return "critical"

	case SeverityBlocker:
		return "blocker"

	default:
		return fmt.Sprintf("Severity(%d)", s)
	}
}

// MarshalText implements [encoding.TextMarshaler].
func (s Severity) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

// Category defines tests category.
//
// Allure checks each test against all the categories in the file,
// from top to bottom. The test gets assigned the first matching category.
// When no matches are found, Allure uses one of the default categories
// if the test is unsuccessful or no category otherwise.
type Category struct {
	// Name of the category.
	Name string `json:"name"`

	// MessageRegex is the regular expression
	// that the test result's message should match.
	MessageRegex string `json:"messageRegex,omitempty"`

	// TraceRegex is the regular expression that
	// the test result's trace should match.
	TraceRegex string `json:"traceRegex,omitempty"`

	// MatchedStatuses are the statuses that
	// the test result should be one of.
	MatchedStatuses []Status `json:"matchedStatuses,omitempty"`
}

// Label holds additional information about the test.
type Label struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// NewLabel returns a new [Label].
func NewLabel(name, value string) Label {
	return Label{Name: name, Value: value}
}

// well-known labels.
const (
	labelSeverity  = "severity"
	labelAllureID  = "allure_id"
	labelTag       = "tag"
	labelSuite     = "suite"
	labelHost      = "host"
	labelFramework = "framework"
	labelLanguage  = "language"
	labelOwner     = "owner"
	labelEpic      = "epic"
	labelFeature   = "feature"
	labelStory     = "story"
)

// LinkType is the type of link.
type LinkType string

// Possible link types.
const (
	LinkTypeIssue LinkType = "issue"
	LinkTypeTMS   LinkType = "tms"
)

// Link to webpage that may be useful for a reader
// investigating a test failure.
type Link struct {
	Name string   `json:"name"`
	URL  string   `json:"url"`
	Type LinkType `json:"type"`
}

// NewLink returns a new [Link] with the given url.
func NewLink(url string) Link {
	return Link{URL: url}
}

// NewLinkf is the same as [NewLink] but
// uses [fmt.Sprintf] to format url.
func NewLinkf(format string, args ...any) Link {
	return Link{URL: fmt.Sprintf(format, args...)}
}

// Issue returns a new [Link] of type issue.
func Issue(issue string) Link {
	return NewLink(issue).Issue()
}

// Issuef returns a new [Link] of type issue.
func Issuef(format string, args ...any) Link {
	return NewLinkf(format, args...).Issue()
}

// TMS returns a new [Link] of type tms.
func TMS(tms string) Link {
	return NewLink(tms).TMS()
}

// TMSf returns a new [Link] of type tms.
func TMSf(format string, args ...any) Link {
	return NewLinkf(format, args...).TMS()
}

// Named returns a link with the given name set.
func (l Link) Named(name string) Link {
	l.Name = name

	return l
}

// Issue returns a link with a type of [LinkTypeIssue].
func (l Link) Issue() Link {
	l.Type = LinkTypeIssue

	return l
}

// TMS returns a link with a type of [LinkTypeTMS].
func (l Link) TMS() Link {
	l.Type = LinkTypeTMS

	return l
}

// Parameter to show in the report.
//
// Allure plugin automatically sets parameters
// for parametrized tests.
type Parameter struct {
	// Name of the parameter.
	Name string
	// Value of the parameter.
	Value any
	// Exclude parameter when comparing
	// the current result with the previous one in the history.
	Exclude bool
	// Mode of the parameter.
	Mode ParameterMode
}

// NewParameter constructs a new [Parameter].
//
//	t.Parameters(allure.NewParameter("arg", 42))
func NewParameter(name string, value any) Parameter {
	return Parameter{
		Name:  name,
		Value: value,
	}
}

// Masked returns a new parameter with mode set to masked.
func (p Parameter) Masked() Parameter {
	return p.withMode(ParameterModeMasked)
}

// Hidden returns a new parameter with mode set to hidden.
func (p Parameter) Hidden() Parameter {
	return p.withMode(ParameterModeHidden)
}

// Excluded returns a new parameter marked as excluded.
//
// Excluded parameters are ignored when comparing
// the current result with the previous one in the history.
func (p Parameter) Excluded() Parameter {
	p.Exclude = true

	return p
}

func (p Parameter) withMode(mode ParameterMode) Parameter {
	p.Mode = mode

	return p
}

func (p Parameter) toInternal() parameter {
	return parameter{
		Name:     p.Name,
		Value:    fmt.Sprintf("%+v", p.Value),
		Excluded: p.Exclude,
		Mode:     p.Mode,
	}
}

type parameter struct {
	Name     string        `json:"name"`
	Value    string        `json:"value"`
	Excluded bool          `json:"excluded"`
	Mode     ParameterMode `json:"mode"`
}

var _ encoding.TextMarshaler = (*ParameterMode)(nil)

// ParameterMode controls how the parameter will be shown in the report.
type ParameterMode int

const (
	// ParameterModeDefault - the parameter and its value
	// will be shown in a table along with other parameters.
	ParameterModeDefault ParameterMode = iota

	// ParameterModeMasked - the parameter will be shown
	// in the table, but its value will be hidden.
	ParameterModeMasked

	// ParameterModeHidden - the parameter and its value
	// will not be shown in the test report.
	ParameterModeHidden
)

func (pm ParameterMode) String() string {
	switch pm {
	case ParameterModeDefault:
		return "default"

	case ParameterModeMasked:
		return "masked"

	case ParameterModeHidden:
		return "hidden"

	default:
		return fmt.Sprintf("ParameterMode(%d)", pm)
	}
}

// MarshalText implements [encoding.TextMarshaler].
func (pm ParameterMode) MarshalText() ([]byte, error) {
	return []byte(pm.String()), nil
}

var _ encoding.TextMarshaler = (*Status)(nil)

// Status is the test status.
//
// See [Test statuses] for more information.
//
// [Test statuses]: https://allurereport.org/docs/test-statuses/
type Status int

const (
	// StatusPassed is a passed test (green) is a test that finished successfully.
	// This usually means that the tested scenario works as expected.
	//
	// https://allurereport.org/docs/test-statuses/#passed
	StatusPassed Status = iota + 1

	// StatusFailed is a failed test (red) is a test that encountered
	// an unexpected behavior in the system under test.
	// This means that the test itself seems to be valid (not [StatusBroken]),
	// but its execution ended with a false assertion.
	//
	// Note that a failure of a step does not necessarily cause a failure of the test as a whole.
	//
	// https://allurereport.org/docs/test-statuses/#failed
	StatusFailed

	// StatusSkipped is a skipped test (gray) is a test that
	// was included in the test plan but then not executed.
	//
	// https://allurereport.org/docs/test-statuses/#skipped
	StatusSkipped

	// StatusBroken is a broken test (yellow) is a test that
	// failed because of a test defect. Unlike Failed, this
	// status means that the test was unable to check the product's behavior
	// as it intended to, therefore, the failure may or may not indicate an actual product defect.
	//
	// https://allurereport.org/docs/test-statuses/#broken
	StatusBroken
)

// Unknown returns true if this status is unknown.
//
// A test gets the "Unknown" status (violet) when the Allure plugin
// did not explicitly set any other status for it.
//
// This may be caused by a bug or an incorrect usage of the Allure plugin.
//
// https://allurereport.org/docs/test-statuses/#unknown
func (s Status) Unknown() bool {
	switch s {
	case StatusFailed, StatusBroken, StatusPassed, StatusSkipped:
		return false

	default:
		return true
	}
}

func (s Status) String() string {
	switch s {
	case StatusFailed:
		return "failed"

	case StatusBroken:
		return "broken"

	case StatusPassed:
		return "passed"

	case StatusSkipped:
		return "skipped"

	default:
		return "unknown"
	}
}

// MarshalText implements [encoding.TextMarshaler].
func (s Status) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

// StatusDetails holds additional information for status.
type StatusDetails struct {
	// Known indicates that the test
	// fails because of a known bug.
	Known bool `json:"known"`

	// Muted indicates that the result
	// must not affect the statistics.
	Muted bool `json:"muted"`

	// Flaky indicates that this test or step is known
	// to be unstable and can may not succeed every time.
	Flaky bool `json:"flaky"`

	// Message is the short text message to display in the
	// test details, such as a name of the exception that led to a failure.
	Message string `json:"message"`

	// Trace is the full stack trace to display in the test details.
	Trace string `json:"trace"`
}

type statusDetails struct {
	Known, Muted, Flaky atomic.Bool

	Message, Trace syncutil.MutexGuarded[string]
}

func (sd *statusDetails) toPublic() StatusDetails {
	return StatusDetails{
		Known:   sd.Known.Load(),
		Muted:   sd.Muted.Load(),
		Flaky:   sd.Flaky.Load(),
		Message: sd.Message.Load(),
		Trace:   sd.Trace.Load(),
	}
}

type attachment struct {
	Name   string    `json:"name"`
	Source string    `json:"source"`
	Type   MediaType `json:"type"`
}

type properties struct {
	GoOS       string
	GoArch     string
	GoVersion  string
	GoCompiler string
	NumCPU     int
}

// MarshalProperties marshals this structure into [.properties] format.
//
// [.properties]: https://en.wikipedia.org/wiki/.properties
func (p properties) MarshalProperties() ([]byte, error) {
	var buf bytes.Buffer

	for _, line := range []struct{ Key, Value string }{
		{Key: "go_os", Value: p.GoOS},
		{Key: "go_arch", Value: p.GoArch},
		{Key: "go_version", Value: p.GoVersion},
		{Key: "go_compiler", Value: p.GoCompiler},
		{Key: "num_cpu", Value: strconv.Itoa(p.NumCPU)},
	} {
		_, err := buf.WriteString(line.Key + " = " + line.Value + "\n")
		if err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

// container describes a set of initialization
// and finalization steps related to certain tests.
type container struct {
	UUID UUID `json:"uuid"`

	// Start is the time when the execution of the initialization steps started,
	// in the UNIX milliseconds timestamp format.
	//
	// Normally, it is equal or close to the earliest start time from the befores.
	Start unixMilli `json:"start"`

	// Stop is the time when the execution of the finalization steps finished,
	// in the UNIX milliseconds timestamp format.
	//
	// Normally, it is equal or close to the latest stop time from the afters.
	Stop unixMilli `json:"stop"`

	Children []UUID `json:"children"`

	// Befores are initialization steps that were executed before the tests.
	//
	// Note that the test report does not necessarily display these steps
	// in the same order as defined in the file.
	Befores []step `json:"befores"`

	// Afters are finalization steps that were executed after the tests.
	//
	// Note that the test report does not necessarily display these steps
	// in the same order as defined in the file.
	Afters []step `json:"afters"`
}

type unixMilli int64

type result struct {
	// UUID is the unique identifier of the test.
	UUID UUID `json:"uuid"`

	// HistoryID is an identifier used by Allure Report.
	// Two runs of the same test with the same set of parameters
	// will always have the same history id.
	HistoryID string `json:"historyId"`

	// TestCaseID is an identifier used by Allure TestOps.
	// Two runs of the same test will always have the same test case id.
	TestCaseID string `json:"testCaseId"`

	// FullName is a unique name of the test.
	FullName string `json:"fullName"`

	// Name is the title of the test.
	Name string `json:"name"`

	// Description of the test in Markdown format.
	Description string `json:"description"`

	// Links added to the test.
	Links []Link `json:"links,omitempty"`

	// Labels added to the test.
	Labels []Label `json:"labels,omitempty"`

	// Parameters added to the test.
	Parameters []parameter `json:"parameters,omitempty"`

	// Attachments added to the test.
	Attachments []attachment `json:"attachments,omitempty"`

	// Status of the test.
	Status Status `json:"status"`

	// StatusDetails of the test.
	StatusDetails StatusDetails `json:"statusDetails"`

	// Start is the time when the execution
	// of the test started, in the UNIX milliseconds timestamp format.
	Start unixMilli `json:"start"`

	// Stop is the time when the execution
	// of the test finished, in the UNIX milliseconds timestamp format.
	Stop unixMilli `json:"stop"`

	// Steps of the test.
	Steps []step `json:"steps,omitempty"`
}

type step struct {
	// Name of the step.
	Name string `json:"name"`

	// Status of the step.
	Status Status `json:"status"`

	// StatusDetails of the step.
	StatusDetails StatusDetails `json:"statusDetails"`

	// Start is the time when the execution
	// of the step started, in the UNIX milliseconds timestamp format.
	Start unixMilli `json:"start"`

	// Stop is the time when the execution
	// of the step finished, in the UNIX milliseconds timestamp format.
	Stop unixMilli `json:"stop"`

	// Steps are sub-steps of the step.
	Steps []step `json:"steps,omitempty"`

	// Attachments added to the step.
	Attachments []attachment `json:"attachments,omitempty"`

	// Parameters added to the step.
	Parameters []parameter `json:"parameters,omitempty"`
}

type testPlan struct {
	Version string         `json:"version"`
	Tests   []testPlanTest `json:"tests"`
}

func (tp testPlan) sets() (ids, selectors map[string]bool) {
	ids = make(map[string]bool, len(tp.Tests))
	selectors = make(map[string]bool, len(tp.Tests))

	for _, t := range tp.Tests {
		if t.Selector != "" {
			selectors[t.Selector] = true
		}

		if t.ID != "" {
			ids[t.ID] = true
		}
	}

	return ids, selectors
}

type testPlanTest struct {
	ID       string `json:"id"`
	Selector string `json:"selector"`
}
