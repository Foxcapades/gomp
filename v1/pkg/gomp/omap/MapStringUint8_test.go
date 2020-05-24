package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringUint8_Put(t *testing.T) {
	Convey("TestMapStringUint8.Put", t, func() {
		var k string = "bd93707c-544c-4a98-a3d9-4fe90e0459bf"
		var v uint8 = 75

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringUint8_Delete(t *testing.T) {
	Convey("TestMapStringUint8.Delete", t, func() {
		var k string = "34efff82-d55f-4599-bfb7-7bcbdda210f6"
		var v uint8 = 122

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringUint8_Has(t *testing.T) {
	Convey("TestMapStringUint8.Has", t, func() {
		var k string = "066c3e11-fcc8-4e2b-8a2c-550af04a2285"
		var v uint8 = 175

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("6c50a50c-9f2c-490d-9b24-29e9fb16280a"+"1d11012c-3559-47c5-bdb3-0775268d5f55"), ShouldBeFalse)
	})
}


func TestMapStringUint8_Get(t *testing.T) {
	Convey("TestMapStringUint8.Get", t, func() {
		var k string = "da3f6692-4580-432d-8470-430e6cc1f673"
		var v uint8 = 231

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("e543812d-2155-4aac-aaf9-a0ece7cf2427" + "0a2c9870-9050-441f-a1bd-357ee1de3314")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringUint8_GetOpt(t *testing.T) {
	Convey("TestMapStringUint8.GetOpt", t, func() {
		var k string = "4c8f4142-a265-4b84-90a9-647407ca71eb"
		var v uint8 = 48

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("7b2ed5b1-48d4-4b79-943b-fe94f28432e9" + "c1e4b6bc-c759-4fa4-ae4f-aace2ad14627")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringUint8_ForEach(t *testing.T) {
	Convey("TestMapStringUint8.ForEach", t, func() {
		var k string = "8d9bf51e-45a0-42ff-95df-e07c71fd3126"
		var v uint8 = 120
		hits := 0

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv uint8) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringUint8_MarshalYAML(t *testing.T) {
	Convey("TestMapStringUint8.MarshalYAML", t, func() {
		var k string = "3b29dbb4-cd05-4175-9b78-243e51dad81d"
		var v uint8 = 240

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringUint8_ToYAML(t *testing.T) {
	Convey("TestMapStringUint8.ToYAML", t, func() {
		var k string = "24a30056-9ae9-447d-8b76-921db4a6bbbd"
		var v uint8 = 27

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringUint8_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringUint8.PutIfNotNil", t, func() {
		var k string = "e18d0a67-a648-4d1d-b392-f958dde60883"
		var v uint8 = 46

		test := omap.NewMapStringUint8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("4f6b282e-3bd2-4f55-832d-2f8355e9ee4d", (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint8 = 1
		So(test.PutIfNotNil("65ec8ae8-8ffe-4fd3-8f71-49300abdbc8a", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringUint8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringUint8.ReplaceIfExists", t, func() {
		var k string = "6eec2047-348e-4580-a412-0468e8e9ebcf"
		var v uint8 = 74
		var x uint8 = 191

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("bd92c7b0-9a84-4f15-9d2b-530f807a730d", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringUint8.ReplaceOrPut", t, func() {
		var k string = "8b1458cc-6170-4a73-b0b9-fcdf19b7e2c0"
		var v uint8 = 130
		var x uint8 = 191

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("2ad0573f-a2ec-4ccb-a43d-e19e134d32d9", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint8_MarshalJSON(t *testing.T) {
	Convey("TestMapStringUint8.MarshalJSON", t, func() {
		var k string = "6dfb7fd4-5d5e-4b68-b4bb-9f03d701a790"
		var v uint8 = 186

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"6dfb7fd4-5d5e-4b68-b4bb-9f03d701a790","value":186}]`)
	})
}
