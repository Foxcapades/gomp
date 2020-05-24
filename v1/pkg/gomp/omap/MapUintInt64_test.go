package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintInt64_Put(t *testing.T) {
	Convey("TestMapUintInt64.Put", t, func() {
		var k uint = 2452190144
		var v int64 = 5898699001791078219

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintInt64_Delete(t *testing.T) {
	Convey("TestMapUintInt64.Delete", t, func() {
		var k uint = 1141072121
		var v int64 = 3180817072411350690

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintInt64_Has(t *testing.T) {
	Convey("TestMapUintInt64.Has", t, func() {
		var k uint = 1690083265
		var v int64 = 7523561836348146193

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(2163054707+2981074016), ShouldBeFalse)
	})
}


func TestMapUintInt64_Get(t *testing.T) {
	Convey("TestMapUintInt64.Get", t, func() {
		var k uint = 912438780
		var v int64 = 313311088499424792

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(921253408 + 1762519563)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintInt64_GetOpt(t *testing.T) {
	Convey("TestMapUintInt64.GetOpt", t, func() {
		var k uint = 2825407454
		var v int64 = 4875495523581452822

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(4246699475 + 1468931082)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintInt64_ForEach(t *testing.T) {
	Convey("TestMapUintInt64.ForEach", t, func() {
		var k uint = 1198174559
		var v int64 = 3198536608070002483
		hits := 0

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv int64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintInt64_MarshalYAML(t *testing.T) {
	Convey("TestMapUintInt64.MarshalYAML", t, func() {
		var k uint = 2020039821
		var v int64 = 3547677592208307633

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintInt64_ToYAML(t *testing.T) {
	Convey("TestMapUintInt64.ToYAML", t, func() {
		var k uint = 2923592222
		var v int64 = 6563437368472663719

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintInt64_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintInt64.PutIfNotNil", t, func() {
		var k uint = 888413658
		var v int64 = 2628699087804777339

		test := omap.NewMapUintInt64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1545268897, (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int64 = 5844611085340946432
		So(test.PutIfNotNil(555536317, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintInt64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintInt64.ReplaceIfExists", t, func() {
		var k uint = 475959011
		var v int64 = 7665564909587715395
		var x int64 = 4362554825192439283

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(3471850850, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintInt64.ReplaceOrPut", t, func() {
		var k uint = 2619302511
		var v int64 = 6044767121497473793
		var x int64 = 8339321978303517467

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(458026199, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt64_MarshalJSON(t *testing.T) {
	Convey("TestMapUintInt64.MarshalJSON", t, func() {
		var k uint = 1005650692
		var v int64 = 8539427686284811488

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1005650692,"value":8539427686284811488}]`)
	})
}

