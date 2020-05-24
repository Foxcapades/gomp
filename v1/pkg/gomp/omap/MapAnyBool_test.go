package omap_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyBool_Put(t *testing.T) {
	Convey("TestMapAnyBool.Put", t, func() {
		var k interface{} = "goodbye"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyBool_Delete(t *testing.T) {
	Convey("TestMapAnyBool.Delete", t, func() {
		var k interface{} = "goodbye"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyBool_Has(t *testing.T) {
	Convey("TestMapAnyBool.Has", t, func() {
		var k interface{} = "goodbye"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
