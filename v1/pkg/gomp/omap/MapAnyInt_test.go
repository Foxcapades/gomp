package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyInt_Put(t *testing.T) {
	Convey("TestMapAnyInt.Put", t, func() {
		var k interface{} = "08c76ecd-b003-4b21-a2be-88aeef1b3a2b"
		var v int = 948921499

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyInt_Delete(t *testing.T) {
	Convey("TestMapAnyInt.Delete", t, func() {
		var k interface{} = "b208ec35-ccd8-4638-ac75-21c61ece983f"
		var v int = 163923651

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyInt_Has(t *testing.T) {
	Convey("TestMapAnyInt.Has", t, func() {
		var k interface{} = "6bd58321-77e1-40a1-9ee1-87097b850c69"
		var v int = 1665080103

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("d404758d-47cf-4f92-996d-5350fe15b2ea"+"13f18c1b-7c78-48b7-b8d6-f7d408cb31e7"), ShouldBeFalse)
	})
}

func TestMapAnyInt_Get(t *testing.T) {
	Convey("TestMapAnyInt.Get", t, func() {
		var k interface{} = "428bf755-f2d4-4f0b-ac00-9e78ffd4dd8c"
		var v int = 1227188268

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("9b70bc44-ce0f-4fac-99de-96f96e60b906" + "525d420a-2e02-457c-8b81-5b3944b63b16")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyInt_GetOpt(t *testing.T) {
	Convey("TestMapAnyInt.GetOpt", t, func() {
		var k interface{} = "ef57fb68-3582-4fa9-b0cc-b9ca1d7e3ad5"
		var v int = 186377453

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("d3000541-0cb4-4536-970d-0a7c34825270" + "ae9ec70a-4798-4a18-bd40-84736d438267")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyInt_ForEach(t *testing.T) {
	Convey("TestMapAnyInt.ForEach", t, func() {
		var k interface{} = "79c3f4a9-a018-4dca-98da-de5b9b4c3880"
		var v int = 21224643
		hits := 0

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv int) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyInt_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyInt.MarshalYAML", t, func() {
		var k interface{} = "86455ce6-b6ee-4910-b202-de09827e2d41"
		var v int = 1514185349

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyInt_ToYAML(t *testing.T) {
	Convey("TestMapAnyInt.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "ea145e1f-12be-4ade-bdfa-f01a6524210f"
			var v int = 1777066728

			test := omap.NewMapAnyInt(1)

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
			var k interface{} = "68746465-ce87-4b05-beda-811da03455ad"
			var v int = 1573816088

			test := omap.NewMapAnyInt(1)
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

func TestMapAnyInt_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyInt.PutIfNotNil", t, func() {
		var k interface{} = "88e526bd-fd1b-4b52-a4c2-2927c33ee28e"
		var v int = 60317687

		test := omap.NewMapAnyInt(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("54826597-7d29-4748-8869-a847b361ee9c", (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int = 599580967
		So(test.PutIfNotNil("4ab90a69-2ef4-47ba-83d9-c378021929a7", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyInt.ReplaceIfExists", t, func() {
		var k interface{} = "69ae6150-afff-4660-b042-d0360333003b"
		var v int = 1314435995
		var x int = 1596966687

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("043fd9c8-e1fb-411e-999a-9b6525672ca5", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyInt.ReplaceOrPut", t, func() {
		var k interface{} = "fb12f48e-f5dc-41ae-b2b3-aafac9bf3d4b"
		var v int = 1499919042
		var x int = 1855292303

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("a6d58f21-566f-44b3-96f8-a98fefebf9be", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyInt.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "109d66a3-0fee-4921-8d0f-c835ab054cfb"
			var v int = 976667463

			test := omap.NewMapAnyInt(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"109d66a3-0fee-4921-8d0f-c835ab054cfb","value":976667463}]`)
		})

		Convey("Unordered", func() {
			var k interface{} = "109d66a3-0fee-4921-8d0f-c835ab054cfb"
			var v int = 976667463

			test := omap.NewMapAnyInt(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"109d66a3-0fee-4921-8d0f-c835ab054cfb":976667463}`)
		})

	})
}
