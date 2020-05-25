package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyInt8_Put(t *testing.T) {
	Convey("TestMapAnyInt8.Put", t, func() {
		var k interface{} = "b35f2c7b-17f3-4721-b2f0-ea65f8ed767b"
		var v int8 = 125

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyInt8_Delete(t *testing.T) {
	Convey("TestMapAnyInt8.Delete", t, func() {
		var k interface{} = "1a931fe4-bc39-43e6-9685-dc85c73b293e"
		var v int8 = 83

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyInt8_Has(t *testing.T) {
	Convey("TestMapAnyInt8.Has", t, func() {
		var k interface{} = "b20e63a9-61af-4fb2-a34f-3be152bc9ea3"
		var v int8 = 60

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("1b587129-056c-4595-a407-56b29808c394"+"d65f6709-95ac-4e3b-b3bd-be4e81d47aa9"), ShouldBeFalse)
	})
}

func TestMapAnyInt8_Get(t *testing.T) {
	Convey("TestMapAnyInt8.Get", t, func() {
		var k interface{} = "544152ee-10f5-4e35-b8eb-009b46d60109"
		var v int8 = 92

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("5f234ba3-75d5-40cd-98e3-bbace51576d1" + "6083d28b-534e-47b8-825f-d6696d0930fa")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyInt8_GetOpt(t *testing.T) {
	Convey("TestMapAnyInt8.GetOpt", t, func() {
		var k interface{} = "f095ed03-f195-43e9-9b09-8d8f0f3d97f6"
		var v int8 = 75

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("b7f0b8e8-0d30-48e3-9481-92afb06bfe06" + "7769e62f-b538-4dbf-9034-01ff60c67990")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyInt8_ForEach(t *testing.T) {
	Convey("TestMapAnyInt8.ForEach", t, func() {
		var k interface{} = "071ea8a4-ce8e-407a-93e3-8f499d175601"
		var v int8 = 6
		hits := 0

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv int8) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyInt8_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyInt8.MarshalYAML", t, func() {
		var k interface{} = "03b475a6-8a8f-49b0-8084-86f82c15de94"
		var v int8 = 89

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyInt8_ToYAML(t *testing.T) {
	Convey("TestMapAnyInt8.ToYAML", t, func() {
		var k interface{} = "d6c7eec9-62f8-48ed-83b2-4c1b289c2826"
		var v int8 = 41

		test := omap.NewMapAnyInt8(1)

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

func TestMapAnyInt8_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyInt8.PutIfNotNil", t, func() {
		var k interface{} = "f1194d58-2318-4ca5-ac5b-ab2f8eb2ec7e"
		var v int8 = 66

		test := omap.NewMapAnyInt8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("bc9a01f4-fcd9-48d4-933b-6f9e4db343d9", (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int8 = 92
		So(test.PutIfNotNil("e17f98cc-b3b1-421b-945f-787fe0921d88", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyInt8.ReplaceIfExists", t, func() {
		var k interface{} = "67875ff0-3dd3-40dd-99e7-40189303ce45"
		var v int8 = 37
		var x int8 = 116

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("ebf99f78-60b5-4dd0-b307-a18f814fad8b", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyInt8.ReplaceOrPut", t, func() {
		var k interface{} = "91737b60-a266-45cd-be28-d711d2084597"
		var v int8 = 87
		var x int8 = 14

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("e91f45b7-049d-4f3c-a139-5d91ec46085e", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt8_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyInt8.MarshalJSON", t, func() {
		var k interface{} = "0b4cea4f-affc-4af8-807e-f8f37be95e5c"
		var v int8 = 107

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"0b4cea4f-affc-4af8-807e-f8f37be95e5c","value":107}]`)
	})
}
