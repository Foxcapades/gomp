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
		var k string = "5fe3fd1e-7177-4728-81e0-e513283a3f08"
		var v interface{} = "72d7515a-abc2-46df-a2ac-fb4f220923ab"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringAny_Delete(t *testing.T) {
	Convey("TestMapStringAny.Delete", t, func() {
		var k string = "d6917fb8-b056-40db-b182-749d62221931"
		var v interface{} = "9dc28980-5c08-4004-bc2c-cdc4570fe2ef"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringAny_Has(t *testing.T) {
	Convey("TestMapStringAny.Has", t, func() {
		var k string = "4a73204c-1e79-45f1-bd4b-8a57c7aa6e9a"
		var v interface{} = "e8804c7e-39ea-4c93-b052-9a0aea823c5e"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("6be7c51b-dac6-430d-b13a-06d57076e194"+"360a4e3b-c453-47f4-8d4a-8765f4ae9331"), ShouldBeFalse)
	})
}

func TestMapStringAny_Get(t *testing.T) {
	Convey("TestMapStringAny.Get", t, func() {
		var k string = "38d89932-3ddb-4362-b24e-f8b7dde5354e"
		var v interface{} = "790d32f3-fcd0-4667-a8f1-a814964ab91c"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("177896f1-aa54-4e58-bcbf-88af15d51934" + "0d1a8e2b-68bc-43f1-9bc6-91d297839ff9")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringAny_GetOpt(t *testing.T) {
	Convey("TestMapStringAny.GetOpt", t, func() {
		var k string = "ba0fcbbe-d7b3-4b77-9235-4b044e433dca"
		var v interface{} = "bb1e8c57-c19a-48f2-a425-29da72cda5fe"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("5f87d4be-e1c1-4d7a-96e4-7d6b22fbafb6" + "883a384e-d270-4db4-b1d5-4e4d92d2a9de")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringAny_ForEach(t *testing.T) {
	Convey("TestMapStringAny.ForEach", t, func() {
		var k string = "4c54e6c1-923d-47dc-85c1-3175fe3e4ec9"
		var v interface{} = "560093a0-689c-472e-8d2f-035b64df0d95"
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
		var k string = "41b0b9ca-6fa9-4b3c-8e27-381e18ef34a1"
		var v interface{} = "62982b28-b165-4625-b270-d563e59fb053"

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
		Convey("Ordered", func() {
			var k string = "a422edaf-9a4f-4851-8145-413d7fc44882"
			var v interface{} = "fdb2c395-f737-4862-93dd-b8e2f53d23b0"

			test := omap.NewMapStringAny(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			c, d := test.ToYAML()
			So(d, ShouldBeNil)
			So(c.Kind, ShouldEqual, yaml.SequenceNode)
			So(c.LongTag(), ShouldEqual, xyml.TagOrderedMap)
			So(len(c.Content), ShouldEqual, 1)
			So(xyml.IsMap(c.Content[0]), ShouldBeTrue)
		})

		Convey("Unordered", func() {
			var k string = "e3a22cb1-27ce-4aa6-afab-42d44e5517a1"
			var v interface{} = "f19fe47b-2ef7-4430-8aa2-44f142d49704"

			test := omap.NewMapStringAny(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			c, d := test.ToYAML()

			So(d, ShouldBeNil)
			So(c.Kind, ShouldEqual, yaml.MappingNode)
			So(c.LongTag(), ShouldEqual, xyml.TagMap)
			So(len(c.Content), ShouldEqual, 2)
		})
	})
}

func TestMapStringAny_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringAny.PutIfNotNil", t, func() {
		var k string = "aa5b6034-05c8-488b-936c-2ec9d3676d4f"
		var v interface{} = "34529acf-83f9-4094-9ac6-1253808aba8d"

		test := omap.NewMapStringAny(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("66ff417c-a964-495c-98c9-e5313c40b725", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "44e4d4d8-45ff-4c45-9ccb-e2a0dd0a7015"
		So(test.PutIfNotNil("c64421eb-2ae2-4bcf-a4dd-368c738a4886", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringAny_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringAny.ReplaceIfExists", t, func() {
		var k string = "adab4fcd-03da-4216-879c-f39aaebe3189"
		var v interface{} = "8224ab80-2582-46c1-82dd-edf9bf5dce58"
		var x interface{} = "9e8acb68-9357-4503-9090-834ed737b14d"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("979cdecf-154b-4d8f-bf5e-95ef15ac8a2c", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringAny_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringAny.ReplaceOrPut", t, func() {
		var k string = "fbd0811d-86a3-48c1-8bad-16d7e539111f"
		var v interface{} = "f9ca5f12-8f8f-42cc-87c0-18d83120d1b9"
		var x interface{} = "786c19b1-212f-4e78-bd63-5eceebd5737c"

		test := omap.NewMapStringAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("ed2ea425-0f88-4b4b-8c24-38b12636c35e", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringAny_MarshalJSON(t *testing.T) {
	Convey("TestMapStringAny.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k string = "e356209b-55bf-49bd-af2e-4f285f8a9ce5"
			var v interface{} = "b9c620ec-2bd5-473f-be52-4b00d38ba7f0"

			test := omap.NewMapStringAny(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"e356209b-55bf-49bd-af2e-4f285f8a9ce5","value":"b9c620ec-2bd5-473f-be52-4b00d38ba7f0"}]`)
		})

		Convey("Unordered", func() {
			var k string = "e356209b-55bf-49bd-af2e-4f285f8a9ce5"
			var v interface{} = "b9c620ec-2bd5-473f-be52-4b00d38ba7f0"

			test := omap.NewMapStringAny(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"e356209b-55bf-49bd-af2e-4f285f8a9ce5":"b9c620ec-2bd5-473f-be52-4b00d38ba7f0"}`)
		})

	})
}
