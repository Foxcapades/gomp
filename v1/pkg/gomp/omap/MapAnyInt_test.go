package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyInt_Put(t *testing.T) {
	Convey("TestMapAnyInt.Put", t, func() {
		var k interface{} = "86d91b74-d517-4cba-97a7-1f738f0f9ab2"
		var v int = 333997230

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyInt_Delete(t *testing.T) {
	Convey("TestMapAnyInt.Delete", t, func() {
		var k interface{} = "75bd5905-4fcc-47ba-b5ba-ce3f02b6f2b0"
		var v int = 655817768

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyInt_Has(t *testing.T) {
	Convey("TestMapAnyInt.Has", t, func() {
		var k interface{} = "ca701836-fe73-47eb-aee0-10ea54c0933d"
		var v int = 1178665654

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("00bd2d95-d282-437d-a13d-b75ed323f194"+"7dfd5191-93b8-4a78-9b55-9b9cc3483451"), ShouldBeFalse)
	})
}


func TestMapAnyInt_Get(t *testing.T) {
	Convey("TestMapAnyInt.Get", t, func() {
		var k interface{} = "24758d89-02e3-413e-b7a7-ba2403e82033"
		var v int = 1808714526

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("ca216cd9-fe71-4c9d-ba80-02413aed1a60" + "f82ee4b4-8f8c-414f-95e7-e0a89fa7e3d0")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyInt_GetOpt(t *testing.T) {
	Convey("TestMapAnyInt.GetOpt", t, func() {
		var k interface{} = "ef2f3e73-9cef-4bc1-aacb-819907bca6de"
		var v int = 54237429

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("66c52be8-898a-47c6-b835-13c37baf4567" + "d0b1366b-b2f1-4246-996d-4631e2e0abd1")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyInt_ForEach(t *testing.T) {
	Convey("TestMapAnyInt.ForEach", t, func() {
		var k interface{} = "4f971668-3862-4ba9-a51d-08ccc710d7e9"
		var v int = 457712475
		hits := 0

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv int) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyInt_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyInt.MarshalYAML", t, func() {
		var k interface{} = "65904be0-6c8a-48c0-b07d-616ff504a6cc"
		var v int = 888820376

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyInt_ToYAML(t *testing.T) {
	Convey("TestMapAnyInt.ToYAML", t, func() {
		var k interface{} = "6fb817df-4928-4f8c-9480-88d43f6dc466"
		var v int = 142627014

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyInt_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyInt.PutIfNotNil", t, func() {
		var k interface{} = "369c5651-c359-4c0a-a6c6-77e395f08a5f"
		var v int = 283320691

		test := omap.NewMapAnyInt(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("fb6d8536-a0e3-4881-a1b9-ccceca519b47", (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int = 1747844274
		So(test.PutIfNotNil("30c8a878-c9a6-4cb9-9a76-ec1e7ae004a1", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyInt.ReplaceIfExists", t, func() {
		var k interface{} = "cbc5c6e1-54d2-4684-b583-7ca047819f76"
		var v int = 1799839269
		var x int = 463647504

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("f62fe795-053d-4f59-875c-fde151062580", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyInt.ReplaceOrPut", t, func() {
		var k interface{} = "e5584e2b-bcf8-4176-8d83-87c8d28d6dcf"
		var v int = 466804320
		var x int = 1204586530

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("2a8b9399-8d6d-4513-9105-a8cf9be41c95", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyInt.MarshalJSON", t, func() {
		var k interface{} = "3d0dc364-4fe2-47ad-8a9f-a3993bdd64c9"
		var v int = 1095803783

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"3d0dc364-4fe2-47ad-8a9f-a3993bdd64c9","value":1095803783}]`)
	})
}
