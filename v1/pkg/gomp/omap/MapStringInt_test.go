package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringInt_Put(t *testing.T) {
	Convey("TestMapStringInt.Put", t, func() {
		var k string = "1a3ee035-5295-42d8-b6fd-319e6f59b62d"
		var v int = 1019295016

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringInt_Delete(t *testing.T) {
	Convey("TestMapStringInt.Delete", t, func() {
		var k string = "98033424-57ca-4c9b-9d01-d99b96919cf6"
		var v int = 1220365328

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringInt_Has(t *testing.T) {
	Convey("TestMapStringInt.Has", t, func() {
		var k string = "1c034759-c167-4f08-a3c4-4d65f4edc116"
		var v int = 1270926168

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("07a42f40-2307-4e02-b458-12866fe4526e"+"6da761cd-8977-4bd6-9ab6-b2163396ed75"), ShouldBeFalse)
	})
}


func TestMapStringInt_Get(t *testing.T) {
	Convey("TestMapStringInt.Get", t, func() {
		var k string = "02601970-205c-48a6-9d4c-ce5445b0e3e2"
		var v int = 1017821721

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("42d6597f-6c4d-4725-95a4-759b6c421cd2" + "6420537f-e7a7-475d-8380-c990ae2696ea")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringInt_GetOpt(t *testing.T) {
	Convey("TestMapStringInt.GetOpt", t, func() {
		var k string = "309f804e-1ab8-4623-88cb-ab7f56ed7541"
		var v int = 211132820

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("bb5f0392-6200-4607-952d-c8ae9ac3dd82" + "3bdf9705-0374-4065-97ee-b923d89d0f95")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringInt_ForEach(t *testing.T) {
	Convey("TestMapStringInt.ForEach", t, func() {
		var k string = "fbcf4e7b-0ac7-4249-9f21-7d3e7366a831"
		var v int = 827031710
		hits := 0

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv int) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringInt_MarshalYAML(t *testing.T) {
	Convey("TestMapStringInt.MarshalYAML", t, func() {
		var k string = "700a3a3a-48cd-4e21-b449-581e0ce3974f"
		var v int = 48631469

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringInt_ToYAML(t *testing.T) {
	Convey("TestMapStringInt.ToYAML", t, func() {
		var k string = "1f267ae8-f284-4cef-b7a7-3ea4b672fa8a"
		var v int = 951750648

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringInt_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringInt.PutIfNotNil", t, func() {
		var k string = "98d77b97-33c7-4a81-a501-ff410f13271d"
		var v int = 1150472413

		test := omap.NewMapStringInt(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("44f05b98-98bf-4ef8-b946-6e349d5ac977", (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int = 1748033704
		So(test.PutIfNotNil("3f02d78f-9831-472c-95bd-03bafcbf77a4", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringInt_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringInt.ReplaceIfExists", t, func() {
		var k string = "1ad5fd49-3ffa-4ffb-bfdf-4f51c25268d6"
		var v int = 1170546076
		var x int = 1150987488

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("fe9c11fc-037e-4de3-8a41-77376548dbf3", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringInt.ReplaceOrPut", t, func() {
		var k string = "89b165e8-5423-46ad-88d3-f940adb0a237"
		var v int = 1541833928
		var x int = 715087382

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("fd36c405-4c96-4cd0-911d-38ee1fe08273", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt_MarshalJSON(t *testing.T) {
	Convey("TestMapStringInt.MarshalJSON", t, func() {
		var k string = "bde7261d-5008-44fd-a343-2d1d876fd9d2"
		var v int = 452127554

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"bde7261d-5008-44fd-a343-2d1d876fd9d2","value":452127554}]`)
	})
}

