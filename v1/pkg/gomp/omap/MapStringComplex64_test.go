package omap_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringComplex64_Put(t *testing.T) {
	Convey("TestMapStringComplex64.Put", t, func() {
		var k string = "hello"
		var v complex64 = 12

		test := omap.NewMapStringComplex64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringComplex64_Delete(t *testing.T) {
	Convey("TestMapStringComplex64.Delete", t, func() {
		var k string = "hello"
		var v complex64 = 12

		test := omap.NewMapStringComplex64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringComplex64_Has(t *testing.T) {
	Convey("TestMapStringComplex64.Has", t, func() {
		var k string = "hello"
		var v complex64 = 12

		test := omap.NewMapStringComplex64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
