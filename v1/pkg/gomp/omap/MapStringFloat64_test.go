package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringFloat64_Put(t *testing.T) {
	Convey("TestMapStringFloat64.Put", t, func() {
		var k string = "a2d4d43c-4b6e-4e9e-be89-f388e4abf28d"
		var v float64 = 0.140

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringFloat64_Delete(t *testing.T) {
	Convey("TestMapStringFloat64.Delete", t, func() {
		var k string = "d0566d04-6c31-490c-8f1b-cf84edc67a6c"
		var v float64 = 0.123

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringFloat64_Has(t *testing.T) {
	Convey("TestMapStringFloat64.Has", t, func() {
		var k string = "1de085c5-fe14-4b1a-a3bd-6c26d24dab1f"
		var v float64 = 0.573

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("483019ba-670d-4613-a726-7a33343c2bd7"+"7b8a2791-1a2a-4cc6-9c99-6e01ebd14eea"), ShouldBeFalse)
	})
}


func TestMapStringFloat64_Get(t *testing.T) {
	Convey("TestMapStringFloat64.Get", t, func() {
		var k string = "1f663a17-3a88-4d96-9274-efe0aa40fbf3"
		var v float64 = 0.083

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("b7d8cca6-5a7e-4f60-80f6-80977637026c" + "f3effc78-fddf-48ea-915b-d12edfb3c337")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringFloat64_GetOpt(t *testing.T) {
	Convey("TestMapStringFloat64.GetOpt", t, func() {
		var k string = "1bb4911f-dce8-4255-95d2-fbfc48be0bb6"
		var v float64 = 0.671

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("386f801d-c917-4c1e-8ceb-4146449adb1c" + "7ee0af44-596c-4e7e-999d-4880d58f4016")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringFloat64_ForEach(t *testing.T) {
	Convey("TestMapStringFloat64.ForEach", t, func() {
		var k string = "adef8f36-70bc-4b4e-866c-6c595c527772"
		var v float64 = 0.601
		hits := 0

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv float64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringFloat64_MarshalYAML(t *testing.T) {
	Convey("TestMapStringFloat64.MarshalYAML", t, func() {
		var k string = "c3e4c8b2-8b98-4d4a-bc78-99a4c32a2f99"
		var v float64 = 0.351

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringFloat64_ToYAML(t *testing.T) {
	Convey("TestMapStringFloat64.ToYAML", t, func() {
		var k string = "c598f41d-f8e3-4455-9b02-884e4540cb98"
		var v float64 = 0.343

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringFloat64_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringFloat64.PutIfNotNil", t, func() {
		var k string = "e8c58d6b-ab97-4f53-9f6d-347e9340089d"
		var v float64 = 0.663

		test := omap.NewMapStringFloat64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("156cb64c-f083-4824-8ce1-dc44059ea176", (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float64 = 0.759
		So(test.PutIfNotNil("0af5b51c-04a8-4963-b9c6-7d048820494c", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringFloat64.ReplaceIfExists", t, func() {
		var k string = "788f6770-7022-4884-87eb-517b5c351760"
		var v float64 = 0.633
		var x float64 = 0.050

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("8de4989b-4cff-4bba-9a73-3f9a8ce2ae1b", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringFloat64.ReplaceOrPut", t, func() {
		var k string = "c081b1c3-f3ee-4587-994e-27599fd53a63"
		var v float64 = 0.261
		var x float64 = 0.465

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("638ff5bc-a031-4ba7-bede-e18446e8643f", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat64_MarshalJSON(t *testing.T) {
	Convey("TestMapStringFloat64.MarshalJSON", t, func() {
		var k string = "06da7130-3e44-4a1f-9320-187b0c21891d"
		var v float64 = 0.538

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"06da7130-3e44-4a1f-9320-187b0c21891d","value":0.538}]`)
	})
}

