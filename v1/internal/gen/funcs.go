package gen

import (
	"math"
	"math/rand"
	"strconv"

	"github.com/google/uuid"
)

// Pad returns a padding string of spaces sized to fill the difference in length
// between string a and string b.
func Pad(a, b string) string {
	cl := 0

	if len(a) < len(b) {
		cl = len(b) - len(a)
	}

	out := make([]byte, cl)

	for i := range out {
		out[i] = ' '
	}

	return string(out)
}

// IsBaseType returns whether or not the given type string is one of Go's built
// in types.
func IsBaseType(k string) bool {
	return k == "bool" ||
		k == "int" ||
		k == "int8" ||
		k == "int16" ||
		k == "int32" ||
		k == "int64" ||
		k == "uint" ||
		k == "uint8" ||
		k == "uint16" ||
		k == "uint32" ||
		k == "uint64" ||
		k == "float32" ||
		k == "float64" ||
		k == "complex64" ||
		k == "complex128" ||
		k == "string" ||
		k == "byte" ||
		k == "rune" ||
		k == "interface{}"
}

func DefaultValue(kind string) string {
	switch kind {
	case "bool":
		return strconv.FormatBool(rand.Intn(1) == 1)
	case "int", "int32", "rune":
		return strconv.FormatInt(int64(rand.Int31n(math.MaxInt32)), 10)
	case "int8":
		return strconv.FormatInt(int64(rand.Intn(math.MaxInt8)), 10)
	case "int16":
		return strconv.FormatInt(int64(rand.Intn(math.MaxInt16)), 10)
	case "int64":
		return strconv.FormatInt(rand.Int63n(math.MaxInt64), 10)
	case "uint", "uint32", "complex64", "complex128":
		return strconv.FormatUint(uint64(rand.Uint32()), 10)
	case "uint8", "byte":
		return strconv.FormatUint(uint64(rand.Intn(math.MaxUint8)), 10)
	case "uint16":
		return strconv.FormatUint(uint64(rand.Intn(math.MaxUint16)), 10)
	case "uint64":
		return strconv.FormatUint(rand.Uint64(), 10)
	case "float32":
		return strconv.FormatFloat(float64(rand.Float32()), 'f', 3, 32)
	case "float64":
		return strconv.FormatFloat(rand.Float64(), 'f', 3, 64)
	case "string":
		return `"` + uuid.New().String() + `"`
	case "interface{}":
		return `"` + uuid.New().String() + `"`
	}
	panic("wat")
}
