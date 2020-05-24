package omap_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntString_Put(t *testing.T) {
	Convey("TestMapIntString.Put", t, func() {
		var k int = 739410870
		var v string = "hello"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntString_Delete(t *testing.T) {
	Convey("TestMapIntString.Delete", t, func() {
		var k int = 543312984
		var v string = "hello"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntString_Has(t *testing.T) {
	Convey("TestMapIntString.Has", t, func() {
		var k int = 464868247
		var v string = "hello"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
