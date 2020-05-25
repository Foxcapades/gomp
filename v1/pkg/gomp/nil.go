package gomp

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

// Deref attempts to dereference the given value if it is a pointer.
func Deref(v interface{}) interface{} {
	r := ValueOf(v)
	if r.Kind() == Ptr {
		return ValueOf(v).Elem().Interface()
	}

	return v
}
