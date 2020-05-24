package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringInt32_Put(t *testing.T) {
	Convey("TestMapStringInt32.Put", t, func() {
		var k string = "3be4ed4c-f4b3-4078-93a0-8e8db94762c0"
		var v int32 = 590316323

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringInt32_Delete(t *testing.T) {
	Convey("TestMapStringInt32.Delete", t, func() {
		var k string = "d349840a-7c56-4eb5-99cf-c226f6bd8d77"
		var v int32 = 328174106

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringInt32_Has(t *testing.T) {
	Convey("TestMapStringInt32.Has", t, func() {
		var k string = "614852de-cec1-416c-978a-9c3d2402939b"
		var v int32 = 404618197

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("a21bb242-02c2-45fa-ad34-0f0395ae91c4"+"490fa49d-ac09-4a9c-8155-b8c42a734e50"), ShouldBeFalse)
	})
}


func TestMapStringInt32_Get(t *testing.T) {
	Convey("TestMapStringInt32.Get", t, func() {
		var k string = "dcdbfa89-9ae7-431d-a24f-d6e30d311044"
		var v int32 = 1186174663

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("09fe163a-0b16-4bb4-82d0-5da22f37b1bc" + "99f11d8e-c204-4524-9714-cee3fe178525")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringInt32_GetOpt(t *testing.T) {
	Convey("TestMapStringInt32.GetOpt", t, func() {
		var k string = "ddde5cc4-84bf-4eb2-b36a-2741dcaa685e"
		var v int32 = 495197215

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("e1add7d1-be82-4e24-ac80-ed620c6d720a" + "b99e3a7c-785d-4955-9efc-8e9edfcf09fd")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringInt32_ForEach(t *testing.T) {
	Convey("TestMapStringInt32.ForEach", t, func() {
		var k string = "8c35cfc2-f92b-463f-97a4-5830bb9afcf6"
		var v int32 = 1075384036
		hits := 0

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv int32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringInt32_MarshalYAML(t *testing.T) {
	Convey("TestMapStringInt32.MarshalYAML", t, func() {
		var k string = "28ed8376-b53b-4037-8772-d41389d1b882"
		var v int32 = 2077484837

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringInt32_ToYAML(t *testing.T) {
	Convey("TestMapStringInt32.ToYAML", t, func() {
		var k string = "01544c8c-ab9d-4c9c-8b3f-6bc3900c432b"
		var v int32 = 1467615655

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringInt32_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringInt32.PutIfNotNil", t, func() {
		var k string = "0ce99e39-4263-42de-93d8-c52d085d7b23"
		var v int32 = 748405658

		test := omap.NewMapStringInt32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("54bbe410-fec2-4b6f-928b-4cb66e2ae76c", (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int32 = 868777064
		So(test.PutIfNotNil("0e6d372f-34dc-4bf5-a6ca-078b5ddc2f1c", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringInt32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringInt32.ReplaceIfExists", t, func() {
		var k string = "fc8f3cfa-da9e-4c3a-8885-a945cfa5f89c"
		var v int32 = 1784034995
		var x int32 = 2071578266

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("9ad8feb7-26dd-4e20-880e-f66fa1b4db3b", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringInt32.ReplaceOrPut", t, func() {
		var k string = "c92180b7-1424-463d-84f7-f44400af67bd"
		var v int32 = 822665019
		var x int32 = 1589081688

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("b0bfd2ad-8424-4de0-8817-266aec5c02d6", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt32_MarshalJSON(t *testing.T) {
	Convey("TestMapStringInt32.MarshalJSON", t, func() {
		var k string = "dc765878-66bd-4d15-a4a2-1fa2e1d32897"
		var v int32 = 1426044116

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"dc765878-66bd-4d15-a4a2-1fa2e1d32897","value":1426044116}]`)
	})
}
