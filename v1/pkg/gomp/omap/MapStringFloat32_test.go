package omap_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringFloat32_Put(t *testing.T) {
	Convey("TestMapStringFloat32.Put", t, func() {
		var k string = "hello"
		var v float32 = 0.524

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringFloat32_Delete(t *testing.T) {
	Convey("TestMapStringFloat32.Delete", t, func() {
		var k string = "hello"
		var v float32 = 0.028

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringFloat32_Has(t *testing.T) {
	Convey("TestMapStringFloat32.Has", t, func() {
		var k string = "hello"
		var v float32 = 0.158

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
