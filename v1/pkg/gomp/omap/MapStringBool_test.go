package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringBool_Put(t *testing.T) {
	Convey("TestMapStringBool.Put", t, func() {
		var k string = "5ebd0c77-cd3f-4c17-bd79-f77ba84a2f52"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringBool_Delete(t *testing.T) {
	Convey("TestMapStringBool.Delete", t, func() {
		var k string = "c6115edc-e866-4c91-9a16-bd9747f09909"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringBool_Has(t *testing.T) {
	Convey("TestMapStringBool.Has", t, func() {
		var k string = "a9ae204b-bbb8-4b05-9b10-843a24f987f2"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("26ea40c8-b374-4333-b938-c74a68cfbe34"+"ab13862c-5051-437d-988b-de713ad51a7d"), ShouldBeFalse)
	})
}


func TestMapStringBool_Get(t *testing.T) {
	Convey("TestMapStringBool.Get", t, func() {
		var k string = "61f30e53-78d6-48f4-98dc-0396f0ad5e4a"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("884ae6bc-edfa-4346-b969-3d7648e1cc59" + "f060e784-02a2-4bb8-8550-89cf2dea8ae5")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringBool_GetOpt(t *testing.T) {
	Convey("TestMapStringBool.GetOpt", t, func() {
		var k string = "e5e485f0-7652-4daf-b067-36fe00bcc294"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("5593c879-9927-4146-80f9-4d033f44f98b" + "8b307c0e-1d22-4f46-b508-88925c87cc2e")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringBool_ForEach(t *testing.T) {
	Convey("TestMapStringBool.ForEach", t, func() {
		var k string = "0dc5800c-6c33-4035-bf21-d036005c174b"
		var v bool = false
		hits := 0

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv bool) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringBool_MarshalYAML(t *testing.T) {
	Convey("TestMapStringBool.MarshalYAML", t, func() {
		var k string = "c9fdf6b2-78f2-4ec0-9cbe-f706e4edaf21"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringBool_ToYAML(t *testing.T) {
	Convey("TestMapStringBool.ToYAML", t, func() {
		var k string = "c5bdddd7-ba6f-42ab-9566-0b30c8e2c5fc"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringBool_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringBool.PutIfNotNil", t, func() {
		var k string = "df79e81e-f17a-4ada-ad0e-ba0a153e9f5f"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*bool)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("b6823368-003d-4317-89da-80fd0fad6ae3", (*bool)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x bool = false
		So(test.PutIfNotNil("0d1197ed-3a45-40d0-ba0e-3a4c92c44aac", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringBool_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringBool.ReplaceIfExists", t, func() {
		var k string = "c8ecbb4a-15a2-49f8-b5c4-578997ce9bb7"
		var v bool = false
		var x bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("82a88edf-d0b4-4712-9063-d4ab4a073c11", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringBool_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringBool.ReplaceOrPut", t, func() {
		var k string = "80c062ae-c2b2-4da5-a9a5-1dc0674739db"
		var v bool = false
		var x bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("e8fb5863-917a-456e-b030-c4c519167e6f", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringBool_MarshalJSON(t *testing.T) {
	Convey("TestMapStringBool.MarshalJSON", t, func() {
		var k string = "9a4d87aa-9e10-445e-a909-e0551c6f681c"
		var v bool = false

		test := omap.NewMapStringBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"9a4d87aa-9e10-445e-a909-e0551c6f681c","value":false}]`)
	})
}

