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
		var k string = "12374751-66bb-484c-806c-aae13f409518"
		var v byte = 3

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringByte_Delete(t *testing.T) {
	Convey("TestMapStringByte.Delete", t, func() {
		var k string = "0072e3d0-4421-4448-bc9e-6046383f1b49"
		var v byte = 5

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringByte_Has(t *testing.T) {
	Convey("TestMapStringByte.Has", t, func() {
		var k string = "c42a3bcb-9da8-47de-bd05-ddf77108895d"
		var v byte = 240

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("e517ffee-6bcb-4de4-82ef-a555d7534e1c"+"ebc31718-7728-484a-9261-b1d01a7aa272"), ShouldBeFalse)
	})
}

func TestMapStringByte_Get(t *testing.T) {
	Convey("TestMapStringByte.Get", t, func() {
		var k string = "370132ca-37c5-4e64-b969-52f0298030bd"
		var v byte = 112

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("7cc08e77-dd88-4e1c-85c3-d33b3201bf6c" + "96e8352e-541d-41d0-9a3b-f86b0d9e766c")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringByte_GetOpt(t *testing.T) {
	Convey("TestMapStringByte.GetOpt", t, func() {
		var k string = "8d3fd3d3-75ac-40f3-9ec4-be39b1c02c95"
		var v byte = 134

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("7179b6be-25ff-4540-b1e6-38d216de7130" + "6f4d12ab-fb49-4d86-9b3b-5bb707dd6b49")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringByte_ForEach(t *testing.T) {
	Convey("TestMapStringByte.ForEach", t, func() {
		var k string = "33859691-1f55-4658-8ae4-28bdf698830e"
		var v byte = 231
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
		var k string = "0426e29a-d511-4444-a602-1518153976a3"
		var v byte = 106

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
		var k string = "3b39950e-62f1-41ee-b454-a8a96221c6c0"
		var v byte = 145

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
		var k string = "a2e84d53-0509-48ba-a460-8a01b0a1b735"
		var v byte = 67

		test := omap.NewMapStringByte(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("e1f796f2-9281-479d-924a-f14e4de42a1f", (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x byte = 232
		So(test.PutIfNotNil("357e96a0-222b-48d4-9ca5-52d107663c52", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringByte_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringByte.ReplaceIfExists", t, func() {
		var k string = "60283a47-2d05-40d9-ad57-86ef51476e87"
		var v byte = 201
		var x byte = 150

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("cb78c7be-2ec2-45cc-9dca-a6a2163d67e8", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringByte_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringByte.ReplaceOrPut", t, func() {
		var k string = "c9e890e4-faa6-476d-a2b3-e214bb41ac08"
		var v byte = 215
		var x byte = 44

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("83ece3a2-5618-4021-b124-2047ab7ecf66", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringByte_MarshalJSON(t *testing.T) {
	Convey("TestMapStringByte.MarshalJSON", t, func() {
		var k string = "4a5798ab-226c-46f0-a63b-5c2a2936aaee"
		var v byte = 44

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"4a5798ab-226c-46f0-a63b-5c2a2936aaee","value":44}]`)
	})
}
