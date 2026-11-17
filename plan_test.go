package allure //nolint:testpackage // Tests exercise private planning and option payloads.

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/ozontech/testo"
	"github.com/ozontech/testo/testoplugin"
	"github.com/ozontech/testo/testoreflect"
	"github.com/stretchr/testify/require"
)

const (
	plannedSampleName = "TestSample"
	regularExecution  = "regular"
	rowOneA           = "1/a"
	rowOneB           = "1/b"
	rowTwoA           = "2/a"
	rowTwoB           = "2/b"
)

// TestPlanAnnotationOptions checks that metadata annotations do not prevent ID selection.
func TestPlanAnnotationOptions(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name string
		opt  testoplugin.Option
	}{
		{name: "labels", opt: WithLabels(NewLabel("component", "sample"))},
		{name: "links", opt: WithLinks(NewLink("https://example.org/reference"))},
		{name: "tags", opt: WithTags("smoke")},
		{name: "owner", opt: WithOwner("sample-owner")},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			test := &plannedTestStub{annotations: []testoplugin.Option{WithID("101"), tc.opt}}

			var selected []testoplugin.PlannedTest

			require.NotPanics(t, func() {
				selected = filterPlannedTests(
					[]testoplugin.PlannedTest{test},
					map[string]bool{"101": true},
					nil,
					false,
				)
			})
			require.Equal(t, []testoplugin.PlannedTest{test}, selected)
		})
	}
}

// TestPlanSelectors checks selector/ID union and inversion without executing runtime options.
func TestPlanSelectors(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name      string
		ids       map[string]bool
		selectors map[string]bool
		inverted  bool
		want      int
	}{
		{name: "selector", selectors: map[string]bool{plannedSampleName: true}, want: 1},
		{
			name:      "selector with unmatched ID",
			ids:       map[string]bool{"999": true},
			selectors: map[string]bool{plannedSampleName: true},
			want:      1,
		},
		{
			name:      "ID with unmatched selector",
			ids:       map[string]bool{"101": true},
			selectors: map[string]bool{"Other": true},
			want:      1,
		},
		{name: "inverted selector", selectors: map[string]bool{plannedSampleName: true}, inverted: true},
		{name: "unmatched selector", selectors: map[string]bool{"Other": true}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			test := &plannedTestStub{annotations: []testoplugin.Option{
				WithID("101"),
				{
					Value: option(
						func(*PluginAllure) { panic("runtime option applied during planning") },
					),
				},
			}}

			selected := filterPlannedTests(
				[]testoplugin.PlannedTest{test},
				tc.ids,
				tc.selectors,
				tc.inverted,
			)

			require.Len(t, selected, tc.want)
		})
	}
}

// TestPlanIDPrecedence checks that planning and runtime application both use the last ID,
// including an empty ID that clears an earlier value.
func TestPlanIDPrecedence(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name       string
		ids        []string
		want       string
		selected   int
		selectedID string
	}{
		{name: "last ID wins", ids: []string{"101", "202"}, want: "202", selected: 1, selectedID: "202"},
		{name: "empty ID clears", ids: []string{"101", ""}, want: "", selectedID: "101"},
		{name: "missing ID", selectedID: "101"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			options := make([]testoplugin.Option, 0, len(tc.ids))

			for _, id := range tc.ids {
				options = append(options, WithID(id))
			}

			var plugin PluginAllure

			plugin.applyOptions(options)
			require.Equal(t, tc.want, plugin.allureID.Load())

			test := &plannedTestStub{annotations: options}
			selected := filterPlannedTests(
				[]testoplugin.PlannedTest{test},
				map[string]bool{tc.selectedID: true},
				nil,
				false,
			)
			require.Len(t, selected, tc.selected)
		})
	}
}

// TestPlanMetadata checks selection through real Testo collection and execution,
// including every parametrized row, and verifies metadata in the emitted reports.
func TestPlanMetadata(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name     string
		plan     testPlan
		inverted bool
		want     []string
		wantIDs  map[string]int
	}{
		{
			name: "regular ID", plan: testPlan{Tests: []testPlanTest{{ID: "101"}}},
			want:    []string{regularExecution},
			wantIDs: map[string]int{"101": 1},
		},
		{
			name: "parametrized ID", plan: testPlan{Tests: []testPlanTest{{ID: "202"}}},
			want: []string{rowOneA, rowOneB, rowTwoA, rowTwoB}, wantIDs: map[string]int{"202": 4},
		},
		{
			name: "inverted regular ID", plan: testPlan{Tests: []testPlanTest{{ID: "101"}}}, inverted: true,
			want: []string{rowOneA, rowOneB, rowTwoA, rowTwoB}, wantIDs: map[string]int{"202": 4},
		},
		{
			name: "inverted parametrized ID", plan: testPlan{Tests: []testPlanTest{{ID: "202"}}}, inverted: true,
			want: []string{regularExecution}, wantIDs: map[string]int{"101": 1},
		},
		{name: "unknown ID", plan: testPlan{Tests: []testPlanTest{{ID: "999"}}}, wantIDs: map[string]int{}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			dir := t.TempDir()
			suite := new(planMetadataSuite)
			testo.RunSuite(t, suite, WithOutputDir(dir),
				testoplugin.Option{Value: planConfig{plan: tc.plan, inverted: tc.inverted}})

			require.ElementsMatch(t, tc.want, suite.ran)

			results := readPlanResults(t, dir)
			require.Len(t, results, len(tc.want))

			ids := make(map[string]int)

			for _, res := range results {
				require.Contains(t, res.Labels, NewLabel("component", "sample"))
				require.Contains(t, res.Labels, NewLabel(labelTag, "smoke"))
				require.Contains(t, res.Labels, NewLabel(labelOwner, "sample-owner"))
				require.Equal(t, []Link{NewLink("https://example.org/reference")}, res.Links)
				require.Equal(t, "passed", res.Status)

				for _, label := range res.Labels {
					if label.Name == labelAllureID {
						ids[label.Value]++
					}
				}
			}

			require.Equal(t, tc.wantIDs, ids)
		})
	}
}

// TestMetadataWithoutPlan checks that all tests run and retain their annotated
// metadata when no explicit plan is supplied.
func TestMetadataWithoutPlan(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	suite := new(planMetadataSuite)
	testo.RunSuite(t, suite, WithOutputDir(dir))

	require.ElementsMatch(
		t,
		[]string{regularExecution, rowOneA, rowOneB, rowTwoA, rowTwoB},
		suite.ran,
	)

	results := readPlanResults(t, dir)
	require.Len(t, results, 5)

	ids := make(map[string]int)

	for _, res := range results {
		require.Contains(t, res.Labels, NewLabel("component", "sample"))
		require.Contains(t, res.Labels, NewLabel(labelTag, "smoke"))
		require.Contains(t, res.Labels, NewLabel(labelOwner, "sample-owner"))
		require.Equal(t, []Link{NewLink("https://example.org/reference")}, res.Links)
		require.Equal(t, "passed", res.Status)

		for _, label := range res.Labels {
			if label.Name == labelAllureID {
				ids[label.Value]++
			}
		}
	}

	require.Equal(t, map[string]int{"101": 1, "202": 4}, ids)
}

// planResult reads only the report fields asserted by the planning tests.
// Status stays a string because the production status type only supports marshaling.
type planResult struct {
	Labels []Label `json:"labels"`
	Links  []Link  `json:"links"`
	Status string  `json:"status"`
}

// readPlanResults decodes test results from a test-owned output directory.
// Container and environment files are excluded from the result count.
func readPlanResults(t *testing.T, dir string) []planResult {
	t.Helper()

	files, err := filepath.Glob(filepath.Join(dir, "*-result.json"))
	require.NoError(t, err)

	results := make([]planResult, 0, len(files))

	for _, file := range files {
		data, err := os.ReadFile(file)
		require.NoError(t, err)

		var res planResult

		require.NoError(t, json.Unmarshal(data, &res))

		results = append(results, res)
	}

	return results
}

// plannedTestStub supplies a stable name and annotations for focused filtering tests.
// The embedded interface satisfies Testo's sealed contract; only the methods below are used.
type plannedTestStub struct {
	testoplugin.PlannedTest

	annotations []testoplugin.Option
}

// Info returns the exact name used by selector fixtures.
func (*plannedTestStub) Info() testoreflect.TestInfo {
	return testoreflect.RegularTestInfo{Name: plannedSampleName}
}

// Annotations returns the option sequence supplied by the scenario.
func (p *plannedTestStub) Annotations() []testoplugin.Option {
	return p.annotations
}

// planConfig passes selection inputs to the fixture without changing process state.
type planConfig struct {
	plan     testPlan
	inverted bool
}

// PluginTestPlan supplies an explicit plan to exercise selection without process environment changes.
type PluginTestPlan struct{}

// Plugin implements [testoplugin.Plugin] for the explicit test plan fixture.
func (*PluginTestPlan) Plugin(
	_ testoplugin.Plugin,
	options ...testoplugin.Option,
) testoplugin.Spec {
	for _, opt := range options {
		if cfg, ok := opt.Value.(planConfig); ok {
			return testoplugin.Spec{Plan: testoplugin.Plan{
				Prepare: func(_ testoreflect.SuiteInfo, tests *[]testoplugin.PlannedTest) {
					ids, selectors := cfg.plan.sets()

					*tests = filterPlannedTests(*tests, ids, selectors, cfg.inverted)
				},
			}}
		}
	}

	return testoplugin.Spec{}
}

type planT struct {
	*testo.T
	*PluginAllure
	*PluginTestPlan
}

// planMetadataSuite records executed rows independently of report generation.
type planMetadataSuite struct {
	testo.Suite[planT]

	ran []string
}

var _ = testo.For((*planMetadataSuite).TestRegular,
	WithID("101"),
	WithLabels(NewLabel("component", "sample")),
	WithLinks(NewLink("https://example.org/reference")),
	WithTags("smoke"),
	WithOwner("sample-owner"),
)

// TestRegular records execution of the ordinary test annotated with ID 101.
func (s *planMetadataSuite) TestRegular(_ planT) {
	s.ran = append(s.ran, regularExecution)
}

// CasesNumber supplies one axis of the four-row Cartesian product.
func (*planMetadataSuite) CasesNumber() []int { return []int{1, 2} }

// CasesLetter supplies the second axis of the four-row Cartesian product.
func (*planMetadataSuite) CasesLetter() []string { return []string{"a", "b"} }

var _ = testo.ForEach((*planMetadataSuite).TestParametrized,
	WithLabels(NewLabel("component", "sample")),
	WithLinks(NewLink("https://example.org/reference")),
	WithTags("smoke"),
	WithOwner("sample-owner"),
	WithID("202"),
)

// TestParametrized records each parameter combination sharing the annotated ID 202.
func (s *planMetadataSuite) TestParametrized(_ planT, p struct {
	Number int
	Letter string
},
) {
	s.ran = append(s.ran, fmt.Sprintf("%d/%s", p.Number, p.Letter))
}
