package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringInt16_Put(t *testing.T) {
	Convey("TestMapStringInt16.Put", t, func() {
		var k string = "d3501dc9-1bf6-4806-b0cc-38d285dd75c4"
		var v int16 = 9206

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringInt16_Delete(t *testing.T) {
	Convey("TestMapStringInt16.Delete", t, func() {
		var k string = "16f68fa8-443a-47f5-9c5d-9bf5943f4c69"
		var v int16 = 16401

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringInt16_Has(t *testing.T) {
	Convey("TestMapStringInt16.Has", t, func() {
		var k string = "14827061-0a67-4ffa-af48-70883f5506dd"
		var v int16 = 3211

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("2d414ec5-b0eb-44c0-aecd-dee57233f602"+"81d611b6-be77-436c-9197-376af917e047"), ShouldBeFalse)
	})
}


func TestMapStringInt16_Get(t *testing.T) {
	Convey("TestMapStringInt16.Get", t, func() {
		var k string = "741aa52a-3ca4-4890-a616-babf85052bec"
		var v int16 = 21731

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("3be56f98-6ad6-4f49-b1fc-9387565f6784"+"f4b5b6f5-09fa-4d09-b0ba-01f21f57dca6")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringInt16_GetOpt(t *testing.T) {
	Convey("TestMapStringInt16.GetOpt", t, func() {
		var k string = "ade4b604-fc0b-495e-b807-047520b917f4"
		var v int16 = 592

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("c6926bd4-adad-4b5e-8b30-cacf8bacdc50"+"26968a2d-a092-41b3-b61d-c4109ede0bdf")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringInt16_ForEach(t *testing.T) {
	Convey("TestMapStringInt16.ForEach", t, func() {
		var k string = "f943c3b2-ebff-4340-a7da-b462dce6d3c5"
		var v int16 = 24728
		hits := 0

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv int16) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringInt16_MarshalYAML(t *testing.T) {
	Convey("TestMapStringInt16.MarshalYAML", t, func() {
		var k string = "c2706269-d727-4ad2-94ea-076bcc673a5b"
		var v int16 = 32624

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringInt16_ToYAML(t *testing.T) {
	Convey("TestMapStringInt16.ToYAML", t, func() {
		var k string = "1c681e50-c0ad-4eed-870c-d49e62b5b997"
		var v int16 = 9071

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringInt16_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringInt16.PutIfNotNil", t, func() {
		var k string = "aaecfedf-e5f3-4f2c-859e-1bc0aeb1f666"
		var v int16 = 26671

		test := omap.NewMapStringInt16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("57ed7541-e4f7-4c1c-8ae5-5a79369609e5", (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int16 = 7848
		So(test.PutIfNotNil("b4092d03-2b9a-4478-87d4-18fbd8ce2da0", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringInt16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringInt16.ReplaceIfExists", t, func() {
		var k string = "f1d1ae35-fba4-4ed1-84c2-4bc1bf4a1033"
		var v int16 = 24901
		var x int16 = 1373

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("0dc2d8f2-5cce-4278-8418-63a305fe735f", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringInt16.ReplaceOrPut", t, func() {
		var k string = "68be9bcd-a18d-4c42-8867-2e1f4921a7c5"
		var v int16 = 7043
		var x int16 = 32258

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("e3abf361-a867-4522-a4e6-f1bab29f84d6", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt16_MarshalJSON(t *testing.T) {
	Convey("TestMapStringInt16.MarshalJSON", t, func() {
		var k string = "e00ca00c-ec96-4420-b684-a90f5db0e3ad"
		var v int16 = 17408

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"e00ca00c-ec96-4420-b684-a90f5db0e3ad","value":17408}]`)
	})
}

