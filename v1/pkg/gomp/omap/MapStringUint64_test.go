package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringUint64_Put(t *testing.T) {
	Convey("TestMapStringUint64.Put", t, func() {
		var k string = "a525e145-4d40-46e3-946f-bbb5093532ce"
		var v uint64 = 9972618714909663997

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringUint64_Delete(t *testing.T) {
	Convey("TestMapStringUint64.Delete", t, func() {
		var k string = "29e813d8-8966-4030-bda9-49fd52883f06"
		var v uint64 = 17360198493881439768

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringUint64_Has(t *testing.T) {
	Convey("TestMapStringUint64.Has", t, func() {
		var k string = "255adad0-6679-4064-8295-6746b503bd85"
		var v uint64 = 1265735805629511407

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("7580f190-5280-480e-862a-cab2c8d0211e"+"dc7682dd-7849-4bf2-87a8-eb4d9b091472"), ShouldBeFalse)
	})
}

func TestMapStringUint64_Get(t *testing.T) {
	Convey("TestMapStringUint64.Get", t, func() {
		var k string = "3a7ad188-42ac-451a-b9e7-cb0f81d89c38"
		var v uint64 = 17808267879619022783

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("5e3d13bc-3859-42db-88ab-df4b5750eb95" + "760c2a03-c230-4389-9511-c3ed00b06e70")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringUint64_GetOpt(t *testing.T) {
	Convey("TestMapStringUint64.GetOpt", t, func() {
		var k string = "2db576c5-eca4-42e5-8c37-5203555c834c"
		var v uint64 = 12911937214394424512

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("09352f34-37af-4833-8e33-8552cd8baf1d" + "5e984086-e476-4fe9-8c87-3adcc1e115e7")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringUint64_ForEach(t *testing.T) {
	Convey("TestMapStringUint64.ForEach", t, func() {
		var k string = "553ea522-ebbf-44c6-a040-e4191eb8c5c5"
		var v uint64 = 8132141189597234880
		hits := 0

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv uint64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringUint64_MarshalYAML(t *testing.T) {
	Convey("TestMapStringUint64.MarshalYAML", t, func() {
		var k string = "b7e5fd71-7ea9-4d01-889d-531294570056"
		var v uint64 = 12468866615456721993

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringUint64_ToYAML(t *testing.T) {
	Convey("TestMapStringUint64.ToYAML", t, func() {
		var k string = "01ce23a1-b0d3-4629-b9cd-6100b3200690"
		var v uint64 = 17907350705951375054

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringUint64_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringUint64.PutIfNotNil", t, func() {
		var k string = "59667711-fc9a-408e-a0ba-7ca207428a3a"
		var v uint64 = 10304094193866346331

		test := omap.NewMapStringUint64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("a8135e0e-fbe3-472a-9404-56cd25ff781f", (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint64 = 5136837769367193113
		So(test.PutIfNotNil("b10bebfd-0e41-4e53-9873-c18e95e9f452", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringUint64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringUint64.ReplaceIfExists", t, func() {
		var k string = "a4b732be-a4db-4499-98e2-de6a1fe7b065"
		var v uint64 = 8825078422992556938
		var x uint64 = 5695226376882192865

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("04d90afc-dcba-4d91-8eb9-fd038490c332", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringUint64.ReplaceOrPut", t, func() {
		var k string = "d4d9eb18-8ceb-47b1-8c33-fc149efcf2bf"
		var v uint64 = 4654947393991697030
		var x uint64 = 14447153000968318863

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("c5a2b4b1-6436-4d48-8928-4f3d97bdaa9f", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint64_MarshalJSON(t *testing.T) {
	Convey("TestMapStringUint64.MarshalJSON", t, func() {
		var k string = "a91b65a1-185c-4f87-a11e-a3a238606870"
		var v uint64 = 322588993718643170

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"a91b65a1-185c-4f87-a11e-a3a238606870","value":322588993718643170}]`)
	})
}
