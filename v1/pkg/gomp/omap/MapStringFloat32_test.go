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
		var k string = "fa239546-8711-4a9b-b139-74cf94272ddc"
		var v float32 = 0.707

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringFloat32_Delete(t *testing.T) {
	Convey("TestMapStringFloat32.Delete", t, func() {
		var k string = "ad1565ab-1b70-40b3-ad50-a1194efc0a7f"
		var v float32 = 0.398

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringFloat32_Has(t *testing.T) {
	Convey("TestMapStringFloat32.Has", t, func() {
		var k string = "e4de88fa-1ed3-4033-b468-fe1a2278b582"
		var v float32 = 0.471

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("9227ab80-fb6f-4bbb-bfdb-5964bd5883c6"+"6754a700-2f6f-463a-8af4-0688a464e8f7"), ShouldBeFalse)
	})
}


func TestMapStringFloat32_Get(t *testing.T) {
	Convey("TestMapStringFloat32.Get", t, func() {
		var k string = "002f04fc-1ad4-4ff1-8b5b-a494d41611b5"
		var v float32 = 0.655

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("165fe328-c747-4e8c-a570-3f78eada06c1"+"9efa704b-bed7-4a9e-82fc-35bac9ba6b07")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringFloat32_GetOpt(t *testing.T) {
	Convey("TestMapStringFloat32.GetOpt", t, func() {
		var k string = "4bfa513a-24eb-48bd-8343-c33f165018df"
		var v float32 = 0.426

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("d520a0aa-080b-47a1-8887-fe6ecdf08745"+"feebe0f1-e027-49b3-a73b-23e7f50de768")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringFloat32_ForEach(t *testing.T) {
	Convey("TestMapStringFloat32.ForEach", t, func() {
		var k string = "4138f726-d7c1-4415-b5d1-9c46551539c1"
		var v float32 = 0.576
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
		var k string = "58eeed05-1656-485f-8d5e-32d630f37720"
		var v float32 = 0.323

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
		var k string = "b2b6cbd2-4003-429c-abeb-278bd5811d8c"
		var v float32 = 0.316

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
		var k string = "8d66d10a-b833-4ed3-9754-946ee818907d"
		var v float32 = 0.222

		test := omap.NewMapStringFloat32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("fd1ac82e-d039-4345-8562-a3b8703c6d29", (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float32 = 0.908
		So(test.PutIfNotNil("acdd5a38-10c1-43aa-ac11-8b99c1281e28", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringFloat32.ReplaceIfExists", t, func() {
		var k string = "a63538cc-8227-4ace-8d6c-9efcea0a1a9d"
		var v float32 = 0.701
		var x float32 = 0.437

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("ce5088b1-58b5-46a5-98ff-9dde8c6841ea", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringFloat32.ReplaceOrPut", t, func() {
		var k string = "76916c59-1c69-4422-bc99-28b210faf607"
		var v float32 = 0.340
		var x float32 = 0.411

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("1382a084-f39c-4b76-9f4c-4c2bce6dae38", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat32_MarshalJSON(t *testing.T) {
	Convey("TestMapStringFloat32.MarshalJSON", t, func() {
		var k string = "79279390-f88c-4fe6-9374-d45adbf8bf50"
		var v float32 = 0.424

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"79279390-f88c-4fe6-9374-d45adbf8bf50","value":0.424}]`)
	})
}

