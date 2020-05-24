package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringUint32_Put(t *testing.T) {
	Convey("TestMapStringUint32.Put", t, func() {
		var k string = "5094a804-ba7b-4289-be0b-8a60b72754d5"
		var v uint32 = 551993475

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringUint32_Delete(t *testing.T) {
	Convey("TestMapStringUint32.Delete", t, func() {
		var k string = "50a08909-a62e-4115-ba5a-d8893286ad2d"
		var v uint32 = 3248873279

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringUint32_Has(t *testing.T) {
	Convey("TestMapStringUint32.Has", t, func() {
		var k string = "f3a6b3f2-9893-4bfd-9074-0dcc2e8b851e"
		var v uint32 = 3950785790

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("a5432e14-98de-4703-8155-402edcdc98f9"+"94f5322b-ae00-49a7-ad4f-411c252238c7"), ShouldBeFalse)
	})
}

func TestMapStringUint32_Get(t *testing.T) {
	Convey("TestMapStringUint32.Get", t, func() {
		var k string = "a03ca22e-111a-4068-8c12-da7d06081a4a"
		var v uint32 = 1759075350

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("b3035d5c-71ea-47d0-ba9b-bbdce0bf9da1" + "0bead6e6-4041-44bf-881c-8abca7992229")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringUint32_GetOpt(t *testing.T) {
	Convey("TestMapStringUint32.GetOpt", t, func() {
		var k string = "09296963-9e46-428b-9210-a485eefde057"
		var v uint32 = 2747851302

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("6ac3b4bc-a9c0-4936-a504-96c40d118568" + "5b901036-a5ca-4228-8cfa-590c3d980987")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringUint32_ForEach(t *testing.T) {
	Convey("TestMapStringUint32.ForEach", t, func() {
		var k string = "6a52b0c7-79d1-4e48-a914-e878117e8151"
		var v uint32 = 3428742101
		hits := 0

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv uint32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringUint32_MarshalYAML(t *testing.T) {
	Convey("TestMapStringUint32.MarshalYAML", t, func() {
		var k string = "09c48243-cc98-4452-9a47-ddc94410567a"
		var v uint32 = 3714199156

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringUint32_ToYAML(t *testing.T) {
	Convey("TestMapStringUint32.ToYAML", t, func() {
		var k string = "4360efd6-5b5a-4338-9b04-9defe9150be3"
		var v uint32 = 1327492233

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringUint32_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringUint32.PutIfNotNil", t, func() {
		var k string = "5cdfa284-66e0-4ffc-b19f-34894c56294e"
		var v uint32 = 501815067

		test := omap.NewMapStringUint32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("123b08a2-f428-4099-8a70-0334450ffceb", (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint32 = 3739404932
		So(test.PutIfNotNil("359851a7-2b4a-4110-9552-97c1a93e6f57", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringUint32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringUint32.ReplaceIfExists", t, func() {
		var k string = "7fec33f9-b5c3-44c2-af9b-1f51bbe22bcd"
		var v uint32 = 1818133480
		var x uint32 = 375859733

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("9550e969-6389-46bc-995a-ee0efd13957e", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringUint32.ReplaceOrPut", t, func() {
		var k string = "01d04a9e-8629-4b87-bcc2-5fdcda0e615e"
		var v uint32 = 3690490778
		var x uint32 = 984959391

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("f0676a9a-f303-47ce-a340-599ce59adef1", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint32_MarshalJSON(t *testing.T) {
	Convey("TestMapStringUint32.MarshalJSON", t, func() {
		var k string = "53683fa6-19ad-4fb6-ba42-144496cecb19"
		var v uint32 = 2896210342

		test := omap.NewMapStringUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"53683fa6-19ad-4fb6-ba42-144496cecb19","value":2896210342}]`)
	})
}
