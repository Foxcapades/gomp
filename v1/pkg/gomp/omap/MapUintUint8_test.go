package omap_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintUint8_Put(t *testing.T) {
	Convey("TestMapUintUint8.Put", t, func() {
		var k uint = 1015018773
		var v uint8 = 32

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintUint8_Delete(t *testing.T) {
	Convey("TestMapUintUint8.Delete", t, func() {
		var k uint = 153565041
		var v uint8 = 35

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintUint8_Has(t *testing.T) {
	Convey("TestMapUintUint8.Has", t, func() {
		var k uint = 2687947845
		var v uint8 = 189

		test := omap.NewMapUintUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
