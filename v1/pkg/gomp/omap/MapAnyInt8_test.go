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
		var k interface{} = "c201c9d8-4494-4349-b386-342ffe1c1b95"
		var v int8 = 104

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyInt8_Delete(t *testing.T) {
	Convey("TestMapAnyInt8.Delete", t, func() {
		var k interface{} = "f2294f44-e8b8-4a06-a2dc-bc2a878241f7"
		var v int8 = 20

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyInt8_Has(t *testing.T) {
	Convey("TestMapAnyInt8.Has", t, func() {
		var k interface{} = "4042fd60-20bc-49f5-a3dd-915503cf2cda"
		var v int8 = 125

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("6392cf63-2535-4088-91e6-a5a52908e0d2"+"af55053b-aea7-4eaa-a1bf-e0afe6c00fb9"), ShouldBeFalse)
	})
}

func TestMapAnyInt8_Get(t *testing.T) {
	Convey("TestMapAnyInt8.Get", t, func() {
		var k interface{} = "b31f5c11-d34d-4da1-affc-0a9e9ad5f481"
		var v int8 = 18

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("057aa108-6732-4e2e-8e56-b9fb6d206967" + "7cd90f67-98ea-455e-bf36-6a9d96a2a73d")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyInt8_GetOpt(t *testing.T) {
	Convey("TestMapAnyInt8.GetOpt", t, func() {
		var k interface{} = "f271c382-018d-42a7-98d1-627cb305b5d1"
		var v int8 = 68

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("3738b08d-50fa-4e30-8456-21d86b4f490c" + "dac4b6d5-0d94-44c3-8d11-97e038f1ee32")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyInt8_ForEach(t *testing.T) {
	Convey("TestMapAnyInt8.ForEach", t, func() {
		var k interface{} = "4e762d3a-4039-475d-be62-7f9f09cf7bdd"
		var v int8 = 123
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
		var k interface{} = "2ffa6424-2993-4529-b288-873bc6647b52"
		var v int8 = 115

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
		var k interface{} = "cdab454d-9643-4bbe-9149-3db0a375bfc8"
		var v int8 = 91

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyInt8_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyInt8.PutIfNotNil", t, func() {
		var k interface{} = "9e36a003-6ccf-4acc-aad0-679755eadf96"
		var v int8 = 5

		test := omap.NewMapAnyInt8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("eb784e11-66d4-47ff-9031-16b38b145454", (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int8 = 7
		So(test.PutIfNotNil("a291738d-c65e-4f96-a306-07b7b5abc11a", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyInt8.ReplaceIfExists", t, func() {
		var k interface{} = "c2408bac-0ec8-4400-a425-f754b946c81d"
		var v int8 = 122
		var x int8 = 65

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("0bab10d2-812b-491f-84ec-e174228ccfbf", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyInt8.ReplaceOrPut", t, func() {
		var k interface{} = "aaa9b6dc-bd63-4d42-bd7e-d192d0b8f40d"
		var v int8 = 47
		var x int8 = 91

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("fdfd6cb2-12a0-424e-9a11-2b604da0fa52", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt8_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyInt8.MarshalJSON", t, func() {
		var k interface{} = "14fc8b1e-dd33-45cf-9a82-71db7b5bcce1"
		var v int8 = 16

		test := omap.NewMapAnyInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"14fc8b1e-dd33-45cf-9a82-71db7b5bcce1","value":16}]`)
	})
}
