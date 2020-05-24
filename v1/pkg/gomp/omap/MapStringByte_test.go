package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringByte_Put(t *testing.T) {
	Convey("TestMapStringByte.Put", t, func() {
		var k string = "2d738d70-d8f4-4b35-82f7-d655155d6e3d"
		var v byte = 99

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringByte_Delete(t *testing.T) {
	Convey("TestMapStringByte.Delete", t, func() {
		var k string = "842214b4-5d8d-4840-b10f-d6287fa47e91"
		var v byte = 200

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringByte_Has(t *testing.T) {
	Convey("TestMapStringByte.Has", t, func() {
		var k string = "55c521f3-c54d-4433-a07c-f13499c8c6ec"
		var v byte = 20

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("8f6e4a70-8fc0-49f3-9af1-e5e36f26b7aa"+"7fe1ffbc-3568-4750-837f-5b59e69bf046"), ShouldBeFalse)
	})
}


func TestMapStringByte_Get(t *testing.T) {
	Convey("TestMapStringByte.Get", t, func() {
		var k string = "51381fc2-0df8-4b52-aca8-d99143bee8e2"
		var v byte = 135

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("bba74ac6-41f8-429c-98e2-434904ebe688" + "66655370-8278-4be9-8026-d0546c0d5e45")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringByte_GetOpt(t *testing.T) {
	Convey("TestMapStringByte.GetOpt", t, func() {
		var k string = "0d74431c-fa47-40ee-9e28-4cd12cdb2c65"
		var v byte = 105

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("0b88e9aa-b77e-407f-9720-7e08cae2eb68" + "29005f20-bfbd-41dc-9084-245e13a157c1")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringByte_ForEach(t *testing.T) {
	Convey("TestMapStringByte.ForEach", t, func() {
		var k string = "6cde7c7e-756b-4107-8ece-264054692ce3"
		var v byte = 87
		hits := 0

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv byte) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringByte_MarshalYAML(t *testing.T) {
	Convey("TestMapStringByte.MarshalYAML", t, func() {
		var k string = "942611be-60bb-4362-b25e-0ebe58adae32"
		var v byte = 12

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringByte_ToYAML(t *testing.T) {
	Convey("TestMapStringByte.ToYAML", t, func() {
		var k string = "bdc7b8c2-cc8e-411d-9cd1-57cca5ac41e7"
		var v byte = 80

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringByte_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringByte.PutIfNotNil", t, func() {
		var k string = "cbba22c3-a4bc-4801-aca5-8aafa4c02402"
		var v byte = 168

		test := omap.NewMapStringByte(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("3a5d54ce-183b-4863-b5d3-bd57bff66e8f", (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x byte = 56
		So(test.PutIfNotNil("c6bd3fa0-531d-4c9f-b3fc-3bc130d3175f", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringByte_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringByte.ReplaceIfExists", t, func() {
		var k string = "cc5ac224-c149-422e-9190-b2ea681fecc4"
		var v byte = 226
		var x byte = 29

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("1607b3ff-3105-48a9-ba25-74f73fd9f06f", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringByte_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringByte.ReplaceOrPut", t, func() {
		var k string = "c268bb60-d8b5-4b23-971b-41a98afce492"
		var v byte = 254
		var x byte = 215

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("cbe5803c-0b45-4c0e-b783-fa18d9383e31", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringByte_MarshalJSON(t *testing.T) {
	Convey("TestMapStringByte.MarshalJSON", t, func() {
		var k string = "1d4b5222-516e-480e-83e8-0e93f6e4b9f0"
		var v byte = 160

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"1d4b5222-516e-480e-83e8-0e93f6e4b9f0","value":160}]`)
	})
}

