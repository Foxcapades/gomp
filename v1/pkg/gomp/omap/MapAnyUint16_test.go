package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyUint16_Put(t *testing.T) {
	Convey("TestMapAnyUint16.Put", t, func() {
		var k interface{} = "c1f7dfd6-608b-4647-b057-4314a3a5ba6d"
		var v uint16 = 59556

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyUint16_Delete(t *testing.T) {
	Convey("TestMapAnyUint16.Delete", t, func() {
		var k interface{} = "983e0ba3-7fbb-40bb-95d9-8c01956f96b8"
		var v uint16 = 57646

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyUint16_Has(t *testing.T) {
	Convey("TestMapAnyUint16.Has", t, func() {
		var k interface{} = "15d40ab1-4abd-4a94-a5db-7ae37dd1aecb"
		var v uint16 = 39704

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("4f4c485d-4e92-4704-9046-c7208692aea9"+"90069836-2385-4167-b416-ac274a5b745b"), ShouldBeFalse)
	})
}


func TestMapAnyUint16_Get(t *testing.T) {
	Convey("TestMapAnyUint16.Get", t, func() {
		var k interface{} = "ae2b9f8e-f5d9-49ec-958a-23b53f35b841"
		var v uint16 = 60412

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("e5bbd9bf-70f0-457c-9c3c-63c856de6f70" + "62e12474-1e6c-4acd-9b6e-30e5a3e3e50d")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyUint16_GetOpt(t *testing.T) {
	Convey("TestMapAnyUint16.GetOpt", t, func() {
		var k interface{} = "f58c9082-6f07-4d45-9ed3-bf6cdb511e33"
		var v uint16 = 44773

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("8f876e83-2d6e-4fc2-a115-6ff7d3b2cd79" + "81bfb8c9-1086-4bec-9449-1a768908a433")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyUint16_ForEach(t *testing.T) {
	Convey("TestMapAnyUint16.ForEach", t, func() {
		var k interface{} = "c8effda4-4745-4bfe-ba91-fb5b90ccc8bb"
		var v uint16 = 62383
		hits := 0

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv uint16) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyUint16_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyUint16.MarshalYAML", t, func() {
		var k interface{} = "2ab851e9-b409-4aa1-b2a5-6628c8d2e394"
		var v uint16 = 17085

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyUint16_ToYAML(t *testing.T) {
	Convey("TestMapAnyUint16.ToYAML", t, func() {
		var k interface{} = "c5701de7-b545-4144-aef4-64192cb4f4f1"
		var v uint16 = 13614

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyUint16_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyUint16.PutIfNotNil", t, func() {
		var k interface{} = "6062e5c2-8c2e-4177-988a-fe0276615da2"
		var v uint16 = 23634

		test := omap.NewMapAnyUint16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("239f1eb0-9c03-4b30-a26c-a77b36aafda8", (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint16 = 54144
		So(test.PutIfNotNil("b5a3c5f7-9180-47fe-90ad-4fc2b9b28692", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyUint16.ReplaceIfExists", t, func() {
		var k interface{} = "5026275a-c165-4f79-b35f-c4536cf786fe"
		var v uint16 = 57873
		var x uint16 = 11824

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("c22add9e-b36d-4d80-a7fa-807243de5f84", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyUint16.ReplaceOrPut", t, func() {
		var k interface{} = "c932c66b-aacb-447c-9b4b-0ba5805619f7"
		var v uint16 = 42207
		var x uint16 = 57013

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("b834cac7-33fd-456e-97a8-35875d5911e4", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint16_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyUint16.MarshalJSON", t, func() {
		var k interface{} = "84ebffee-762d-4d15-b58f-ba21624f5d08"
		var v uint16 = 22438

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"84ebffee-762d-4d15-b58f-ba21624f5d08","value":22438}]`)
	})
}
