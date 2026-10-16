package allure

import (
	"path"
	"sync"
	"testing"

	"github.com/ozontech/testo"
	"github.com/stretchr/testify/require"
)

type skipResult struct {
	Skipped bool
	Reached bool
}

type SkipStepSuite struct {
	testo.Suite[T]

	mu      sync.Mutex
	results map[string]skipResult
}

// track records whether the test was skipped once it completes.
func (s *SkipStepSuite) track(t T) {
	t.Cleanup(func() {
		s.update(t, func(r *skipResult) { r.Skipped = t.Skipped() })
	})
}

// reach records that the test continued after the step.
func (s *SkipStepSuite) reach(t T) {
	s.update(t, func(r *skipResult) { r.Reached = true })
}

func (s *SkipStepSuite) update(t T, f func(r *skipResult)) {
	s.mu.Lock()
	defer s.mu.Unlock()

	name := path.Base(t.Name())

	r := s.results[name]
	f(&r)
	s.results[name] = r
}

func (s *SkipStepSuite) TestSkip(t T) {
	s.track(t)

	Step(t, "skipped step", func(t T) {
		t.Skipf("skip %s", "reason")
	})

	s.reach(t)
}

func (s *SkipStepSuite) TestNestedSkip(t T) {
	s.track(t)

	Step(t, "outer step", func(t T) {
		Step(t, "inner step", func(t T) {
			t.Skip("skip")
		})
	})

	s.reach(t)
}

func (s *SkipStepSuite) TestSetupSkip(t T) {
	s.track(t)

	Setup(t, "skipped setup", func(t T) {
		t.SkipNow()
	})

	s.reach(t)
}

func (s *SkipStepSuite) TestTearDownSkip(t T) {
	s.track(t)

	TearDown(t, "skipped teardown", func(t T) {
		t.SkipNow()
	})

	s.reach(t)
}

func (s *SkipStepSuite) TestParallelStepSkip(t T) {
	s.track(t)

	Step(t, "parallel step", func(t T) {
		t.Parallel()
		t.Skip("skip")
	})

	s.reach(t)
}

func TestStepSkip(t *testing.T) {
	t.Chdir(t.TempDir())

	s := &SkipStepSuite{results: make(map[string]skipResult)}

	testo.RunSuite(t, s)

	require.Equal(t, map[string]skipResult{
		"TestSkip":         {Skipped: true},
		"TestNestedSkip":   {Skipped: true},
		"TestSetupSkip":    {Skipped: true},
		"TestTearDownSkip": {Skipped: true},
		// skip of a parallel step is not propagated
		"TestParallelStepSkip": {Reached: true},
	}, s.results)
}
