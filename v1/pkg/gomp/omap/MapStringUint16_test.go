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
		var k string = "a89423b4-0c32-4e81-af4d-77d8b95f5140"
		var v uint16 = 54614

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringUint16_Delete(t *testing.T) {
	Convey("TestMapStringUint16.Delete", t, func() {
		var k string = "5f4c10d4-da23-4f67-9387-4e095fb957f6"
		var v uint16 = 55572

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringUint16_Has(t *testing.T) {
	Convey("TestMapStringUint16.Has", t, func() {
		var k string = "03502f42-7360-4263-89da-fa26aba1a379"
		var v uint16 = 43204

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("7c29872a-9226-4a94-a320-e8f7027b4b27"+"4d7f6d1f-eef6-49c7-879b-1ac559878a1d"), ShouldBeFalse)
	})
}

func TestMapStringUint16_Get(t *testing.T) {
	Convey("TestMapStringUint16.Get", t, func() {
		var k string = "dbd26eef-e7f4-426a-aad0-fe3966053874"
		var v uint16 = 41845

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("14779a79-9df3-4696-abd1-28489c61a90d" + "a6c36ce3-69b2-42fc-9e6c-68dfc103ac10")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringUint16_GetOpt(t *testing.T) {
	Convey("TestMapStringUint16.GetOpt", t, func() {
		var k string = "1d575fad-1ee0-418a-b70c-8098bfded3c8"
		var v uint16 = 2772

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("b472d0ad-5b31-4ec0-a3cb-192a24fb78c2" + "4bd9e57f-8b72-41e0-98dc-f988271100a9")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringUint16_ForEach(t *testing.T) {
	Convey("TestMapStringUint16.ForEach", t, func() {
		var k string = "8249b440-f77c-4972-bf8f-6232af90ea74"
		var v uint16 = 47423
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
		var k string = "feb214fb-a590-4c98-9146-bd3a42b60b20"
		var v uint16 = 26667

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
		var k string = "133ea728-8777-490b-8530-9c8a77fba94c"
		var v uint16 = 37679

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
		var k string = "eb6586f2-39e7-4821-8447-183b48cc3a8b"
		var v uint16 = 11212

		test := omap.NewMapStringUint16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("0f5b18ad-3867-4a31-9bf8-d8e646887d8b", (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint16 = 38584
		So(test.PutIfNotNil("777f2670-1dfd-4208-83e7-03fb64149d9e", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringUint16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringUint16.ReplaceIfExists", t, func() {
		var k string = "2e9a2213-21cb-4fab-9032-258ac21d3e03"
		var v uint16 = 6286
		var x uint16 = 21023

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("bddefeb7-34f5-4624-8f36-32eca8da0d9c", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringUint16.ReplaceOrPut", t, func() {
		var k string = "29cb811f-3f80-44bd-9ac4-0e87b0aa2503"
		var v uint16 = 15824
		var x uint16 = 64928

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("bf4c5ed8-3c6d-4457-9f5f-b82adf947fca", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint16_MarshalJSON(t *testing.T) {
	Convey("TestMapStringUint16.MarshalJSON", t, func() {
		var k string = "5ed15794-47a9-4aaf-8ca5-f398e198939a"
		var v uint16 = 18030

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"5ed15794-47a9-4aaf-8ca5-f398e198939a","value":18030}]`)
	})
}
