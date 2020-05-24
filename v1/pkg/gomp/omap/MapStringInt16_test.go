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
		var k string = "23079bb5-7aaf-46e4-8fb3-1d73193d3cae"
		var v int16 = 5226

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringInt16_Delete(t *testing.T) {
	Convey("TestMapStringInt16.Delete", t, func() {
		var k string = "39bbdbe8-9802-4607-ae39-cbd5e3d51719"
		var v int16 = 1060

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringInt16_Has(t *testing.T) {
	Convey("TestMapStringInt16.Has", t, func() {
		var k string = "cdcb38f6-6ca4-4f0e-8f3c-f2929eb77f60"
		var v int16 = 27197

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("3306579b-9ec8-418f-8fc5-828a1398661f"+"e8678bd1-8978-4a1f-9dd2-e13d3d693e76"), ShouldBeFalse)
	})
}


func TestMapStringInt16_Get(t *testing.T) {
	Convey("TestMapStringInt16.Get", t, func() {
		var k string = "69202f79-2908-4935-bb96-7d51fcbfd15a"
		var v int16 = 26917

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("072fc075-6b68-4d85-9d14-543fcd32acfa" + "23e92b7f-7331-4e0b-9bb7-7f1b31e43e34")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringInt16_GetOpt(t *testing.T) {
	Convey("TestMapStringInt16.GetOpt", t, func() {
		var k string = "674efcce-2035-4eff-8b07-4449dc3976b6"
		var v int16 = 13267

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("6ac5eb15-07d2-4263-a435-e068b27359d9" + "ed398f0c-413f-4f12-bf64-f4e61c739ab5")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringInt16_ForEach(t *testing.T) {
	Convey("TestMapStringInt16.ForEach", t, func() {
		var k string = "42aebee8-32b7-423e-93ba-0da86c2dd60f"
		var v int16 = 29604
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
		var k string = "f573071b-10f7-4ce1-81a7-6f3eedc284b6"
		var v int16 = 6799

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
		var k string = "77ed900e-0450-4c10-a1a8-68ffc62cb5f4"
		var v int16 = 20844

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
		var k string = "75125e80-3c76-417a-9e6f-3fc2440e405d"
		var v int16 = 20314

		test := omap.NewMapStringInt16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("b994ef17-307f-4fe7-9d03-875ea94d2043", (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int16 = 20404
		So(test.PutIfNotNil("8cbd5b20-b14c-4217-93a1-8884c7ef8106", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringInt16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringInt16.ReplaceIfExists", t, func() {
		var k string = "5a8211f5-d002-45e5-95ed-a58b3bde1bd4"
		var v int16 = 10617
		var x int16 = 642

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("ef523efc-4571-44b8-a7bd-7b6babee86f1", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringInt16.ReplaceOrPut", t, func() {
		var k string = "eb15fa01-acad-4697-8e3d-4a7d5b1a90cb"
		var v int16 = 11663
		var x int16 = 21854

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("137208f9-2e7e-4e40-b20f-c12da18d22e9", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt16_MarshalJSON(t *testing.T) {
	Convey("TestMapStringInt16.MarshalJSON", t, func() {
		var k string = "9371853f-f30b-4ae1-a783-042fc785a9f3"
		var v int16 = 9110

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"9371853f-f30b-4ae1-a783-042fc785a9f3","value":9110}]`)
	})
}
