package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringByte_Put(t *testing.T) {
	Convey("TestMapStringByte.Put", t, func() {
		var k string = "ad2f5341-e157-4ac4-a658-6b2189603fdf"
		var v byte = 134

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringByte_Delete(t *testing.T) {
	Convey("TestMapStringByte.Delete", t, func() {
		var k string = "c3c851f4-929f-4a62-89a0-2aa819710134"
		var v byte = 214

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringByte_Has(t *testing.T) {
	Convey("TestMapStringByte.Has", t, func() {
		var k string = "c6e4ce5c-4881-448f-baf2-4f7731bf5215"
		var v byte = 19

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("c40cb381-8865-412a-aa69-d0bbecffd3c7"+"d4d0ec6b-87ab-4700-9969-1b8b2fc72719"), ShouldBeFalse)
	})
}

func TestMapStringByte_Get(t *testing.T) {
	Convey("TestMapStringByte.Get", t, func() {
		var k string = "780fefb6-5bfa-4e47-8663-bdf82ad6b28b"
		var v byte = 160

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("784c69bb-60b8-44fc-bf7d-daabfc3a9383" + "7fb3b301-2d97-469f-978a-a593f96e3753")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringByte_GetOpt(t *testing.T) {
	Convey("TestMapStringByte.GetOpt", t, func() {
		var k string = "b9c5eec4-c5a6-4463-ae43-a6de01e8c18f"
		var v byte = 97

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("5f32b54d-103e-4c6a-b088-5574a977fcbc" + "398ce747-ffb1-4e53-bcab-51c5d6a08a1d")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringByte_ForEach(t *testing.T) {
	Convey("TestMapStringByte.ForEach", t, func() {
		var k string = "bd6be56e-f40b-4a0a-bb68-e0a74c9acf9e"
		var v byte = 2
		hits := 0

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv byte) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringByte_MarshalYAML(t *testing.T) {
	Convey("TestMapStringByte.MarshalYAML", t, func() {
		var k string = "087a45c6-f93b-4b51-aeac-aadaaa8bf4da"
		var v byte = 140

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringByte_ToYAML(t *testing.T) {
	Convey("TestMapStringByte.ToYAML", t, func() {
		var k string = "c96387e9-302c-43ff-9aad-44d187315900"
		var v byte = 5

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringByte_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringByte.PutIfNotNil", t, func() {
		var k string = "e62c2f26-be5d-4f55-a4fa-ec8a259fd04b"
		var v byte = 10

		test := omap.NewMapStringByte(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("00a57612-6928-4b28-9627-e847095dd10e", (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x byte = 219
		So(test.PutIfNotNil("407c5682-8016-4289-adc1-ce4b386f11e6", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringByte_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringByte.ReplaceIfExists", t, func() {
		var k string = "e9f16a66-5ff8-4634-9022-c69a7a2a48a3"
		var v byte = 249
		var x byte = 220

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("9498b597-b7b2-4a5d-8bc8-12ddc8b091cd", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringByte_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringByte.ReplaceOrPut", t, func() {
		var k string = "904c70bd-516d-4e89-a529-8dabaabc7e6a"
		var v byte = 39
		var x byte = 123

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("3b38e7c3-fb9c-4b18-af32-445f85e6b1c7", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringByte_MarshalJSON(t *testing.T) {
	Convey("TestMapStringByte.MarshalJSON", t, func() {
		var k string = "32def4ef-7781-4217-8757-f20a20bd68b5"
		var v byte = 40

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"32def4ef-7781-4217-8757-f20a20bd68b5","value":40}]`)
	})
}
