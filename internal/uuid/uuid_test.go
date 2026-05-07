package uuid

import "testing"

func TestUnique(t *testing.T) {
	t.Parallel()

	const count = 100_000

	seen := make(map[UUID]bool, count)

	for range count {
		value := New()

		if seen[value] {
			t.Errorf("duplicate uuid generated: %s", value.String())
		}

		seen[value] = true
	}
}
