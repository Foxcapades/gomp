package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyUint64_Put(t *testing.T) {
	Convey("TestMapAnyUint64.Put", t, func() {
		var k interface{} = "b12f166b-b436-43c6-a13d-d4ae52b262ff"
		var v uint64 = 3963315077986104770

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyUint64_Delete(t *testing.T) {
	Convey("TestMapAnyUint64.Delete", t, func() {
		var k interface{} = "eee58742-8c98-4794-9d88-a655b6f15496"
		var v uint64 = 7922807834861930642

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyUint64_Has(t *testing.T) {
	Convey("TestMapAnyUint64.Has", t, func() {
		var k interface{} = "5cbab265-7b1a-4931-bace-3ca44fe2fcc1"
		var v uint64 = 13393242131644069724

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("544cc236-48d8-4a10-a240-e2835d1a3b12"+"f2b0c2e9-3c54-41aa-81c0-14586166a318"), ShouldBeFalse)
	})
}

func TestMapAnyUint64_Get(t *testing.T) {
	Convey("TestMapAnyUint64.Get", t, func() {
		var k interface{} = "707cff3a-5515-4139-b0b4-98cef12f7e56"
		var v uint64 = 6654874352450798209

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("e8e3c3d5-c815-4216-b85e-085520ebf226" + "aae11871-e2bf-4d64-af3e-d74280bd9e75")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyUint64_GetOpt(t *testing.T) {
	Convey("TestMapAnyUint64.GetOpt", t, func() {
		var k interface{} = "890415f0-9c9c-4437-a8e6-c336516b07f6"
		var v uint64 = 2228244207485341634

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("e661f402-afea-47b7-bc21-73b5cbec9684" + "5e245eb6-4a2b-4fce-b964-0025198c30e8")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyUint64_ForEach(t *testing.T) {
	Convey("TestMapAnyUint64.ForEach", t, func() {
		var k interface{} = "185f3d8d-f634-40d0-a05b-5f9e3cb28d3d"
		var v uint64 = 11535971362814566731
		hits := 0

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv uint64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyUint64_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyUint64.MarshalYAML", t, func() {
		var k interface{} = "0e050491-1255-4f5a-8690-c61bb9d8a837"
		var v uint64 = 72821772561136090

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyUint64_ToYAML(t *testing.T) {
	Convey("TestMapAnyUint64.ToYAML", t, func() {
		var k interface{} = "b866f86d-ead6-41a7-ab76-73c618d36353"
		var v uint64 = 7620591171096332365

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.SequenceNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 1)
		So(xyml.IsMap(c.Content[0]), ShouldBeTrue)
	})
}

func TestMapAnyUint64_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyUint64.PutIfNotNil", t, func() {
		var k interface{} = "6d875b42-f4d6-4846-95d4-e91eb594b791"
		var v uint64 = 14799586150054959445

		test := omap.NewMapAnyUint64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("3e2b1ce5-099c-4c6d-9eb4-3d0d2aa4850f", (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint64 = 738095586491946336
		So(test.PutIfNotNil("fb564852-a3d2-4bb9-b8a4-afad1c30e4a7", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyUint64.ReplaceIfExists", t, func() {
		var k interface{} = "5403813b-248d-491f-b11b-fbc3e43e28ef"
		var v uint64 = 18384310251159491334
		var x uint64 = 12667740073535672707

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("b0e33ff0-2bc6-4f39-a509-a7b7fd92861c", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyUint64.ReplaceOrPut", t, func() {
		var k interface{} = "64ad743f-b4f1-42de-9b0f-dec3cc4d625b"
		var v uint64 = 10981510661408728526
		var x uint64 = 13217035417252128182

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("5ad72066-63ac-4e2a-907d-9dbc4fa7d875", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint64_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyUint64.MarshalJSON", t, func() {
		var k interface{} = "5d88f506-da62-447a-acb6-f7e6ab1bd5ac"
		var v uint64 = 14878609253644713957

		test := omap.NewMapAnyUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"5d88f506-da62-447a-acb6-f7e6ab1bd5ac","value":14878609253644713957}]`)
	})
}
