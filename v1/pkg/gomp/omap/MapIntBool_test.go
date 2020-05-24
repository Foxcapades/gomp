package omap_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntBool_Put(t *testing.T) {
	Convey("TestMapIntBool.Put", t, func() {
		var k int = 342847743
		var v bool = false

		test := omap.NewMapIntBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntBool_Delete(t *testing.T) {
	Convey("TestMapIntBool.Delete", t, func() {
		var k int = 692801166
		var v bool = false

		test := omap.NewMapIntBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntBool_Has(t *testing.T) {
	Convey("TestMapIntBool.Has", t, func() {
		var k int = 1225894535
		var v bool = false

		test := omap.NewMapIntBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
