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
		var k interface{} = "b84f7947-c60b-4f56-97fc-62d112f3e7ad"
		var v uint8 = 10

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyUint8_Delete(t *testing.T) {
	Convey("TestMapAnyUint8.Delete", t, func() {
		var k interface{} = "95a788af-3a8d-4500-9045-ecf4076724fe"
		var v uint8 = 95

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyUint8_Has(t *testing.T) {
	Convey("TestMapAnyUint8.Has", t, func() {
		var k interface{} = "38a103fa-150c-4da3-9c04-a3dab4915ebc"
		var v uint8 = 172

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("6f13eb38-0ef6-4a73-ba96-9236e2a516d9"+"8ea5867e-7e7b-459c-990d-700583300e36"), ShouldBeFalse)
	})
}


func TestMapAnyUint8_Get(t *testing.T) {
	Convey("TestMapAnyUint8.Get", t, func() {
		var k interface{} = "785902a7-4ea5-40c1-b46a-1991541da652"
		var v uint8 = 48

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("430c5bba-58fb-429b-afb3-2c342437ff63"+"aad06d02-f058-4aeb-8026-b520ecdcf4c4")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyUint8_GetOpt(t *testing.T) {
	Convey("TestMapAnyUint8.GetOpt", t, func() {
		var k interface{} = "fd2882ca-1296-4ae3-8dbc-8466355c973f"
		var v uint8 = 104

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("9c6946e6-05ff-4925-bcde-f63723f223a2"+"c3407793-7c45-476f-902d-dec59884e13f")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyUint8_ForEach(t *testing.T) {
	Convey("TestMapAnyUint8.ForEach", t, func() {
		var k interface{} = "12295ab2-2275-42d2-9f03-cc4f9391d693"
		var v uint8 = 71
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
		var k interface{} = "8943266e-9e78-49d2-b249-f66af721750b"
		var v uint8 = 18

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
		var k interface{} = "79e4c25b-dc30-4520-804b-a493a3bc2352"
		var v uint8 = 251

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
		var k interface{} = "03adea2c-2495-49d6-ae3f-01401978368f"
		var v uint8 = 18

		test := omap.NewMapAnyUint8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("f4e159d6-bd3a-4367-b054-a9497e3186ef", (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint8 = 217
		So(test.PutIfNotNil("0f3b3b1b-156c-4b35-9aa6-2bfbfb359136", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyUint8.ReplaceIfExists", t, func() {
		var k interface{} = "2807ab79-372f-4580-924b-d53fcd738fee"
		var v uint8 = 232
		var x uint8 = 169

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("4bff0be9-cc4d-47cb-a297-cffcf1e71548", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyUint8.ReplaceOrPut", t, func() {
		var k interface{} = "2fd2af8a-7891-4266-9eba-fdfea8a3a0b3"
		var v uint8 = 168
		var x uint8 = 95

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("bfb1cb1a-e61e-461f-8faa-7fe448d9ef09", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint8_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyUint8.MarshalJSON", t, func() {
		var k interface{} = "1c9f1343-67f1-4eea-819c-7c6dd61caf93"
		var v uint8 = 2

		test := omap.NewMapAnyUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"1c9f1343-67f1-4eea-819c-7c6dd61caf93","value":2}]`)
	})
}

