package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringRune_Put(t *testing.T) {
	Convey("TestMapStringRune.Put", t, func() {
		var k string = "59da8201-3451-4017-b552-e5fc5b216a4b"
		var v rune = 1892690142

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringRune_Delete(t *testing.T) {
	Convey("TestMapStringRune.Delete", t, func() {
		var k string = "733d316f-85c3-486b-b0b7-2c74e1dbee5c"
		var v rune = 1216863360

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringRune_Has(t *testing.T) {
	Convey("TestMapStringRune.Has", t, func() {
		var k string = "fef8ff35-4540-40ea-94ea-fe0465d981ad"
		var v rune = 1738133608

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("fb5d8bf0-4a57-454f-9217-c292a9149837"+"b65cbbc8-99b2-4d1b-90e1-087383f22840"), ShouldBeFalse)
	})
}

func TestMapStringRune_Get(t *testing.T) {
	Convey("TestMapStringRune.Get", t, func() {
		var k string = "973502e3-a2f9-4e1c-9114-5f8c803fee72"
		var v rune = 1009677704

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("8930dbca-7457-4f78-877a-2d4a559b0c1a" + "e6579a8f-9b0e-44d1-a83b-b4cbf6d19e6a")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringRune_GetOpt(t *testing.T) {
	Convey("TestMapStringRune.GetOpt", t, func() {
		var k string = "e0d2b03b-17ba-4743-8169-bc81dc616ac7"
		var v rune = 1947342465

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("18289137-c4dd-4e40-954f-6e4e87984ac4" + "99d5eac7-e8a7-4a9f-a93d-4a9b7d38eb7b")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringRune_ForEach(t *testing.T) {
	Convey("TestMapStringRune.ForEach", t, func() {
		var k string = "665a9454-fcb2-4a2a-9d6c-54df3275c4b1"
		var v rune = 1032374665
		hits := 0

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv rune) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringRune_MarshalYAML(t *testing.T) {
	Convey("TestMapStringRune.MarshalYAML", t, func() {
		var k string = "3fce6a2c-2810-47cb-9666-445ca8735fc8"
		var v rune = 1403586752

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringRune_ToYAML(t *testing.T) {
	Convey("TestMapStringRune.ToYAML", t, func() {
		var k string = "9d48ed15-20c0-4ec4-a92a-b7100df09d3d"
		var v rune = 1371657246

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringRune_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringRune.PutIfNotNil", t, func() {
		var k string = "c61b87e9-9bb8-4b8a-a076-d10334807c0d"
		var v rune = 1666367309

		test := omap.NewMapStringRune(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("b458fe2b-1a8f-4bf7-8c54-27f23db807ae", (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x rune = 1413077743
		So(test.PutIfNotNil("90c4159e-0803-4c18-8f19-e69c514ac5c4", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringRune_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringRune.ReplaceIfExists", t, func() {
		var k string = "4f621354-97b6-42a9-9e36-7bff0d000e33"
		var v rune = 1203445776
		var x rune = 731402340

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("e23ea9bb-7767-410b-be77-c7120d8b8b6c", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringRune_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringRune.ReplaceOrPut", t, func() {
		var k string = "c0835df9-e678-4d0e-82cb-137e7d19bc8f"
		var v rune = 1042657947
		var x rune = 1102658991

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("2e6bc274-8805-446d-bd9d-01f3bcea4012", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringRune_MarshalJSON(t *testing.T) {
	Convey("TestMapStringRune.MarshalJSON", t, func() {
		var k string = "e7404704-7539-48d6-a17f-2964712afce3"
		var v rune = 1459873248

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"e7404704-7539-48d6-a17f-2964712afce3","value":1459873248}]`)
	})
}
