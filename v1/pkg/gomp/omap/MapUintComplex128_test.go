package omap_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintComplex128_Put(t *testing.T) {
	Convey("TestMapUintComplex128.Put", t, func() {
		var k uint = 3847445738
		var v complex128 = 12

		test := omap.NewMapUintComplex128(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintComplex128_Delete(t *testing.T) {
	Convey("TestMapUintComplex128.Delete", t, func() {
		var k uint = 4115336141
		var v complex128 = 12

		test := omap.NewMapUintComplex128(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintComplex128_Has(t *testing.T) {
	Convey("TestMapUintComplex128.Has", t, func() {
		var k uint = 80372672
		var v complex128 = 12

		test := omap.NewMapUintComplex128(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
