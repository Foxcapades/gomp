package omap_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintInt16_Put(t *testing.T) {
	Convey("TestMapUintInt16.Put", t, func() {
		var k uint = 1077015772
		var v int16 = 23081

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintInt16_Delete(t *testing.T) {
	Convey("TestMapUintInt16.Delete", t, func() {
		var k uint = 4182790799
		var v int16 = 26852

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintInt16_Has(t *testing.T) {
	Convey("TestMapUintInt16.Has", t, func() {
		var k uint = 92486895
		var v int16 = 5710

		test := omap.NewMapUintInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
