package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringUint8_Put(t *testing.T) {
	Convey("TestMapStringUint8.Put", t, func() {
		var k string = "8272ec7c-72d9-48bd-920a-76e05f90247a"
		var v uint8 = 83

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringUint8_Delete(t *testing.T) {
	Convey("TestMapStringUint8.Delete", t, func() {
		var k string = "9032e44d-0a46-470d-93ad-ad740f06a231"
		var v uint8 = 181

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringUint8_Has(t *testing.T) {
	Convey("TestMapStringUint8.Has", t, func() {
		var k string = "540b626d-8006-4170-82dd-5259074a9645"
		var v uint8 = 237

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("e613b73f-0e06-4789-b98b-bcafebbd48cb"+"3d353cce-a830-43fe-9ea1-1b4dfcc84e97"), ShouldBeFalse)
	})
}


func TestMapStringUint8_Get(t *testing.T) {
	Convey("TestMapStringUint8.Get", t, func() {
		var k string = "d2168010-5976-4efb-ba19-c5b7669b31cc"
		var v uint8 = 151

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("ac2eb086-c184-4d89-ac3a-8b996f0f4b97" + "ec13ed1a-c3c0-4056-93e0-6faf490f5d50")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringUint8_GetOpt(t *testing.T) {
	Convey("TestMapStringUint8.GetOpt", t, func() {
		var k string = "2c87718b-150f-4463-a9be-55b919ad3978"
		var v uint8 = 242

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("93a8c208-693d-4ec9-8734-04065ef4583b" + "3615577c-6d2a-40fc-b757-2943d70cd4ca")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringUint8_ForEach(t *testing.T) {
	Convey("TestMapStringUint8.ForEach", t, func() {
		var k string = "c9a80446-a73a-46ac-b840-cb46b8cb1f95"
		var v uint8 = 30
		hits := 0

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv uint8) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringUint8_MarshalYAML(t *testing.T) {
	Convey("TestMapStringUint8.MarshalYAML", t, func() {
		var k string = "0e9aae96-3c7f-459f-94ac-b7d27b56fc7b"
		var v uint8 = 54

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringUint8_ToYAML(t *testing.T) {
	Convey("TestMapStringUint8.ToYAML", t, func() {
		var k string = "638d65f9-e7dc-410b-9f49-46cff17540ef"
		var v uint8 = 29

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringUint8_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringUint8.PutIfNotNil", t, func() {
		var k string = "455bc7e8-050a-4e9c-b640-a6a43835ce7c"
		var v uint8 = 17

		test := omap.NewMapStringUint8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("30ad51f7-870b-4461-928c-cc9d98bf4220", (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint8 = 20
		So(test.PutIfNotNil("153ae026-53fb-48ef-a55e-a43770807e4c", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringUint8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringUint8.ReplaceIfExists", t, func() {
		var k string = "45292c7e-1503-43bb-aead-5d206d3c116c"
		var v uint8 = 27
		var x uint8 = 198

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("d33ecfaa-0bfd-4886-9809-9f90dd317e41", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringUint8.ReplaceOrPut", t, func() {
		var k string = "ed3c4e2c-bb54-4466-8444-e79ad3094ac3"
		var v uint8 = 90
		var x uint8 = 25

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("7fe6af3e-8e8a-4846-b35b-ffa94bb1734b", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint8_MarshalJSON(t *testing.T) {
	Convey("TestMapStringUint8.MarshalJSON", t, func() {
		var k string = "9d0e68f8-547d-43bd-b04e-d3f1d2a3f86c"
		var v uint8 = 213

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"9d0e68f8-547d-43bd-b04e-d3f1d2a3f86c","value":213}]`)
	})
}

