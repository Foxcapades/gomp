package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintInt_Put(t *testing.T) {
	Convey("TestMapUintInt.Put", t, func() {
		var k uint = 3596943298
		var v int = 1654747505

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintInt_Delete(t *testing.T) {
	Convey("TestMapUintInt.Delete", t, func() {
		var k uint = 3272436862
		var v int = 1145152472

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintInt_Has(t *testing.T) {
	Convey("TestMapUintInt.Has", t, func() {
		var k uint = 2197835931
		var v int = 1660285077

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(2506109358+551842997), ShouldBeFalse)
	})
}

func TestMapUintInt_Get(t *testing.T) {
	Convey("TestMapUintInt.Get", t, func() {
		var k uint = 2084956778
		var v int = 545471172

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1237469255 + 212807899)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintInt_GetOpt(t *testing.T) {
	Convey("TestMapUintInt.GetOpt", t, func() {
		var k uint = 2774125220
		var v int = 1146836436

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(2061670942 + 2955208301)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintInt_ForEach(t *testing.T) {
	Convey("TestMapUintInt.ForEach", t, func() {
		var k uint = 3804938813
		var v int = 546747089
		hits := 0

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv int) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintInt_MarshalYAML(t *testing.T) {
	Convey("TestMapUintInt.MarshalYAML", t, func() {
		var k uint = 1350837971
		var v int = 1547995531

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintInt_ToYAML(t *testing.T) {
	Convey("TestMapUintInt.ToYAML", t, func() {
		var k uint = 51734242
		var v int = 937457911

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintInt_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintInt.PutIfNotNil", t, func() {
		var k uint = 3754825742
		var v int = 726425721

		test := omap.NewMapUintInt(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(2837207180, (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int = 1980413993
		So(test.PutIfNotNil(1036370243, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintInt_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintInt.ReplaceIfExists", t, func() {
		var k uint = 928364108
		var v int = 1677519683
		var x int = 2021353402

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(4046175959, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintInt.ReplaceOrPut", t, func() {
		var k uint = 2480547610
		var v int = 1119372695
		var x int = 910610578

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1614211929, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt_MarshalJSON(t *testing.T) {
	Convey("TestMapUintInt.MarshalJSON", t, func() {
		var k uint = 1004247937
		var v int = 1106641614

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1004247937,"value":1106641614}]`)
	})
}
