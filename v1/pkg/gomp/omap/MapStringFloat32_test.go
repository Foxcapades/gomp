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
		var k string = "5bc21c5f-2651-4bc7-9f61-6f0734772fc8"
		var v float32 = 0.735

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringFloat32_Delete(t *testing.T) {
	Convey("TestMapStringFloat32.Delete", t, func() {
		var k string = "8fe00cdf-db85-41ac-a513-acea28479bd6"
		var v float32 = 0.119

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringFloat32_Has(t *testing.T) {
	Convey("TestMapStringFloat32.Has", t, func() {
		var k string = "a659f7c1-1b76-4089-9b8e-754d07d381bd"
		var v float32 = 0.245

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("c7ededa8-e8c3-4ad3-8e69-30a51d86c10f"+"ca65d83d-1e4f-4d58-88b8-d22a9be158cb"), ShouldBeFalse)
	})
}


func TestMapStringFloat32_Get(t *testing.T) {
	Convey("TestMapStringFloat32.Get", t, func() {
		var k string = "ad30bcc8-2e11-4203-bea1-4016e53e46f0"
		var v float32 = 0.559

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("4122f8ce-8a92-460c-9634-4e940a7aa999" + "b4524da5-d84b-4e37-b1c8-2c9809077293")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringFloat32_GetOpt(t *testing.T) {
	Convey("TestMapStringFloat32.GetOpt", t, func() {
		var k string = "ca386315-3d84-4862-939f-5f344b5b82af"
		var v float32 = 0.936

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("0b0023d8-4873-483d-b2c7-ffaaef9d8e72" + "019ac540-0404-4a14-87ee-32d2d2b2a3c6")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringFloat32_ForEach(t *testing.T) {
	Convey("TestMapStringFloat32.ForEach", t, func() {
		var k string = "51322fb4-1283-45ff-9f72-55e2bff2e275"
		var v float32 = 0.099
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
		var k string = "3721e3f7-bc8d-445a-8ff2-120566b36af8"
		var v float32 = 0.080

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
		var k string = "6a2c3a4b-5c49-4621-9f8a-9fb78f3069f9"
		var v float32 = 0.695

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringFloat32_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringFloat32.PutIfNotNil", t, func() {
		var k string = "b9e0567a-30f6-42bc-82cb-c43bc9f8e61b"
		var v float32 = 0.111

		test := omap.NewMapStringFloat32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("6bc9b09b-43f3-41a2-b486-ae15e9b98b24", (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float32 = 0.277
		So(test.PutIfNotNil("06a1b114-563a-4ab2-940f-83947d4a63a5", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringFloat32.ReplaceIfExists", t, func() {
		var k string = "884060b7-ee6b-4e41-8bbc-c8b2fdb771af"
		var v float32 = 0.373
		var x float32 = 0.761

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("b51396d3-bbf2-4feb-829f-13153c8f7107", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringFloat32.ReplaceOrPut", t, func() {
		var k string = "1f2c2717-7d45-43f1-a22c-6a9c4297ef55"
		var v float32 = 0.110
		var x float32 = 0.957

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("d21f4447-7521-4425-afa8-091021e211ac", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat32_MarshalJSON(t *testing.T) {
	Convey("TestMapStringFloat32.MarshalJSON", t, func() {
		var k string = "7bd49e35-dd18-4a55-8139-7ec37592b5d5"
		var v float32 = 0.100

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"7bd49e35-dd18-4a55-8139-7ec37592b5d5","value":0.1}]`)
	})
}
