package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringInt8_Put(t *testing.T) {
	Convey("TestMapStringInt8.Put", t, func() {
		var k string = "198e24cf-c9e3-4669-9c11-0c364cdc0a3f"
		var v int8 = 22

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringInt8_Delete(t *testing.T) {
	Convey("TestMapStringInt8.Delete", t, func() {
		var k string = "7b55f870-38b8-4139-a107-c37002dbe269"
		var v int8 = 107

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringInt8_Has(t *testing.T) {
	Convey("TestMapStringInt8.Has", t, func() {
		var k string = "b86c2ed7-e3ce-401c-98aa-f4733ccf6ac4"
		var v int8 = 21

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("bb08a91b-096c-4d99-9f9c-443f86aef821"+"33a67e2a-6042-41bb-9e51-eb53614a4b28"), ShouldBeFalse)
	})
}

func TestMapStringInt8_Get(t *testing.T) {
	Convey("TestMapStringInt8.Get", t, func() {
		var k string = "e72d88f1-413e-41e9-ac92-9c005fc628d8"
		var v int8 = 120

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("4cc2fcb2-3785-40f3-86ab-5572404c86fb" + "96c27753-7789-4eae-a93b-123ce84b03a8")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringInt8_GetOpt(t *testing.T) {
	Convey("TestMapStringInt8.GetOpt", t, func() {
		var k string = "48df1b5e-f8cc-4fa5-b58b-e3115b445c4f"
		var v int8 = 27

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("e11857c6-cd81-4a49-8836-3df3b791ced9" + "a17bedc5-844b-4f31-b4a4-fdf1dc4a8c57")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringInt8_ForEach(t *testing.T) {
	Convey("TestMapStringInt8.ForEach", t, func() {
		var k string = "865947d8-cbde-4c62-bfec-e5aa9a649777"
		var v int8 = 78
		hits := 0

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv int8) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringInt8_MarshalYAML(t *testing.T) {
	Convey("TestMapStringInt8.MarshalYAML", t, func() {
		var k string = "49ba8ae9-dda0-4902-9ce3-fdd0f35558d7"
		var v int8 = 16

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringInt8_ToYAML(t *testing.T) {
	Convey("TestMapStringInt8.ToYAML", t, func() {
		var k string = "122e7366-548d-466f-b98d-47873ad2cb4c"
		var v int8 = 126

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringInt8_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringInt8.PutIfNotNil", t, func() {
		var k string = "fe99031a-4d97-4e82-9f8b-e5d684e10f85"
		var v int8 = 95

		test := omap.NewMapStringInt8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("9a6599d8-f7ea-4c2b-8444-5e4183e97ffd", (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int8 = 114
		So(test.PutIfNotNil("0de494dd-23a4-4f06-b013-ad6f8fc1b65f", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringInt8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringInt8.ReplaceIfExists", t, func() {
		var k string = "ad7e6239-5590-409e-bf73-42c286a6e45a"
		var v int8 = 39
		var x int8 = 18

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("cc753683-82c5-404e-a913-f87190bfe18a", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringInt8.ReplaceOrPut", t, func() {
		var k string = "7b46e754-887f-45ee-a2e5-7acb748b8080"
		var v int8 = 80
		var x int8 = 50

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("77cf1748-c83b-47f2-877f-352c9006a499", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt8_MarshalJSON(t *testing.T) {
	Convey("TestMapStringInt8.MarshalJSON", t, func() {
		var k string = "625f3eca-3e91-455d-9a82-e4aee7e95faf"
		var v int8 = 125

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"625f3eca-3e91-455d-9a82-e4aee7e95faf","value":125}]`)
	})
}
