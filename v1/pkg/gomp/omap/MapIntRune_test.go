package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntRune_Put(t *testing.T) {
	Convey("TestMapIntRune.Put", t, func() {
		var k int = 1199553348
		var v rune = 2024191595

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntRune_Delete(t *testing.T) {
	Convey("TestMapIntRune.Delete", t, func() {
		var k int = 1108811375
		var v rune = 1208904054

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntRune_Has(t *testing.T) {
	Convey("TestMapIntRune.Has", t, func() {
		var k int = 2027932974
		var v rune = 1194511992

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1674061121+1472134625), ShouldBeFalse)
	})
}

func TestMapIntRune_Get(t *testing.T) {
	Convey("TestMapIntRune.Get", t, func() {
		var k int = 1747075241
		var v rune = 1943583037

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(366733951 + 1902239082)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntRune_GetOpt(t *testing.T) {
	Convey("TestMapIntRune.GetOpt", t, func() {
		var k int = 403958151
		var v rune = 885892190

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1869942249 + 1474000632)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntRune_ForEach(t *testing.T) {
	Convey("TestMapIntRune.ForEach", t, func() {
		var k int = 10158505
		var v rune = 837989684
		hits := 0

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv rune) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntRune_MarshalYAML(t *testing.T) {
	Convey("TestMapIntRune.MarshalYAML", t, func() {
		var k int = 1411095005
		var v rune = 6214365

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntRune_ToYAML(t *testing.T) {
	Convey("TestMapIntRune.ToYAML", t, func() {
		var k int = 837834552
		var v rune = 393353023

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntRune_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntRune.PutIfNotNil", t, func() {
		var k int = 980994076
		var v rune = 762431570

		test := omap.NewMapIntRune(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1962025074, (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x rune = 905965035
		So(test.PutIfNotNil(756354472, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntRune_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntRune.ReplaceIfExists", t, func() {
		var k int = 1639999830
		var v rune = 843054852
		var x rune = 604652666

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1527067325, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntRune_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntRune.ReplaceOrPut", t, func() {
		var k int = 238627627
		var v rune = 528827482
		var x rune = 1147530026

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1235911474, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntRune_MarshalJSON(t *testing.T) {
	Convey("TestMapIntRune.MarshalJSON", t, func() {
		var k int = 1571207557
		var v rune = 1625647650

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1571207557,"value":1625647650}]`)
	})
}
