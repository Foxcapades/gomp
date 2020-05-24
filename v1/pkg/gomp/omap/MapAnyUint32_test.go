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
		var k interface{} = "c79989a5-8981-4533-ad33-f0c3db65619d"
		var v uint32 = 889618560

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyUint32_Delete(t *testing.T) {
	Convey("TestMapAnyUint32.Delete", t, func() {
		var k interface{} = "528218b0-aabf-4e30-bf08-fc0d6327b199"
		var v uint32 = 2971904297

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyUint32_Has(t *testing.T) {
	Convey("TestMapAnyUint32.Has", t, func() {
		var k interface{} = "892a3c8d-4f00-451a-9a39-17b9f5366083"
		var v uint32 = 3214172112

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("e83b7e13-d0e1-4371-b112-74c7b674a2ee"+"76d2b3d8-5a54-42f6-8cbe-33b10b278d0f"), ShouldBeFalse)
	})
}


func TestMapAnyUint32_Get(t *testing.T) {
	Convey("TestMapAnyUint32.Get", t, func() {
		var k interface{} = "93e3ddcb-3b17-4056-84b7-de41d49945d7"
		var v uint32 = 3115205132

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("cfabf5ef-9262-4f23-9d26-48eef493c050" + "46d06a80-7d64-46d5-8aa1-8b79c55d8735")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyUint32_GetOpt(t *testing.T) {
	Convey("TestMapAnyUint32.GetOpt", t, func() {
		var k interface{} = "8d62e0bb-bac0-4a29-851d-09bc9bec60dc"
		var v uint32 = 1658090264

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("202a7032-5d2c-4e12-bf5c-c6bbc60ac000" + "ca671bf4-2bb5-4052-b482-bc0ccf50f7c8")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyUint32_ForEach(t *testing.T) {
	Convey("TestMapAnyUint32.ForEach", t, func() {
		var k interface{} = "963d2e0b-ed89-4d25-8b90-b56c8068fce2"
		var v uint32 = 4094293190
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
		var k interface{} = "44269116-fa3f-4f24-8e3d-4966c10d9b47"
		var v uint32 = 2169042811

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
		var k interface{} = "58e06964-0403-4034-a84b-7f3180ab952b"
		var v uint32 = 808640462

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
		var k interface{} = "c1110f44-5392-4e3e-8f3e-91cb58addc20"
		var v uint32 = 4246912579

		test := omap.NewMapAnyUint32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("d9ac0059-0952-4377-91a2-067472b00805", (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint32 = 1179512108
		So(test.PutIfNotNil("b4faaf50-c7dd-4f5d-831c-d25a32b58ba3", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyUint32.ReplaceIfExists", t, func() {
		var k interface{} = "97229c87-aad8-45d9-9e51-465eef7eb3af"
		var v uint32 = 1874143671
		var x uint32 = 1075997266

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("4b973385-e6f7-4b3f-bb16-9e90f0d0f055", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyUint32.ReplaceOrPut", t, func() {
		var k interface{} = "0b695b18-2409-4bac-a2e6-377da45281cb"
		var v uint32 = 3994546130
		var x uint32 = 3725236094

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("2032d521-7580-425e-bbaf-0d7c0587b97d", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint32_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyUint32.MarshalJSON", t, func() {
		var k interface{} = "87be020f-ca72-49b9-b8da-aa8c997d9d8a"
		var v uint32 = 877814125

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"87be020f-ca72-49b9-b8da-aa8c997d9d8a","value":877814125}]`)
	})
}

