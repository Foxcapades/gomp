package omap_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringBool_Put(t *testing.T) {
	Convey("TestMapStringBool.Put", t, func() {
		var k string = "hello"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringBool_Delete(t *testing.T) {
	Convey("TestMapStringBool.Delete", t, func() {
		var k string = "hello"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringBool_Has(t *testing.T) {
	Convey("TestMapStringBool.Has", t, func() {
		var k string = "hello"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
