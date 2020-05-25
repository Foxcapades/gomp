package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringFloat32_Put(t *testing.T) {
	Convey("TestMapStringFloat32.Put", t, func() {
		var k string = "98fb1975-21fb-4455-bd92-5902c285aa3c"
		var v float32 = 0.924

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringFloat32_Delete(t *testing.T) {
	Convey("TestMapStringFloat32.Delete", t, func() {
		var k string = "c591c4be-b3b0-46c9-ae36-30f2463e24c6"
		var v float32 = 0.344

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringFloat32_Has(t *testing.T) {
	Convey("TestMapStringFloat32.Has", t, func() {
		var k string = "1e97f11b-a7d8-4f0a-89b7-63a340e45c1b"
		var v float32 = 0.530

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("aaeb66c3-4054-4cfc-8f44-04f0c4ad3261"+"e4a44b38-eef5-4237-841e-47d181db04f1"), ShouldBeFalse)
	})
}

func TestMapStringFloat32_Get(t *testing.T) {
	Convey("TestMapStringFloat32.Get", t, func() {
		var k string = "42724b2f-cb4c-4353-ba41-e411f9cfa614"
		var v float32 = 0.984

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("3e4bcd57-25e8-4490-bae7-e6bc6a620d29" + "e3f857dc-3db0-4c6d-aaa0-31fef00cd330")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringFloat32_GetOpt(t *testing.T) {
	Convey("TestMapStringFloat32.GetOpt", t, func() {
		var k string = "ce9f8c6e-6326-496d-b200-ab954fe83dba"
		var v float32 = 0.267

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("8eb44946-e2fc-46e1-8e74-7342ce93d619" + "40431933-8f28-4f3f-9827-def3328304ae")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringFloat32_ForEach(t *testing.T) {
	Convey("TestMapStringFloat32.ForEach", t, func() {
		var k string = "9750bd47-b8c6-48e6-9429-4657a82c50e5"
		var v float32 = 0.320
		hits := 0

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv float32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringFloat32_MarshalYAML(t *testing.T) {
	Convey("TestMapStringFloat32.MarshalYAML", t, func() {
		var k string = "7fbcc883-703d-4618-b9ee-a52e7f5a955b"
		var v float32 = 0.628

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringFloat32_ToYAML(t *testing.T) {
	Convey("TestMapStringFloat32.ToYAML", t, func() {
		var k string = "1ed47c71-2f6b-4829-ba6f-e9b6a6d94170"
		var v float32 = 0.773

		test := omap.NewMapStringFloat32(1)

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

func TestMapStringFloat32_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringFloat32.PutIfNotNil", t, func() {
		var k string = "cb8d535f-aec3-4770-a6d2-0c92e01637cd"
		var v float32 = 0.800

		test := omap.NewMapStringFloat32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("5b4d296a-aa12-410c-bdaf-cfdd58239106", (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float32 = 0.036
		So(test.PutIfNotNil("1352371f-df32-44a4-bf2d-22c29116fd64", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringFloat32.ReplaceIfExists", t, func() {
		var k string = "672741f5-9d94-49fb-b6ba-8dc857eacc07"
		var v float32 = 0.654
		var x float32 = 0.369

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("f0a49df2-4e8f-4f12-a78c-70a899b64850", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringFloat32.ReplaceOrPut", t, func() {
		var k string = "cc210a9a-4a3c-45c6-8591-bef58d351261"
		var v float32 = 0.210
		var x float32 = 0.041

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("58c221be-4627-478e-9410-db8cd9ef709e", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat32_MarshalJSON(t *testing.T) {
	Convey("TestMapStringFloat32.MarshalJSON", t, func() {
		var k string = "7ea4a7eb-fb1b-4416-ae97-f5ce07076397"
		var v float32 = 0.024

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"7ea4a7eb-fb1b-4416-ae97-f5ce07076397","value":0.024}]`)
	})
}
