package omap_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntComplex64_Put(t *testing.T) {
	Convey("TestMapIntComplex64.Put", t, func() {
		var k int = 211277578
		var v complex64 = 12

		test := omap.NewMapIntComplex64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntComplex64_Delete(t *testing.T) {
	Convey("TestMapIntComplex64.Delete", t, func() {
		var k int = 1117508154
		var v complex64 = 12

		test := omap.NewMapIntComplex64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntComplex64_Has(t *testing.T) {
	Convey("TestMapIntComplex64.Has", t, func() {
		var k int = 214167822
		var v complex64 = 12

		test := omap.NewMapIntComplex64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
