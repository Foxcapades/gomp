package omap_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringComplex128_Put(t *testing.T) {
	Convey("TestMapStringComplex128.Put", t, func() {
		var k string = "hello"
		var v complex128 = 12

		test := omap.NewMapStringComplex128(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringComplex128_Delete(t *testing.T) {
	Convey("TestMapStringComplex128.Delete", t, func() {
		var k string = "hello"
		var v complex128 = 12

		test := omap.NewMapStringComplex128(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringComplex128_Has(t *testing.T) {
	Convey("TestMapStringComplex128.Has", t, func() {
		var k string = "hello"
		var v complex128 = 12

		test := omap.NewMapStringComplex128(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
