package set

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	values := []int{
		1, 99, 842, -124, 99, 1, 10, 842,
	}

	want := []int{
		1, 99, 842, -124, 10,
	}

	var set Set[int]

	set.Add(values...)

	require.Equal(t, want, set.ClonedSlice())
}
