package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringUint16_Put(t *testing.T) {
	Convey("TestMapStringUint16.Put", t, func() {
		var k string = "38120e08-6101-4924-94ad-ff1af6231543"
		var v uint16 = 12655

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringUint16_Delete(t *testing.T) {
	Convey("TestMapStringUint16.Delete", t, func() {
		var k string = "3feba41a-12e0-436d-b513-560b705a09bd"
		var v uint16 = 57557

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringUint16_Has(t *testing.T) {
	Convey("TestMapStringUint16.Has", t, func() {
		var k string = "69fa21d8-a421-48ba-9b3c-21a4ef465e30"
		var v uint16 = 30739

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("012fcdec-d830-44d6-99d8-40134fed3c6b"+"b475f366-1d0d-4f10-8a48-b27f20c009a9"), ShouldBeFalse)
	})
}

func TestMapStringUint16_Get(t *testing.T) {
	Convey("TestMapStringUint16.Get", t, func() {
		var k string = "cd9963bf-ed6d-406d-b00d-46bb4904c863"
		var v uint16 = 22678

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("6750ad3f-1a89-496d-89a6-c7268f04eb9a" + "367f768e-8f7c-4130-94c2-0995f0252bfb")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringUint16_GetOpt(t *testing.T) {
	Convey("TestMapStringUint16.GetOpt", t, func() {
		var k string = "9a7e585c-dfbb-4d70-86ed-ac61e9023eaf"
		var v uint16 = 21156

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("7a1ce3e6-9cf3-4af3-b422-d2b15db8a704" + "4fbb02e2-2725-4fd2-b6ed-10806e227153")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringUint16_ForEach(t *testing.T) {
	Convey("TestMapStringUint16.ForEach", t, func() {
		var k string = "8f1077e6-0c81-4d09-8194-71f623a25d14"
		var v uint16 = 29883
		hits := 0

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv uint16) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringUint16_MarshalYAML(t *testing.T) {
	Convey("TestMapStringUint16.MarshalYAML", t, func() {
		var k string = "8e7fc9f5-6fc8-4616-bd64-875e906b7caf"
		var v uint16 = 29243

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringUint16_ToYAML(t *testing.T) {
	Convey("TestMapStringUint16.ToYAML", t, func() {
		var k string = "aee0e214-11f0-415a-93d8-7f2c7c669cd3"
		var v uint16 = 37282

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringUint16_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringUint16.PutIfNotNil", t, func() {
		var k string = "f6cde3c9-94df-45b8-a0ad-5d023290d597"
		var v uint16 = 5896

		test := omap.NewMapStringUint16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("90fef73e-1a1c-4f21-b01f-34bfb405b12b", (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint16 = 15950
		So(test.PutIfNotNil("74309d85-b37b-4697-8dda-56be1bf3f5e1", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringUint16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringUint16.ReplaceIfExists", t, func() {
		var k string = "98310d8c-5989-4e66-ae36-b19a48d36960"
		var v uint16 = 19719
		var x uint16 = 18765

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("ab354696-8eb9-455b-b359-9baa55de149d", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringUint16.ReplaceOrPut", t, func() {
		var k string = "35be7237-9fd1-4be1-a00d-ec02f014cd42"
		var v uint16 = 49160
		var x uint16 = 2395

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("d7db25e2-520b-479d-bf33-950158b3ce62", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint16_MarshalJSON(t *testing.T) {
	Convey("TestMapStringUint16.MarshalJSON", t, func() {
		var k string = "4e293160-24dd-4994-887c-666281216b1d"
		var v uint16 = 2503

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"4e293160-24dd-4994-887c-666281216b1d","value":2503}]`)
	})
}
