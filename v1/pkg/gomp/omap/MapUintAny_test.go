package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintAny_Put(t *testing.T) {
	Convey("TestMapUintAny.Put", t, func() {
		var k uint = 3657217943
		var v interface{} = "452ec987-8e79-4d2f-839c-6a7a1906d66d"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintAny_Delete(t *testing.T) {
	Convey("TestMapUintAny.Delete", t, func() {
		var k uint = 2084827865
		var v interface{} = "d830c1db-77ea-4b45-bac9-fbffb3168710"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintAny_Has(t *testing.T) {
	Convey("TestMapUintAny.Has", t, func() {
		var k uint = 3266189709
		var v interface{} = "b4686d6d-270e-46df-ba76-b438326e346c"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(444253654+3367587263), ShouldBeFalse)
	})
}


func TestMapUintAny_Get(t *testing.T) {
	Convey("TestMapUintAny.Get", t, func() {
		var k uint = 2433850904
		var v interface{} = "fd4afeb3-322e-4845-bcd6-5f0a05ab029c"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(2140993018+2837810556)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintAny_GetOpt(t *testing.T) {
	Convey("TestMapUintAny.GetOpt", t, func() {
		var k uint = 1966901341
		var v interface{} = "2d27bf54-b569-41ea-ba61-5bd159aeaf90"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(2666507142+2137420122)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintAny_ForEach(t *testing.T) {
	Convey("TestMapUintAny.ForEach", t, func() {
		var k uint = 2120412087
		var v interface{} = "40635a36-0134-4de2-8bad-d7923198c217"
		hits := 0

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv interface{}) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintAny_MarshalYAML(t *testing.T) {
	Convey("TestMapUintAny.MarshalYAML", t, func() {
		var k uint = 1919946868
		var v interface{} = "cf632244-1cd4-4731-9cac-c0920f07744c"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintAny_ToYAML(t *testing.T) {
	Convey("TestMapUintAny.ToYAML", t, func() {
		var k uint = 3583725097
		var v interface{} = "2bd0d269-560e-4af7-8d77-a0d49d412b12"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintAny_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintAny.PutIfNotNil", t, func() {
		var k uint = 2998537565
		var v interface{} = "1c030d63-7d6b-4d5b-af3f-d9ccf2b10ecb"

		test := omap.NewMapUintAny(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(3508047366, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "a27a7096-7cef-402e-aa96-bc9d25896b51"
		So(test.PutIfNotNil(2909489844, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintAny_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintAny.ReplaceIfExists", t, func() {
		var k uint = 3737612551
		var v interface{} = "34022f52-f76a-4b44-92a0-55b079eda92a"
		var x interface{} = "94c67cec-c304-494c-a296-b0c3a8000bf3"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(239868637, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintAny_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintAny.ReplaceOrPut", t, func() {
		var k uint = 3214485754
		var v interface{} = "8b9303f3-5b4f-4da5-9bf2-71b0ee6078a1"
		var x interface{} = "9a82eaac-d21f-4c09-b7e5-110440e178e6"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1348184158, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintAny_MarshalJSON(t *testing.T) {
	Convey("TestMapUintAny.MarshalJSON", t, func() {
		var k uint = 347902334
		var v interface{} = "b3eaf260-fbf0-46c3-8b29-9c949a6a71ac"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":347902334,"value":"b3eaf260-fbf0-46c3-8b29-9c949a6a71ac"}]`)
	})
}

