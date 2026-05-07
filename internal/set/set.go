package set

import (
	"slices"
)

type Set[T comparable] struct {
	seen   map[T]bool
	values []T
}

func (s Set[T]) ClonedSlice() []T {
	return slices.Clone(s.values)
}

func (s *Set[T]) Add(values ...T) {
	for _, v := range values {
		s.add(v)
	}
}

func (s *Set[T]) add(value T) {
	if s.seen == nil {
		s.seen = map[T]bool{value: true}
		s.values = append(s.values, value)

		return
	}

	if s.seen[value] {
		return
	}

	s.seen[value] = true

	s.values = append(s.values, value)
}
