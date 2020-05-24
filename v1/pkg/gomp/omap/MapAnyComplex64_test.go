package omap_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyComplex64_Put(t *testing.T) {
	Convey("TestMapAnyComplex64.Put", t, func() {
		var k interface{} = "goodbye"
		var v complex64 = 12

		test := omap.NewMapAnyComplex64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyComplex64_Delete(t *testing.T) {
	Convey("TestMapAnyComplex64.Delete", t, func() {
		var k interface{} = "goodbye"
		var v complex64 = 12

		test := omap.NewMapAnyComplex64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyComplex64_Has(t *testing.T) {
	Convey("TestMapAnyComplex64.Has", t, func() {
		var k interface{} = "goodbye"
		var v complex64 = 12

		test := omap.NewMapAnyComplex64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
