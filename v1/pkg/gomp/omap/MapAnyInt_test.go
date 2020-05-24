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
		var k interface{} = "fdc2f9bb-d4f1-41e4-933a-13e711b9c5c1"
		var v int = 6294427

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyInt_Delete(t *testing.T) {
	Convey("TestMapAnyInt.Delete", t, func() {
		var k interface{} = "6460a458-d9d2-4ee3-9cdb-7eff1c765922"
		var v int = 922645150

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyInt_Has(t *testing.T) {
	Convey("TestMapAnyInt.Has", t, func() {
		var k interface{} = "ff4bdc19-75b5-4dbc-b99d-e0de4c1430ed"
		var v int = 613009498

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("6ad6de5d-325b-4363-8d53-4ea011ed115d"+"0651a44a-7855-4bb4-a10f-381c0bb5ff07"), ShouldBeFalse)
	})
}

func TestMapAnyInt_Get(t *testing.T) {
	Convey("TestMapAnyInt.Get", t, func() {
		var k interface{} = "0cbc7a45-14da-4f1f-af91-8bb0b79c0f56"
		var v int = 2014535282

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("35630b34-8119-4b66-bdb3-2110be6f2a53" + "0deb7db6-5923-4a7e-8232-a86d506e9394")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyInt_GetOpt(t *testing.T) {
	Convey("TestMapAnyInt.GetOpt", t, func() {
		var k interface{} = "f461eb5a-64ee-4a38-a107-ba79cf1755f3"
		var v int = 1846115304

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("99cc4134-c1d3-454e-b00a-08d90d6375f8" + "f3818af5-efaa-45f9-ac4a-1cebea33e32e")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyInt_ForEach(t *testing.T) {
	Convey("TestMapAnyInt.ForEach", t, func() {
		var k interface{} = "20a7f629-6cd7-4101-aad0-c1a972c9631e"
		var v int = 993027470
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
		var k interface{} = "14f4bf27-6d95-48f7-8655-3307debb7470"
		var v int = 1368221880

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
		var k interface{} = "08ca522d-c748-492d-af5d-93d75b910792"
		var v int = 704158015

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
		var k interface{} = "2065c9e2-71bc-4d6e-9655-60ea1ccc7a7a"
		var v int = 1984187360

		test := omap.NewMapAnyInt(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("bab22638-f991-4dc4-80fd-71e752cee4a8", (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int = 371527026
		So(test.PutIfNotNil("feba5f06-6106-440c-8cdd-9e5954cf819e", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyInt.ReplaceIfExists", t, func() {
		var k interface{} = "b20d228c-35d1-4b3e-bda3-158ce881e16b"
		var v int = 1340639846
		var x int = 966260141

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("9456bd35-87ae-4313-a04d-fe3e8092ee19", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyInt.ReplaceOrPut", t, func() {
		var k interface{} = "b8a139de-26e1-456b-ad52-b94efa53a80b"
		var v int = 681219282
		var x int = 807302847

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("1912f333-bb6d-4ad7-863b-bbb98932fb14", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyInt.MarshalJSON", t, func() {
		var k interface{} = "c1aeee8d-9bac-4fdd-bda3-9e116260b658"
		var v int = 44152411

		test := omap.NewMapAnyInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"c1aeee8d-9bac-4fdd-bda3-9e116260b658","value":44152411}]`)
	})
}
