// Package stacktrace provides support for gathering stack traces
// efficiently.
//
// Simplified version of [uber-go/zap] implementation.
//
// [uber-go/zap]: https://github.com/uber-go/zap/blob/07077a697f639389cc998ff91b8885feb25f520d/internal/stacktrace/stack.go#L4
package stacktrace

import (
	"bytes"
	"runtime"
	"strconv"
	"sync"
)

var stackPool = sync.Pool{
	New: func() any {
		return &Stack{
			storage: make([]uintptr, 64),
		}
	},
}

var bufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

// Stack is a captured stack trace.
type Stack struct {
	pcs    []uintptr // program counters; always a subslice of storage
	frames *runtime.Frames

	// The size of pcs varies depending on requirements:
	// it will be one if the only the first frame was requested,
	// and otherwise it will reflect the depth of the call stack.
	//
	// storage decouples the slice we need (pcs) from the slice we pool.
	// We will always allocate a reasonably large storage, but we'll use
	// only as much of it as we need.
	storage []uintptr
}

// Depth specifies how deep of a stack trace should be captured.
type Depth int

const (
	// First captures only the first frame.
	First Depth = iota

	// Full captures the entire call stack, allocating more
	// storage for it if needed.
	Full
)

// capture captures a stack trace of the specified depth, skipping
// the provided number of frames. Value of zero for skip identifies the caller of
// capture.
//
// The caller must call Free on the returned stacktrace after using it.
func capture(skip int, depth Depth) *Stack {
	stack := stackPool.Get().(*Stack)

	switch depth {
	case First:
		stack.pcs = stack.storage[:1]
	case Full:
		stack.pcs = stack.storage
	}

	// Unlike other "skip"-based APIs, skip=0 identifies runtime.Callers
	// itself. +2 to skip [Capture] and runtime.Callers.
	numFrames := runtime.Callers(
		skip+2,
		stack.pcs,
	)

	// runtime.Callers truncates the recorded stacktrace if there is no
	// room in the provided slice. For the full stack trace, keep expanding
	// storage until there are fewer frames than there is room.
	if depth == Full {
		pcs := stack.pcs
		for numFrames == len(pcs) {
			pcs = make([]uintptr, len(pcs)*2)
			numFrames = runtime.Callers(skip+2, pcs)
		}

		// Discard old storage instead of returning it to the pool.
		// This will adjust the pool size over time if stack traces are
		// consistently very deep.
		stack.storage = pcs
		stack.pcs = pcs[:numFrames]
	} else {
		stack.pcs = stack.pcs[:numFrames]
	}

	stack.frames = runtime.CallersFrames(stack.pcs)

	return stack
}

// Free releases resources associated with this stacktrace
// and returns it back to the pool.
func (st *Stack) Free() {
	st.frames = nil
	st.pcs = nil
	stackPool.Put(st)
}

// Count reports the total number of frames in this stacktrace.
// Count DOES NOT change as Next is called.
func (st *Stack) Count() int {
	return len(st.pcs)
}

// Next returns the next frame in the stack trace,
// and a boolean indicating whether there are more after it.
func (st *Stack) Next() (runtime.Frame, bool) {
	return st.frames.Next()
}

// Take returns a string representation of the current stacktrace.
//
// skip is the number of frames to skip before recording the stack trace.
// skip=0 identifies the caller of Take.
func Take(skip int) string {
	stack := capture(skip+1, Full)
	defer stack.Free()

	buffer := bufferPool.Get().(*bytes.Buffer)
	defer buffer.Reset()

	stackfmt := NewFormatter(buffer)
	stackfmt.FormatStack(stack)

	return buffer.String()
}

// Formatter formats a stack trace into a readable string representation.
type Formatter struct {
	b        *bytes.Buffer
	nonEmpty bool // whether we've written at least one frame already
}

// NewFormatter builds a new Formatter.
func NewFormatter(b *bytes.Buffer) Formatter {
	return Formatter{b: b}
}

// FormatStack formats all remaining frames in the provided stacktrace -- minus
// the final runtime.main/runtime.goexit frame.
func (sf *Formatter) FormatStack(stack *Stack) {
	// Note: On the last iteration, frames.Next() returns false, with a valid
	// frame, but we ignore this frame. The last frame is a runtime frame which
	// adds noise, since it's only either runtime.main or runtime.goexit.
	for frame, more := stack.Next(); more; frame, more = stack.Next() {
		sf.FormatFrame(frame)
	}
}

// FormatFrame formats the given frame.
func (sf *Formatter) FormatFrame(frame runtime.Frame) {
	if sf.nonEmpty {
		sf.b.WriteByte('\n')
	}

	sf.nonEmpty = true
	sf.b.WriteString(frame.Function)
	sf.b.WriteByte('\n')
	sf.b.WriteByte('\t')
	sf.b.WriteString(frame.File)
	sf.b.WriteByte(':')
	sf.b.WriteString(strconv.Itoa(frame.Line))
}
