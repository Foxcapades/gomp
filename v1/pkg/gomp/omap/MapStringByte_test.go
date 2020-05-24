package omap_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringByte_Put(t *testing.T) {
	Convey("TestMapStringByte.Put", t, func() {
		var k string = "hello"
		var v byte = 156

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringByte_Delete(t *testing.T) {
	Convey("TestMapStringByte.Delete", t, func() {
		var k string = "hello"
		var v byte = 103

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringByte_Has(t *testing.T) {
	Convey("TestMapStringByte.Has", t, func() {
		var k string = "hello"
		var v byte = 220

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
