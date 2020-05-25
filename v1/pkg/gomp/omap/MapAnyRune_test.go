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
		var k interface{} = "a7ea2706-cc3c-4255-9d2b-cbe0782acd56"
		var v rune = 601290484

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyRune_Delete(t *testing.T) {
	Convey("TestMapAnyRune.Delete", t, func() {
		var k interface{} = "361a90b9-8b47-489e-b1d3-684c684e1e88"
		var v rune = 250185391

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyRune_Has(t *testing.T) {
	Convey("TestMapAnyRune.Has", t, func() {
		var k interface{} = "93dde637-bdb8-479a-838f-1e5b1115c547"
		var v rune = 2147215844

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("2ed2c068-63db-41f7-9833-59642abc9825"+"c2bfe299-9af7-44bf-8277-f99560d327b7"), ShouldBeFalse)
	})
}

func TestMapAnyRune_Get(t *testing.T) {
	Convey("TestMapAnyRune.Get", t, func() {
		var k interface{} = "1dda4cc0-c141-485a-94ac-d03ceb628c05"
		var v rune = 2070388577

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("b571f699-d79a-4cb3-937b-2d3e150941fc" + "081c5e65-d6cc-4d0b-be0f-67b6e6b0e5c0")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyRune_GetOpt(t *testing.T) {
	Convey("TestMapAnyRune.GetOpt", t, func() {
		var k interface{} = "c29a0fb4-8f3b-4db3-8ccc-805fe6139020"
		var v rune = 397754323

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("36b5cdf6-340c-4bd5-97c4-fc549dcf2db8" + "b2b7290c-edc9-4b05-9655-717acb2979ff")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyRune_ForEach(t *testing.T) {
	Convey("TestMapAnyRune.ForEach", t, func() {
		var k interface{} = "350308ac-b6a1-4663-86ba-a7fb0ba71d0f"
		var v rune = 681270503
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
		var k interface{} = "6352cdca-fcaa-496f-bc35-5aea85211a39"
		var v rune = 931263441

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
		var k interface{} = "96a8ae8c-31e6-4084-b8c9-3619d92480c5"
		var v rune = 697385851

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.SequenceNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 1)
		So(xyml.IsMap(c.Content[0]), ShouldBeTrue)
	})
}

func TestMapAnyRune_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyRune.PutIfNotNil", t, func() {
		var k interface{} = "11ac8837-e367-40ed-b1a2-a54952195324"
		var v rune = 1034581053

		test := omap.NewMapAnyRune(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("c9637fc1-88ba-4fe7-920b-e1ca9e990d51", (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x rune = 1813701192
		So(test.PutIfNotNil("40dd4b10-d52d-46df-8893-11b05111c6ba", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyRune_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyRune.ReplaceIfExists", t, func() {
		var k interface{} = "d1410bee-8443-486c-ae80-391243034b2d"
		var v rune = 1894959925
		var x rune = 1654796424

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("40a49c2e-e9ef-495e-ab38-79147e1e11de", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyRune_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyRune.ReplaceOrPut", t, func() {
		var k interface{} = "c9102942-c45b-421f-ba5b-492e94602734"
		var v rune = 1422424427
		var x rune = 1968860615

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("93c8b733-c685-46c2-af19-7c7599eb0088", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyRune_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyRune.MarshalJSON", t, func() {
		var k interface{} = "98b7ed6b-b622-44d1-a844-70753c6324d7"
		var v rune = 578054574

		test := omap.NewMapAnyRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"98b7ed6b-b622-44d1-a844-70753c6324d7","value":578054574}]`)
	})
}
