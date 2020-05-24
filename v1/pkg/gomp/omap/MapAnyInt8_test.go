package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyInt8_Put(t *testing.T) {
	Convey("TestMapAnyInt8.Put", t, func() {
		var k interface{} = "ebd6156c-8f02-4854-b915-49bb6b966a9b"
		var v int8 = 11

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyInt8_Delete(t *testing.T) {
	Convey("TestMapAnyInt8.Delete", t, func() {
		var k interface{} = "6304f352-82c2-4779-8dbd-38a47d304f94"
		var v int8 = 20

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyInt8_Has(t *testing.T) {
	Convey("TestMapAnyInt8.Has", t, func() {
		var k interface{} = "b34bcaca-385e-4c91-b351-e470f8785a59"
		var v int8 = 76

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("e0e4cd33-30d9-461f-9a51-71ab0dd62e50"+"d56df754-964c-49a9-8d4b-f5ca511cdc40"), ShouldBeFalse)
	})
}


func TestMapAnyInt8_Get(t *testing.T) {
	Convey("TestMapAnyInt8.Get", t, func() {
		var k interface{} = "ad1bce79-1075-4a90-8aa4-199d69551165"
		var v int8 = 112

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("cb205e8c-3deb-4cb5-ab88-ccb533f3c9ae" + "2da0dda6-5124-4234-beeb-13d0403ea3f8")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyInt8_GetOpt(t *testing.T) {
	Convey("TestMapAnyInt8.GetOpt", t, func() {
		var k interface{} = "130841a1-b7f3-4496-91ad-6eed70398dc4"
		var v int8 = 56

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("24d7bbec-b773-4e2e-af4e-213068d3d102" + "20f22547-3253-4eb3-b9aa-1812c16eaf23")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyInt8_ForEach(t *testing.T) {
	Convey("TestMapAnyInt8.ForEach", t, func() {
		var k interface{} = "53f34c6b-371c-4f1c-b97c-920895d409a1"
		var v int8 = 106
		hits := 0

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv int8) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyInt8_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyInt8.MarshalYAML", t, func() {
		var k interface{} = "5eeebb64-c7d4-4ff8-92cc-5fad0c4af1af"
		var v int8 = 54

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyInt8_ToYAML(t *testing.T) {
	Convey("TestMapAnyInt8.ToYAML", t, func() {
		var k interface{} = "a5bb934d-9ea2-46d1-bac8-185358978461"
		var v int8 = 91

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyInt8_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyInt8.PutIfNotNil", t, func() {
		var k interface{} = "be466d51-9da0-434b-a2e5-07a23bbd3cf0"
		var v int8 = 30

		test := omap.NewMapAnyInt8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("e9276d5f-91f0-462f-ba21-d84f07ca4c26", (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int8 = 108
		So(test.PutIfNotNil("e6b3408d-8ca8-485b-88e7-578bc75d14aa", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyInt8.ReplaceIfExists", t, func() {
		var k interface{} = "396ef66b-1ec2-4b96-b8aa-256e85b2627b"
		var v int8 = 74
		var x int8 = 28

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("3ccfb9f4-6c79-41f1-bb2e-910ea59fdfd2", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyInt8.ReplaceOrPut", t, func() {
		var k interface{} = "1cdf21ff-2bfe-4c13-9e80-91891ea8f67b"
		var v int8 = 57
		var x int8 = 100

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("d9ca6056-f24e-4142-b2ec-90cdbf42273a", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt8_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyInt8.MarshalJSON", t, func() {
		var k interface{} = "5ea2e6cf-e4eb-4bf4-b1d2-e667f33dd9b6"
		var v int8 = 82

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"5ea2e6cf-e4eb-4bf4-b1d2-e667f33dd9b6","value":82}]`)
	})
}

