package omap_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringInt8_Put(t *testing.T) {
	Convey("TestMapStringInt8.Put", t, func() {
		var k string = "hello"
		var v int8 = 36

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringInt8_Delete(t *testing.T) {
	Convey("TestMapStringInt8.Delete", t, func() {
		var k string = "hello"
		var v int8 = 44

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringInt8_Has(t *testing.T) {
	Convey("TestMapStringInt8.Has", t, func() {
		var k string = "hello"
		var v int8 = 21

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
