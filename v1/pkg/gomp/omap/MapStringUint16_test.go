package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringUint16_Put(t *testing.T) {
	Convey("TestMapStringUint16.Put", t, func() {
		var k string = "e95a8a5e-8f3d-44bb-8da2-1a049d53bb1e"
		var v uint16 = 50986

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringUint16_Delete(t *testing.T) {
	Convey("TestMapStringUint16.Delete", t, func() {
		var k string = "d23079ee-4799-43b6-9e65-89f4db928cbd"
		var v uint16 = 10770

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringUint16_Has(t *testing.T) {
	Convey("TestMapStringUint16.Has", t, func() {
		var k string = "eb2d99c6-b861-4745-a658-5801b1d6897c"
		var v uint16 = 15848

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("1caf4a9c-1d5e-49a9-92a7-32ab657904ce"+"188fec46-cc8f-498b-8ecd-b28e37919342"), ShouldBeFalse)
	})
}

func TestMapStringUint16_Get(t *testing.T) {
	Convey("TestMapStringUint16.Get", t, func() {
		var k string = "aa262226-be18-42ab-a0fc-1b0659d6f3f8"
		var v uint16 = 21638

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("c5294748-808d-48c1-801a-6d39c6c6463c" + "9389f08a-3146-44f7-a965-931d33172574")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringUint16_GetOpt(t *testing.T) {
	Convey("TestMapStringUint16.GetOpt", t, func() {
		var k string = "848b2a51-4e11-4d74-9c49-cc5e9726f543"
		var v uint16 = 8600

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("f83c7526-82f4-43f6-9059-72c322625bb6" + "a4d258dd-f201-4503-96e6-6d259140ab24")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringUint16_ForEach(t *testing.T) {
	Convey("TestMapStringUint16.ForEach", t, func() {
		var k string = "90cc5475-de5a-4e89-9209-2e699b38b2ad"
		var v uint16 = 31515
		hits := 0

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv uint16) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringUint16_MarshalYAML(t *testing.T) {
	Convey("TestMapStringUint16.MarshalYAML", t, func() {
		var k string = "4f9da562-d059-408b-870a-6e74ffcebd7b"
		var v uint16 = 39177

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringUint16_ToYAML(t *testing.T) {
	Convey("TestMapStringUint16.ToYAML", t, func() {
		var k string = "861b30a7-a718-4d1a-b96d-d11e3e0f6410"
		var v uint16 = 63714

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringUint16_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringUint16.PutIfNotNil", t, func() {
		var k string = "2a94b5e7-dba2-4747-abe7-8edb68e3bfac"
		var v uint16 = 38156

		test := omap.NewMapStringUint16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("15fb6a3d-c88d-40f5-9101-116370ee69da", (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint16 = 1407
		So(test.PutIfNotNil("5e7dfee1-f10b-45ba-8800-edfc6679b713", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringUint16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringUint16.ReplaceIfExists", t, func() {
		var k string = "14b58f0d-1156-482d-b382-521731811080"
		var v uint16 = 14752
		var x uint16 = 62761

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("ff1cefac-20c3-4f50-ad98-6123b0af3160", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringUint16.ReplaceOrPut", t, func() {
		var k string = "1cc8be95-caaf-4920-a02b-c8de568b29dc"
		var v uint16 = 5729
		var x uint16 = 7422

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("f03da1e2-57e3-41f7-8802-88777e21e09b", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint16_MarshalJSON(t *testing.T) {
	Convey("TestMapStringUint16.MarshalJSON", t, func() {
		var k string = "fc684ae5-f09c-43c4-8016-4fecb5924029"
		var v uint16 = 56911

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"fc684ae5-f09c-43c4-8016-4fecb5924029","value":56911}]`)
	})
}
