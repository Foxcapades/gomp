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
		var k string = "4fcbafde-7194-4c78-8dd9-399b18aa970c"
		var v uint8 = 130

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringUint8_Delete(t *testing.T) {
	Convey("TestMapStringUint8.Delete", t, func() {
		var k string = "1e4ddd64-95dd-42bb-a73e-9483ce1fbce7"
		var v uint8 = 134

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringUint8_Has(t *testing.T) {
	Convey("TestMapStringUint8.Has", t, func() {
		var k string = "a4e5ef9a-ec3e-4d81-a154-3bd5741801fb"
		var v uint8 = 34

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("57fac86d-9040-4096-8bb3-c1c213af6270"+"7b2d0b65-fd7e-406a-95b3-62e0ed94314c"), ShouldBeFalse)
	})
}

func TestMapStringUint8_Get(t *testing.T) {
	Convey("TestMapStringUint8.Get", t, func() {
		var k string = "7ebe9e43-851b-40a8-ab98-2b9299994206"
		var v uint8 = 178

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("2f011068-db30-4bd7-a9c7-eec7eee81b29" + "d31ac33b-a009-4304-ad2e-6b6a9f310bec")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringUint8_GetOpt(t *testing.T) {
	Convey("TestMapStringUint8.GetOpt", t, func() {
		var k string = "167bb726-b9c8-4f3d-a574-7f751dc32fcc"
		var v uint8 = 84

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("60fe7544-0765-4310-884f-39df37880baf" + "87d36e4d-fc7a-49c2-b854-ffe80a1427f8")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringUint8_ForEach(t *testing.T) {
	Convey("TestMapStringUint8.ForEach", t, func() {
		var k string = "e8875ffe-6d69-463a-a520-02105e4f6989"
		var v uint8 = 252
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
		var k string = "d004f8ec-24c5-4e50-90cf-148d4c0808b5"
		var v uint8 = 26

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
		var k string = "2f73e0bb-2e55-4f87-baed-8f755a2e5a00"
		var v uint8 = 71

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
		var k string = "58e9b0bb-137e-4e95-bc6b-35f8ce86c95c"
		var v uint8 = 141

		test := omap.NewMapStringUint8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("2bd5d7e7-33a6-4328-8449-016b4ad049b1", (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint8 = 216
		So(test.PutIfNotNil("a28d66d2-aa8b-4bcd-8c51-2ec05abbedee", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringUint8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringUint8.ReplaceIfExists", t, func() {
		var k string = "1ed78f47-6c90-457f-8b7e-24d436490b15"
		var v uint8 = 126
		var x uint8 = 168

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("6d642229-08a2-4f3b-a0d0-d467d5ed1a7a", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringUint8.ReplaceOrPut", t, func() {
		var k string = "a7dbe8f2-5f82-42a8-9e7d-bc6758a5b9c7"
		var v uint8 = 84
		var x uint8 = 60

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("d9178dc5-30ab-409b-b8b5-d9904ae758d6", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint8_MarshalJSON(t *testing.T) {
	Convey("TestMapStringUint8.MarshalJSON", t, func() {
		var k string = "bf1539bd-f7f9-435e-9cc1-0c81f7942260"
		var v uint8 = 140

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"bf1539bd-f7f9-435e-9cc1-0c81f7942260","value":140}]`)
	})
}
