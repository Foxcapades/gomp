package omap_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintByte_Put(t *testing.T) {
	Convey("TestMapUintByte.Put", t, func() {
		var k uint = 184347049
		var v byte = 153

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintByte_Delete(t *testing.T) {
	Convey("TestMapUintByte.Delete", t, func() {
		var k uint = 1494043213
		var v byte = 68

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintByte_Has(t *testing.T) {
	Convey("TestMapUintByte.Has", t, func() {
		var k uint = 3607546704
		var v byte = 85

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
