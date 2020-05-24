package omap_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringAny_Put(t *testing.T) {
	Convey("TestMapStringAny.Put", t, func() {
		var k string = "hello"
		var v interface{} = "goodbye"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringAny_Delete(t *testing.T) {
	Convey("TestMapStringAny.Delete", t, func() {
		var k string = "hello"
		var v interface{} = "goodbye"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringAny_Has(t *testing.T) {
	Convey("TestMapStringAny.Has", t, func() {
		var k string = "hello"
		var v interface{} = "goodbye"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
