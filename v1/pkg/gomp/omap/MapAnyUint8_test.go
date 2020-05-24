package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyUint8_Put(t *testing.T) {
	Convey("TestMapAnyUint8.Put", t, func() {
		var k interface{} = "bf046111-899e-4bf6-9656-f59f50e27509"
		var v uint8 = 22

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyUint8_Delete(t *testing.T) {
	Convey("TestMapAnyUint8.Delete", t, func() {
		var k interface{} = "f4c71d46-5a36-4c8d-8e26-94e9de45f550"
		var v uint8 = 118

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyUint8_Has(t *testing.T) {
	Convey("TestMapAnyUint8.Has", t, func() {
		var k interface{} = "7b8f560d-3eb3-4f7a-8328-dc19489bbbd5"
		var v uint8 = 176

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("38ee5122-5f98-4134-8f85-57803d8c086b"+"3a204383-f82e-4cad-94aa-aef21a2ec01b"), ShouldBeFalse)
	})
}

func TestMapAnyUint8_Get(t *testing.T) {
	Convey("TestMapAnyUint8.Get", t, func() {
		var k interface{} = "16ccc725-11c8-40d0-a07e-0229425aad40"
		var v uint8 = 10

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("002fd53b-d22e-4c76-82b8-dce8414b35c2" + "3818a1cf-84b6-45ca-8a82-d1dc19da22db")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyUint8_GetOpt(t *testing.T) {
	Convey("TestMapAnyUint8.GetOpt", t, func() {
		var k interface{} = "bec5a7e8-ec84-4895-891e-703283d77d80"
		var v uint8 = 99

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("34d8c034-4333-4ffc-b1b6-fc1fc20d7630" + "6b72aebf-418e-4a4c-b533-cf7f52bc25ef")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyUint8_ForEach(t *testing.T) {
	Convey("TestMapAnyUint8.ForEach", t, func() {
		var k interface{} = "84dc6968-c63e-47a6-b523-0f777a1d5549"
		var v uint8 = 39
		hits := 0

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv uint8) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyUint8_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyUint8.MarshalYAML", t, func() {
		var k interface{} = "8fd77492-1fc3-4535-acf8-7462415ebfdf"
		var v uint8 = 143

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyUint8_ToYAML(t *testing.T) {
	Convey("TestMapAnyUint8.ToYAML", t, func() {
		var k interface{} = "66ab5ecc-cba5-45b8-80c2-5da2402ee998"
		var v uint8 = 40

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyUint8_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyUint8.PutIfNotNil", t, func() {
		var k interface{} = "f5146df4-4768-46e0-94b7-5e421604e540"
		var v uint8 = 31

		test := omap.NewMapAnyUint8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("7a88a0f5-97b8-48c5-85f5-c401ada6261b", (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint8 = 187
		So(test.PutIfNotNil("64357375-8d5d-4113-92c0-e55f569db473", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyUint8.ReplaceIfExists", t, func() {
		var k interface{} = "ff2b865c-c144-441d-ba30-c6cd63709130"
		var v uint8 = 112
		var x uint8 = 91

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("2af2349a-67c6-486a-8f53-4663c9ca2975", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyUint8.ReplaceOrPut", t, func() {
		var k interface{} = "5a09a295-fc72-4a4c-a99a-4fab24313263"
		var v uint8 = 10
		var x uint8 = 202

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("5e37ab0d-77a0-41e7-bc61-877bb6d51a6d", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint8_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyUint8.MarshalJSON", t, func() {
		var k interface{} = "b51bfa9e-6b91-4b1f-8812-3924b2a20e79"
		var v uint8 = 164

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"b51bfa9e-6b91-4b1f-8812-3924b2a20e79","value":164}]`)
	})
}
