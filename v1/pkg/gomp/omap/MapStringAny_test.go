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
		var k string = "4dbab37a-daf2-4a27-be7a-c9c799b2d92e"
		var v interface{} = "3210d6e1-29e6-4082-ac18-4995e43dfb4b"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringAny_Delete(t *testing.T) {
	Convey("TestMapStringAny.Delete", t, func() {
		var k string = "fdb6007d-f7cd-4e9c-994f-f4d981361c50"
		var v interface{} = "0011eb2b-5312-450f-83b7-f6f8a6f07b91"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringAny_Has(t *testing.T) {
	Convey("TestMapStringAny.Has", t, func() {
		var k string = "37bf6520-9885-4a10-9f97-33dc18d3e38e"
		var v interface{} = "31f62c4f-6b00-4e75-ae03-b3f0f14fb5c0"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("68b888c3-f910-4b84-a765-ac5438205057"+"37e29944-9080-4723-b9e1-a138de6d1817"), ShouldBeFalse)
	})
}

func TestMapStringAny_Get(t *testing.T) {
	Convey("TestMapStringAny.Get", t, func() {
		var k string = "e59afcf8-deee-4eeb-a2d8-de208dc2b64f"
		var v interface{} = "1047b40c-f2ac-4d81-97d1-993624ac1795"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("d6b6400c-b8b9-4500-a49a-0f0b33baf946" + "b06ce3b1-26c5-4cde-8020-8a6b5ecce09d")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringAny_GetOpt(t *testing.T) {
	Convey("TestMapStringAny.GetOpt", t, func() {
		var k string = "9899b919-fa21-4737-95e7-286a0da2952b"
		var v interface{} = "797a8fb8-297a-4a3e-b09e-9892d9bd1011"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("bdf133cd-f8a0-4250-a0ba-d0bbe633cbaf" + "dd06ac61-22bb-41da-be81-040a34a88b1d")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringAny_ForEach(t *testing.T) {
	Convey("TestMapStringAny.ForEach", t, func() {
		var k string = "4cf3dd24-cdf6-47eb-99fb-5a0294578a7d"
		var v interface{} = "809bf542-5bac-41f5-b5cd-b330722cbad0"
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
		var k string = "d3bfeb96-1550-4125-b91d-f40ee87397a8"
		var v interface{} = "94644c0a-e874-4cf1-a553-89dfa814d6f4"

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
		var k string = "7b4e9a78-def2-4540-af79-d6cd8332fb2d"
		var v interface{} = "f418b25f-a59b-4fcc-ad28-d281a87554c7"

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
		var k string = "c2e9f5f8-0a66-4f3f-8cf9-bd7627e68df3"
		var v interface{} = "0171518d-4346-4289-8a37-3f631210c28f"

		test := omap.NewMapStringAny(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("a64f56fd-c140-4f19-9b7e-fb91e0031910", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "02d0b16e-9efd-40c1-8f89-c4b818d2874a"
		So(test.PutIfNotNil("5510d017-696c-4633-8bf1-6b543174077d", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringAny_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringAny.ReplaceIfExists", t, func() {
		var k string = "4efff580-8093-4bc2-9031-92c032cfca02"
		var v interface{} = "1626dcc6-6d1e-4196-8742-d43130a219f4"
		var x interface{} = "4700ae5e-c3ae-493c-9891-ece4fb606f6b"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("34468ca4-61ba-4ba5-9ec6-ae897c4d1fe2", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringAny_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringAny.ReplaceOrPut", t, func() {
		var k string = "e15706a8-da18-4781-a3b5-18d48f983ee6"
		var v interface{} = "ef574948-f2d9-411b-8357-3aa24d0433ff"
		var x interface{} = "7d86e202-1d84-4f2e-a45d-5bc62989ce35"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("25c216a2-da0c-4bd2-ad9e-a75df8da981b", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringAny_MarshalJSON(t *testing.T) {
	Convey("TestMapStringAny.MarshalJSON", t, func() {
		var k string = "1fb76911-3ba9-42ab-a7a4-1cc29095843f"
		var v interface{} = "3aa3ee4a-0d06-468e-9a30-4686e5bdba8f"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"1fb76911-3ba9-42ab-a7a4-1cc29095843f","value":"3aa3ee4a-0d06-468e-9a30-4686e5bdba8f"}]`)
	})
}
