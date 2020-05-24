package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyInt32_Put(t *testing.T) {
	Convey("TestMapAnyInt32.Put", t, func() {
		var k interface{} = "7d552c37-4f86-42b4-b0bf-35f691158aab"
		var v int32 = 602739342

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyInt32_Delete(t *testing.T) {
	Convey("TestMapAnyInt32.Delete", t, func() {
		var k interface{} = "c87c1a14-5bf6-4260-bc38-0a92d7b1b3c3"
		var v int32 = 652306406

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyInt32_Has(t *testing.T) {
	Convey("TestMapAnyInt32.Has", t, func() {
		var k interface{} = "15393136-07b9-4533-8d97-076403c09fbe"
		var v int32 = 1862600504

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("a81ccb4d-37dc-4900-9730-e29d015c25a5"+"b91fa3d0-1f6c-4a91-a878-787ec77c5b49"), ShouldBeFalse)
	})
}


func TestMapAnyInt32_Get(t *testing.T) {
	Convey("TestMapAnyInt32.Get", t, func() {
		var k interface{} = "60f2bd26-75f3-416d-8388-4550c5dfa603"
		var v int32 = 1834152446

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("bada029f-dc5d-4493-9303-dcf4490df05b" + "a25e7fe1-7ade-4bfe-9a5c-b78cfc004b52")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyInt32_GetOpt(t *testing.T) {
	Convey("TestMapAnyInt32.GetOpt", t, func() {
		var k interface{} = "9a682a30-121a-42ad-8e85-36ebab456e20"
		var v int32 = 1860529691

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("52ab25d9-a491-4512-9546-d398f33e30ed" + "6b2043fe-1bcf-473f-af31-7cd9d7086de0")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyInt32_ForEach(t *testing.T) {
	Convey("TestMapAnyInt32.ForEach", t, func() {
		var k interface{} = "8ec14486-93e3-4080-b67b-afbac10d2ba9"
		var v int32 = 297299456
		hits := 0

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv int32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyInt32_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyInt32.MarshalYAML", t, func() {
		var k interface{} = "fa799611-988e-49b1-94b3-03c22e4f537e"
		var v int32 = 1137813766

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyInt32_ToYAML(t *testing.T) {
	Convey("TestMapAnyInt32.ToYAML", t, func() {
		var k interface{} = "cce6d319-00e2-4aeb-b484-3f446bdc9750"
		var v int32 = 206432068

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyInt32_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyInt32.PutIfNotNil", t, func() {
		var k interface{} = "930c5797-7246-4470-b1bf-c46f1377bba5"
		var v int32 = 1381765629

		test := omap.NewMapAnyInt32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("92492e43-b8e0-4450-b909-ce8a7beb5b05", (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int32 = 1676218231
		So(test.PutIfNotNil("7c918df3-f0cc-4090-abff-fb41613bdb0c", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyInt32.ReplaceIfExists", t, func() {
		var k interface{} = "26511aed-ecf8-474a-be0a-ced5bdb487c1"
		var v int32 = 1754634524
		var x int32 = 810404031

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("aa0be88e-4a5c-45dd-b773-b63c26efa4d8", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyInt32.ReplaceOrPut", t, func() {
		var k interface{} = "cdb8baa4-0f58-4379-bb1d-7d0df900d85d"
		var v int32 = 490678721
		var x int32 = 1742673478

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("18603af3-22d0-4a3c-b6fb-221702036ce0", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt32_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyInt32.MarshalJSON", t, func() {
		var k interface{} = "90e08290-d938-4953-a8c8-c98b388a5ed4"
		var v int32 = 1384420758

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"90e08290-d938-4953-a8c8-c98b388a5ed4","value":1384420758}]`)
	})
}
