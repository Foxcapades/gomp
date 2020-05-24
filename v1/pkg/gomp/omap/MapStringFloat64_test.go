package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringFloat64_Put(t *testing.T) {
	Convey("TestMapStringFloat64.Put", t, func() {
		var k string = "9624aa4f-dc4f-49d4-828b-fabf2451eb26"
		var v float64 = 0.631

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringFloat64_Delete(t *testing.T) {
	Convey("TestMapStringFloat64.Delete", t, func() {
		var k string = "cfbb5ef3-d055-4fa5-8292-2faff39838ca"
		var v float64 = 0.879

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringFloat64_Has(t *testing.T) {
	Convey("TestMapStringFloat64.Has", t, func() {
		var k string = "965300b4-3106-41d5-8c96-d5f3b7826a35"
		var v float64 = 0.182

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("c3f2517f-8b58-4dd6-a81d-6625530e719d"+"8babd327-bc2b-4257-8aa4-8af284c8f9d6"), ShouldBeFalse)
	})
}

func TestMapStringFloat64_Get(t *testing.T) {
	Convey("TestMapStringFloat64.Get", t, func() {
		var k string = "c1fa552c-8087-4071-b542-62fb061339b3"
		var v float64 = 0.343

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("f8031ea3-87fc-4068-a0b4-6cb56f12e86c" + "79ecc4cb-9204-4d03-8536-e815c4ab3db2")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringFloat64_GetOpt(t *testing.T) {
	Convey("TestMapStringFloat64.GetOpt", t, func() {
		var k string = "40f1116e-b389-4370-9816-1aac2ef1487f"
		var v float64 = 0.922

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("dc2e046f-a3e9-4a74-bf1d-a844b521d71f" + "dbdb27f1-1ebf-4376-846a-45eb6929d19d")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringFloat64_ForEach(t *testing.T) {
	Convey("TestMapStringFloat64.ForEach", t, func() {
		var k string = "e8fae116-0553-48ce-8f76-4207a5ed7096"
		var v float64 = 0.501
		hits := 0

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv float64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringFloat64_MarshalYAML(t *testing.T) {
	Convey("TestMapStringFloat64.MarshalYAML", t, func() {
		var k string = "95b95453-9db9-473b-9e45-88b543eee045"
		var v float64 = 0.990

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringFloat64_ToYAML(t *testing.T) {
	Convey("TestMapStringFloat64.ToYAML", t, func() {
		var k string = "a34a67bd-3696-4380-9a6e-9e2acff61b20"
		var v float64 = 0.060

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringFloat64_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringFloat64.PutIfNotNil", t, func() {
		var k string = "e57bca58-846f-46c0-bef9-1d33af41fe30"
		var v float64 = 0.926

		test := omap.NewMapStringFloat64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("8676b77f-6b01-4387-b8b2-01105ddfb937", (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float64 = 0.231
		So(test.PutIfNotNil("c9d821ce-4bc8-4e23-bcc1-e35927c5e73a", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringFloat64.ReplaceIfExists", t, func() {
		var k string = "b75dcda8-db06-424c-9cfc-8cdb34ad0557"
		var v float64 = 0.020
		var x float64 = 0.377

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("c385c5d2-a07a-4667-9b4f-d6dcd4a35d9a", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringFloat64.ReplaceOrPut", t, func() {
		var k string = "9cd18ca4-edb9-4ceb-bbdd-0d044a1da6e9"
		var v float64 = 0.545
		var x float64 = 0.884

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("11625347-8ce5-4e86-b446-b8c590f708d8", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat64_MarshalJSON(t *testing.T) {
	Convey("TestMapStringFloat64.MarshalJSON", t, func() {
		var k string = "858da67f-5ebf-4500-8e2b-c88e13e965c8"
		var v float64 = 0.701

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"858da67f-5ebf-4500-8e2b-c88e13e965c8","value":0.701}]`)
	})
}
