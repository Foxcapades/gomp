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
		var k string = "88b79c07-1bcb-4d54-9a05-e7f4e4579b32"
		var v byte = 201

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringByte_Delete(t *testing.T) {
	Convey("TestMapStringByte.Delete", t, func() {
		var k string = "1fcb1794-4d65-4fb9-a38c-ff3465e53bc7"
		var v byte = 129

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringByte_Has(t *testing.T) {
	Convey("TestMapStringByte.Has", t, func() {
		var k string = "88297e81-db25-4f3c-8d51-82f4191097b9"
		var v byte = 42

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("e9a1cd0e-a535-4aac-9d29-17e2d056d380"+"670667b6-57f8-48d3-92d1-72cff02ae87b"), ShouldBeFalse)
	})
}

func TestMapStringByte_Get(t *testing.T) {
	Convey("TestMapStringByte.Get", t, func() {
		var k string = "81f25016-32bb-43be-a0c2-b7e0305316dc"
		var v byte = 170

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("2ceaaede-ec17-4a77-ae74-2aa3beed7a99" + "8aacffba-9a50-4058-94ff-9f381ac9d695")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringByte_GetOpt(t *testing.T) {
	Convey("TestMapStringByte.GetOpt", t, func() {
		var k string = "6cff77f2-cafc-4cf3-aaf3-7c9bf1ca7ad4"
		var v byte = 242

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("0ee30aa1-6951-40ed-b4c4-6ae8e163bb6a" + "e02aed48-e03c-4b8a-8fe4-80df00141726")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringByte_ForEach(t *testing.T) {
	Convey("TestMapStringByte.ForEach", t, func() {
		var k string = "33bd34f9-3de0-4485-9fd0-4d9198f5a264"
		var v byte = 128
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
		var k string = "5e7aca88-55c5-4471-8c78-27d5d3d1ad0f"
		var v byte = 72

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
		var k string = "ad16eec0-6be1-4062-86dc-53ff88080405"
		var v byte = 126

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
		var k string = "b480b965-43d7-4fb2-9e73-48cf564080f4"
		var v byte = 129

		test := omap.NewMapStringByte(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("fbe03c93-a4de-49a7-9bbb-32dd3aa66b37", (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x byte = 226
		So(test.PutIfNotNil("c7bc095c-fd6e-4503-ae73-192cff8b9c1d", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringByte_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringByte.ReplaceIfExists", t, func() {
		var k string = "6c9903ee-3105-4967-bce2-cc7e12947a9c"
		var v byte = 245
		var x byte = 13

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("15762607-8907-40d5-84a4-d7776af3ceb2", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringByte_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringByte.ReplaceOrPut", t, func() {
		var k string = "aee060ca-27cf-4fe1-bb86-74b7ddc82c4d"
		var v byte = 145
		var x byte = 94

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("95998483-3da0-49fb-b69b-d3865c9fd861", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringByte_MarshalJSON(t *testing.T) {
	Convey("TestMapStringByte.MarshalJSON", t, func() {
		var k string = "30171186-5de9-4c59-924f-d3e16f719695"
		var v byte = 221

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"30171186-5de9-4c59-924f-d3e16f719695","value":221}]`)
	})
}
