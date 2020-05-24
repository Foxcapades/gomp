package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyString_Put(t *testing.T) {
	Convey("TestMapAnyString.Put", t, func() {
		var k interface{} = "8853d749-29bb-480d-b1cd-77ecc8f39163"
		var v string = "b17b0a82-6b21-46fc-b126-40e7648a6221"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyString_Delete(t *testing.T) {
	Convey("TestMapAnyString.Delete", t, func() {
		var k interface{} = "7677380c-b364-40ff-86e5-fba637aa13ba"
		var v string = "2e6d8927-ec95-4e49-b573-fe52eff982e3"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyString_Has(t *testing.T) {
	Convey("TestMapAnyString.Has", t, func() {
		var k interface{} = "a928d51f-93df-414d-99d6-62b21dbaf35c"
		var v string = "f0456c3e-6860-45bf-8359-733bbe2509bd"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("aaedf5a2-2a51-405a-a0cb-c097004f58b8"+"31d02cc9-d2b9-4f6b-81cd-b59aacec6997"), ShouldBeFalse)
	})
}

func TestMapAnyString_Get(t *testing.T) {
	Convey("TestMapAnyString.Get", t, func() {
		var k interface{} = "501ee183-b448-4dfe-bf9e-9e073c1a247a"
		var v string = "8c129dc7-62a4-43a2-8b9e-fa03cc7f92ee"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("36a1a903-d578-48b7-9317-a17a3558ed6d" + "0cd1fa28-e840-4c49-af22-9546292253f0")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyString_GetOpt(t *testing.T) {
	Convey("TestMapAnyString.GetOpt", t, func() {
		var k interface{} = "f6bff88a-1fba-4ff7-b517-7b16867743d6"
		var v string = "5ac67f82-45de-4571-b3af-30f661866bf6"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("8bd48e56-7e3b-4df2-a63c-c5636239820f" + "38da2584-9d0e-4f7d-a114-31957558e99e")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyString_ForEach(t *testing.T) {
	Convey("TestMapAnyString.ForEach", t, func() {
		var k interface{} = "18c3e9ba-1e22-4dab-b458-86fe2eef440c"
		var v string = "4f868f22-dcc2-4531-a87a-bd1784c32593"
		hits := 0

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv string) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyString_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyString.MarshalYAML", t, func() {
		var k interface{} = "3a9a5d19-0523-449b-ace3-1e0828be2d60"
		var v string = "520ac68b-d020-4863-83b0-276f8b550eca"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyString_ToYAML(t *testing.T) {
	Convey("TestMapAnyString.ToYAML", t, func() {
		var k interface{} = "b54b7575-7f38-40eb-984c-2663d9f0f76e"
		var v string = "a907bdbe-3ae9-49bd-b17a-425de885db1a"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyString_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyString.PutIfNotNil", t, func() {
		var k interface{} = "48c61952-2a66-4bf4-8b84-53426b9622ab"
		var v string = "700e3bfe-0263-4331-b7f6-e1774324b1ab"

		test := omap.NewMapAnyString(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("17c21e2e-39fa-462d-a99b-95342b71e083", (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "46e52cd1-4cce-4cdc-ad41-95c2fdcec37a"
		So(test.PutIfNotNil("c9a34da9-ebbf-4e41-a93f-64cef98ba263", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyString_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyString.ReplaceIfExists", t, func() {
		var k interface{} = "84f2f797-d799-4eca-be7e-897ca1f30215"
		var v string = "f2f3310b-e54a-455b-bb4f-829d9ef7ec1f"
		var x string = "22e00862-6fec-4a34-b01e-feaa9213cb3f"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("a71bf403-bdf3-4a39-aa6e-738ea94f571b", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyString_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyString.ReplaceOrPut", t, func() {
		var k interface{} = "46e74e1a-62e7-4901-be07-a981b0cc686d"
		var v string = "7074c51a-9e71-46ae-96c4-f34cd09044fa"
		var x string = "13c8ed85-c0f8-4084-a051-9cbf02a24b5a"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("a0b78e88-80a5-49a7-a166-8d8f0ffe9e8b", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyString_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyString.MarshalJSON", t, func() {
		var k interface{} = "72e56da0-a74e-41d7-92a3-fcc0c0685a56"
		var v string = "ba64c1f1-53b9-4795-8cb4-7b837ce92ef6"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"72e56da0-a74e-41d7-92a3-fcc0c0685a56","value":"ba64c1f1-53b9-4795-8cb4-7b837ce92ef6"}]`)
	})
}
