package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAny_Put(t *testing.T) {
	Convey("TestMapAny.Put", t, func() {
		var k interface{} = "74b999a8-ce86-48c9-989e-99e38f7c0797"
		var v interface{} = "39bee635-1c92-4c8a-93e5-5881324489e8"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAny_Delete(t *testing.T) {
	Convey("TestMapAny.Delete", t, func() {
		var k interface{} = "519e7f99-7dcb-4167-9ba2-8bde18ea2677"
		var v interface{} = "b1bdf10d-0d2b-44e3-8165-4967d9106432"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAny_Has(t *testing.T) {
	Convey("TestMapAny.Has", t, func() {
		var k interface{} = "b17d77ce-3176-4617-afb2-5f7ddb526109"
		var v interface{} = "11daa0d3-ec07-4f78-8a89-ec788b24b8de"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("af8b2a48-f8cd-4348-aa67-abe8ad8a7d97"+"397cfdd4-3187-4367-ba15-7edfa4750eaa"), ShouldBeFalse)
	})
}

func TestMapAny_Get(t *testing.T) {
	Convey("TestMapAny.Get", t, func() {
		var k interface{} = "9e36085d-d144-446a-816c-c2f49e8ab22f"
		var v interface{} = "1fd7cc6c-b056-4911-a4eb-4313d477e2bb"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("9230c1bc-ef6c-4695-abba-489b47142ba1" + "dab5f4e0-5b69-48cb-88d1-0d11df063d47")
		So(b, ShouldBeFalse)
	})
}

func TestMapAny_GetOpt(t *testing.T) {
	Convey("TestMapAny.GetOpt", t, func() {
		var k interface{} = "4b9eef5b-c87b-435b-8749-e2786f6d4511"
		var v interface{} = "d5bd4fad-c615-4ada-ac28-77b269c3f01a"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("84b84cd7-e30c-4e12-bbe5-6380ad7746de" + "01a159ca-39fb-4bf4-a1e3-57f8425ba93f")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAny_ForEach(t *testing.T) {
	Convey("TestMapAny.ForEach", t, func() {
		var k interface{} = "26250201-6b3e-4ad2-9a7f-da1178b9c6d6"
		var v interface{} = "413d89b5-17d7-4f12-bdde-fcb01b7efdfc"
		hits := 0

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv interface{}) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAny_MarshalYAML(t *testing.T) {
	Convey("TestMapAny.MarshalYAML", t, func() {
		var k interface{} = "b4493209-e737-4a50-be20-a30471109283"
		var v interface{} = "e3b649be-77c8-47f1-a790-9f5595d071f2"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAny_ToYAML(t *testing.T) {
	Convey("TestMapAny.ToYAML", t, func() {
		var k interface{} = "974d46e3-9f6b-401b-b330-2640bbf3f90e"
		var v interface{} = "e526594e-8a38-4d99-8bf8-a093f705c7e9"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAny_PutIfNotNil(t *testing.T) {
	Convey("TestMapAny.PutIfNotNil", t, func() {
		var k interface{} = "59e1b4ea-48a7-41b4-b2cb-73cb0ca15842"
		var v interface{} = "8234c94c-a835-4047-9c54-69f511350556"

		test := omap.NewMapAny(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("4f3ceec2-a210-4c2f-87ff-76cb5638029c", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "67ddaeda-d767-4a9f-b8c3-a79780769280"
		So(test.PutIfNotNil("de6036be-01a8-4c67-a567-fa8aae641695", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAny_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAny.ReplaceIfExists", t, func() {
		var k interface{} = "bdddfa30-c570-4229-93db-1f6a4cbb1d42"
		var v interface{} = "fe6727dc-f9be-4391-9fee-e170e3a12181"
		var x interface{} = "4d455990-4a16-49e8-a348-50a5472ffe6c"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("ed1e526c-dce2-4b2b-9765-bc8dfeed0407", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAny_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAny.ReplaceOrPut", t, func() {
		var k interface{} = "621d632b-e85c-434e-b088-239cf8d1d171"
		var v interface{} = "10442535-7bf3-4651-a7ec-ab76c50439b4"
		var x interface{} = "14ceb99b-c9ce-4b32-875e-fda4c1518d81"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("9bf02b76-b003-47b1-a164-dfcb49a7fedb", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAny_MarshalJSON(t *testing.T) {
	Convey("TestMapAny.MarshalJSON", t, func() {
		var k interface{} = "ba1f4a07-9369-45af-a9ca-12ca5a779758"
		var v interface{} = "87eaf773-3be4-46e4-be14-5262b0e2ca8f"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"ba1f4a07-9369-45af-a9ca-12ca5a779758","value":"87eaf773-3be4-46e4-be14-5262b0e2ca8f"}]`)
	})
}
