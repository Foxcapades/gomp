package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringInt64_Put(t *testing.T) {
	Convey("TestMapStringInt64.Put", t, func() {
		var k string = "c073090a-31a9-4e85-b66f-ea3f65d88268"
		var v int64 = 645597993881127551

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringInt64_Delete(t *testing.T) {
	Convey("TestMapStringInt64.Delete", t, func() {
		var k string = "b1562f44-ac47-467b-992c-1cfc7dc836e2"
		var v int64 = 5907470870181063053

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringInt64_Has(t *testing.T) {
	Convey("TestMapStringInt64.Has", t, func() {
		var k string = "8d4a92f1-3e46-49d2-a058-58432a4b8aa3"
		var v int64 = 7419461738555197303

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("c2fd20db-0e1d-4db6-abc6-eb2c4ec5cdd1"+"07486da3-70a7-4b54-9945-5d2a62e7928f"), ShouldBeFalse)
	})
}


func TestMapStringInt64_Get(t *testing.T) {
	Convey("TestMapStringInt64.Get", t, func() {
		var k string = "f42fef69-90a9-4327-803a-13ad2d7a9396"
		var v int64 = 750570677677681460

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("2e2bb8df-c01d-4026-8526-ce571f147f45"+"2345cd09-41f8-471a-acb4-25e832e05e2f")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringInt64_GetOpt(t *testing.T) {
	Convey("TestMapStringInt64.GetOpt", t, func() {
		var k string = "57fce4c6-eaa9-4fd2-83cc-7450c63bc0d6"
		var v int64 = 1929037833464083161

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("b7dd8741-7810-46db-963f-4f146f44bcaa"+"1e2a11f6-ec7c-4b3a-b7bf-769a4f000217")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringInt64_ForEach(t *testing.T) {
	Convey("TestMapStringInt64.ForEach", t, func() {
		var k string = "82691e8a-a221-4933-87c4-2c2606068788"
		var v int64 = 6897416173785842563
		hits := 0

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv int64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringInt64_MarshalYAML(t *testing.T) {
	Convey("TestMapStringInt64.MarshalYAML", t, func() {
		var k string = "b9801be9-fbd3-4a53-8432-ecd6e49e74b0"
		var v int64 = 661120049752708039

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringInt64_ToYAML(t *testing.T) {
	Convey("TestMapStringInt64.ToYAML", t, func() {
		var k string = "a62d3fa8-45ff-4874-8873-ce3f377e7600"
		var v int64 = 8744937918720231070

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringInt64_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringInt64.PutIfNotNil", t, func() {
		var k string = "0d2c5377-1a05-42b6-839b-454852f71ed6"
		var v int64 = 467161858300813250

		test := omap.NewMapStringInt64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("763a1db0-e2f5-4119-9e4f-c0e26b140b64", (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int64 = 1058687617453660662
		So(test.PutIfNotNil("10ccc85c-6dfc-40b5-9bf3-3d3827960aee", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringInt64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringInt64.ReplaceIfExists", t, func() {
		var k string = "c45d93b7-99a8-4f2b-aae5-d2b7b4b91c79"
		var v int64 = 3333126119590041980
		var x int64 = 4459735513674988711

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("433523c7-fe7d-41e1-b207-5c5b7a372aa6", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringInt64.ReplaceOrPut", t, func() {
		var k string = "53be8304-3cc6-4155-b183-b141ac15a532"
		var v int64 = 8940112674682615184
		var x int64 = 5347049067880602720

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("64a21030-81af-496c-a393-23f7cea73c53", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt64_MarshalJSON(t *testing.T) {
	Convey("TestMapStringInt64.MarshalJSON", t, func() {
		var k string = "8099e560-ae2a-4ae3-ba81-cc1b5992f38a"
		var v int64 = 2482850133057057846

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"8099e560-ae2a-4ae3-ba81-cc1b5992f38a","value":2482850133057057846}]`)
	})
}

