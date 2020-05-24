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
		var k interface{} = "d9a982e8-a4f2-4cdc-a587-b91ef128e92a"
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
		var k interface{} = "0bae4533-63c4-42fc-b9af-ad4ebbf6fe19"
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
		var k interface{} = "d4f79145-cced-4e16-9099-f47665997a73"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("db663939-74ad-4466-ab20-d24283b95667"+"d6c2593f-0678-4477-a2d5-7fdf40cda280"), ShouldBeFalse)
	})
}

func TestMapAnyBool_Get(t *testing.T) {
	Convey("TestMapAnyBool.Get", t, func() {
		var k interface{} = "8d111fb1-0532-4ce1-8b20-2989f1256389"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("449cfea8-a4c6-4da5-a374-b3466b304af0" + "801f7fc0-7f70-4f9a-8e7f-74f5eadfa488")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyBool_GetOpt(t *testing.T) {
	Convey("TestMapAnyBool.GetOpt", t, func() {
		var k interface{} = "99e2e44d-dafa-4498-b1b8-7c9d624fc6b9"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("de686d2b-a89d-4375-9b64-d82f15f41122" + "86408dcb-1c0d-4f27-b0ef-14831a9d6a37")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyBool_ForEach(t *testing.T) {
	Convey("TestMapAnyBool.ForEach", t, func() {
		var k interface{} = "59768548-8020-405b-9ca5-fbf9a8d76eb4"
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
		var k interface{} = "7e4667d2-5593-4aaf-a7ba-5187514d1be5"
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
		var k interface{} = "ee48c45b-8bf9-40ef-95e7-893bdb4dc2ac"
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
		var k interface{} = "720b1782-fc6b-4454-9347-15e3ff62503c"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*bool)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("2134a05c-cde5-42f3-96af-4a75ecdc49e8", (*bool)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x bool = false
		So(test.PutIfNotNil("73fb131e-1b24-47e7-a091-c5dee44397d6", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyBool_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyBool.ReplaceIfExists", t, func() {
		var k interface{} = "a5e5b016-fdb4-4a78-ada4-39e1e003eaf2"
		var v bool = false
		var x bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("13e22ac7-76bf-4db2-a118-1d9caf3874d2", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyBool_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyBool.ReplaceOrPut", t, func() {
		var k interface{} = "8a255ed0-220c-4fdc-bc3a-c8c2e4aac032"
		var v bool = false
		var x bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("fe806099-910a-423f-87e9-57fba3b0b2ce", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyBool_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyBool.MarshalJSON", t, func() {
		var k interface{} = "6def4b6e-9bfe-443b-8bc9-8d9a54bc73a5"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"6def4b6e-9bfe-443b-8bc9-8d9a54bc73a5","value":false}]`)
	})
}
