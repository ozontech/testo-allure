package allurehash

import (
	"cmp"
	"encoding/binary"
	"hash/fnv"
	"io"
	"reflect"
	"slices"
)

// Hash returns a FNV-1a hash of the given value.
//
// Guarantees:
//   - Hashes are consistent between function calls and program execution.
//   - Hashes of two values are the same if values are equal based on their exported fields.
//   - Hashes of two functions are the same if functions share the same signature and name or its absence.
//   - Hashes of two different pointers to the same (equal) values are the same.
//
// Pointers are always dereferenced.
// Pointer cycles are detected and won't halt the execution.
func Hash(v any) uint64 {
	h := fnv.New64a()

	hash(h, reflect.ValueOf(v), make(map[uintptr]bool))

	return h.Sum64()
}

// NOTE(metafates): do not change hashing algorithm as it will break consistency guarantee.
//
//nolint:cyclop,funlen // huge switch
func hash(w io.Writer, val reflect.Value, visited map[uintptr]bool) {
	hashString(w, "<value>")
	defer hashString(w, "</value>")

	if !val.IsValid() {
		hashString(w, "invalid")

		return
	}

	hashString(w, "valid")
	hashType(w, val.Type(), make(map[reflect.Type]bool))

	// NOTE(metafates): when (if?) new types are added they should be added here.
	switch val.Kind() {
	case reflect.Bool:
		hashBool(w, val.Bool())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		hashInt(w, val.Int())

	case reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr:
		hashUint(w, val.Uint())

	case reflect.Float32, reflect.Float64:
		hashFloat(w, val.Float())

	case reflect.Complex64, reflect.Complex128:
		hashComplex(w, val.Complex())

	case reflect.String:
		hashString(w, val.String())

	case reflect.Pointer:
		hashPointer(w, val, visited)

	case reflect.Struct:
		hashStruct(w, val, visited)

	case reflect.Array:
		hashArray(w, val, visited)

	case reflect.Slice:
		hashSlice(w, val, visited)

	case reflect.Map:
		hashMap(w, val, visited)

	case reflect.Interface:
		hashInterface(w, val, visited)

	case reflect.Func:
		hashFunc(w, val)

	case reflect.Chan:
		hashChan(w, val)

	default:
		hashString(w, "unsupported")
	}
}

func hashType(w io.Writer, typ reflect.Type, visited map[reflect.Type]bool) {
	hashString(w, "<type>")
	defer hashString(w, "</type>")

	hashString(w, typ.Kind().String())

	if typ.Kind() == reflect.Pointer {
		hashType(w, typ.Elem(), visited)

		return
	}

	if visited[typ] {
		hashString(w, "cycle")

		return
	}

	visited[typ] = true

	switch typ.Kind() {
	case reflect.Array:
		hashInt(w, int64(typ.Len()))
		hashType(w, typ.Elem(), visited)

	case reflect.Slice:
		hashType(w, typ.Elem(), visited)

	case reflect.Chan:
		hashString(w, typ.ChanDir().String())
		hashType(w, typ.Elem(), visited)

	case reflect.Map:
		hashType(w, typ.Key(), visited)
		hashType(w, typ.Elem(), visited)

	case reflect.Func:
		hashFuncType(w, typ, visited)

	case reflect.Struct:
		hashStructType(w, typ, visited)
	}

	visited[typ] = false
}

func hashStructType(w io.Writer, typ reflect.Type, visited map[reflect.Type]bool) {
	for i := range typ.NumField() {
		field := typ.Field(i)

		if !field.IsExported() {
			continue
		}

		hashString(w, "<field>")

		hashBool(w, field.Anonymous)
		hashString(w, field.Name)
		hashString(w, string(field.Tag))
		hashType(w, field.Type, visited)

		hashString(w, "</field>")
	}
}

func hashFuncType(w io.Writer, typ reflect.Type, visited map[reflect.Type]bool) {
	for i := range typ.NumIn() {
		in := typ.In(i)

		hashString(w, "<in>")
		hashType(w, in, visited)
		hashString(w, "</in>")
	}

	for i := range typ.NumOut() {
		in := typ.Out(i)

		hashString(w, "<out>")
		hashType(w, in, visited)
		hashString(w, "</out>")
	}
}

func hashPointer(w io.Writer, val reflect.Value, visited map[uintptr]bool) {
	if val.IsNil() {
		hashString(w, "nil")

		return
	}

	ptr := val.Pointer()
	if visited[ptr] {
		hashString(w, "cycle")

		return
	}

	visited[ptr] = true
	hash(w, val.Elem(), visited)
}

func hashStruct(w io.Writer, val reflect.Value, visited map[uintptr]bool) {
	for i := range val.NumField() {
		field := val.Field(i)

		if field.CanInterface() {
			hash(w, field, visited)
		}
	}
}

func hashArray(w io.Writer, val reflect.Value, visited map[uintptr]bool) {
	for i := range val.Len() {
		hash(w, val.Index(i), visited)
	}
}

func hashInterface(w io.Writer, val reflect.Value, visited map[uintptr]bool) {
	if val.IsNil() {
		hashString(w, "nil")

		return
	}

	hash(w, val.Elem(), visited)
}

func hashSlice(w io.Writer, val reflect.Value, visited map[uintptr]bool) {
	if val.IsNil() {
		hashString(w, "nil")

		return
	}

	hashInt(w, int64(val.Len()))

	for i := range val.Len() {
		hash(w, val.Index(i), visited)
	}
}

func hashMap(w io.Writer, val reflect.Value, visited map[uintptr]bool) {
	if val.IsNil() {
		hashString(w, "nil")

		return
	}

	keys := val.MapKeys()

	slices.SortFunc(keys, func(a, b reflect.Value) int {
		return cmp.Compare(
			Hash(a.Interface()),
			Hash(b.Interface()),
		)
	})

	for _, k := range keys {
		hash(w, k, visited)
		hash(w, val.MapIndex(k), visited)
	}
}

func hashFunc(w io.Writer, val reflect.Value) {
	if val.IsNil() {
		hashString(w, "nil-func")

		return
	}

	pc := val.Pointer()
	if pc == 0 {
		hashString(w, "nil-func")

		return
	}

	hashString(w, "non-nil-func")
}

func hashChan(w io.Writer, val reflect.Value) {
	if val.IsNil() {
		hashString(w, "nil-chan")

		return
	}

	hashString(w, "non-nil-chan")
}

func hashBool(w io.Writer, val bool) {
	binaryWrite(w, val)
}

func hashInt(w io.Writer, val int64) {
	binaryWrite(w, val)
}

func hashUint(w io.Writer, val uint64) {
	binaryWrite(w, val)
}

func hashFloat(w io.Writer, val float64) {
	binaryWrite(w, val)
}

func hashComplex(w io.Writer, val complex128) {
	binaryWrite(w, val)
}

func hashString(w io.Writer, val string) {
	binaryWrite(w, []byte(val))
}

func binaryWrite(w io.Writer, data any) {
	_ = binary.Write(w, binary.BigEndian, data)
}
