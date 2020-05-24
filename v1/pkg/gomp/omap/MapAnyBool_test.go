package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyBool_Put(t *testing.T) {
	Convey("TestMapAnyBool.Put", t, func() {
		var k interface{} = "361d7d4d-d7d6-41a6-b3dd-45da2cad0580"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyBool_Delete(t *testing.T) {
	Convey("TestMapAnyBool.Delete", t, func() {
		var k interface{} = "d604259e-07b4-4eb8-b1c3-6b70875f73ab"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyBool_Has(t *testing.T) {
	Convey("TestMapAnyBool.Has", t, func() {
		var k interface{} = "385694b8-0b44-4b20-9b35-ac552a46664d"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("1e4da884-aeb9-462f-aeda-c332974dc38d"+"67217948-27d1-46b3-aa21-cbb01ea8bec0"), ShouldBeFalse)
	})
}


func TestMapAnyBool_Get(t *testing.T) {
	Convey("TestMapAnyBool.Get", t, func() {
		var k interface{} = "2de74335-45f4-462b-a0fc-95d752c71e4b"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("afe52a3c-8945-4ddf-b140-ee6bde23f4c8" + "b005cf58-644e-4336-8dba-57bd110bdd8d")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyBool_GetOpt(t *testing.T) {
	Convey("TestMapAnyBool.GetOpt", t, func() {
		var k interface{} = "c815d94d-25c9-4123-a261-499a0a8ad38d"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("3c13d675-08c5-4596-a5c1-9838c0b12f75" + "0bbb5ab5-2966-4a11-b1d3-973d41f752bc")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyBool_ForEach(t *testing.T) {
	Convey("TestMapAnyBool.ForEach", t, func() {
		var k interface{} = "2d6c31b9-da51-4b22-a9ef-9f57539034ef"
		var v bool = false
		hits := 0

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv bool) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyBool_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyBool.MarshalYAML", t, func() {
		var k interface{} = "56b3c84f-d546-4495-a89a-76dcb2c5cc1b"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyBool_ToYAML(t *testing.T) {
	Convey("TestMapAnyBool.ToYAML", t, func() {
		var k interface{} = "2bb10140-3e67-49eb-8d89-72a9dba996d8"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyBool_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyBool.PutIfNotNil", t, func() {
		var k interface{} = "3d611c51-d790-4318-bff3-77328680ae53"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*bool)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("4834efd1-f106-45aa-8e67-28583c8bc1a5", (*bool)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x bool = false
		So(test.PutIfNotNil("c23a15fc-7e8c-49e4-a750-3d3308341d0b", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyBool_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyBool.ReplaceIfExists", t, func() {
		var k interface{} = "69cf6ed6-7902-4a6a-8ed5-92e31b368de2"
		var v bool = false
		var x bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("903e56e5-d39b-4083-91da-57f0a0eab820", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyBool_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyBool.ReplaceOrPut", t, func() {
		var k interface{} = "c575e93e-d474-4c42-b444-39354ae0787c"
		var v bool = false
		var x bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("2631c716-ac30-4137-a796-650ab7304852", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyBool_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyBool.MarshalJSON", t, func() {
		var k interface{} = "f0a991f0-7a9b-44f6-88e9-b32d73938d25"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"f0a991f0-7a9b-44f6-88e9-b32d73938d25","value":false}]`)
	})
}
