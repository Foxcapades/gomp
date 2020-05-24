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
		var k interface{} = "378973ef-22c8-495b-9ab5-c6195c69ce0c"
		var v byte = 192

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyByte_Delete(t *testing.T) {
	Convey("TestMapAnyByte.Delete", t, func() {
		var k interface{} = "550f04de-53be-4dc0-9d7f-65eb14f98ff9"
		var v byte = 120

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyByte_Has(t *testing.T) {
	Convey("TestMapAnyByte.Has", t, func() {
		var k interface{} = "1f9e4b43-b8a2-45d4-b1ef-37eec795f4e6"
		var v byte = 150

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("5c102351-8ce4-4c4b-8020-19375f90cff3"+"ac4db53a-e2a2-4d72-930b-945f2a9efcda"), ShouldBeFalse)
	})
}


func TestMapAnyByte_Get(t *testing.T) {
	Convey("TestMapAnyByte.Get", t, func() {
		var k interface{} = "d9e946bd-519d-411c-a2af-025a0f4e88c9"
		var v byte = 237

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("fc6debaa-27e6-4edf-a969-933462260744"+"7c8086cc-94ee-4fac-8212-6064bc9e4a1b")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyByte_GetOpt(t *testing.T) {
	Convey("TestMapAnyByte.GetOpt", t, func() {
		var k interface{} = "4d5011d5-fbc1-45dd-b00c-104898564d39"
		var v byte = 14

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("c20bac87-b520-44b3-b362-d61307150975"+"7ad42589-f78f-4ba7-8a8a-8b1642357461")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyByte_ForEach(t *testing.T) {
	Convey("TestMapAnyByte.ForEach", t, func() {
		var k interface{} = "2ddc0ec3-7eee-454f-8287-92c998547961"
		var v byte = 49
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
		var k interface{} = "d9446323-a6cc-492b-938a-0695fb5b3ea8"
		var v byte = 179

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
		var k interface{} = "7f90a600-c683-4914-a84b-8a76659606fa"
		var v byte = 179

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyByte_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyByte.PutIfNotNil", t, func() {
		var k interface{} = "6c35032f-83fc-484e-95bd-699aeea34274"
		var v byte = 104

		test := omap.NewMapAnyByte(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("e717b78b-1d7e-4a88-97b2-57c616688fa2", (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x byte = 218
		So(test.PutIfNotNil("2f2d53f3-6285-4be0-87ad-f6887d62aa17", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyByte_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyByte.ReplaceIfExists", t, func() {
		var k interface{} = "6dd01418-0db0-458a-9dbb-5c6cf7e02eb9"
		var v byte = 207
		var x byte = 87

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("e0b1b09c-e2a2-40b1-b9c0-b79ecee75933", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyByte_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyByte.ReplaceOrPut", t, func() {
		var k interface{} = "59d48521-126d-4d02-a6af-8ed678cc3299"
		var v byte = 246
		var x byte = 167

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("3d9c2dfb-4894-4263-b04b-e76882bc6fc6", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyByte_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyByte.MarshalJSON", t, func() {
		var k interface{} = "a7b9a382-632b-47a4-99fc-722401bef3bf"
		var v byte = 253

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"a7b9a382-632b-47a4-99fc-722401bef3bf","value":253}]`)
	})
}

