package omap_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyComplex128_Put(t *testing.T) {
	Convey("TestMapAnyComplex128.Put", t, func() {
		var k interface{} = "goodbye"
		var v complex128 = 12

		test := omap.NewMapAnyComplex128(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyComplex128_Delete(t *testing.T) {
	Convey("TestMapAnyComplex128.Delete", t, func() {
		var k interface{} = "goodbye"
		var v complex128 = 12

		test := omap.NewMapAnyComplex128(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyComplex128_Has(t *testing.T) {
	Convey("TestMapAnyComplex128.Has", t, func() {
		var k interface{} = "goodbye"
		var v complex128 = 12

		test := omap.NewMapAnyComplex128(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
