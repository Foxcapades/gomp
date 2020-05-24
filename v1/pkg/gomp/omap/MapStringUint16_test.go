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
		var k string = "edeb6ef7-90aa-47e9-b670-aa56148e2b33"
		var v uint16 = 61776

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringUint16_Delete(t *testing.T) {
	Convey("TestMapStringUint16.Delete", t, func() {
		var k string = "75d590f4-aabe-45d0-b0e0-9ca347c57a2b"
		var v uint16 = 41960

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringUint16_Has(t *testing.T) {
	Convey("TestMapStringUint16.Has", t, func() {
		var k string = "fdcb1452-dda0-4078-8923-e3ca6257417f"
		var v uint16 = 60631

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("b2805424-01e4-4cba-900d-0ad2d0e7f620"+"5145dcd8-1dae-49b3-b1f1-1bee05d96c60"), ShouldBeFalse)
	})
}


func TestMapStringUint16_Get(t *testing.T) {
	Convey("TestMapStringUint16.Get", t, func() {
		var k string = "d7c6ee92-391a-42a8-80b6-0f1a043cb4a7"
		var v uint16 = 60043

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("6808a04d-0e69-4a6b-91bb-2089fb7060fd" + "412b915f-44fd-48a1-87a4-39ea894d68a6")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringUint16_GetOpt(t *testing.T) {
	Convey("TestMapStringUint16.GetOpt", t, func() {
		var k string = "2268e482-59fc-48c9-b257-e30c19b0b2b8"
		var v uint16 = 47906

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("40fc2d12-a9fa-4c44-89d6-262939474db3" + "190f71b9-4b95-49cf-82c2-59b899c53f5a")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringUint16_ForEach(t *testing.T) {
	Convey("TestMapStringUint16.ForEach", t, func() {
		var k string = "d6e896ae-359a-4542-b106-32d5fbe741fa"
		var v uint16 = 40634
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
		var k string = "7a4b387f-5a91-4b42-94fe-3e57b85c2689"
		var v uint16 = 59809

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
		var k string = "cbe29690-2859-4df1-a964-d1b1a736f277"
		var v uint16 = 32862

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
		var k string = "a1b9ca34-7723-459c-be9e-4bdc672a8173"
		var v uint16 = 44752

		test := omap.NewMapStringUint16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("68cd8e49-b27f-400f-9e2f-199218f16695", (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint16 = 61267
		So(test.PutIfNotNil("8917bc4c-89f2-4cb0-bbcb-63b2f9b52b2a", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringUint16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringUint16.ReplaceIfExists", t, func() {
		var k string = "5b8d4b02-ad6d-4de7-8f2c-265a8185b784"
		var v uint16 = 17166
		var x uint16 = 9956

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("65d7ee63-2a3e-4cbb-bae1-3a616b704c03", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringUint16.ReplaceOrPut", t, func() {
		var k string = "db59a4c0-79cd-45da-934a-0dbf1e0eea0c"
		var v uint16 = 31290
		var x uint16 = 52659

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("e0c3de4f-453f-46e0-84f1-6b776c6ea704", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint16_MarshalJSON(t *testing.T) {
	Convey("TestMapStringUint16.MarshalJSON", t, func() {
		var k string = "a794582b-c636-4a67-9528-8290a35a7f42"
		var v uint16 = 38487

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"a794582b-c636-4a67-9528-8290a35a7f42","value":38487}]`)
	})
}

