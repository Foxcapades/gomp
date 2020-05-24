package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAny_Put(t *testing.T) {
	Convey("TestMapAny.Put", t, func() {
		var k interface{} = "35f6afb1-9bc6-4f1e-885a-516b9083057c"
		var v interface{} = "7bebc909-5042-423b-84fb-b88d6c117cc8"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAny_Delete(t *testing.T) {
	Convey("TestMapAny.Delete", t, func() {
		var k interface{} = "5728487d-a3fc-4053-a909-3432a9ffdb53"
		var v interface{} = "29692a48-5d51-40ce-9922-15191f88a141"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAny_Has(t *testing.T) {
	Convey("TestMapAny.Has", t, func() {
		var k interface{} = "f28e33bc-e8bc-40f8-afa8-cb8ddf353cd7"
		var v interface{} = "1c738638-7060-42a7-8379-7304cfe5ace7"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("504b5177-0fb5-4666-9b79-beba45018f07"+"9a7a44f9-e9ec-4f45-a719-e34322abc5e2"), ShouldBeFalse)
	})
}


func TestMapAny_Get(t *testing.T) {
	Convey("TestMapAny.Get", t, func() {
		var k interface{} = "2190ad21-1469-4c4f-a308-d8caea5607dc"
		var v interface{} = "679f91a3-6704-4aee-936d-22333fa84da9"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("df6a6d51-d5f4-4537-9901-4b96a6a8c7f2" + "19e47827-9180-43dc-b306-df0bd21b961c")
		So(b, ShouldBeFalse)
	})
}

func TestMapAny_GetOpt(t *testing.T) {
	Convey("TestMapAny.GetOpt", t, func() {
		var k interface{} = "f17ed68c-7332-4536-81c6-874ad15ee614"
		var v interface{} = "1277d4c0-3ac6-4a6f-bb42-c0b740a63dd1"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("99a70a8c-a95d-4aa3-9cb5-5e25785bccad" + "9151483b-7b99-4211-8cd9-67c25c114fa6")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAny_ForEach(t *testing.T) {
	Convey("TestMapAny.ForEach", t, func() {
		var k interface{} = "594738c4-cf4c-4705-80ba-9473145c7c34"
		var v interface{} = "eb1ebafa-455f-4b7e-8bc3-1fda6be8fe9a"
		hits := 0

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv interface{}) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAny_MarshalYAML(t *testing.T) {
	Convey("TestMapAny.MarshalYAML", t, func() {
		var k interface{} = "63fd5973-3507-4759-ad16-d8c5c804236f"
		var v interface{} = "5db16a26-60e6-4af1-ad24-295939636ee8"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAny_ToYAML(t *testing.T) {
	Convey("TestMapAny.ToYAML", t, func() {
		var k interface{} = "2b2bb874-decf-441f-a48a-ad0f5de156f6"
		var v interface{} = "937d780e-d054-44c4-ac5c-c3f16109d499"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAny_PutIfNotNil(t *testing.T) {
	Convey("TestMapAny.PutIfNotNil", t, func() {
		var k interface{} = "43e0b166-c0cd-4eff-b781-e12a1094683f"
		var v interface{} = "ee6de792-1678-4765-b345-cef8680f781b"

		test := omap.NewMapAny(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("1e9dcd07-792b-4254-9174-8fa3c59e3b51", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "51d5f6ff-7b42-4f70-b40f-51f97a634347"
		So(test.PutIfNotNil("52e8a0e4-1da7-460f-abfc-629ea2342b79", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAny_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAny.ReplaceIfExists", t, func() {
		var k interface{} = "996a7b7a-bda6-4140-b607-fe145292f510"
		var v interface{} = "cea018ad-c2b2-436f-961e-d4c624e3d393"
		var x interface{} = "736ec7b3-64cc-49d1-92e8-0f2b98d2f3be"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("d8e1bdda-66c0-4f92-9d18-b706b2871942", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAny_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAny.ReplaceOrPut", t, func() {
		var k interface{} = "8c25a719-2197-4904-93dd-4514adc1be2c"
		var v interface{} = "3903625c-944c-4f12-8d32-96bd1748a674"
		var x interface{} = "f40105e9-28c0-44a6-bbcb-9b065c539ebc"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("3c7f3c6e-f1a9-4f66-ab74-39066931c5a1", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAny_MarshalJSON(t *testing.T) {
	Convey("TestMapAny.MarshalJSON", t, func() {
		var k interface{} = "94fd0201-df9d-4c09-b74c-40088dc161dc"
		var v interface{} = "f0d0813d-5f10-4c08-a2b0-205af138c1a2"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"94fd0201-df9d-4c09-b74c-40088dc161dc","value":"f0d0813d-5f10-4c08-a2b0-205af138c1a2"}]`)
	})
}

