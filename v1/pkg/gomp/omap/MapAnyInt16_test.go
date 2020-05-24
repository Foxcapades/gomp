package omap_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyInt16_Put(t *testing.T) {
	Convey("TestMapAnyInt16.Put", t, func() {
		var k interface{} = "goodbye"
		var v int16 = 4697

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyInt16_Delete(t *testing.T) {
	Convey("TestMapAnyInt16.Delete", t, func() {
		var k interface{} = "goodbye"
		var v int16 = 22251

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyInt16_Has(t *testing.T) {
	Convey("TestMapAnyInt16.Has", t, func() {
		var k interface{} = "goodbye"
		var v int16 = 32137

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
