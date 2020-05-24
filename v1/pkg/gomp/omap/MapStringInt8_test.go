package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringInt8_Put(t *testing.T) {
	Convey("TestMapStringInt8.Put", t, func() {
		var k string = "8c097782-a573-4bc7-bc25-5b23e81a60f0"
		var v int8 = 39

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringInt8_Delete(t *testing.T) {
	Convey("TestMapStringInt8.Delete", t, func() {
		var k string = "07f0d2ab-c75b-4953-8c7b-ad84fbe5cdbb"
		var v int8 = 124

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringInt8_Has(t *testing.T) {
	Convey("TestMapStringInt8.Has", t, func() {
		var k string = "bb50a36e-8e3d-44ff-9040-7f3f4d4a0e40"
		var v int8 = 89

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("fefd221c-ff97-4446-8a61-b424730338d1"+"41094bab-3c91-40fb-b29f-2cc14fab7b12"), ShouldBeFalse)
	})
}

func TestMapStringInt8_Get(t *testing.T) {
	Convey("TestMapStringInt8.Get", t, func() {
		var k string = "56e2260e-f6aa-4f63-936a-f3a1f9de6dcb"
		var v int8 = 35

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("0ddc8079-866d-478e-8172-16b5cee78bda" + "133ebb47-f8ce-404f-9e88-feb26ca3284d")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringInt8_GetOpt(t *testing.T) {
	Convey("TestMapStringInt8.GetOpt", t, func() {
		var k string = "ffe1df26-5941-442c-b431-5f7108c6a0e5"
		var v int8 = 47

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("c211f822-6393-4485-af24-25681a149ff0" + "f95c06c8-789d-4d2c-b185-ce1f1e95df02")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringInt8_ForEach(t *testing.T) {
	Convey("TestMapStringInt8.ForEach", t, func() {
		var k string = "92d1115f-c3a2-480f-956a-a1acfd1a0701"
		var v int8 = 68
		hits := 0

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv int8) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringInt8_MarshalYAML(t *testing.T) {
	Convey("TestMapStringInt8.MarshalYAML", t, func() {
		var k string = "586ec358-d81d-4ad4-94b3-1ea7bfe044eb"
		var v int8 = 32

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringInt8_ToYAML(t *testing.T) {
	Convey("TestMapStringInt8.ToYAML", t, func() {
		var k string = "a4baabdd-6341-44ce-867f-b56b00826c64"
		var v int8 = 65

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringInt8_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringInt8.PutIfNotNil", t, func() {
		var k string = "bacfae13-16e1-46f7-b998-ed95e552068d"
		var v int8 = 31

		test := omap.NewMapStringInt8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("0e2ab75a-fdc1-442c-9bb0-6ae8ef53f9dd", (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int8 = 50
		So(test.PutIfNotNil("54759d0e-6001-4e05-8ca6-7aacc5d91ff8", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringInt8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringInt8.ReplaceIfExists", t, func() {
		var k string = "f84e9dc3-1d5e-49aa-8891-6c39e6beefd9"
		var v int8 = 57
		var x int8 = 29

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("1a82da6c-b4ed-4cee-a795-f3856cbf1b3a", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringInt8.ReplaceOrPut", t, func() {
		var k string = "b593ed5d-94d2-42f5-b0fa-a4ddba9df024"
		var v int8 = 19
		var x int8 = 62

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("bf7c97b7-bf86-45b4-a06c-728560dffd00", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt8_MarshalJSON(t *testing.T) {
	Convey("TestMapStringInt8.MarshalJSON", t, func() {
		var k string = "e4111251-bad5-4c8e-940c-69a300fc9264"
		var v int8 = 123

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"e4111251-bad5-4c8e-940c-69a300fc9264","value":123}]`)
	})
}
