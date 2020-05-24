package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntAny_Put(t *testing.T) {
	Convey("TestMapIntAny.Put", t, func() {
		var k int = 722934209
		var v interface{} = "99eafcc6-0855-445b-b642-d418916a1ed5"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntAny_Delete(t *testing.T) {
	Convey("TestMapIntAny.Delete", t, func() {
		var k int = 144449532
		var v interface{} = "0e6a48e5-3152-475c-9c23-1f553d7f5a4d"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntAny_Has(t *testing.T) {
	Convey("TestMapIntAny.Has", t, func() {
		var k int = 1997075662
		var v interface{} = "66956236-b9ba-48e4-a982-8cc7a53dc785"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1175422475+1214464122), ShouldBeFalse)
	})
}


func TestMapIntAny_Get(t *testing.T) {
	Convey("TestMapIntAny.Get", t, func() {
		var k int = 953636794
		var v interface{} = "79e284a3-02bf-46e3-83ce-70c7be64c64d"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1768421663 + 10112067)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntAny_GetOpt(t *testing.T) {
	Convey("TestMapIntAny.GetOpt", t, func() {
		var k int = 1569873779
		var v interface{} = "a7b238c4-3458-475b-b805-f5c623428030"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(615686008 + 969549194)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntAny_ForEach(t *testing.T) {
	Convey("TestMapIntAny.ForEach", t, func() {
		var k int = 787489468
		var v interface{} = "ff93b3a9-b073-400e-a429-3de19c45aa5e"
		hits := 0

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv interface{}) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntAny_MarshalYAML(t *testing.T) {
	Convey("TestMapIntAny.MarshalYAML", t, func() {
		var k int = 1443797912
		var v interface{} = "62aad9df-86d9-4aca-a21d-c71dac6f82d6"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntAny_ToYAML(t *testing.T) {
	Convey("TestMapIntAny.ToYAML", t, func() {
		var k int = 1910860730
		var v interface{} = "a1169da5-f370-4813-b114-3fd1ab2e3cf0"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntAny_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntAny.PutIfNotNil", t, func() {
		var k int = 639374089
		var v interface{} = "ca2d5714-4eff-4052-bf70-8bbc025adc33"

		test := omap.NewMapIntAny(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1159606022, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "30f68b40-1183-49fa-a63d-196c26de64ec"
		So(test.PutIfNotNil(384704042, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntAny_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntAny.ReplaceIfExists", t, func() {
		var k int = 685438202
		var v interface{} = "3e85b017-06e5-4e6e-936a-742fdf490425"
		var x interface{} = "3749fea9-d330-4b36-b2e5-8356ddbc15a9"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(139724871, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntAny_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntAny.ReplaceOrPut", t, func() {
		var k int = 1937494504
		var v interface{} = "859729a9-5332-4647-b19c-6a7dbef38f64"
		var x interface{} = "7d57ebfe-5d36-413f-8c35-60f9d69791db"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(211097457, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntAny_MarshalJSON(t *testing.T) {
	Convey("TestMapIntAny.MarshalJSON", t, func() {
		var k int = 1744583016
		var v interface{} = "c06ba8b9-8311-45e0-bee0-cc52824f485d"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1744583016,"value":"c06ba8b9-8311-45e0-bee0-cc52824f485d"}]`)
	})
}
