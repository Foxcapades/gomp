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
		var k interface{} = "21b822da-21c9-46a0-95c1-89bdff4842b9"
		var v uint32 = 595632742

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyUint32_Delete(t *testing.T) {
	Convey("TestMapAnyUint32.Delete", t, func() {
		var k interface{} = "6d7585d7-8925-40c3-96bf-f9f9aff7b48f"
		var v uint32 = 772221370

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyUint32_Has(t *testing.T) {
	Convey("TestMapAnyUint32.Has", t, func() {
		var k interface{} = "ecf64985-5c2a-4457-906b-4de423238806"
		var v uint32 = 1063996698

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("52babc5e-c17d-4753-b46e-90ec5cee9798"+"05d165e9-ce41-4e7c-a3a3-c5b281193990"), ShouldBeFalse)
	})
}


func TestMapAnyUint32_Get(t *testing.T) {
	Convey("TestMapAnyUint32.Get", t, func() {
		var k interface{} = "0d6dee2c-4aa2-4a96-be0a-c1544c766c06"
		var v uint32 = 2495491639

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("81ed376b-a9b4-4492-b73b-2b6c820920b1" + "840df9d4-2dce-4b1e-841f-44e0d71716f7")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyUint32_GetOpt(t *testing.T) {
	Convey("TestMapAnyUint32.GetOpt", t, func() {
		var k interface{} = "de719693-561b-4cec-a744-af584dbcf462"
		var v uint32 = 485939169

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("414af345-4b0b-40f5-adee-5cea8f6fe95b" + "6e794eb0-00bf-4c87-b30e-e6bc9ac4f089")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyUint32_ForEach(t *testing.T) {
	Convey("TestMapAnyUint32.ForEach", t, func() {
		var k interface{} = "acf65538-1a52-4e6a-be57-f2ef4170c53c"
		var v uint32 = 227170175
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
		var k interface{} = "e5b47dd8-11c0-4f40-8269-d4529800aabb"
		var v uint32 = 2590771777

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
		var k interface{} = "93a4b366-6b24-424f-990a-ba3770464d97"
		var v uint32 = 1653989431

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
		var k interface{} = "4d3da2c5-f39f-47c6-9260-f09f6a47f4eb"
		var v uint32 = 183864003

		test := omap.NewMapAnyUint32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("b6cd7780-7480-4b11-bf65-b8f65bad79dc", (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint32 = 1544782544
		So(test.PutIfNotNil("a04f47f4-45ab-4888-8e69-fcc3cbcf1da5", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyUint32.ReplaceIfExists", t, func() {
		var k interface{} = "5fc5b5c6-d6fc-4077-8654-19112b836ae6"
		var v uint32 = 3013626192
		var x uint32 = 222077777

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("1698e041-2431-48fc-b7c0-fd61c0df42ff", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyUint32.ReplaceOrPut", t, func() {
		var k interface{} = "134f2dc3-99ad-4419-a456-a3fcdf66dd28"
		var v uint32 = 2259420024
		var x uint32 = 2967891712

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("55efd198-c44d-47cd-955c-5611157fce48", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint32_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyUint32.MarshalJSON", t, func() {
		var k interface{} = "47eeb04c-1cf6-4f96-b089-2b0a26e09643"
		var v uint32 = 1458868856

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"47eeb04c-1cf6-4f96-b089-2b0a26e09643","value":1458868856}]`)
	})
}
