package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyUint16_Put(t *testing.T) {
	Convey("TestMapAnyUint16.Put", t, func() {
		var k interface{} = "57d3f930-fc96-4127-b816-baf0ee8fb3e9"
		var v uint16 = 26127

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyUint16_Delete(t *testing.T) {
	Convey("TestMapAnyUint16.Delete", t, func() {
		var k interface{} = "ed67daa9-e0a0-42b0-8638-da9c0897b410"
		var v uint16 = 37750

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyUint16_Has(t *testing.T) {
	Convey("TestMapAnyUint16.Has", t, func() {
		var k interface{} = "6adeef83-3181-45f8-b871-8fa2540c36b4"
		var v uint16 = 42718

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("5d44f96c-643f-4fd1-a589-e0c4eba64a04"+"755ea10c-064f-455d-92c0-9b4aecc5c1ad"), ShouldBeFalse)
	})
}


func TestMapAnyUint16_Get(t *testing.T) {
	Convey("TestMapAnyUint16.Get", t, func() {
		var k interface{} = "49d445d0-2cfa-4f06-acb9-c6c765adbf6d"
		var v uint16 = 62184

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("dcd5f0a3-527c-4a03-93db-61abe8ef5bcc" + "d02e7b3c-b350-4e41-a6be-a1af3aaba21b")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyUint16_GetOpt(t *testing.T) {
	Convey("TestMapAnyUint16.GetOpt", t, func() {
		var k interface{} = "15853da7-0bcb-466c-bdcd-4ef65cc96aba"
		var v uint16 = 1761

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("f2f199b4-7f65-40c0-89c7-22859083e8cd" + "4b9acf62-24dd-4e61-92ca-3e0b70491e31")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyUint16_ForEach(t *testing.T) {
	Convey("TestMapAnyUint16.ForEach", t, func() {
		var k interface{} = "770ef086-9145-46a8-9190-d832f6692f6e"
		var v uint16 = 4334
		hits := 0

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv uint16) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyUint16_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyUint16.MarshalYAML", t, func() {
		var k interface{} = "8f603e16-396f-4218-b27d-1d389681c9bf"
		var v uint16 = 24900

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyUint16_ToYAML(t *testing.T) {
	Convey("TestMapAnyUint16.ToYAML", t, func() {
		var k interface{} = "9f0bc42c-b99c-4f1d-a29f-baa66e3c51ca"
		var v uint16 = 7087

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyUint16_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyUint16.PutIfNotNil", t, func() {
		var k interface{} = "bf17ff87-c462-493c-b59e-db776b41af67"
		var v uint16 = 44555

		test := omap.NewMapAnyUint16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("e47a9066-62b2-45c8-aeda-bdbdabdbd8d7", (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint16 = 65221
		So(test.PutIfNotNil("145aa8a1-e721-4bf0-925a-135ade892fb4", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyUint16.ReplaceIfExists", t, func() {
		var k interface{} = "ab913015-1b70-4490-919b-c7fefa3af81e"
		var v uint16 = 4229
		var x uint16 = 52506

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("68233239-e7d6-4ae3-95a4-a7dedc3f1879", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyUint16.ReplaceOrPut", t, func() {
		var k interface{} = "ff76ceb6-20a8-4b4a-a040-01fa76523bdf"
		var v uint16 = 209
		var x uint16 = 23437

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("b94a4d8b-8237-4a5a-825f-eb80f21c4349", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint16_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyUint16.MarshalJSON", t, func() {
		var k interface{} = "4741ce78-fbf5-414e-9b47-bf93f8f5139f"
		var v uint16 = 54413

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"4741ce78-fbf5-414e-9b47-bf93f8f5139f","value":54413}]`)
	})
}

