package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyUint32_Put(t *testing.T) {
	Convey("TestMapAnyUint32.Put", t, func() {
		var k interface{} = "f6896bbd-192a-4994-b15d-7d5e82f88a6b"
		var v uint32 = 1872033643

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyUint32_Delete(t *testing.T) {
	Convey("TestMapAnyUint32.Delete", t, func() {
		var k interface{} = "5d398040-d06b-49cd-9dc1-e5f2397c9022"
		var v uint32 = 224853009

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyUint32_Has(t *testing.T) {
	Convey("TestMapAnyUint32.Has", t, func() {
		var k interface{} = "6eb4634e-26d0-481a-8eff-7bd6cbd03f65"
		var v uint32 = 926057676

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("4f5c27ea-3471-426e-803f-6c4a7c6844c4"+"3c4c311f-788a-4b4c-8228-0887935cc826"), ShouldBeFalse)
	})
}


func TestMapAnyUint32_Get(t *testing.T) {
	Convey("TestMapAnyUint32.Get", t, func() {
		var k interface{} = "0a881580-c750-4af1-ba9f-82f997abc24c"
		var v uint32 = 1417086542

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("5864495b-cc5d-484c-9b66-700078c91f2d"+"6f40de8c-0238-4755-a77c-8724c6c76912")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyUint32_GetOpt(t *testing.T) {
	Convey("TestMapAnyUint32.GetOpt", t, func() {
		var k interface{} = "6bea0359-210a-4a2f-8a39-e16bc5525201"
		var v uint32 = 2595038371

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("9c6b6c94-bf1d-464b-a79c-56c84c83a83a"+"780c885e-eef4-4462-94a8-6e6f7497a7c9")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyUint32_ForEach(t *testing.T) {
	Convey("TestMapAnyUint32.ForEach", t, func() {
		var k interface{} = "c3513afb-523b-4964-b9ff-8e810cb554f1"
		var v uint32 = 473216263
		hits := 0

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv uint32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyUint32_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyUint32.MarshalYAML", t, func() {
		var k interface{} = "a9aa8f54-4476-4bbb-9ecd-0c271968cc64"
		var v uint32 = 3320033066

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyUint32_ToYAML(t *testing.T) {
	Convey("TestMapAnyUint32.ToYAML", t, func() {
		var k interface{} = "dc0ce659-620b-417c-82a4-2ec1572573b4"
		var v uint32 = 3569032170

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyUint32_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyUint32.PutIfNotNil", t, func() {
		var k interface{} = "c2136fb4-01c6-45c4-bc3c-81788789838a"
		var v uint32 = 1702054099

		test := omap.NewMapAnyUint32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("f602c682-dba5-4716-9462-96167e84b8f6", (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint32 = 1621057833
		So(test.PutIfNotNil("e91af3d7-4db9-4309-947d-3faf8d2e9859", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyUint32.ReplaceIfExists", t, func() {
		var k interface{} = "856c3aa1-f2a2-4844-8a1d-41ac2745127a"
		var v uint32 = 2879077055
		var x uint32 = 345511212

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("5a426719-11f7-478a-9010-89b4caa56c8a", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyUint32.ReplaceOrPut", t, func() {
		var k interface{} = "4d2958f7-3f60-4ce1-82cb-143e78df58fa"
		var v uint32 = 4126659624
		var x uint32 = 804298207

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("96e30bf0-4507-4011-a4ae-02764bf63cd4", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint32_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyUint32.MarshalJSON", t, func() {
		var k interface{} = "e399366f-2d99-43b3-9073-f3ac18ad6266"
		var v uint32 = 1495529044

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"e399366f-2d99-43b3-9073-f3ac18ad6266","value":1495529044}]`)
	})
}

