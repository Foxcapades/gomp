package omap_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringUint8_Put(t *testing.T) {
	Convey("TestMapStringUint8.Put", t, func() {
		var k string = "hello"
		var v uint8 = 167

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringUint8_Delete(t *testing.T) {
	Convey("TestMapStringUint8.Delete", t, func() {
		var k string = "hello"
		var v uint8 = 131

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringUint8_Has(t *testing.T) {
	Convey("TestMapStringUint8.Has", t, func() {
		var k string = "hello"
		var v uint8 = 50

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
