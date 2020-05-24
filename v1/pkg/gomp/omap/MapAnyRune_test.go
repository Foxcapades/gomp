package omap_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyRune_Put(t *testing.T) {
	Convey("TestMapAnyRune.Put", t, func() {
		var k interface{} = "goodbye"
		var v rune = 1384138643

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyRune_Delete(t *testing.T) {
	Convey("TestMapAnyRune.Delete", t, func() {
		var k interface{} = "goodbye"
		var v rune = 183653891

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyRune_Has(t *testing.T) {
	Convey("TestMapAnyRune.Has", t, func() {
		var k interface{} = "goodbye"
		var v rune = 1437902002

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
