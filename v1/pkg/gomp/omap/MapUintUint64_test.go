package omap_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintUint64_Put(t *testing.T) {
	Convey("TestMapUintUint64.Put", t, func() {
		var k uint = 2714434079
		var v uint64 = 894060311800635659

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintUint64_Delete(t *testing.T) {
	Convey("TestMapUintUint64.Delete", t, func() {
		var k uint = 63683067
		var v uint64 = 5384925032239196514

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintUint64_Has(t *testing.T) {
	Convey("TestMapUintUint64.Has", t, func() {
		var k uint = 295305609
		var v uint64 = 9207450753580197605

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
