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
		var k string = "02a5c300-8884-4c76-b955-568bb21c6db2"
		var v uint8 = 188

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringUint8_Delete(t *testing.T) {
	Convey("TestMapStringUint8.Delete", t, func() {
		var k string = "2e2d7d0c-0676-4586-889c-0565a5c6b14b"
		var v uint8 = 28

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringUint8_Has(t *testing.T) {
	Convey("TestMapStringUint8.Has", t, func() {
		var k string = "45dc6269-0758-4cac-9735-fc2ca9651670"
		var v uint8 = 122

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("13cffcae-0122-4a21-b24a-d458ac6ef9fc"+"c653e957-e8d1-4e92-b48d-6ab045c5ae10"), ShouldBeFalse)
	})
}

func TestMapStringUint8_Get(t *testing.T) {
	Convey("TestMapStringUint8.Get", t, func() {
		var k string = "90f00c94-c48f-4837-b6fc-23f0cdeaf2c7"
		var v uint8 = 128

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("ae768a8c-561c-46a9-9140-77ef3495be7f" + "4bd16413-08f5-486b-bc52-c37445c4440e")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringUint8_GetOpt(t *testing.T) {
	Convey("TestMapStringUint8.GetOpt", t, func() {
		var k string = "6e157efe-4188-46d2-9279-95630e1f4852"
		var v uint8 = 242

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("2f84a86c-abc8-4b76-8e67-9bb9e0f9d077" + "23bb6991-d3ad-4865-844a-b0685dbe29f8")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringUint8_ForEach(t *testing.T) {
	Convey("TestMapStringUint8.ForEach", t, func() {
		var k string = "622199d4-9c70-4700-850d-09b96c3bf59b"
		var v uint8 = 37
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
		var k string = "68297513-d69a-4ee9-b222-436064107a07"
		var v uint8 = 56

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
		var k string = "94d5dc20-a0d6-461d-9e85-b32fd7f81783"
		var v uint8 = 221

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
		var k string = "36a3cc62-95de-4dfc-ab90-3c3f9f3523a3"
		var v uint8 = 163

		test := omap.NewMapStringUint8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("dafc1462-5fe0-45c6-a92b-5017d47aeeac", (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint8 = 172
		So(test.PutIfNotNil("1f2665cd-7d0e-4bf9-8a1a-1b9f5ae1d2a8", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringUint8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringUint8.ReplaceIfExists", t, func() {
		var k string = "b43a8da5-b8ef-475e-be2d-666bce502259"
		var v uint8 = 10
		var x uint8 = 39

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("a625fb6d-dac6-4407-ba48-a903921afa4d", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringUint8.ReplaceOrPut", t, func() {
		var k string = "7d24cd51-4bed-4891-9b58-427f2eddc8c2"
		var v uint8 = 124
		var x uint8 = 232

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("8914ebd8-8289-42e6-afeb-3551b017d39f", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint8_MarshalJSON(t *testing.T) {
	Convey("TestMapStringUint8.MarshalJSON", t, func() {
		var k string = "769764f9-23fc-4c24-9814-d5b50e2232fa"
		var v uint8 = 200

		test := omap.NewMapStringUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"769764f9-23fc-4c24-9814-d5b50e2232fa","value":200}]`)
	})
}
