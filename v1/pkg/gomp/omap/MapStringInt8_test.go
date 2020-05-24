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
		var k string = "47cca687-48c0-4509-b0d6-970bbfab3b19"
		var v int8 = 125

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringInt8_Delete(t *testing.T) {
	Convey("TestMapStringInt8.Delete", t, func() {
		var k string = "9e7266f0-5303-428d-afc9-9873d1021b6b"
		var v int8 = 18

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringInt8_Has(t *testing.T) {
	Convey("TestMapStringInt8.Has", t, func() {
		var k string = "a86c3218-212a-43c4-907d-d37d26877804"
		var v int8 = 4

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("02f0d853-f7ec-4035-9a5b-2cdb80c41f1e"+"eb0f8619-2352-42a7-aa71-2d7546be627a"), ShouldBeFalse)
	})
}


func TestMapStringInt8_Get(t *testing.T) {
	Convey("TestMapStringInt8.Get", t, func() {
		var k string = "d64f693e-07d8-4a26-aedb-3ff9e1753094"
		var v int8 = 78

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("a4457c20-1e4c-4578-8543-aa7eff5392ce"+"45d7032d-78b3-4be0-a78b-cfc5f5ad8bc2")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringInt8_GetOpt(t *testing.T) {
	Convey("TestMapStringInt8.GetOpt", t, func() {
		var k string = "f3201628-2f34-4ead-a31e-9ae644766881"
		var v int8 = 23

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("267cb9f6-2969-47e3-83cc-5e5f8cb984d3"+"05d99527-8aa1-4d5c-99fa-69a92effdea9")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringInt8_ForEach(t *testing.T) {
	Convey("TestMapStringInt8.ForEach", t, func() {
		var k string = "62a1545a-45fb-4883-b714-e62e53a9824f"
		var v int8 = 75
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
		var k string = "a5ddcf02-30cb-465b-a3f1-f56bcc8f47cf"
		var v int8 = 40

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
		var k string = "cf930ab1-881f-4cd5-b5eb-9caeaabe08f0"
		var v int8 = 57

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
		var k string = "bbd34543-7702-409a-928a-0c1a60f2e285"
		var v int8 = 29

		test := omap.NewMapStringInt8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("03f0bb86-8098-47bb-986e-72154c4416a1", (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int8 = 24
		So(test.PutIfNotNil("51c56231-e64a-4a72-8f45-92695dae59b7", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringInt8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringInt8.ReplaceIfExists", t, func() {
		var k string = "f6d1f1c1-6a38-48c8-82c2-294425e1a95e"
		var v int8 = 49
		var x int8 = 32

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("cc0ad233-badd-46a2-ac9d-7b8eccd2d44a", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringInt8.ReplaceOrPut", t, func() {
		var k string = "13fbf577-834e-429c-86c2-4861a6fb0a02"
		var v int8 = 55
		var x int8 = 124

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("c7b9e6b2-24be-4367-a224-7271c65db292", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt8_MarshalJSON(t *testing.T) {
	Convey("TestMapStringInt8.MarshalJSON", t, func() {
		var k string = "b0738c5c-1c1c-408c-8e2d-2e079e06faff"
		var v int8 = 4

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"b0738c5c-1c1c-408c-8e2d-2e079e06faff","value":4}]`)
	})
}

