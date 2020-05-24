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
		var k interface{} = "8faff414-a6ad-4e35-bea5-6e44d762e014"
		var v interface{} = "09a4e9d9-c5bc-49d7-b675-9f49ef446d87"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAny_Delete(t *testing.T) {
	Convey("TestMapAny.Delete", t, func() {
		var k interface{} = "c267570f-7793-46b4-b425-ceecaff5727b"
		var v interface{} = "a9a30e3a-4449-47d2-b4d0-d76ba6521715"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAny_Has(t *testing.T) {
	Convey("TestMapAny.Has", t, func() {
		var k interface{} = "2d0c0781-0f3a-44d8-b5db-291665d99f98"
		var v interface{} = "c838242d-3a9b-44f8-996d-f25f4fa33055"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("17b207c2-5b4a-4a45-b8e9-e2c34d41c5f3"+"1cd319a3-254e-4073-9917-e60ce6d02e82"), ShouldBeFalse)
	})
}

func TestMapAny_Get(t *testing.T) {
	Convey("TestMapAny.Get", t, func() {
		var k interface{} = "a29bacb2-a0d3-4896-a410-95345cde2499"
		var v interface{} = "303d2783-a941-491a-aeaf-78d3b36c4d3f"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("96f40de2-0ee1-41e5-92d7-5aee9a76841a" + "36d59f5b-ed5b-4c5a-aa5f-972851c1b12a")
		So(b, ShouldBeFalse)
	})
}

func TestMapAny_GetOpt(t *testing.T) {
	Convey("TestMapAny.GetOpt", t, func() {
		var k interface{} = "b8d1e9a0-3c9c-411a-a8d1-d1a950e51183"
		var v interface{} = "452be1d0-4d18-4cba-85ab-a8d0c7b996eb"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("c2bda074-2e43-49ee-9976-30c5dc35e489" + "51f992c5-e4dd-4dd1-b7f1-bcccda2e301c")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAny_ForEach(t *testing.T) {
	Convey("TestMapAny.ForEach", t, func() {
		var k interface{} = "7bb86c7d-e259-4962-a387-fc64a7b49c88"
		var v interface{} = "f24d3d32-6f84-44a2-94f4-d880bd241a73"
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
		var k interface{} = "64dce3f9-eec6-4ccd-96f7-841cd5b02df4"
		var v interface{} = "7e37ebf5-9237-4903-83ae-4046fc1e4916"

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
		var k interface{} = "d7cf19f7-5096-4d88-93c4-d993b208dc61"
		var v interface{} = "40927918-5f71-404f-8b55-66ffe4b16781"

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
		var k interface{} = "7093469b-671f-43dc-922d-6d6ecc2f30e4"
		var v interface{} = "aa22606c-24d4-47ab-8ee6-7b6d238210a7"

		test := omap.NewMapAny(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("2ed98abf-2810-4591-beff-c16e6f2a1450", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "c3c83c39-b779-4904-8d2b-20878610437e"
		So(test.PutIfNotNil("bcdad311-c5ca-4772-a2bc-de58dadff098", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAny_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAny.ReplaceIfExists", t, func() {
		var k interface{} = "102bf43d-7e41-4d20-9c24-d3cb2dbf34fe"
		var v interface{} = "05b7931f-a3f3-4d7a-a99c-501bf2460fab"
		var x interface{} = "0763c7d4-87d0-4762-ad56-62ced8920a4d"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("8407822d-eacf-46bb-b27b-d222d3a5f94f", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAny_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAny.ReplaceOrPut", t, func() {
		var k interface{} = "3e99a248-c3bb-4c84-8f91-cec7857d6894"
		var v interface{} = "272cde75-fd6d-426b-879b-858dba6929d1"
		var x interface{} = "8fac80d0-7b7d-45db-8a33-e214a586c689"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("e88a65f5-6277-4843-88cd-4f5c09c37efc", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAny_MarshalJSON(t *testing.T) {
	Convey("TestMapAny.MarshalJSON", t, func() {
		var k interface{} = "1a8cc8db-046c-437e-b31b-51136f541add"
		var v interface{} = "ed26ac45-dc40-4370-badd-316152a77417"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"1a8cc8db-046c-437e-b31b-51136f541add","value":"ed26ac45-dc40-4370-badd-316152a77417"}]`)
	})
}
