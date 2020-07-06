package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyByte_Put(t *testing.T) {
	Convey("TestMapAnyByte.Put", t, func() {
		var k interface{} = "695aa5a0-4010-4a1c-891f-35e6dd63b475"
		var v byte = 28

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyByte_Delete(t *testing.T) {
	Convey("TestMapAnyByte.Delete", t, func() {
		var k interface{} = "ee732d80-1056-4b7e-ae26-cf675a62ef0e"
		var v byte = 47

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyByte_Has(t *testing.T) {
	Convey("TestMapAnyByte.Has", t, func() {
		var k interface{} = "43e418e3-71e5-470a-9d7c-e7f9918010e5"
		var v byte = 130

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("a614c4ff-788a-44c6-844d-20e42ddcfa02"+"91e1d560-3717-4724-b60c-e25c3b001083"), ShouldBeFalse)
	})
}

func TestMapAnyByte_Get(t *testing.T) {
	Convey("TestMapAnyByte.Get", t, func() {
		var k interface{} = "16e87093-893c-40cf-af3c-8da97692713e"
		var v byte = 66

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("818f4e73-b5e4-4440-ad6d-b5962c93b255" + "007ff617-5ef2-4ddc-af99-62077d8ae41c")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyByte_GetOpt(t *testing.T) {
	Convey("TestMapAnyByte.GetOpt", t, func() {
		var k interface{} = "fbf234bc-4072-4cab-8c4b-6f8b8dda8f69"
		var v byte = 154

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("5d97568d-29f3-48ee-826e-832632d52bbf" + "ca85132d-b49e-4b7d-af59-f8dd0f80339d")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyByte_ForEach(t *testing.T) {
	Convey("TestMapAnyByte.ForEach", t, func() {
		var k interface{} = "2dce373a-909f-4133-b4e8-f95ed573685e"
		var v byte = 2
		hits := 0

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv byte) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyByte_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyByte.MarshalYAML", t, func() {
		var k interface{} = "34316d4f-49c8-4c9d-95a6-a4011f722574"
		var v byte = 160

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyByte_ToYAML(t *testing.T) {
	Convey("TestMapAnyByte.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "80ef05f6-626a-4178-b61e-09c8006c027f"
			var v byte = 238

			test := omap.NewMapAnyByte(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			c, d := test.ToYAML()
			So(d, ShouldBeNil)
			So(c.Kind, ShouldEqual, yaml.SequenceNode)
			So(c.LongTag(), ShouldEqual, xyml.TagOrderedMap)
			So(len(c.Content), ShouldEqual, 1)
			So(xyml.IsMap(c.Content[0]), ShouldBeTrue)
		})

		Convey("Unordered", func() {
			var k interface{} = "9096f936-d3a5-4d62-814c-860d55bc7fae"
			var v byte = 62

			test := omap.NewMapAnyByte(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			c, d := test.ToYAML()

			So(d, ShouldBeNil)
			So(c.Kind, ShouldEqual, yaml.MappingNode)
			So(c.LongTag(), ShouldEqual, xyml.TagMap)
			So(len(c.Content), ShouldEqual, 2)
		})
	})
}

func TestMapAnyByte_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyByte.PutIfNotNil", t, func() {
		var k interface{} = "87d19c85-a569-4d21-a598-772e0bf73fea"
		var v byte = 237

		test := omap.NewMapAnyByte(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("9bf25b90-b224-4d40-9045-748282887a58", (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x byte = 176
		So(test.PutIfNotNil("22ed99d4-fe45-4186-876c-dc29c66c27b4", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyByte_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyByte.ReplaceIfExists", t, func() {
		var k interface{} = "f3afabd2-5557-4f69-95bc-e531e0d4c79b"
		var v byte = 253
		var x byte = 104

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("42dd6396-351a-4649-898d-6afe92cc60d2", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyByte_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyByte.ReplaceOrPut", t, func() {
		var k interface{} = "864b8455-0152-4b84-b053-dfe399014aff"
		var v byte = 162
		var x byte = 113

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("c8c37808-66e5-4133-9006-0a67ac860ed7", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyByte_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyByte.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "cd9fad09-1b70-4bac-9ba9-fd91b0eb055b"
			var v byte = 253

			test := omap.NewMapAnyByte(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"cd9fad09-1b70-4bac-9ba9-fd91b0eb055b","value":253}]`)
		})

		Convey("Unordered", func() {
			var k interface{} = "cd9fad09-1b70-4bac-9ba9-fd91b0eb055b"
			var v byte = 253

			test := omap.NewMapAnyByte(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"cd9fad09-1b70-4bac-9ba9-fd91b0eb055b":253}`)
		})

	})
}
