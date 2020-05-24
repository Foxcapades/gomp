package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringInt16_Put(t *testing.T) {
	Convey("TestMapStringInt16.Put", t, func() {
		var k string = "870258bf-002f-414e-9e52-d26eb3aebf52"
		var v int16 = 30882

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringInt16_Delete(t *testing.T) {
	Convey("TestMapStringInt16.Delete", t, func() {
		var k string = "2f550dbd-7076-4346-9250-5d49317901d7"
		var v int16 = 10547

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringInt16_Has(t *testing.T) {
	Convey("TestMapStringInt16.Has", t, func() {
		var k string = "fb0ab061-3ceb-442d-8afc-6adaf1c2064b"
		var v int16 = 20139

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("f36f1ee0-6f13-47a5-b96d-987d7c47d3e0"+"5842aace-748a-4ee0-a1d3-1e5135bba31b"), ShouldBeFalse)
	})
}

func TestMapStringInt16_Get(t *testing.T) {
	Convey("TestMapStringInt16.Get", t, func() {
		var k string = "0f236863-b69b-4b76-b176-8264819d8b52"
		var v int16 = 31965

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("8f9c46e6-2f32-4cd6-93db-724184087969" + "693e6514-b0a6-4821-a37a-f9bc237fcbd7")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringInt16_GetOpt(t *testing.T) {
	Convey("TestMapStringInt16.GetOpt", t, func() {
		var k string = "e13d1670-d516-4a3a-9848-892463f7b3a4"
		var v int16 = 17661

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("f4b92ddd-65df-481b-b169-ea6db127eb7e" + "f664ab47-c156-4457-b600-5da495dc1074")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringInt16_ForEach(t *testing.T) {
	Convey("TestMapStringInt16.ForEach", t, func() {
		var k string = "9a851fd1-b59d-4aa9-9144-87aa92aa64d0"
		var v int16 = 25466
		hits := 0

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv int16) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringInt16_MarshalYAML(t *testing.T) {
	Convey("TestMapStringInt16.MarshalYAML", t, func() {
		var k string = "dd814b52-3c14-45e5-96e7-83ce09acf1af"
		var v int16 = 14871

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringInt16_ToYAML(t *testing.T) {
	Convey("TestMapStringInt16.ToYAML", t, func() {
		var k string = "9417ba85-7ff1-41a3-91b9-9b3344c69d8f"
		var v int16 = 20523

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringInt16_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringInt16.PutIfNotNil", t, func() {
		var k string = "f838ae00-f434-40da-9348-5180f36d22a4"
		var v int16 = 27684

		test := omap.NewMapStringInt16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("0dd1f5aa-2fc3-4993-bf22-aea342cea1e2", (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int16 = 30276
		So(test.PutIfNotNil("1f5b95d2-1955-4145-a00c-093b9239742b", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringInt16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringInt16.ReplaceIfExists", t, func() {
		var k string = "b0b215fa-c234-4123-8b72-ff1ffe9b9f2e"
		var v int16 = 24343
		var x int16 = 21768

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("ea32b4cc-97b2-44bc-85e5-501749111cdb", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringInt16.ReplaceOrPut", t, func() {
		var k string = "8d273121-2b49-40b0-9483-3f1a9f4e9599"
		var v int16 = 19413
		var x int16 = 15363

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("fd671d89-d0fd-49c3-8446-49c27afcd189", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt16_MarshalJSON(t *testing.T) {
	Convey("TestMapStringInt16.MarshalJSON", t, func() {
		var k string = "4f6c8a6c-1c64-480b-a6be-2d5607abe6e0"
		var v int16 = 10800

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"4f6c8a6c-1c64-480b-a6be-2d5607abe6e0","value":10800}]`)
	})
}
