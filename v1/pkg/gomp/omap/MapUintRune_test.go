package omap_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintRune_Put(t *testing.T) {
	Convey("TestMapUintRune.Put", t, func() {
		var k uint = 534137245
		var v rune = 560870371

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintRune_Delete(t *testing.T) {
	Convey("TestMapUintRune.Delete", t, func() {
		var k uint = 3586072234
		var v rune = 676038151

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintRune_Has(t *testing.T) {
	Convey("TestMapUintRune.Has", t, func() {
		var k uint = 32990530
		var v rune = 1932198682

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(k+k), ShouldBeFalse)
	})
}
