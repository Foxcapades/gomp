package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyInt16_Put(t *testing.T) {
	Convey("TestMapAnyInt16.Put", t, func() {
		var k interface{} = "29316da7-2126-478c-9c0d-613b3cd9227e"
		var v int16 = 18429

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyInt16_Delete(t *testing.T) {
	Convey("TestMapAnyInt16.Delete", t, func() {
		var k interface{} = "dfdb8194-5de0-46ec-a851-3a1b9e817d43"
		var v int16 = 21335

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyInt16_Has(t *testing.T) {
	Convey("TestMapAnyInt16.Has", t, func() {
		var k interface{} = "e9c70b0a-8a7e-4b69-95cf-cb73ddf5175a"
		var v int16 = 11763

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("2b623efb-e251-4c5c-9fe1-7cb52a1f1742"+"fe2be722-aac5-4044-8e56-48b21ef6368e"), ShouldBeFalse)
	})
}


func TestMapAnyInt16_Get(t *testing.T) {
	Convey("TestMapAnyInt16.Get", t, func() {
		var k interface{} = "b6b1b369-8e06-4a38-85fe-b42a16b8aa28"
		var v int16 = 21055

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("b5b44bf1-e1ab-44ce-b06b-6242de0f0dd0"+"b034b018-a43e-4515-a563-bf4271599ac3")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyInt16_GetOpt(t *testing.T) {
	Convey("TestMapAnyInt16.GetOpt", t, func() {
		var k interface{} = "cd43b7ce-1326-4d80-859c-1cdb315461b3"
		var v int16 = 14954

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("0c706af7-1444-402e-b8e3-1a96cdba2091"+"ed5deabe-c3b8-4bde-bdbe-3ed3eee8ab4b")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyInt16_ForEach(t *testing.T) {
	Convey("TestMapAnyInt16.ForEach", t, func() {
		var k interface{} = "b2d42033-168b-4247-a855-7e9269f4d737"
		var v int16 = 9485
		hits := 0

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv int16) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyInt16_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyInt16.MarshalYAML", t, func() {
		var k interface{} = "46d5ab84-8113-4ff4-8a42-496be754813a"
		var v int16 = 16526

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyInt16_ToYAML(t *testing.T) {
	Convey("TestMapAnyInt16.ToYAML", t, func() {
		var k interface{} = "31635abd-123f-4a07-abd9-a88d6c0b3837"
		var v int16 = 23016

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyInt16_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyInt16.PutIfNotNil", t, func() {
		var k interface{} = "0c90d6fd-c469-4614-a721-b5aff23a15c8"
		var v int16 = 20137

		test := omap.NewMapAnyInt16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("8368fae4-d968-4fe8-b40b-9c9ebb8a79bc", (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int16 = 21193
		So(test.PutIfNotNil("79ecf0c3-f40a-46ea-a741-db59654c0358", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyInt16.ReplaceIfExists", t, func() {
		var k interface{} = "435c36f6-c877-4643-a243-16120c211f03"
		var v int16 = 22711
		var x int16 = 15083

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("eb5ad4be-1653-4a4a-99f4-fcc50b8224a2", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyInt16.ReplaceOrPut", t, func() {
		var k interface{} = "de1cef68-6e93-4e8d-98a9-20aff42ad459"
		var v int16 = 30149
		var x int16 = 8940

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("8e2cf4cc-5213-45a6-9e7f-b992a973fff3", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt16_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyInt16.MarshalJSON", t, func() {
		var k interface{} = "2eaa75e7-1f3a-423a-8156-9d12b5e19cc8"
		var v int16 = 10114

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"2eaa75e7-1f3a-423a-8156-9d12b5e19cc8","value":10114}]`)
	})
}

