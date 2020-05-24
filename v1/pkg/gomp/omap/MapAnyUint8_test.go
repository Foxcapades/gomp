package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyUint8_Put(t *testing.T) {
	Convey("TestMapAnyUint8.Put", t, func() {
		var k interface{} = "a83147f0-6dfc-454a-aafc-b7871e3600bb"
		var v uint8 = 164

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyUint8_Delete(t *testing.T) {
	Convey("TestMapAnyUint8.Delete", t, func() {
		var k interface{} = "8aa184fd-8455-4055-8a66-97ba47dc6985"
		var v uint8 = 91

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyUint8_Has(t *testing.T) {
	Convey("TestMapAnyUint8.Has", t, func() {
		var k interface{} = "9efd74ff-0d73-4c62-9a9f-21cc88e7b3a0"
		var v uint8 = 225

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("06c12ec9-e4de-4496-a50f-aa53f8882ad2"+"8e543e85-3993-4ed5-b5bc-d6eaeedb930a"), ShouldBeFalse)
	})
}

func TestMapAnyUint8_Get(t *testing.T) {
	Convey("TestMapAnyUint8.Get", t, func() {
		var k interface{} = "9a9c3d11-4c0b-4c83-994a-80dcc453b1e7"
		var v uint8 = 69

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("01102dfb-5de0-4ee1-8d51-1cc7d9d3e002" + "b2cb8910-6b6e-4fd9-a25d-7147247db93f")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyUint8_GetOpt(t *testing.T) {
	Convey("TestMapAnyUint8.GetOpt", t, func() {
		var k interface{} = "aed35df7-7d75-4a9c-8ab1-1a2ba256ca39"
		var v uint8 = 20

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("5984312f-9b24-4e26-8983-0273246d8c35" + "08382c6c-1ae8-40dc-9db8-87a21f46dd53")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyUint8_ForEach(t *testing.T) {
	Convey("TestMapAnyUint8.ForEach", t, func() {
		var k interface{} = "273a1bad-2e94-4a22-bab9-f2712cf47b08"
		var v uint8 = 87
		hits := 0

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv uint8) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyUint8_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyUint8.MarshalYAML", t, func() {
		var k interface{} = "b4df079f-e8e4-4049-bddb-e9c4f2386732"
		var v uint8 = 115

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyUint8_ToYAML(t *testing.T) {
	Convey("TestMapAnyUint8.ToYAML", t, func() {
		var k interface{} = "82e8e663-1a16-4935-9df5-50bc0b13e30d"
		var v uint8 = 40

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyUint8_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyUint8.PutIfNotNil", t, func() {
		var k interface{} = "c32a5e75-dfe5-4a14-a2d8-1a02fa0bc71b"
		var v uint8 = 241

		test := omap.NewMapAnyUint8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("a5a2c1c3-1cfd-42ad-ab1e-e75401095638", (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint8 = 21
		So(test.PutIfNotNil("bd3c6be1-b93e-42b8-bc5c-ea7d9d91ecd4", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyUint8.ReplaceIfExists", t, func() {
		var k interface{} = "9dadad9e-6ad9-4b63-a64f-679e5ef018e6"
		var v uint8 = 0
		var x uint8 = 205

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("d77b3201-384d-474a-93e9-23f867797618", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyUint8.ReplaceOrPut", t, func() {
		var k interface{} = "82264b4d-59ff-4820-994b-9a9025f8e455"
		var v uint8 = 27
		var x uint8 = 246

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("1e5069a5-26fc-413d-8542-3c0e63ba084f", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint8_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyUint8.MarshalJSON", t, func() {
		var k interface{} = "3b61f320-aa04-4531-8917-9deead01dbc7"
		var v uint8 = 86

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"3b61f320-aa04-4531-8917-9deead01dbc7","value":86}]`)
	})
}
