package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyUint8_Put(t *testing.T) {
	Convey("TestMapAnyUint8.Put", t, func() {
		var k interface{} = "a1224064-aca3-4a0f-9d9a-3fa316bc7765"
		var v uint8 = 182

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyUint8_Delete(t *testing.T) {
	Convey("TestMapAnyUint8.Delete", t, func() {
		var k interface{} = "ea714a9a-3c3a-421e-9f6b-e57a4fe7bdd6"
		var v uint8 = 246

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyUint8_Has(t *testing.T) {
	Convey("TestMapAnyUint8.Has", t, func() {
		var k interface{} = "b55cfd7d-ea7b-4437-8067-2f1915f76137"
		var v uint8 = 105

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("c7243da9-71e8-4cc2-b302-bab8977e53e7"+"5c573878-45eb-4eac-bd5e-72a6a8e5773a"), ShouldBeFalse)
	})
}

func TestMapAnyUint8_Get(t *testing.T) {
	Convey("TestMapAnyUint8.Get", t, func() {
		var k interface{} = "e4c73b4b-80d1-4046-8503-6d8b657ce99f"
		var v uint8 = 250

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("fdeca953-88c9-49b6-9a6f-12debce0aca6" + "79d9603c-3eee-4bdb-91bb-768fc1469b96")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyUint8_GetOpt(t *testing.T) {
	Convey("TestMapAnyUint8.GetOpt", t, func() {
		var k interface{} = "9d86a021-0943-4ea4-9e0b-2faba3004c69"
		var v uint8 = 126

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("55fd3bbf-2e1c-4c7a-9136-e42c2f5a80d5" + "5e204ff4-e1de-405a-865e-169ebcf41409")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyUint8_ForEach(t *testing.T) {
	Convey("TestMapAnyUint8.ForEach", t, func() {
		var k interface{} = "2327f48c-1bf2-4249-bdde-e15b7fdc9dae"
		var v uint8 = 85
		hits := 0

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv uint8) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyUint8_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyUint8.MarshalYAML", t, func() {
		var k interface{} = "022c54d8-af9c-432e-ab9a-ddb3f9224458"
		var v uint8 = 77

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyUint8_ToYAML(t *testing.T) {
	Convey("TestMapAnyUint8.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "39da3f8c-47f8-4de9-b1e7-77041d05c040"
			var v uint8 = 203

			test := omap.NewMapAnyUint8(1)

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
			var k interface{} = "088ca0a9-f6d3-434d-86e5-abef3e044807"
			var v uint8 = 23

			test := omap.NewMapAnyUint8(1)
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

func TestMapAnyUint8_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyUint8.PutIfNotNil", t, func() {
		var k interface{} = "0a45cf5e-94c8-43ae-8caa-f30fcdf45294"
		var v uint8 = 241

		test := omap.NewMapAnyUint8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("f1b37bb7-b77a-406c-b117-67577fa78f1f", (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint8 = 239
		So(test.PutIfNotNil("89253949-f34b-42c1-81d0-f446aebcda92", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyUint8.ReplaceIfExists", t, func() {
		var k interface{} = "5bcefb7c-413e-4ba5-a4de-8059a720c686"
		var v uint8 = 106
		var x uint8 = 86

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("e8d1538c-52ca-46c4-98ce-ba05d61a0561", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyUint8.ReplaceOrPut", t, func() {
		var k interface{} = "742c86db-f55c-4bb0-88df-6d384e7dcf91"
		var v uint8 = 232
		var x uint8 = 6

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("fd4bf74a-c410-4bfc-8e09-da42be50431a", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint8_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyUint8.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "d425a123-b955-45d9-8fc1-abd52c2de8aa"
			var v uint8 = 36

			test := omap.NewMapAnyUint8(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"d425a123-b955-45d9-8fc1-abd52c2de8aa","value":36}]`)
		})

		Convey("Unordered", func() {
			var k interface{} = "d425a123-b955-45d9-8fc1-abd52c2de8aa"
			var v uint8 = 36

			test := omap.NewMapAnyUint8(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"d425a123-b955-45d9-8fc1-abd52c2de8aa":36}`)
		})

	})
}
