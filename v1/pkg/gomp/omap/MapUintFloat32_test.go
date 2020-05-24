package omap_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintFloat32_Put(t *testing.T) {
	Convey("TestMapUintFloat32.Put", t, func() {
		var k uint = 2788243018
		var v float32 = 0.985

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintFloat32_Delete(t *testing.T) {
	Convey("TestMapUintFloat32.Delete", t, func() {
		var k uint = 3585463438
		var v float32 = 0.332

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintFloat32_Has(t *testing.T) {
	Convey("TestMapUintFloat32.Has", t, func() {
		var k uint = 2840662080
		var v float32 = 0.956

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
