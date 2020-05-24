package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyByte_Put(t *testing.T) {
	Convey("TestMapAnyByte.Put", t, func() {
		var k interface{} = "cfdbf55f-99ac-4911-b8e5-7528f4becde9"
		var v byte = 135

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyByte_Delete(t *testing.T) {
	Convey("TestMapAnyByte.Delete", t, func() {
		var k interface{} = "c6f2b46c-bb27-49b2-b333-d8f5b0f0a578"
		var v byte = 163

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyByte_Has(t *testing.T) {
	Convey("TestMapAnyByte.Has", t, func() {
		var k interface{} = "8606e5db-d8aa-481a-9708-8d9fd4ecb196"
		var v byte = 112

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("b16b9774-40f8-4269-ad05-889edbc94d23"+"58374f77-2718-4fbf-abff-85ad4b8196b7"), ShouldBeFalse)
	})
}

func TestMapAnyByte_Get(t *testing.T) {
	Convey("TestMapAnyByte.Get", t, func() {
		var k interface{} = "8d45f9ee-fdf6-4de2-ab42-3dc4b2098026"
		var v byte = 210

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("cd0509a4-9059-41e7-9d2e-ca67e9f3b966" + "ee515341-bac1-4545-95b1-7c05c8d18664")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyByte_GetOpt(t *testing.T) {
	Convey("TestMapAnyByte.GetOpt", t, func() {
		var k interface{} = "557b08f2-5ca5-498f-93ae-6ea9173181eb"
		var v byte = 98

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("48fe8ca3-6e19-4cd2-969e-45a9ac4a8254" + "e5873710-28f3-4620-ab15-0f843595ec51")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyByte_ForEach(t *testing.T) {
	Convey("TestMapAnyByte.ForEach", t, func() {
		var k interface{} = "bab2eb53-dd66-43e3-9c9e-bceb3bee046a"
		var v byte = 45
		hits := 0

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv byte) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyByte_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyByte.MarshalYAML", t, func() {
		var k interface{} = "84078a75-8dd1-4f01-a087-827ca925fddd"
		var v byte = 162

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyByte_ToYAML(t *testing.T) {
	Convey("TestMapAnyByte.ToYAML", t, func() {
		var k interface{} = "1d090d55-49b2-4047-97a8-411a1d97f28d"
		var v byte = 164

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyByte_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyByte.PutIfNotNil", t, func() {
		var k interface{} = "c0d81e6a-38cd-4ab2-86b7-facaa7cf9163"
		var v byte = 193

		test := omap.NewMapAnyByte(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("d1de8904-e380-4b77-b915-fd709f7be671", (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x byte = 189
		So(test.PutIfNotNil("838d2692-4234-4997-9655-b08d00122209", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyByte_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyByte.ReplaceIfExists", t, func() {
		var k interface{} = "19d731e0-475b-45da-a55a-3feca164c5ba"
		var v byte = 199
		var x byte = 213

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("a55851cf-32bc-4e45-ba67-05f78de87f53", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyByte_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyByte.ReplaceOrPut", t, func() {
		var k interface{} = "130b8eed-7b34-4521-8d4f-229b2adb3c06"
		var v byte = 19
		var x byte = 239

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("b854f561-cb8a-45a4-926e-ee60235ca5c7", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyByte_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyByte.MarshalJSON", t, func() {
		var k interface{} = "009958f9-a655-49a6-ac46-4b81b2615715"
		var v byte = 214

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"009958f9-a655-49a6-ac46-4b81b2615715","value":214}]`)
	})
}
