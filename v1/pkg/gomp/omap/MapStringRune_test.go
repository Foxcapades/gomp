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
		var k string = "9ec6103d-8833-4f43-a006-edaf45c5c8f5"
		var v rune = 517649287

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringRune_Delete(t *testing.T) {
	Convey("TestMapStringRune.Delete", t, func() {
		var k string = "53a946bd-399f-4d36-ba9d-a7e17e940a40"
		var v rune = 1549940064

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringRune_Has(t *testing.T) {
	Convey("TestMapStringRune.Has", t, func() {
		var k string = "5988eabf-30a1-4f11-9fc0-9e3d91bfaefb"
		var v rune = 1247813399

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("2969f2dc-e50a-4cb4-b4d3-efafd758718d"+"d9ecbd64-88ab-44e3-a651-b0a0e795cbf5"), ShouldBeFalse)
	})
}


func TestMapStringRune_Get(t *testing.T) {
	Convey("TestMapStringRune.Get", t, func() {
		var k string = "ce04f8fc-0876-458d-9d1e-d38463d15f05"
		var v rune = 986728103

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("f457d2e4-a7aa-40c2-84a9-5e555425b9f3"+"91f28dbf-e25e-44a2-a362-ed28a2330a43")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringRune_GetOpt(t *testing.T) {
	Convey("TestMapStringRune.GetOpt", t, func() {
		var k string = "8860faca-6c76-4fca-ab45-0b23c82de009"
		var v rune = 1089559628

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("197dd704-fed5-459e-b809-b7313671d0b0"+"761beab8-6e4d-4375-8ac9-d61e6004f674")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringRune_ForEach(t *testing.T) {
	Convey("TestMapStringRune.ForEach", t, func() {
		var k string = "192049d2-014f-4c9e-b968-c5549a0ff5e2"
		var v rune = 380272228
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
		var k string = "91754da9-7c05-4dd4-a1c1-0ecf899090d1"
		var v rune = 1969486313

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
		var k string = "f8e1eda9-8491-4048-b0f8-faaabe41a276"
		var v rune = 2017730898

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
		var k string = "43ebcb6b-f899-459d-bb37-944642e581a3"
		var v rune = 1087891930

		test := omap.NewMapStringRune(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("4893bc14-e10b-455a-a2bd-4a612c00f1f3", (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x rune = 808449631
		So(test.PutIfNotNil("ae903d13-c2b1-4fcb-bf4b-84998256e09a", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringRune_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringRune.ReplaceIfExists", t, func() {
		var k string = "56c0c9ca-631b-42c2-a9f7-89c1d54d4d6e"
		var v rune = 107283001
		var x rune = 1486795035

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("949c1261-45d5-4e2a-86f1-d88e4f946eb3", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringRune_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringRune.ReplaceOrPut", t, func() {
		var k string = "713c8b81-91cc-4919-a2b3-ba37ff1d4cd9"
		var v rune = 740188922
		var x rune = 1296574855

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("a88bc037-8d0b-4fd9-a3d8-c8ab48ea7a67", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringRune_MarshalJSON(t *testing.T) {
	Convey("TestMapStringRune.MarshalJSON", t, func() {
		var k string = "7b22848c-b1d9-4cd6-986f-afff28784222"
		var v rune = 1840723318

		test := omap.NewMapStringRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"7b22848c-b1d9-4cd6-986f-afff28784222","value":1840723318}]`)
	})
}

