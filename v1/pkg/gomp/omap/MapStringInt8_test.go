package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringInt8_Put(t *testing.T) {
	Convey("TestMapStringInt8.Put", t, func() {
		var k string = "6d1dc19d-69b6-457d-8d13-ed89942619cc"
		var v int8 = 72

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringInt8_Delete(t *testing.T) {
	Convey("TestMapStringInt8.Delete", t, func() {
		var k string = "59215518-c098-4f85-8e69-a803a0521424"
		var v int8 = 58

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringInt8_Has(t *testing.T) {
	Convey("TestMapStringInt8.Has", t, func() {
		var k string = "ce23f7b8-b6c7-45a2-a1a9-ca27a2881aa8"
		var v int8 = 120

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("d5691e4b-21f8-4973-9371-1db491e46803"+"a5e29aad-dfcf-4f2a-9199-ce7199d78160"), ShouldBeFalse)
	})
}

func TestMapStringInt8_Get(t *testing.T) {
	Convey("TestMapStringInt8.Get", t, func() {
		var k string = "9f64a523-2228-43aa-b703-bb4ba3fd9562"
		var v int8 = 84

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("5bda0385-843f-4cbe-b04d-9c22fc6687a2" + "57b6c021-7722-4779-9456-581dd129aa34")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringInt8_GetOpt(t *testing.T) {
	Convey("TestMapStringInt8.GetOpt", t, func() {
		var k string = "dc4fc8a8-25e4-4417-9d15-65b074dc33fd"
		var v int8 = 45

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("d4227978-9979-4b68-995a-2d211fa92fdb" + "15f3fba3-d4f9-4b32-8149-a812dc2fe67e")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringInt8_ForEach(t *testing.T) {
	Convey("TestMapStringInt8.ForEach", t, func() {
		var k string = "c43c91ac-f556-4cc5-b51c-44ce3e53a6af"
		var v int8 = 7
		hits := 0

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv int8) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringInt8_MarshalYAML(t *testing.T) {
	Convey("TestMapStringInt8.MarshalYAML", t, func() {
		var k string = "6ec95fee-1213-46d4-8d7b-3cc910f2736a"
		var v int8 = 65

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringInt8_ToYAML(t *testing.T) {
	Convey("TestMapStringInt8.ToYAML", t, func() {
		var k string = "bf42ae99-ac89-44e6-a068-08f9bcfc629e"
		var v int8 = 69

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.SequenceNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 1)
		So(xyml.IsMap(c.Content[0]), ShouldBeTrue)
	})
}

func TestMapStringInt8_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringInt8.PutIfNotNil", t, func() {
		var k string = "0440ea62-1c65-4f0c-968c-3c18f75d7040"
		var v int8 = 17

		test := omap.NewMapStringInt8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("54f12e77-d0fc-4dbf-b5d4-92bcd9f806ff", (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int8 = 109
		So(test.PutIfNotNil("3e746e10-356d-437a-b96b-7c69f8081e87", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringInt8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringInt8.ReplaceIfExists", t, func() {
		var k string = "5b498943-0218-4ce4-9e1f-2b7dc2550a5b"
		var v int8 = 67
		var x int8 = 85

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("8ad9f7fe-94c8-4ccd-8902-8d34eef1f812", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringInt8.ReplaceOrPut", t, func() {
		var k string = "6cd739e1-196d-4687-8e48-a232443bff01"
		var v int8 = 14
		var x int8 = 117

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("e815e52d-b17c-4046-ae4c-fedb7ec2974e", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt8_MarshalJSON(t *testing.T) {
	Convey("TestMapStringInt8.MarshalJSON", t, func() {
		var k string = "8f0597d2-2f01-413f-bf0d-9324b24c31b8"
		var v int8 = 82

		test := omap.NewMapStringInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"8f0597d2-2f01-413f-bf0d-9324b24c31b8","value":82}]`)
	})
}
