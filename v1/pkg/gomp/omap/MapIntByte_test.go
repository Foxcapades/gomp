package omap_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntByte_Put(t *testing.T) {
	Convey("TestMapIntByte.Put", t, func() {
		var k int = 1469254904
		var v byte = 7

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntByte_Delete(t *testing.T) {
	Convey("TestMapIntByte.Delete", t, func() {
		var k int = 1126354657
		var v byte = 225

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntByte_Has(t *testing.T) {
	Convey("TestMapIntByte.Has", t, func() {
		var k int = 1538389371
		var v byte = 29

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
