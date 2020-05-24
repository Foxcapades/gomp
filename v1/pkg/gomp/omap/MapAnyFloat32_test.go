package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyFloat32_Put(t *testing.T) {
	Convey("TestMapAnyFloat32.Put", t, func() {
		var k interface{} = "69d793c5-1590-4364-bc4b-0576b1beb81a"
		var v float32 = 0.541

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyFloat32_Delete(t *testing.T) {
	Convey("TestMapAnyFloat32.Delete", t, func() {
		var k interface{} = "ee0383c2-9b3e-4450-aa04-bfc508702bbf"
		var v float32 = 0.579

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyFloat32_Has(t *testing.T) {
	Convey("TestMapAnyFloat32.Has", t, func() {
		var k interface{} = "b03b97ef-97cb-4877-b706-e9449991dd59"
		var v float32 = 0.627

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("59dd9e0f-ee1f-4fd2-8dd5-87fcdd81b431"+"e4be0602-a7a7-4af3-8b08-017d7137db9c"), ShouldBeFalse)
	})
}

func TestMapAnyFloat32_Get(t *testing.T) {
	Convey("TestMapAnyFloat32.Get", t, func() {
		var k interface{} = "487705f0-71cb-476a-9d78-af834ff4a65b"
		var v float32 = 0.736

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("ed74f448-a694-4f82-a9fe-0d9e68f6db36" + "7f6bd149-4289-498d-89b9-a8002dff643a")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyFloat32_GetOpt(t *testing.T) {
	Convey("TestMapAnyFloat32.GetOpt", t, func() {
		var k interface{} = "93a7bd7d-9586-4069-bc58-46f1fc773f71"
		var v float32 = 0.727

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("4f34c675-b82c-4828-913e-df1fe304ec77" + "b064d76f-8392-4cb0-93ef-0f85de5722c4")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyFloat32_ForEach(t *testing.T) {
	Convey("TestMapAnyFloat32.ForEach", t, func() {
		var k interface{} = "fd362035-f92a-4f18-8c94-edf63797b89e"
		var v float32 = 0.632
		hits := 0

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv float32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyFloat32_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyFloat32.MarshalYAML", t, func() {
		var k interface{} = "2dfaff31-cfb6-4cf7-8ab3-1924768bd586"
		var v float32 = 0.213

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyFloat32_ToYAML(t *testing.T) {
	Convey("TestMapAnyFloat32.ToYAML", t, func() {
		var k interface{} = "1fe38e0d-eacc-499b-96ae-11d9aff0cd73"
		var v float32 = 0.329

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyFloat32_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyFloat32.PutIfNotNil", t, func() {
		var k interface{} = "dbc35daa-68fc-4eef-a375-c222c104fcee"
		var v float32 = 0.026

		test := omap.NewMapAnyFloat32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("da62cc4e-5258-4847-af6e-e080741b27c9", (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float32 = 0.565
		So(test.PutIfNotNil("d6d53e30-e961-4dbf-8d10-8b8718e940bb", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyFloat32.ReplaceIfExists", t, func() {
		var k interface{} = "f15f5395-626e-4851-9e8b-31674798ea23"
		var v float32 = 0.615
		var x float32 = 0.829

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("415ab52d-166b-40da-84a7-54185f38a0ab", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyFloat32.ReplaceOrPut", t, func() {
		var k interface{} = "f80ce108-4007-4770-bb7e-68198065c0c4"
		var v float32 = 0.615
		var x float32 = 0.780

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("047457d6-0aa4-41ac-a23d-952fe8de63f7", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat32_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyFloat32.MarshalJSON", t, func() {
		var k interface{} = "2053a5a6-c0b0-41af-a80a-9fd5524d493e"
		var v float32 = 0.811

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"2053a5a6-c0b0-41af-a80a-9fd5524d493e","value":0.811}]`)
	})
}
