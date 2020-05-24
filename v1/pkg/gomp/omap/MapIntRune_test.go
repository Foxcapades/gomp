package omap_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntRune_Put(t *testing.T) {
	Convey("TestMapIntRune.Put", t, func() {
		var k int = 27543430
		var v rune = 65889513

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntRune_Delete(t *testing.T) {
	Convey("TestMapIntRune.Delete", t, func() {
		var k int = 210519700
		var v rune = 792661359

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntRune_Has(t *testing.T) {
	Convey("TestMapIntRune.Has", t, func() {
		var k int = 1774796720
		var v rune = 746640783

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
