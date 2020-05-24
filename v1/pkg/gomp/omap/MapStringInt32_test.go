package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringInt32_Put(t *testing.T) {
	Convey("TestMapStringInt32.Put", t, func() {
		var k string = "60a87222-491b-41a4-a3ff-ae98220bc355"
		var v int32 = 1071646288

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringInt32_Delete(t *testing.T) {
	Convey("TestMapStringInt32.Delete", t, func() {
		var k string = "11391a2d-2e24-420c-8fcb-078829cfd652"
		var v int32 = 262914611

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringInt32_Has(t *testing.T) {
	Convey("TestMapStringInt32.Has", t, func() {
		var k string = "b30adbb2-3f98-476d-a852-5066f2606238"
		var v int32 = 641540474

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("3682a383-2fb1-4f47-8eb8-5a3e74c01ae6"+"3bcea138-9bf0-45b2-966c-9b75a18bcd5f"), ShouldBeFalse)
	})
}

func TestMapStringInt32_Get(t *testing.T) {
	Convey("TestMapStringInt32.Get", t, func() {
		var k string = "d0bd6291-afde-42d7-ba4e-1c7267b7ff6f"
		var v int32 = 16294296

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("7d1ae39f-9563-4659-b95e-8fb80f9d99aa" + "d29ed1ab-48d7-4242-8984-fd145bf24ce0")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringInt32_GetOpt(t *testing.T) {
	Convey("TestMapStringInt32.GetOpt", t, func() {
		var k string = "4b9e7867-4284-4531-96ba-760ca34557f4"
		var v int32 = 731035209

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("4f0af72d-c70f-4256-93c8-4c83d9f958b7" + "226db64f-c28f-4c09-8db8-4eede626d3f8")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringInt32_ForEach(t *testing.T) {
	Convey("TestMapStringInt32.ForEach", t, func() {
		var k string = "293f9b63-166b-4a01-b777-facc353da812"
		var v int32 = 1865706530
		hits := 0

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv int32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringInt32_MarshalYAML(t *testing.T) {
	Convey("TestMapStringInt32.MarshalYAML", t, func() {
		var k string = "211bd724-444d-442b-bb7c-67a593f74816"
		var v int32 = 397585540

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringInt32_ToYAML(t *testing.T) {
	Convey("TestMapStringInt32.ToYAML", t, func() {
		var k string = "7b0b589f-3ad3-49d5-8f40-1a2040b57d47"
		var v int32 = 897605058

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringInt32_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringInt32.PutIfNotNil", t, func() {
		var k string = "a8036e9f-5785-4515-adea-efdd9e7e0ab3"
		var v int32 = 1526612659

		test := omap.NewMapStringInt32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("13453669-f7b9-4913-8725-33c91d19f593", (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int32 = 1324224966
		So(test.PutIfNotNil("35223d67-4c54-4469-b37c-030a7b9a9317", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringInt32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringInt32.ReplaceIfExists", t, func() {
		var k string = "c1071ca9-453c-448a-ae33-75a6e89fa327"
		var v int32 = 1901243824
		var x int32 = 46850805

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("b6e6ff73-84a3-4221-81f3-dc8169d8cc10", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringInt32.ReplaceOrPut", t, func() {
		var k string = "8a735c6b-b763-426d-a103-a2dabf27c563"
		var v int32 = 1481646500
		var x int32 = 902608294

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("c5767c5b-2f91-44c1-bdc9-736814177a0c", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt32_MarshalJSON(t *testing.T) {
	Convey("TestMapStringInt32.MarshalJSON", t, func() {
		var k string = "1cd80ba7-4f09-43b1-b70e-9137922ebde5"
		var v int32 = 53640243

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"1cd80ba7-4f09-43b1-b70e-9137922ebde5","value":53640243}]`)
	})
}
