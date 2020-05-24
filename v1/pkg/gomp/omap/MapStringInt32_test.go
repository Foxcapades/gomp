package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringInt32_Put(t *testing.T) {
	Convey("TestMapStringInt32.Put", t, func() {
		var k string = "9397284b-78c4-4a2c-ac8e-d4a933e4be1a"
		var v int32 = 1996051196

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringInt32_Delete(t *testing.T) {
	Convey("TestMapStringInt32.Delete", t, func() {
		var k string = "e98eb5cf-2a3a-4a02-a481-e1c871e014fc"
		var v int32 = 95823941

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringInt32_Has(t *testing.T) {
	Convey("TestMapStringInt32.Has", t, func() {
		var k string = "5653a79a-93db-4018-9a24-ad4ba4473102"
		var v int32 = 2024782935

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("ac209401-8580-45c6-a2db-991af1cbdbaf"+"f32988b2-ba59-49c3-8f2c-dcb01a439a1e"), ShouldBeFalse)
	})
}

func TestMapStringInt32_Get(t *testing.T) {
	Convey("TestMapStringInt32.Get", t, func() {
		var k string = "4cf681db-a269-44ec-8559-afd82b4fde44"
		var v int32 = 556865090

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("96313f4f-4f9a-47b2-a7d1-3ef5a24f25e6" + "c8c35f11-e608-4ec8-b245-a66630f4c7fb")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringInt32_GetOpt(t *testing.T) {
	Convey("TestMapStringInt32.GetOpt", t, func() {
		var k string = "f2a1c566-80c4-42a4-8800-fbf9e285ef17"
		var v int32 = 498141268

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("fd5149bd-3363-473b-a606-5b635a922174" + "ff076b1d-ac15-4251-ba2b-6b605bf422f1")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringInt32_ForEach(t *testing.T) {
	Convey("TestMapStringInt32.ForEach", t, func() {
		var k string = "42de6dbd-9542-42e7-9e50-68043b555d07"
		var v int32 = 250681120
		hits := 0

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv int32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringInt32_MarshalYAML(t *testing.T) {
	Convey("TestMapStringInt32.MarshalYAML", t, func() {
		var k string = "329782e7-a60e-493f-986b-39e621d6a50a"
		var v int32 = 1506328231

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringInt32_ToYAML(t *testing.T) {
	Convey("TestMapStringInt32.ToYAML", t, func() {
		var k string = "a1d62457-0d5f-4952-a354-584d897bc1e3"
		var v int32 = 44027599

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringInt32_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringInt32.PutIfNotNil", t, func() {
		var k string = "89ea0235-f86f-4812-b375-491fc10aa653"
		var v int32 = 224243951

		test := omap.NewMapStringInt32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("83736722-74fc-47b8-9612-b45a901e2fa1", (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int32 = 2079223844
		So(test.PutIfNotNil("cdd4d817-1bc4-4945-a8cb-f779129127e3", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringInt32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringInt32.ReplaceIfExists", t, func() {
		var k string = "1242f20b-8e8d-4f64-a617-11e19a5829a9"
		var v int32 = 1999564811
		var x int32 = 645218607

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("45daa41b-8734-44fe-98b5-ce36327315dd", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringInt32.ReplaceOrPut", t, func() {
		var k string = "7a9c4089-8d0e-44d7-a8b8-f463ffd588d2"
		var v int32 = 407501087
		var x int32 = 1751433949

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("cc3c3244-8a8d-4b68-9502-15ebe5d4dad8", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt32_MarshalJSON(t *testing.T) {
	Convey("TestMapStringInt32.MarshalJSON", t, func() {
		var k string = "9b6cf9c9-d0a9-4d46-8ae4-7fc9eeb71d73"
		var v int32 = 1091425522

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"9b6cf9c9-d0a9-4d46-8ae4-7fc9eeb71d73","value":1091425522}]`)
	})
}
