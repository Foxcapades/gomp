package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintString_Put(t *testing.T) {
	Convey("TestMapUintString.Put", t, func() {
		var k uint = 3055239480
		var v string = "e6a776f6-3c15-4af6-94ea-95ae07a7774f"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintString_Delete(t *testing.T) {
	Convey("TestMapUintString.Delete", t, func() {
		var k uint = 4184044078
		var v string = "dbc9ada3-cd2c-44a3-a02e-86c4d9aeb6d0"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintString_Has(t *testing.T) {
	Convey("TestMapUintString.Has", t, func() {
		var k uint = 1768039393
		var v string = "02be6363-d952-443b-9e21-88cf686c4c84"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(3389839557+2554516647), ShouldBeFalse)
	})
}

func TestMapUintString_Get(t *testing.T) {
	Convey("TestMapUintString.Get", t, func() {
		var k uint = 3787427904
		var v string = "f72347ef-1143-424a-8ec6-a5a8a65515bf"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(4153906849 + 124397549)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintString_GetOpt(t *testing.T) {
	Convey("TestMapUintString.GetOpt", t, func() {
		var k uint = 3248399408
		var v string = "87f941be-eb33-49d3-8e5b-583ae1a2ddb0"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1234243545 + 3524857633)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintString_ForEach(t *testing.T) {
	Convey("TestMapUintString.ForEach", t, func() {
		var k uint = 2772144718
		var v string = "47dd0655-6f4e-4e93-a6d7-aa0b4463a51a"
		hits := 0

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv string) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintString_MarshalYAML(t *testing.T) {
	Convey("TestMapUintString.MarshalYAML", t, func() {
		var k uint = 939878258
		var v string = "36c39a27-5b50-4f39-b014-842666098159"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintString_ToYAML(t *testing.T) {
	Convey("TestMapUintString.ToYAML", t, func() {
		var k uint = 2116984329
		var v string = "99c95122-7912-4ddf-bbf3-5ca53f14298a"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintString_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintString.PutIfNotNil", t, func() {
		var k uint = 3145077284
		var v string = "89cb15b3-7b4f-48fc-a9eb-900b0e6a53c6"

		test := omap.NewMapUintString(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1828174578, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "a3b808c5-366d-4938-bb8a-62fbe5b90022"
		So(test.PutIfNotNil(4199188091, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintString_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintString.ReplaceIfExists", t, func() {
		var k uint = 2927248668
		var v string = "f88053f3-eb5c-44e9-8c60-8c7a64536356"
		var x string = "c1491346-ecb3-4bb3-9594-4a62fa122c8d"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1969456714, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintString_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintString.ReplaceOrPut", t, func() {
		var k uint = 3751650141
		var v string = "acb4501a-59ec-4e46-9f33-25f4e0d57b4e"
		var x string = "53311e74-ac2c-44a3-8f76-2b65a69e0e67"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1198965276, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintString_MarshalJSON(t *testing.T) {
	Convey("TestMapUintString.MarshalJSON", t, func() {
		var k uint = 3042265488
		var v string = "4baba6d5-8b32-4af2-9876-445dbbd831c0"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":3042265488,"value":"4baba6d5-8b32-4af2-9876-445dbbd831c0"}]`)
	})
}
