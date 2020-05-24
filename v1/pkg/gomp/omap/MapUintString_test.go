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
		var k uint = 302020621
		var v string = "50fb81aa-7dab-4e83-9f00-e4ec62beea49"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintString_Delete(t *testing.T) {
	Convey("TestMapUintString.Delete", t, func() {
		var k uint = 293030097
		var v string = "8b78c8ee-1066-43fb-9559-a853485be87e"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintString_Has(t *testing.T) {
	Convey("TestMapUintString.Has", t, func() {
		var k uint = 380021541
		var v string = "53c3c24b-e69c-4e71-a41d-2bfb923a67f8"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(3418581196+3012227885), ShouldBeFalse)
	})
}


func TestMapUintString_Get(t *testing.T) {
	Convey("TestMapUintString.Get", t, func() {
		var k uint = 1359140893
		var v string = "a66bd752-ef20-49a2-9a31-cf8fdef63751"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(3831500228 + 3259041319)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintString_GetOpt(t *testing.T) {
	Convey("TestMapUintString.GetOpt", t, func() {
		var k uint = 565153067
		var v string = "e7505df9-6a7c-4565-b3b5-415184ad8ff0"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(656433830 + 112844074)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintString_ForEach(t *testing.T) {
	Convey("TestMapUintString.ForEach", t, func() {
		var k uint = 3475470896
		var v string = "b01e78fd-5704-440d-a7d8-b01854abd0de"
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
		var k uint = 3005642192
		var v string = "c0813db7-d03f-498d-94ff-116be9634d8e"

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
		var k uint = 727409156
		var v string = "a30a7b80-406b-4a34-a799-3dbcc35b9b54"

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
		var k uint = 2760918352
		var v string = "c4f5c1b2-f57e-40e6-bebf-8990aab9460e"

		test := omap.NewMapUintString(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(612032041, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "78cbd181-8014-4fb6-ae15-64c651bd8c41"
		So(test.PutIfNotNil(4025553399, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintString_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintString.ReplaceIfExists", t, func() {
		var k uint = 2694880091
		var v string = "b908319c-68db-4b8c-9c42-a041478acd2c"
		var x string = "8775e10a-5091-4cc5-a285-f71f263136fb"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1163467686, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintString_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintString.ReplaceOrPut", t, func() {
		var k uint = 3224366534
		var v string = "917f4587-2e4b-4e51-8705-f0bd4b334488"
		var x string = "61c2661b-5475-4825-ab49-ba477ab3182e"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1135452063, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintString_MarshalJSON(t *testing.T) {
	Convey("TestMapUintString.MarshalJSON", t, func() {
		var k uint = 2913970218
		var v string = "0f318a6d-e472-49b3-a204-606146cdb0f7"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":2913970218,"value":"0f318a6d-e472-49b3-a204-606146cdb0f7"}]`)
	})
}

