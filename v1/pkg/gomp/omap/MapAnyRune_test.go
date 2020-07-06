package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyRune_Put(t *testing.T) {
	Convey("TestMapAnyRune.Put", t, func() {
		var k interface{} = "6e44273a-61ed-4582-8914-ad7aa0a95af4"
		var v rune = 1651965525

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyRune_Delete(t *testing.T) {
	Convey("TestMapAnyRune.Delete", t, func() {
		var k interface{} = "b9d4a613-d916-4dc4-a94a-e14305469699"
		var v rune = 1924576605

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyRune_Has(t *testing.T) {
	Convey("TestMapAnyRune.Has", t, func() {
		var k interface{} = "59a74b70-bceb-43b8-aaea-f204615b8403"
		var v rune = 723654501

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("24b38a87-e231-4f0c-b803-3b9453209935"+"a383a677-6784-44ca-bb00-1059d668318c"), ShouldBeFalse)
	})
}

func TestMapAnyRune_Get(t *testing.T) {
	Convey("TestMapAnyRune.Get", t, func() {
		var k interface{} = "a313b113-ab34-4b6d-ab96-50e7b40cae10"
		var v rune = 544861190

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("17e04922-9a67-4842-806b-458f24d6e241" + "3ae5afd9-63ad-46bd-abb3-09ac9c19547d")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyRune_GetOpt(t *testing.T) {
	Convey("TestMapAnyRune.GetOpt", t, func() {
		var k interface{} = "8823b862-adf1-48d8-a652-7f85731c967e"
		var v rune = 1269287753

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("bdd66cb4-daba-427d-bac8-44b2b924535a" + "1211bb4a-8fe4-4ee0-8ddb-bfde9767ae6c")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyRune_ForEach(t *testing.T) {
	Convey("TestMapAnyRune.ForEach", t, func() {
		var k interface{} = "de530d6d-9c14-44b4-8168-a9ee40cde997"
		var v rune = 659356362
		hits := 0

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv rune) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyRune_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyRune.MarshalYAML", t, func() {
		var k interface{} = "d6f076dd-2f9f-43ff-9972-8e348d609e0e"
		var v rune = 1894838732

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyRune_ToYAML(t *testing.T) {
	Convey("TestMapAnyRune.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "2c46a5c3-eaee-499c-bcfe-0b14d12c5d99"
			var v rune = 787068507

			test := omap.NewMapAnyRune(1)

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
			var k interface{} = "97f9d7ad-7b1d-4fd1-9f23-5fb86a836fe6"
			var v rune = 1187626370

			test := omap.NewMapAnyRune(1)
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

func TestMapAnyRune_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyRune.PutIfNotNil", t, func() {
		var k interface{} = "572cf906-849f-475a-b447-22f69a86380b"
		var v rune = 729490352

		test := omap.NewMapAnyRune(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("4517e5ac-cc75-4bab-aa51-b3749521be7d", (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x rune = 1508072063
		So(test.PutIfNotNil("5269893e-9408-4c9d-b1ce-10267c6d7e21", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyRune_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyRune.ReplaceIfExists", t, func() {
		var k interface{} = "71d6b230-5dc6-4434-a192-bd6145f91a1a"
		var v rune = 890857426
		var x rune = 1982002300

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("49d2eaa3-adb3-48ef-b44c-8481a8cee97d", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyRune_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyRune.ReplaceOrPut", t, func() {
		var k interface{} = "2b350777-5629-4419-bf41-7bebe5b98487"
		var v rune = 730369882
		var x rune = 1454741733

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("94abaffc-76ca-4733-a48b-e01873a2527a", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyRune_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyRune.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "d98499e5-1dd4-4f41-9408-dc042db9623f"
			var v rune = 1620553493

			test := omap.NewMapAnyRune(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"d98499e5-1dd4-4f41-9408-dc042db9623f","value":1620553493}]`)
		})

		Convey("Unordered", func() {
			var k interface{} = "d98499e5-1dd4-4f41-9408-dc042db9623f"
			var v rune = 1620553493

			test := omap.NewMapAnyRune(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"d98499e5-1dd4-4f41-9408-dc042db9623f":1620553493}`)
		})

	})
}
