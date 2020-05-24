package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntInt64_Put(t *testing.T) {
	Convey("TestMapIntInt64.Put", t, func() {
		var k int = 362586459
		var v int64 = 22014573323504945

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntInt64_Delete(t *testing.T) {
	Convey("TestMapIntInt64.Delete", t, func() {
		var k int = 1154972996
		var v int64 = 2842966509374737262

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntInt64_Has(t *testing.T) {
	Convey("TestMapIntInt64.Has", t, func() {
		var k int = 171724086
		var v int64 = 2266961426411858450

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(980607911+1025354947), ShouldBeFalse)
	})
}


func TestMapIntInt64_Get(t *testing.T) {
	Convey("TestMapIntInt64.Get", t, func() {
		var k int = 220285554
		var v int64 = 3439619164949677506

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(994221534 + 1432621208)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntInt64_GetOpt(t *testing.T) {
	Convey("TestMapIntInt64.GetOpt", t, func() {
		var k int = 1828305876
		var v int64 = 398408573531553945

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(2105808331 + 602228782)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntInt64_ForEach(t *testing.T) {
	Convey("TestMapIntInt64.ForEach", t, func() {
		var k int = 18562722
		var v int64 = 5386199069205414211
		hits := 0

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv int64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntInt64_MarshalYAML(t *testing.T) {
	Convey("TestMapIntInt64.MarshalYAML", t, func() {
		var k int = 1644582064
		var v int64 = 5954204502145351202

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntInt64_ToYAML(t *testing.T) {
	Convey("TestMapIntInt64.ToYAML", t, func() {
		var k int = 207400424
		var v int64 = 5104713625973888867

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntInt64_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntInt64.PutIfNotNil", t, func() {
		var k int = 56782891
		var v int64 = 4747170212777515733

		test := omap.NewMapIntInt64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(165169366, (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int64 = 8151068107523268593
		So(test.PutIfNotNil(643914263, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntInt64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntInt64.ReplaceIfExists", t, func() {
		var k int = 1597848027
		var v int64 = 8052373372414473872
		var x int64 = 6876942032325764812

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1992925523, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntInt64.ReplaceOrPut", t, func() {
		var k int = 28932175
		var v int64 = 1640113887179313487
		var x int64 = 857717721688495636

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1454989406, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt64_MarshalJSON(t *testing.T) {
	Convey("TestMapIntInt64.MarshalJSON", t, func() {
		var k int = 1760763254
		var v int64 = 859308227665263190

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1760763254,"value":859308227665263190}]`)
	})
}
