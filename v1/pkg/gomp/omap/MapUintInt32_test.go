package omap_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintInt32_Put(t *testing.T) {
	Convey("TestMapUintInt32.Put", t, func() {
		var k uint = 399303777
		var v int32 = 1386916611

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintInt32_Delete(t *testing.T) {
	Convey("TestMapUintInt32.Delete", t, func() {
		var k uint = 1339538206
		var v int32 = 963069888

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintInt32_Has(t *testing.T) {
	Convey("TestMapUintInt32.Has", t, func() {
		var k uint = 2092676637
		var v int32 = 177123756

		test := omap.NewMapUintInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
