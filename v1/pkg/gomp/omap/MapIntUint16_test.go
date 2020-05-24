package omap_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntUint16_Put(t *testing.T) {
	Convey("TestMapIntUint16.Put", t, func() {
		var k int = 1157944885
		var v uint16 = 26225

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntUint16_Delete(t *testing.T) {
	Convey("TestMapIntUint16.Delete", t, func() {
		var k int = 1612251387
		var v uint16 = 9559

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntUint16_Has(t *testing.T) {
	Convey("TestMapIntUint16.Has", t, func() {
		var k int = 1617401528
		var v uint16 = 56708

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
