package omap_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintComplex64_Put(t *testing.T) {
	Convey("TestMapUintComplex64.Put", t, func() {
		var k uint = 1792400843
		var v complex64 = 12

		test := omap.NewMapUintComplex64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintComplex64_Delete(t *testing.T) {
	Convey("TestMapUintComplex64.Delete", t, func() {
		var k uint = 3086064781
		var v complex64 = 12

		test := omap.NewMapUintComplex64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintComplex64_Has(t *testing.T) {
	Convey("TestMapUintComplex64.Has", t, func() {
		var k uint = 1746921148
		var v complex64 = 12

		test := omap.NewMapUintComplex64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
