// Package allure provides Allure provider as a plugin for testo.
package allure

import (
	"cmp"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"slices"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unicode"

	"github.com/ozontech/testo"
	"github.com/ozontech/testo-allure/internal/allurehash"
	"github.com/ozontech/testo-allure/internal/set"
	"github.com/ozontech/testo-allure/internal/stacktrace"
	"github.com/ozontech/testo-allure/internal/syncutil"
	"github.com/ozontech/testo-allure/internal/uuid"
	"github.com/ozontech/testo/testoplugin"
	"github.com/ozontech/testo/testoreflect"
)

const (
	permFile os.FileMode = 0o600
	permDir  os.FileMode = 0o750
)

// heuristic time windows when we decide to
// write afterXxx hooks just before test timeout is triggered.
const (
	suiteDeadlineWindow = 25 * time.Millisecond
	testDeadlineWindow  = 50 * time.Millisecond
	stepDeadlineWindow  = 75 * time.Millisecond
)

//go:generate go tool ifacemaker -f $GOFILE -o interface.go -s PluginAllure -i Interface -p $GOPACKAGE -e Plugin -y "Interface defines allure plugin interface.\nUseful for writing helpers which require allure methods but can't rely on concrete type." -x -e panicked -e status -e asResult -e parameters -e links -e attachments -e allRawAttachments -e title -e asStep -e timeBoundaries -e steps -e containers -e beforeEach -e afterEach -e hooks -e addMessage -e addTrace -e overrides -e results -e resultsGroupParametrized -e afterAll -e writeResults -e writeContainers -e writeAttachments -e writeAttachment -e writeProperties -e writeCategories -e labels -e attachmentPath -e baseName -e testCaseID -e historyID -e resultsFlattenParametrized -e statusDetails -e suiteName -e plugin -e beforeAll -e cleanup -e writeReport -e plan -e applyOptions -e fullName -e createOutputDir -e asContainer -e beforeEachSub -e afterEachSub -e propagatedStatusDetails -e hookDescendants -e descendants -e testChildren -e hasTestNeighbors -e subtest -e attach -e parentSuiteName -e realStatus

var _ Interface = (*PluginAllure)(nil)

type timeBoundary struct {
	Start, Stop time.Time
}

func (tb timeBoundary) Duration() time.Duration {
	return tb.Stop.Sub(tb.Start)
}

var _ testoplugin.Plugin = (*PluginAllure)(nil)

// PluginAllure defines allure plugin.
type PluginAllure struct {
	*testo.T

	parent *PluginAllure

	uuid UUID

	timeTest      syncutil.MutexGuarded[timeBoundary]
	timeBeforeAll syncutil.MutexGuarded[timeBoundary]
	timeAfterAll  syncutil.MutexGuarded[timeBoundary]

	// used for BeforeAll hooks to set stop once when first test is run
	setBeforeAllStopOnce sync.Once

	// This is tricky.
	//
	// When test calls for t.Parallel() it is paused
	// and resumed when parallel tests are run.
	//
	// However, since we register start time before t.Parallel() call,
	// test duration becomes incorrect because it includes this pause which we don't want.
	//
	// Therefore, we remember how long the test took before t.Parallel() was called
	// so that we can adjust start and stop time with this to make overall test duration correct.
	//
	// See [PluginAllure.timeBoundaries].
	beforeParallel time.Duration

	rawLabels        syncutil.MutexGuarded[set.Set[Label]]
	rawParameters    syncutil.MutexGuarded[[]Parameter]
	rawLinks         syncutil.MutexGuarded[set.Set[Link]]
	description      syncutil.AtomicValue[string]
	rawStatusDetails statusDetails
	categories       []Category
	attachments      syncutil.MutexGuarded[[]attachment]

	children syncutil.MutexGuarded[[]*PluginAllure]

	excluded               bool
	inverted               bool
	outputDir              string
	stage                  stage
	inAssertion            atomic.Bool
	inStep                 atomic.Bool
	deduplicateAttachments bool
	groupHooks             bool
	handleTimeouts         bool

	owner           syncutil.AtomicValue[string]
	epic            syncutil.AtomicValue[string]
	feature         syncutil.AtomicValue[string]
	story           syncutil.AtomicValue[string]
	severity        syncutil.AtomicInt[Severity]
	titleOverwrite  syncutil.AtomicValue[string]
	allureID        syncutil.AtomicValue[string]
	statusOverwrite syncutil.AtomicInt[Status]

	linkTransformer LinkTransformerFunc

	queuedSetups    syncutil.MutexGuarded[[]*PluginAllure]
	queuedTearDowns syncutil.MutexGuarded[[]*PluginAllure]

	maxAttachmentSize int64

	testsStarted atomic.Bool
	testTimedOut atomic.Bool
	stepTimedOut atomic.Bool

	timedOut atomic.Bool

	running atomic.Bool
}

// Plugin implements [testoplugin.Plugin].
func (a *PluginAllure) Plugin(
	parent testoplugin.Plugin,
	options ...testoplugin.Option,
) testoplugin.Spec {
	return a.plugin(parent.(*PluginAllure), options...)
}

// Require returns a new [Requirements] instance.
func (a *PluginAllure) Require() Requirements {
	return Requirements{t: a}
}

// Assert returns a new [Assertions] instance.
func (a *PluginAllure) Assert() Assertions {
	return Assertions{t: a}
}

// Title sets a human-readable title of the test.
//
// If not provided, function or subtest name is used instead.
//
//	t.Title("My Test")
func (a *PluginAllure) Title(title string) {
	a.Helper()

	a.titleOverwrite.Store(title)

	a.Log(title)
}

// Titlef is the same as [PluginAllure.Title] but
// uses [fmt.Sprintf] to set a title.
//
//	t.Titlef("Request to %s", url)
func (a *PluginAllure) Titlef(format string, args ...any) {
	a.Helper()

	a.Title(fmt.Sprintf(format, args...))
}

// Description sets an arbitrary text describing the test in
// more details than the title could fit.
//
// The description will be treated as a Markdown text,
// so you can you apply some basic formatting in it.
// HTML tags are not allowed in such a text and will
// be removed when building the report.
//
//	t.Description("Test description with **markdown** _support_!")
func (a *PluginAllure) Description(desc string) {
	a.description.Store(trimLines(desc))
}

// Descriptionf is the same as [PluginAllure.Description] but
// uses [fmt.Sprintf] to set a description.
func (a *PluginAllure) Descriptionf(format string, args ...any) {
	a.Description(fmt.Sprintf(format, args...))
}

// Links adds a list of links to webpages that may be useful for a reader investigating a test failure.
// You can provide as many links as needed.
//
// There are three types of links:
//   - a standard web link, e.g., a link to the description of the feature being tested;
//   - a link to an issue in the product's issue tracker;
//   - a link to the test description in a test management system (TMS).
func (a *PluginAllure) Links(links ...Link) {
	if len(links) == 0 {
		return
	}

	a.rawLinks.Modify(func(value *set.Set[Link]) {
		value.Add(links...)
	})
}

// Labels adds given labels to the test result.
//
// A test result can have multiple labels with the same name.
// For example, this is often the case when a test result has multiple tags.
//
// Consider using helper methods such as [PluginAllure.Tags] or [PluginAllure.Severity]
// instead of using labels directly.
//
//	t.Labels(allure.NewLabel("name", "value"), allure.NewLabel("otherLabel", "42"))
func (a *PluginAllure) Labels(labels ...Label) {
	a.Helper()

	if len(labels) == 0 {
		return
	}

	a.rawLabels.Modify(func(value *set.Set[Label]) {
		value.Add(labels...)
	})
}

// Tags adds short terms the test is related to.
// Usually it's a good idea to list relevant
// features that are being tested.
//
// Tags can then be used for [filtering].
//
//	t.Tags("heavy", "other tag")
//
// [filtering]: https://allurereport.org/docs/sorting-and-filtering/#filter-tests-by-tags
func (a *PluginAllure) Tags(tags ...string) {
	a.Helper()

	if len(tags) == 0 {
		return
	}

	labels := make([]Label, 0, len(tags))

	for _, tag := range tags {
		labels = append(labels, Label{Name: labelTag, Value: tag})
	}

	a.Labels(labels...)
}

// ID specifies unique identifier of this test in Allure TestOps' database.
//
// If Allure TestOps discovers ID in a test result, it ignores all
// the information related to testCaseId and links the test result to a particular test case.
//
// See [Cooking the AllureID] for more information.
//
// [Cooking the AllureID]: https://help.qameta.io/support/solutions/articles/101000480600-cooking-the-allureid
func (a *PluginAllure) ID(id string) {
	a.allureID.Store(id)
}

// Status overwrites the status with which the test or step finished.
//
// Unless this method was called with known status,
// allure plugin will automatically derive a status with the following rules:
//
//   - [StatusBroken] if the test panicked.
//   - [StatusFailed] if the test has failed.
//   - [StatusSkipped] if the test was skipped.
//   - [StatusPassed] otherwise.
func (a *PluginAllure) Status(status Status) {
	a.statusOverwrite.Store(status)
}

// Parameters adds parameters to show for this report in the result.
//
// Allure plugin automatically sets parameters for parametrized tests.
// This behavior can be configured with [WithAutoParameters] option.
//
//	t.Parameters(allure.NewParameter("login", "john doe"), allure.NewParameter("age", 42))
func (a *PluginAllure) Parameters(parameters ...Parameter) {
	if len(parameters) == 0 {
		return
	}

	a.rawParameters.Modify(func(value *[]Parameter) {
		*value = append(*value, parameters...)
	})
}

// Owner sets the team member who is responsible for the test's stability.
// For example, this can be the test's author, the
// leading developer of the feature being tested, etc.
//
//	t.Owner("John Doe")
func (a *PluginAllure) Owner(owner string) {
	a.owner.Store(owner)
}

// Severity sets a value indicating how important the test is.
// This may give the future reader an idea of how
// to prioritize the investigations of different test failures.
//
//	t.Severity(allure.SeverityCritical)
func (a *PluginAllure) Severity(severity Severity) {
	a.severity.Store(severity)
}

// Epic linked to this test.
func (a *PluginAllure) Epic(epic string) {
	a.epic.Store(epic)
}

// Feature linked to this test.
func (a *PluginAllure) Feature(feature string) {
	a.feature.Store(feature)
}

// Story linked to this test.
func (a *PluginAllure) Story(story string) {
	a.story.Store(story)
}

// Flaky indicates that this test or step is known
// to be unstable and may not succeed every time.
func (a *PluginAllure) Flaky() {
	a.rawStatusDetails.Flaky.Store(true)
}

// Muted indicates that the result
// must not affect the statistics.
func (a *PluginAllure) Muted() {
	a.rawStatusDetails.Muted.Store(true)
}

// Known indicates that the test
// fails because of a known bug.
func (a *PluginAllure) Known() {
	a.rawStatusDetails.Known.Store(true)
}

// Attach an attachment.
//
// If option [WithMaxAttachmentSize] is specified, passed
// attachment is automatically trimmed of its suffix.
//
// Trimmed attachments are always of type [TextPlain] with suffix
// message added stating that an attachment exceeds a size limit.
//
// See [Bytes] and [File] to create an attachment.
//
//	t.Attach("login page", allure.Bytes([]byte(...)))
//	t.Attach("server response", allure.Bytes(`{"key": "value"}`))
func (a *PluginAllure) Attach(name string, at Attachment) {
	a.Helper()

	if a.excluded {
		return
	}

	if a.maxAttachmentSize <= 0 {
		a.attach(name, at)

		return
	}

	if size, ok := at.SizeHint(); ok && size <= a.maxAttachmentSize {
		a.attach(name, at)

		return
	}

	// fast path (most common).
	if b, ok := at.(AttachmentBytes); ok {
		trimmed := trimmedAttachment(
			b.Data,
			b.Type(),
			a.maxAttachmentSize,
		)

		a.attach(name, trimmed)

		return
	}

	r, err := at.Open()
	if err != nil {
		a.attach(name, at)

		return
	}

	defer func() { _ = r.Close() }()

	// add one extra byte so that [trimmedAttachment] trims it,
	// yet we don't load more data in memory than needed.
	data, err := io.ReadAll(io.LimitReader(r, a.maxAttachmentSize+1))
	if err != nil {
		a.attach(name, at)

		return
	}

	trimmed := trimmedAttachment(data, at.Type(), a.maxAttachmentSize)

	a.attach(name, trimmed)
}

func (a *PluginAllure) attach(name string, at Attachment) {
	a.Helper()

	if err := mkdir(a.outputDir); err != nil {
		a.Logf("allure: failed to create output dir: %v", err)

		return
	}

	source, mediaType, err := globalAttachmentWriter.Write(
		a.outputDir,
		at,
		a.deduplicateAttachments,
	)
	if err != nil {
		a.Logf("allure: failed to write attachment %q: %v", name, err)

		return
	}

	a.attachments.Modify(func(value *[]attachment) {
		*value = append(*value, attachment{
			Name:   name,
			Source: source,
			Type:   cmp.Or(at.Type(), mediaType),
		})
	})
}

func (a *PluginAllure) applyOptions(options []testoplugin.Option) {
	for _, o := range options {
		if o, ok := o.Value.(option); ok {
			o(a)
		}
	}
}

func (a *PluginAllure) plugin(
	parent *PluginAllure,
	options ...testoplugin.Option,
) testoplugin.Spec {
	a.parent = parent
	a.uuid = uuid.New().String()
	a.outputDir = *flagDir
	a.inverted = *flagInvert
	a.groupHooks = true
	a.handleTimeouts = true

	a.applyOptions(options)

	if parent != nil {
		parent.children.Modify(func(value *[]*PluginAllure) {
			*value = append(*value, a)
		})
	}

	return testoplugin.Spec{
		Plan:      a.plan(),
		Hooks:     a.hooks(),
		Overrides: a.overrides(),
	}
}

func (a *PluginAllure) panicked() bool {
	return testo.Reflect(a.T).Panic != nil
}

func (a *PluginAllure) status() Status {
	if status := a.statusOverwrite.Load(); !status.Unknown() {
		return status
	}

	// if test is still running it means we are forming
	// a report before full completion, which is possible only
	// during timeout.
	if a.running.Load() {
		return StatusBroken
	}

	return a.realStatus()
}

func (a *PluginAllure) realStatus() Status {
	if a.panicked() {
		return StatusBroken
	}

	if a.Failed() {
		r := testo.Reflect(a)

		// treat top level hooks (Before,AfterAll) differently
		if r.Test.GetLevel() == 0 && r.FailureSource != testoreflect.TestFailureSourceSelf {
			return StatusPassed
		}

		return StatusFailed
	}

	if a.Skipped() {
		return StatusSkipped
	}

	return StatusPassed
}

func (a *PluginAllure) asResult() result {
	start, stop := a.timeBoundaries()

	var descendants []*PluginAllure

	if testo.Reflect(a).Test.GetLevel() == 0 {
		descendants = a.hookDescendants()
	} else {
		descendants = a.descendants()
	}

	return result{
		UUID:          a.uuid,
		FullName:      a.fullName(),
		TestCaseID:    a.testCaseID(),
		HistoryID:     a.historyID(),
		Name:          a.title(),
		Description:   a.description.Load(),
		Links:         a.links(),
		Parameters:    a.parameters(),
		Labels:        a.labels(),
		Status:        a.status(),
		StatusDetails: a.propagatedStatusDetails(descendants),
		Attachments:   a.attachments.Load(),
		Start:         start,
		Stop:          stop,
		Steps:         a.steps(),
	}
}

func (a *PluginAllure) hookDescendants() []*PluginAllure {
	descendants := make([]*PluginAllure, 0, len(a.children.Load()))

	for _, child := range a.children.Load() {
		if !child.subtest() {
			continue
		}

		descendants = append(descendants, child)
		descendants = append(descendants, child.descendants()...)
	}

	return descendants
}

func (a *PluginAllure) descendants() []*PluginAllure {
	descendants := make([]*PluginAllure, 0, len(a.children.Load()))

	for _, child := range a.children.Load() {
		descendants = append(descendants, child)
		descendants = append(descendants, child.descendants()...)
	}

	return descendants
}

func (a *PluginAllure) fullName() string {
	return fullName(a.Name())
}

func fullName(name string) string {
	parent, base := path.Split(name)

	// remove (if any) test index.
	// e.g. "Suite/MyTest#01" -> "Suite/MyTest"
	idx := strings.LastIndex(base, "#")
	if idx != -1 {
		base = base[:idx]
	}

	return parent + base
}

func (a *PluginAllure) statusDetails() StatusDetails {
	return a.rawStatusDetails.toPublic()
}

func (a *PluginAllure) parameters() []parameter {
	raw := a.rawParameters.Load()

	params := make([]parameter, 0, len(raw))

	for _, p := range raw {
		params = append(params, p.toInternal())
	}

	return params
}

func (a *PluginAllure) links() []Link {
	links := a.rawLinks.Load().ClonedSlice()

	if a.linkTransformer == nil {
		return links
	}

	for i, l := range links {
		links[i] = a.linkTransformer(l)
	}

	return links
}

func (a *PluginAllure) title() string {
	overwrite := a.titleOverwrite.Load()

	if overwrite != "" {
		return overwrite
	}

	inspect := testo.Reflect(a.T)

	switch info := inspect.Test.(type) {
	case testoreflect.RegularTestInfo:
		return info.RawBaseName

	case testoreflect.ParametrizedTestInfo:
		return info.BaseName

	default:
		return testBaseName(a.Name())
	}
}

func (a *PluginAllure) asStep() step {
	start, stop := a.timeBoundaries()

	return step{
		Name:          a.title(),
		Status:        a.status(),
		StatusDetails: a.statusDetails(),
		Start:         start,
		Stop:          stop,
		Steps:         a.steps(),
		Parameters:    a.parameters(),
		Attachments:   a.attachments.Load(),
	}
}

func (a *PluginAllure) timeBoundaries() (start, stop unixMilli) {
	test := a.timeTest.Load()

	// when timed-out
	if test.Stop.IsZero() {
		test.Stop = time.Now()
	}

	start = unixMilli(test.Start.Add(-a.beforeParallel).UnixMilli())
	stop = unixMilli(test.Stop.UnixMilli())

	return start, stop
}

func (a *PluginAllure) subtest() bool {
	r := testo.Reflect(a)

	if t, ok := r.Test.(testoreflect.RegularTestInfo); ok {
		return t.IsSubtest
	}

	return false
}

func (a *PluginAllure) steps() []step {
	steps := make([]step, 0, len(a.children.Load()))

	level := testo.Reflect(a).Test.GetLevel()

	isSubtest := a.subtest()

	for _, c := range a.children.Load() {
		if c.excluded {
			continue
		}

		if isSubtest {
			steps = append(steps, c.asStep())

			continue
		}

		switch level {
		case 0, 1:
			if c.stage == stageTest {
				steps = append(steps, c.asStep())
			}

		default:
			steps = append(steps, c.asStep())
		}
	}

	return steps
}

func sharedContainer(befores, afters []step, children []UUID) container {
	var stop unixMilli

	start := unixMilli(math.MaxInt64)

	for _, b := range befores {
		start = min(start, b.Start)
		stop = max(stop, b.Stop)
	}

	for _, a := range afters {
		start = min(start, a.Start)
		stop = max(stop, a.Stop)
	}

	return container{
		UUID:     uuid.New().String(),
		Start:    start,
		Stop:     stop,
		Children: children,
		Befores:  befores,
		Afters:   afters,
	}
}

func (a *PluginAllure) asContainer() (container, bool) {
	var befores, afters []step

	start := unixMilli(math.MaxInt64)

	var stop unixMilli

	for _, child := range a.children.Load() {
		if child.excluded {
			continue
		}

		switch child.stage {
		case stageSetup:
			s := child.asStep()

			befores = append(befores, s)
			start = min(start, s.Start)
			stop = max(stop, s.Stop)

		case stageTearDown:
			s := child.asStep()

			afters = append(afters, s)
			start = min(start, s.Start)
			stop = max(stop, s.Stop)
		}
	}

	if len(befores) == 0 && len(afters) == 0 {
		return container{}, false
	}

	return container{
		UUID:     uuid.New().String(),
		Start:    start,
		Stop:     stop,
		Children: []UUID{a.uuid},
		Befores:  befores,
		Afters:   afters,
	}, true
}

func (a *PluginAllure) beforeAll() {
	if deadline, ok := a.Deadline(); ok && a.handleTimeouts {
		ctx, cancel := context.WithCancel(context.Background())

		a.Cleanup(func() { <-ctx.Done() })

		go func() {
			defer cancel()

			maxDuration := time.Until(deadline) - suiteDeadlineWindow

			select {
			case <-a.Context().Done():
				a.afterAll()

			case <-time.After(maxDuration):
				a.timedOut.Store(true)

				if !a.testTimedOut.Load() && !a.stepTimedOut.Load() {
					// fix current status so that call to error won't actually
					// mark this as failed (it'll be marked as broken later).
					a.Status(a.realStatus())

					a.Error("timed out")
				}

				a.afterAll()
			}
		}()
	} else {
		a.Cleanup(a.afterAll)
	}

	if err := writeCategories(a.outputDir, a.categories); err != nil {
		a.Logf("failed to write categories: %v", err)
	}

	a.running.Store(true)

	a.timeBeforeAll.Modify(func(value *timeBoundary) {
		value.Start = time.Now()
	})
}

//nolint:funlen,cyclop,gocyclo,gocognit // splitting would make less readable, probably
func (a *PluginAllure) afterAll() {
	a.running.Store(false)

	now := time.Now()

	// in case it was not set from tests (no tests)
	a.setBeforeAllStopOnce.Do(func() {
		a.timeBeforeAll.Modify(func(value *timeBoundary) {
			value.Stop = now
		})
	})

	a.timeAfterAll.Modify(func(value *timeBoundary) {
		value.Stop = now
	})

	if !a.Failed() && a.Skipped() {
		return
	}

	setups := a.queuedSetups.Load()
	tearDowns := a.queuedTearDowns.Load()

	tests := a.testChildren()

	standalone := func(hook string, steps []*PluginAllure, timing timeBoundary, status *Status) {
		res := a.asResult()

		if status != nil {
			res.Status = *status
		}

		if res.Status == StatusPassed {
			res.StatusDetails.Message = ""
			res.StatusDetails.Trace = ""
		}

		res.Name = hook

		res.FullName += "/" + strings.ReplaceAll(hook, " ", "_")

		res.Steps = make([]step, 0, len(setups))
		res.TestCaseID = testCaseID(res.FullName)
		res.HistoryID = res.TestCaseID
		res.UUID = uuid.New().String()
		res.Start = unixMilli(timing.Start.UnixMilli())
		res.Stop = unixMilli(timing.Stop.UnixMilli())

		for _, s := range steps {
			asStep := s.asStep()

			res.Steps = append(res.Steps, asStep)

			if res.Status == StatusPassed && asStep.Status == StatusFailed {
				res.Status = StatusFailed
			}
		}

		if err := writeResult(a.outputDir, res); err != nil {
			a.Logf("allure: failed to write result file: %v", err)
		}
	}

	{
		hooks := testo.Reflect(a).Suite.Hooks

		const (
			hookBeforeAll       = "Before All"
			hookAfterAll        = "After All"
			hooksBeforeAfterAll = "Before & After All"
		)

		switch {
		case !hooks.MissedBeforeAll && !hooks.MissedAfterAll && a.groupHooks:
			beforeAll := a.timeBeforeAll.Load()
			afterAll := a.timeAfterAll.Load()

			stop := beforeAll.Stop

			if !stop.Equal(afterAll.Stop) && !afterAll.Start.IsZero() {
				stop = stop.Add(afterAll.Duration())
			}

			if stop.IsZero() {
				stop = time.Now()
			}

			name := hooksBeforeAfterAll

			if a.timedOut.Load() && !a.testsStarted.Load() {
				name = hookBeforeAll
			}

			if a.testTimedOut.Load() {
				name = hookBeforeAll
			}

			all := make([]*PluginAllure, 0, len(setups)+len(tearDowns))

			all = append(all, setups...)
			all = append(all, tearDowns...)

			var status *Status

			if a.timedOut.Load() && !a.testTimedOut.Load() {
				s := StatusBroken

				status = &s
			}

			standalone(name, all, timeBoundary{
				Start: beforeAll.Start,
				// even though this is not technically correct,
				// as actual end time would equal to just timeAfterAll.Stop
				// this timings are primarily used for duration, not for stop & end stamps.
				// so to make duration more accurate, we should just add duration of AfterAll.
				Stop: stop,
			}, status)

		case !hooks.MissedBeforeAll:
			var status *Status

			if a.timedOut.Load() && !a.testsStarted.Load() {
				s := StatusBroken

				status = &s
			}

			standalone(hookBeforeAll, setups, a.timeBeforeAll.Load(), status)

			if hooks.MissedAfterAll {
				break
			}

			fallthrough

		case !hooks.MissedAfterAll:
			if a.testTimedOut.Load() {
				break
			}

			var status *Status

			if a.timedOut.Load() {
				s := StatusBroken

				status = &s
			}

			standalone(hookAfterAll, tearDowns, a.timeAfterAll.Load(), status)
		}
	}

	if len(tests) == 0 {
		return
	}

	befores := make([]step, 0, len(setups))
	afters := make([]step, 0, len(tearDowns))
	children := make([]UUID, 0, len(tests))

	fillSteps := func(from []*PluginAllure, to *[]step) {
		for _, f := range from {
			s := f.asStep()

			*to = append(*to, s)
		}
	}

	fillSteps(setups, &befores)
	fillSteps(tearDowns, &afters)

	for _, test := range tests {
		children = append(children, test.uuid)
	}

	if len(befores)+len(afters) == 0 {
		return
	}

	c := sharedContainer(befores, afters, children)

	if err := writeContainer(a.outputDir, c); err != nil {
		a.Logf("allure: failed to write container file: %v", err)
	}
}

//nolint:cyclop,funlen // TODO: split into subfunctions.
func (a *PluginAllure) beforeEach() {
	if a.parent != nil {
		a.parent.testsStarted.Store(true)
	}

	if deadline, ok := a.Deadline(); ok && a.handleTimeouts {
		ctx, cancel := context.WithCancel(context.Background())

		a.Cleanup(func() { <-ctx.Done() }) // halt test completion

		go func() {
			defer cancel()

			maxDuration := time.Until(deadline) - testDeadlineWindow

			// test runner immediately kills testing functions at specified deadline.
			// even cleanup functions won't run, therefore we need to
			// write a test report for the current test just before this test is killed.

			select {
			case <-a.Context().Done():
				a.afterEach()

			case <-time.After(maxDuration):
				a.timedOut.Store(true)

				if a.parent != nil {
					a.parent.testTimedOut.Store(true)
				}

				if !a.stepTimedOut.Load() {
					a.Error("timed out")
				}

				a.Status(StatusBroken)
				a.afterEach()
			}
		}()
	} else {
		a.Cleanup(a.afterEach)
	}

	if a.parent != nil {
		a.parent.setBeforeAllStopOnce.Do(func() {
			now := time.Now()

			a.parent.timeBeforeAll.Modify(func(value *timeBoundary) {
				value.Stop = now
			})
		})
	}

	inspect := testo.Reflect(a.T)

	if p, ok := inspect.Test.(testoreflect.ParametrizedTestInfo); ok {
		params := make([]Parameter, 0, len(p.Params))

		for name, value := range p.Params {
			params = append(params, NewParameter(name, value))
		}

		slices.SortStableFunc(params, func(a, b Parameter) int {
			return cmp.Compare(a.Name, b.Name)
		})

		a.Parameters(params...)
	}

	a.running.Store(true)

	a.timeTest.Modify(func(value *timeBoundary) {
		value.Start = time.Now()
	})
}

func (a *PluginAllure) beforeEachSub() {
	if deadline, ok := a.Deadline(); ok && a.handleTimeouts {
		ctx, cancel := context.WithCancel(context.Background())

		a.Cleanup(func() { <-ctx.Done() }) // halt test completion

		reflection := testo.Reflect(a)

		level := reflection.Test.GetLevel()

		const perLevel = 5 * time.Millisecond

		go func() {
			defer cancel()

			maxDuration := time.Until(deadline) - stepDeadlineWindow - perLevel*time.Duration(level)

			select {
			// can't use a.Context() because it is overridden for steps
			// to inherit Done() channel from parent, meaning its done
			// when parent is done which we don't want here.
			case <-reflection.TestingT.Context().Done():
				a.afterEachSub()

			case <-time.After(maxDuration):
				a.timedOut.Store(true)

				if a.parent != nil {
					a.parent.stepTimedOut.Store(true)
				}

				if !a.stepTimedOut.Load() {
					a.Error("timed out")
				}

				a.Status(StatusBroken)
				a.afterEachSub()
			}
		}()
	} else {
		a.Cleanup(a.afterEachSub)
	}

	a.running.Store(true)

	a.timeTest.Modify(func(value *timeBoundary) {
		value.Start = time.Now()
	})
}

var propertiesWritten sync.Map

func (a *PluginAllure) afterEach() {
	a.Helper()

	a.running.Store(false)

	now := time.Now()

	a.timeTest.Modify(func(value *timeBoundary) {
		value.Stop = now
	})

	inspect := testo.Reflect(a.T)

	if inspect.Panic != nil {
		a.rawStatusDetails.Message.Modify(func(value *string) {
			*value += fmt.Sprintf("panic: %v", inspect.Panic.Value)
		})

		a.rawStatusDetails.Trace.Store(inspect.Panic.Trace)
	}

	if a.excluded {
		return
	}

	if err := writeResult(a.outputDir, a.asResult()); err != nil {
		a.Logf("allure: failed to write result file: %v", err)
	}

	if c, ok := a.asContainer(); ok {
		if err := writeContainer(a.outputDir, c); err != nil {
			a.Logf("allure: failed to write container file: %v", err)
		}
	}

	if err := writeProperties(a.outputDir, newProperties()); err != nil {
		a.Logf("allure: failed to write properties file: %v", err)
	}

	if a.parent != nil {
		now := time.Now()

		a.parent.timeAfterAll.Modify(func(value *timeBoundary) {
			if value.Start.Compare(now) < 0 {
				value.Start = now
			}
		})
	}
}

func (a *PluginAllure) afterEachSub() {
	a.Helper()

	a.running.Store(false)

	a.timeTest.Modify(func(value *timeBoundary) {
		value.Stop = time.Now()
	})

	inspect := testo.Reflect(a.T)

	if inspect.Panic != nil {
		a.rawStatusDetails.Message.Modify(func(value *string) {
			*value += fmt.Sprintf("panic: %v", inspect.Panic.Value)
		})

		a.rawStatusDetails.Trace.Store(inspect.Panic.Trace)
	}

	if a.excluded || inspect.Test.GetLevel() != 1 {
		return
	}

	// in after all
	if a.hasTestNeighbors() {
		a.parent.queuedTearDowns.Modify(func(value *[]*PluginAllure) {
			*value = append(*value, a)
		})

		return
	}

	// in before all
	a.parent.queuedSetups.Modify(func(value *[]*PluginAllure) {
		*value = append(*value, a)
	})
}

func (a *PluginAllure) testChildren() []*PluginAllure {
	children := make([]*PluginAllure, 0, len(a.children.Load()))

	for _, child := range a.children.Load() {
		switch t := testo.Reflect(child).Test.(type) {
		case testoreflect.ParametrizedTestInfo:
			children = append(children, child)

		case testoreflect.RegularTestInfo:
			if !t.IsSubtest {
				children = append(children, child)
			}
		}
	}

	return children
}

func (a *PluginAllure) hasTestNeighbors() bool {
	if a.parent == nil {
		return false
	}

	for _, neighbor := range a.parent.children.Load() {
		if neighbor.uuid == a.uuid {
			continue
		}

		switch t := testo.Reflect(neighbor).Test.(type) {
		case testoreflect.ParametrizedTestInfo:
			return true

		case testoreflect.RegularTestInfo:
			if !t.IsSubtest {
				return true
			}
		}
	}

	return false
}

func (a *PluginAllure) plan() testoplugin.Plan {
	a.Helper()

	return testoplugin.Plan{
		Prepare: func(_ testoreflect.SuiteInfo, tests *[]testoplugin.PlannedTest) {
			a.Helper()

			planPath := os.Getenv("ALLURE_TESTPLAN_PATH")
			if planPath == "" {
				return
			}

			data, err := os.ReadFile(planPath)
			if err != nil {
				return
			}

			var plan testPlan

			err = json.Unmarshal(data, &plan)
			if err != nil {
				return
			}

			ids, selectors := plan.sets()

			// If test plan does not specify any tests we should ignore it,
			// rather than skipping all the tests.
			//
			// This is what pytest plugin does:
			// https://github.com/allure-framework/allure-python/blob/21571c8a7a7792a4dcd63b0240bb6e91bef0d4a6/allure-pytest/src/plugin.py#L202
			if len(ids) == 0 && len(selectors) == 0 {
				return
			}

			total := len(*tests)

			*tests = slices.DeleteFunc(*tests, func(pt testoplugin.PlannedTest) bool {
				isSelected := selectors[pt.Info().GetName()]

				if !isSelected {
					var tmp PluginAllure

					tmp.applyOptions(pt.Annotations())

					isSelected = ids[tmp.allureID.Load()]
				}

				if a.inverted {
					return isSelected
				}

				return !isSelected
			})

			a.Logf(
				"allure: using test plan v%s from %q, tests excluded: %d",
				plan.Version,
				planPath,
				total-len(*tests),
			)
		},
	}
}

func (a *PluginAllure) hooks() testoplugin.Hooks {
	return testoplugin.Hooks{
		BeforeAll:     testoplugin.Hook{Func: a.beforeAll},
		BeforeEach:    testoplugin.Hook{Func: a.beforeEach, Priority: testoplugin.TryFirst},
		BeforeEachSub: testoplugin.Hook{Func: a.beforeEachSub},
	}
}

func (a *PluginAllure) addMessage(msg string) {
	a.rawStatusDetails.Message.Modify(func(value *string) {
		if *value == "" {
			*value = msg
		} else {
			*value += "\n\n" + msg
		}
	})
}

func (a *PluginAllure) addTrace(trace string) {
	a.rawStatusDetails.Trace.Modify(func(value *string) {
		if *value == "" {
			*value = trace
		} else {
			*value += "\n\n" + trace
		}
	})
}

func captureTrace[F ~func()](a *PluginAllure) testoplugin.Override[F] {
	return func(f F) F {
		return func() {
			a.Helper()

			if !a.inAssertion.Load() {
				a.addTrace(stacktrace.Take(1))
			}

			f()
		}
	}
}

func captureOutput[F ~func(...any)](a *PluginAllure) testoplugin.Override[F] {
	return func(f F) F {
		return func(args ...any) {
			a.Helper()

			msg := fmt.Sprint(args...)

			if a.inAssertion.Load() {
				msg = transformTestifyErrorMsg(msg)
			}

			a.addMessage(trimLines(msg))

			f(msg)
		}
	}
}

var trimTestifyErrorTraceRegex = regexp.MustCompile(`(?sU)Error Trace:.+\s*Error:`)

func transformTestifyErrorMsg(s string) string {
	s = trimTestifyErrorTrace(s)
	s = fixTestifyErrorMsg(s)

	return s
}

func fixTestifyErrorMsg(s string) string {
	lines := strings.Split(s, "\n")

	for i, l := range lines {
		l = strings.TrimSpace(l)

		const (
			errPrefix  = "Error: "
			testPrefix = "Test: "
		)

		switch {
		case strings.HasPrefix(l, errPrefix):
			const limit = 2000

			if len(l) > limit {
				l = l[:limit] + "..."
			}

			idx := strings.IndexFunc(
				l[len(errPrefix):],
				func(r rune) bool { return !unicode.IsSpace(r) },
			)

			if idx >= 0 {
				lines[i] = l[:len(errPrefix)] + l[len(errPrefix)+idx:]
			}

		case strings.HasPrefix(l, testPrefix):
			idx := strings.IndexFunc(
				l[len(testPrefix):],
				func(r rune) bool { return !unicode.IsSpace(r) },
			)

			if idx >= 0 {
				lines[i] = l[:len(testPrefix)] + " " + l[len(testPrefix)+idx:]
			}
		}
	}

	return strings.Join(lines, "\n")
}

func trimTestifyErrorTrace(s string) string {
	return trimTestifyErrorTraceRegex.ReplaceAllString(s, "Error:")
}

func trimCallerLine(s string) string {
	lines := strings.Split(s, "\n")

	lines = slices.DeleteFunc(lines, func(l string) bool {
		return l == "" || strings.HasPrefix(l, "Caller: ")
	})

	return strings.Join(lines, "\n")
}

func (a *PluginAllure) overrides() testoplugin.Overrides {
	return testoplugin.Overrides{
		Skip:  captureOutput[testoplugin.FuncSkip](a),
		Error: captureOutput[testoplugin.FuncError](a),
		Fatal: captureOutput[testoplugin.FuncFatal](a),

		FailNow: captureTrace[testoplugin.FuncFailNow](a),
		Fail:    captureTrace[testoplugin.FuncFail](a),

		Context: func(f testoplugin.FuncContext) testoplugin.FuncContext {
			if !a.inStep.Load() {
				return f
			}

			return func() context.Context {
				var parentCtx context.Context

				// parent being nil in step
				// should be unreachable, but just in case
				if a.parent != nil {
					parentCtx = a.parent.Context()
				}

				ctx := context.WithoutCancel(f())

				if parentCtx == nil {
					return ctx
				}

				ctx, cancel := context.WithCancel(ctx)

				if deadline, ok := parentCtx.Deadline(); ok {
					ctx, cancel = context.WithDeadline(ctx, deadline)
				}

				// cancel this context as soon as parent is cancelled.
				go func() {
					<-parentCtx.Done()

					cancel()
				}()

				return ctx
			}
		},

		Parallel: func(f testoplugin.FuncParallel) testoplugin.FuncParallel {
			return func() {
				// if start is zero it means we are inside a BeforeEach hook of other plugin.
				// in that case, real test has not started yet, so we shouldn't compute beforeParallel timing.
				test := a.timeTest.Load()
				if !test.Start.IsZero() {
					a.beforeParallel = time.Since(test.Start)
				}

				f()

				a.timeTest.Modify(func(value *timeBoundary) {
					value.Start = time.Now()
				})
			}
		},
	}
}

func writeResult(dir string, res result) error {
	if err := mkdir(dir); err != nil {
		return fmt.Errorf("create dir: %w", err)
	}

	marshalled, err := json.Marshal(res)
	if err != nil {
		return fmt.Errorf("marshal allure test result file: %w", err)
	}

	resultPath := filepath.Join(dir, res.UUID+"-result.json")

	err = os.WriteFile(
		resultPath,
		marshalled,
		permFile,
	)
	if err != nil {
		return fmt.Errorf("write test result file for %q at %q: %w", res.FullName, resultPath, err)
	}

	return nil
}

func writeContainer(dir string, c container) error {
	if err := mkdir(dir); err != nil {
		return fmt.Errorf("mkdir: %w", err)
	}

	marshalled, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("marshal container: %w", err)
	}

	containerPath := filepath.Join(dir, c.UUID+"-container.json")

	err = os.WriteFile(
		containerPath,
		marshalled,
		permFile,
	)
	if err != nil {
		return fmt.Errorf("write container file to %q: %w", containerPath, err)
	}

	return nil
}

func mkdir(dir string) error {
	err := os.MkdirAll(dir, permDir)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}

	return nil
}

func writeProperties(dir string, p properties) error {
	if _, loaded := propertiesWritten.LoadOrStore(dir, true); loaded {
		return nil
	}

	marshalled, err := p.MarshalProperties()
	if err != nil {
		return fmt.Errorf("marshal properties: %w", err)
	}

	const filename = "environment.properties"

	propertiesPath := filepath.Join(dir, filename)

	err = os.WriteFile(propertiesPath, marshalled, permFile)
	if err != nil {
		return fmt.Errorf("write properties file at %s: %w", propertiesPath, err)
	}

	return nil
}

var writeCategoriesMutex sync.Mutex

func writeCategories(dir string, categories []Category) error {
	if len(categories) == 0 {
		return nil
	}

	// If multiple suites are run in parallel, there exists a small
	// chance that they will finish at the same time.
	// In that case categories file won't be written properly.
	//
	// TODO(metafates): is this enough or should we use [file locks]?
	//
	// [file locks]: https://pkg.go.dev/cmd/go/internal/lockedfile/internal/filelock#Lock
	writeCategoriesMutex.Lock()
	defer writeCategoriesMutex.Unlock()

	// This is tricky.
	// We could already have categories file written
	// by other suite, so we need to append to it.
	// But also we have to remain categories unique.
	categoriesPath := filepath.Join(dir, "categories.json")

	readExisting := func() []Category {
		file, err := os.ReadFile(categoriesPath)
		if err != nil {
			return nil
		}

		var out []Category

		// if json is malformed we should ignore it and overwrite.
		_ = json.Unmarshal(file, &out)

		return out
	}

	categories = slices.Clone(categories)

	categories = append(readExisting(), categories...)

	categories = uniqueCategories(categories)

	if len(categories) == 0 {
		return nil
	}

	marshalled, err := json.Marshal(categories)
	if err != nil {
		return fmt.Errorf("marshal categories: %w", err)
	}

	err = os.WriteFile(categoriesPath, marshalled, permFile)
	if err != nil {
		return fmt.Errorf("write categories file at %q: %w", categoriesPath, err)
	}

	return nil
}

func (a *PluginAllure) labels() []Label {
	labels := a.rawLabels.Load().ClonedSlice()

	labelsSet := make(map[string]bool, len(labels))

	for _, l := range labels {
		labelsSet[l.Name] = true
	}

	hostname, _ := os.Hostname()

	// these labels we should not add if user already did so.
	// because these labels are added implicitly without user interaction.
	for _, l := range []Label{
		{Name: labelParentSuite, Value: a.parentSuiteName()},
		{Name: labelSuite, Value: a.suiteName()},
		{Name: labelHost, Value: hostname},
		{Name: labelLanguage, Value: "go"},
		{Name: labelFramework, Value: "testo"},
	} {
		if l.Value != "" && !labelsSet[l.Name] {
			labels = append(labels, l)
		}
	}

	// these labels take precedence over user added raw labels.
	// because these labels are added explicitly by the user.
	for _, l := range []Label{
		{Name: labelOwner, Value: a.owner.Load()},
		{Name: labelEpic, Value: a.epic.Load()},
		{Name: labelFeature, Value: a.feature.Load()},
		{Name: labelStory, Value: a.story.Load()},
		{Name: labelSeverity, Value: a.severity.Load().String()},
		{Name: labelAllureID, Value: a.allureID.Load()},
	} {
		if l.Value == "" {
			continue
		}

		if labelsSet[l.Name] {
			labels = slices.DeleteFunc(labels, func(label Label) bool {
				return label.Name == l.Name
			})
		}

		labels = append(labels, l)
	}

	return labels
}

// unwrapAllure returns underlying [PluginAllure] instance.
func (a *PluginAllure) unwrapAllure() *PluginAllure { return a }

func (a *PluginAllure) testCaseID() string {
	return testCaseID(a.fullName())
}

func testCaseID(fullName string) string {
	id := fmt.Sprintf("%016x", allurehash.Hash(fullName))

	return id
}

func (a *PluginAllure) historyID() string {
	raw := a.rawParameters.Load()

	params := make([]Parameter, 0, len(raw))

	for _, p := range raw {
		if p.Exclude {
			continue
		}

		params = append(params, p)
	}

	slices.SortStableFunc(params, func(a, b Parameter) int {
		return cmp.Compare(a.Name, b.Name)
	})

	type Signature struct {
		Name   string
		Params []Parameter
	}

	id := fmt.Sprintf("%016x", allurehash.Hash(Signature{
		Name:   a.fullName(),
		Params: params,
	}))

	return id
}

func (a *PluginAllure) parentSuiteName() string {
	parent := testo.Reflect(a).Suite.Parent
	if parent == nil {
		return ""
	}

	if stringer, ok := parent.Value.(fmt.Stringer); ok {
		return stringer.String()
	}

	return parent.Name
}

func (a *PluginAllure) suiteName() string {
	s := testo.Reflect(a).Suite

	if stringer, ok := s.Value.(fmt.Stringer); ok {
		return stringer.String()
	}

	return s.Name
}

func newProperties() properties {
	return properties{
		GoOS:       runtime.GOOS,
		GoArch:     runtime.GOARCH,
		GoVersion:  runtime.Version(),
		GoCompiler: runtime.Compiler,
		NumCPU:     runtime.NumCPU(),
	}
}

func trimLines(s string) string {
	s = strings.TrimSpace(s)

	lines := strings.Split(s, "\n")

	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}

	return strings.Join(lines, "\n")
}

func uniqueCategories(categories []Category) []Category {
	unique := uniqueBy(categories, func(c Category) string {
		return c.Name
	})

	return unique
}

func uniqueBy[S ~[]T, K comparable, T any](s S, f func(T) K) S {
	seen := make(map[K]bool, len(s))

	res := make(S, 0, len(s))

	for _, e := range s {
		key := f(e)

		if !seen[key] {
			seen[key] = true

			res = append(res, e)
		}
	}

	return res
}

func testBaseName(testName string) string {
	idx := strings.LastIndex(testName, "/")
	if idx == -1 {
		return testName
	}

	return testName[idx+1:]
}

func (a *PluginAllure) propagatedStatusDetails(descendants []*PluginAllure) StatusDetails {
	var details []StatusDetails

	for _, p := range append([]*PluginAllure{a}, descendants...) {
		d := p.statusDetails()

		if p.inAssertion.Load() {
			d.Message = trimCallerLine(d.Message)
		}

		if d.Message != "" {
			details = append(details, d)
		}
	}

	messages := make([]string, 0, len(details))
	traces := make([]string, 0, len(details))

	for _, d := range details {
		messages = append(messages, d.Message)

		if d.Trace == "" {
			continue
		}

		if len(details) == 1 {
			traces = append(traces, d.Trace)
		} else {
			traces = append(traces, fmt.Sprintf("%s\n\n%s", d.Message, d.Trace))
		}
	}

	return StatusDetails{
		Message: strings.Join(messages, "\n\n\n"),
		Trace:   strings.Join(traces, "\n\n\n"),
	}
}
