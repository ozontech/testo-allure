package allurehash_test

import (
	"errors"
	"math"
	"reflect"
	"testing"

	"github.com/ozontech/testo-allure/internal/allurehash"
)

// Named functions for func testing
func funcA() {}

func TestHash(t *testing.T) {
	t.Parallel()

	// Pointer cycle setup
	type Cycle struct{ c *Cycle }
	var cycle1, cycle2 Cycle
	cycle1.c = &cycle1
	cycle2.c = &cycle2

	tests := []struct {
		name string
		v    any
		hash uint64
	}{
		// Basic types
		{"invalid", reflect.ValueOf(nil), 5334426602648574892},
		{"bool true", true, 784737815042604214},
		{"bool false", false, 10019070845011441513},
		{"int zero", 0, 13486526013545714072},
		{"int positive", 42, 8778961447134563034},
		{"int negative", -42, 6771251166012109539},
		{"uint zero", uint(0), 10304770384034042703},
		{"uint positive", uint(42), 11372966129978798985},
		{"float32 zero", float32(0), 12959678047180102526},
		{"float32 positive", float32(3.14), 8243353157485405451},
		{"float32 negative", float32(-3.14), 16342336385654720651},
		{"float32 NaN", float32(math.NaN()), 17675879962654930305},
		{"float32 Inf", float32(math.Inf(1)), 15695357267467037065},
		{"float64 zero", 0.0, 5504752136553225259},
		{"float64 positive", math.Pi, 12034510501260225033},
		{"float64 negative", -math.Pi, 1686749655719988617},
		{"float64 NaN", math.NaN(), 14311142514948334007},
		{"float64 Inf", math.Inf(1), 8574410181311368948},
		{"complex64 zero", complex64(0), 2505988549892606393},
		{"complex64 non-zero", complex64(1 + 2i), 8526642000527936238},
		{"complex128 zero", 0 + 0i, 8280107630189209024},
		{"complex128 non-zero", 1 + 2i, 17080548450716985839},
		{"string empty", "", 7743079449923625704},
		{"string non-empty", "hello", 4172708510790326704},

		// Pointers
		{"pointer nil", (*int)(nil), 3109597232765749772},
		{"pointer non-nil", new(int), 16384614761578734128},
		{"pointer cycle", cycle1, 5334426602648574892},

		// Structs
		{"struct empty", struct{}{}, 5334426602648574892},
		{"struct fields", struct{ A int }{42}, 1490305942443437939},
		{"struct nested", struct{ S struct{ A int } }{struct{ A int }{42}}, 11711704002059199536},

		// Slices
		{"slice nil", []int(nil), 8978264485201051520},
		{"slice empty", []int{}, 2198672732742028149},
		{"slice non-empty", []int{1, 2}, 1145491437765248442},

		// Arrays
		{"array empty", [0]int{}, 17769674858550263560},
		{"array non-empty", [2]int{1, 2}, 7405073706681676343},

		// Maps
		{"map nil", map[string]int(nil), 10348868668137447816},
		{"map empty", map[string]int{}, 6157603024903860861},
		{"map non-empty", map[string]int{"a": 1}, 7016608754986550801},
		{"map multiple keys", map[string]int{"a": 1, "b": 2}, 11698675498283957817},

		// Interfaces
		{"interface nil", error(nil), 3733343879303597645},
		{"interface non-nil", errors.New("error"), 10870157907507224536},
		{"interface nil with new", new(error), 8681769911282877815},

		// Functions
		{"func nil", (func())(nil), 13804434912164354801},
		{"func non-nil named", funcA, 10287895637114379171},
		{"func non-nil anonymous with return", func() int { return 0 }, 507643743327676908},

		// Channels
		{"int channel both", make(chan int), 2234621693301695195},
		{"int channel recv", make(<-chan int), 16380748977679313594},
		{"int channel send", make(chan<- int), 9482263965925405102},
		{"nil int channel send", (chan<- int)(nil), 15505706936240131312},
		{"nil int channel both", (chan int)(nil), 4078323793842941209},
		{"nil int channel recv", (<-chan int)(nil), 10414342112142925756},
	}

	for _, tt := range tests {
		t.Run("consistent: "+tt.name, func(t *testing.T) {
			t.Parallel()

			have := allurehash.Hash(tt.v)

			if have != tt.hash {
				t.Errorf("Hash(%v) (= %v) != %v", tt.v, have, tt.hash)
			}
		})
	}

	t.Run("ensure unique", func(t *testing.T) {
		t.Parallel()

		seen := make(map[uint64]string, len(tests))

		for _, tt := range tests {
			name, ok := seen[tt.hash]

			switch name {
			case "invalid", "pointer cycle", "struct empty":
				continue
			}

			if ok {
				t.Errorf("duplicate hash value for %q and %q: %v", name, tt.name, tt.hash)
			} else {
				seen[tt.hash] = tt.name
			}
		}
	})

	// Distinct value tests
	distinctTests := []struct {
		name string
		a    any
		b    any
	}{
		{"bool true vs false", true, false},
		{"int vs uint", 42, uint(42)},
		{"different ints", 42, 43},
		{"float32 vs float64", float32(3.14), 3.14},
		{"different floats", 3.14, 3.15},
		{"NaN vs Inf", math.NaN(), math.Inf(1)},
		{"different strings", "hello", "world"},
		{"pointer nil vs non-nil", (*int)(nil), new(int)},
		{"slice nil vs empty", []int(nil), []int{}},
		{"slice different lengths", []int{1}, []int{1, 2}},
		{"slice same length different values", []int{1, 2}, []int{2, 1}},
		{"array different", [2]int{1, 2}, [2]int{2, 1}},
		{"map nil vs empty", map[string]int(nil), map[string]int{}},
		{"map different keys", map[string]int{"a": 1}, map[string]int{"b": 1}},
		{"map same keys different values", map[string]int{"a": 1}, map[string]int{"a": 2}},
		{"interface nil vs non-nil", error(nil), errors.New("error")},
		{"func nil vs non-nil", (func())(nil), funcA},
		{"struct different fields", struct{ A int }{}, struct{ B int }{}},
		{"struct same field different values", struct{ A int }{1}, struct{ A int }{2}},
		{"channel vs other", make(chan int), 42},
	}

	for _, tt := range distinctTests {
		t.Run("distinct: "+tt.name, func(t *testing.T) {
			t.Parallel()

			if allurehash.Hash(tt.a) == allurehash.Hash(tt.b) {
				t.Errorf("Hash(%v) == Hash(%v)", tt.a, tt.b)
			}
		})
	}

	// Special cases
	t.Run("identical cycles have same hash", func(t *testing.T) {
		t.Parallel()

		if allurehash.Hash(cycle1) != allurehash.Hash(cycle2) {
			t.Error("Identical cycles should have same hash")
		}
	})
}
