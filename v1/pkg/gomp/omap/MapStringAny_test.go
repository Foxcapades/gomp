package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringAny_Put(t *testing.T) {
	Convey("TestMapStringAny.Put", t, func() {
		var k string = "22c52df9-091a-4e2a-812e-af0f17f65fb5"
		var v interface{} = "83907914-d6c3-46db-9d86-b6376b40470a"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringAny_Delete(t *testing.T) {
	Convey("TestMapStringAny.Delete", t, func() {
		var k string = "c9d7d410-0d5a-48b2-9ab5-f09056a7962d"
		var v interface{} = "cacff04e-f6e7-4244-8ceb-a68b801b1225"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringAny_Has(t *testing.T) {
	Convey("TestMapStringAny.Has", t, func() {
		var k string = "d1e9b727-52dd-41f3-8ae7-5a665cbb51f4"
		var v interface{} = "ace3c473-3d99-4e98-bfcc-65b3b336f679"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("339bd09d-b26b-4893-968e-f47b0dde2141"+"43100477-c957-4c59-aa22-a81c6c4f2f68"), ShouldBeFalse)
	})
}

func TestMapStringAny_Get(t *testing.T) {
	Convey("TestMapStringAny.Get", t, func() {
		var k string = "bd407634-411a-4c75-8b2a-6ef33fa030d5"
		var v interface{} = "6ccc26ae-6110-498d-ae4a-ccd6670ed0fc"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("007cf144-38b2-4de6-bd0f-fcf12e351ca4" + "3ec574ab-0f27-4ab0-8062-8d6fba9850ce")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringAny_GetOpt(t *testing.T) {
	Convey("TestMapStringAny.GetOpt", t, func() {
		var k string = "3583e148-7728-4f11-a334-283bc60586ed"
		var v interface{} = "354992d3-2ec5-4699-b318-cc4ab4889590"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("76a14a1f-c778-49b3-8f94-56c6c3972887" + "ac018bbe-8f74-4f4f-b1bc-91ffc31a1567")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringAny_ForEach(t *testing.T) {
	Convey("TestMapStringAny.ForEach", t, func() {
		var k string = "83f10bb8-61ac-41fa-b7b4-6ea01ffa0ffb"
		var v interface{} = "4a51b6bc-dca7-4c5f-9517-8028307f669f"
		hits := 0

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv interface{}) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringAny_MarshalYAML(t *testing.T) {
	Convey("TestMapStringAny.MarshalYAML", t, func() {
		var k string = "8c26bb5a-ce10-40ca-85bf-857d27d70f21"
		var v interface{} = "b629f412-2d2c-417f-a26d-6818ee836b0f"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringAny_ToYAML(t *testing.T) {
	Convey("TestMapStringAny.ToYAML", t, func() {
		var k string = "327fde9f-cbf3-426e-93f4-5ddc56955f9f"
		var v interface{} = "dfd62946-d7ef-432e-bfb7-3a2269c79706"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringAny_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringAny.PutIfNotNil", t, func() {
		var k string = "908d3c0d-3ee2-4bc8-ac0a-6211d7f5020c"
		var v interface{} = "87316041-f22f-4e59-946a-7e1c347f4bd4"

		test := omap.NewMapStringAny(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("dd7a8435-5f60-4c48-a6bb-9d7a5865f680", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "7dc4d6f8-3560-42d6-a264-4bd3287650ed"
		So(test.PutIfNotNil("eae6099d-641e-41f5-a7d1-46c074628d27", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringAny_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringAny.ReplaceIfExists", t, func() {
		var k string = "786dee05-98ff-4d7c-ad93-8613439b07bc"
		var v interface{} = "689d2165-4119-483a-af03-da3960121deb"
		var x interface{} = "b10264f5-d15f-4a64-9748-dacb7d418899"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("b1cf5d32-f20b-40e5-8eac-42f9f2a8b97b", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringAny_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringAny.ReplaceOrPut", t, func() {
		var k string = "2b07dec6-2bdf-4785-b960-1fbbdabbd5bb"
		var v interface{} = "4395e856-fa23-4a99-badd-279bf37e1dc6"
		var x interface{} = "e6d8fce7-e2d8-45b5-902b-a1a6ab1a30c1"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("87e17732-e89d-48ff-a666-2a5ba682484f", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringAny_MarshalJSON(t *testing.T) {
	Convey("TestMapStringAny.MarshalJSON", t, func() {
		var k string = "a9260b82-8723-4aa0-85ff-2baa67280d64"
		var v interface{} = "f82d334a-d2ff-48b9-bda0-c2304ffa305d"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"a9260b82-8723-4aa0-85ff-2baa67280d64","value":"f82d334a-d2ff-48b9-bda0-c2304ffa305d"}]`)
	})
}
