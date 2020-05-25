package util

import . "reflect"

// IsNil returns whether or not the given value is nil.
func IsNil(v interface{}) bool {
	if v == nil {
		return true
	}

	r := ValueOf(v)

	switch r.Kind() {
	case Ptr, Map, Slice, Array, Chan, Func:
		return r.IsNil()
	default:
		return false
	}
}

func Deref(v interface{}) interface{} {
	return ValueOf(v).Elem().Interface()
}
