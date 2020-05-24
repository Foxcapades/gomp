package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringString_Put(t *testing.T) {
	Convey("TestMapStringString.Put", t, func() {
		var k string = "eaffc695-fae6-480f-8bad-1e69492fdfe3"
		var v string = "3f2205f1-8ad5-420d-bbdd-9273640b5d33"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringString_Delete(t *testing.T) {
	Convey("TestMapStringString.Delete", t, func() {
		var k string = "c1b4f9e6-419d-4552-92a7-8caf49d571f0"
		var v string = "9ba7e794-2c04-4334-b2e1-7f5dfc1b780a"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringString_Has(t *testing.T) {
	Convey("TestMapStringString.Has", t, func() {
		var k string = "03319544-ccba-4abf-8e5b-b619a7a20f67"
		var v string = "70658dce-22ea-42c4-8f64-be616b99f1da"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("603ae8a2-b7b5-4f62-812b-b109e4fd50c2"+"f63edd6f-c54e-4344-9275-637de0acc0e5"), ShouldBeFalse)
	})
}

func TestMapStringString_Get(t *testing.T) {
	Convey("TestMapStringString.Get", t, func() {
		var k string = "a5eb6fdf-d6e0-432c-b768-9a3acaee2bf1"
		var v string = "3c81e769-4a1c-4aa8-a9d8-15c2f1035196"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("b034041e-c02e-406c-93e7-c9b1fd0a5e10" + "3bb26353-f2fc-4982-b3ad-d004f90bcab8")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringString_GetOpt(t *testing.T) {
	Convey("TestMapStringString.GetOpt", t, func() {
		var k string = "5bcc8f26-a0ce-4672-af5c-20371ddea8a8"
		var v string = "41e31fcb-25f1-419b-a816-88bfd9c08464"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("99cbb0eb-e72d-4be6-a637-8d4c16d5e524" + "8bdf5e2c-7d32-4883-8d6f-4a7e8b4506e3")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringString_ForEach(t *testing.T) {
	Convey("TestMapStringString.ForEach", t, func() {
		var k string = "a1fe80a1-1bda-46c7-aa5b-4445cbe08fef"
		var v string = "6ce003f1-33a8-4aaa-bdfd-ed581b602895"
		hits := 0

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv string) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringString_MarshalYAML(t *testing.T) {
	Convey("TestMapStringString.MarshalYAML", t, func() {
		var k string = "69bb2ee1-f541-4821-b2d3-f7cde2b87737"
		var v string = "2234f766-efdc-43a1-af2b-3080c3005889"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringString_ToYAML(t *testing.T) {
	Convey("TestMapStringString.ToYAML", t, func() {
		var k string = "a41dc54a-06cc-446b-8546-d3a914042472"
		var v string = "4908f2bd-a832-42bf-881f-7b8e8b22b73c"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringString_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringString.PutIfNotNil", t, func() {
		var k string = "5aac0ee7-2df1-4086-a111-6c4cf67f0336"
		var v string = "8d4acc4f-f431-4d39-b72c-55d687b373d5"

		test := omap.NewMapStringString(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("9fb1bc74-3374-487b-a57a-4aaf36103cc5", (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "d784aff4-eb91-455e-93d3-4e23dae0747c"
		So(test.PutIfNotNil("2a0a1e85-e6c8-4aaf-bb10-aa0d90510e43", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringString_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringString.ReplaceIfExists", t, func() {
		var k string = "63b285e3-8bfe-44d1-833e-142652810a99"
		var v string = "33d22746-12c6-4aec-9fc4-8132fde90532"
		var x string = "df61bcd1-c3fe-43a8-8354-d299131a3536"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("f8a9ad73-ea7d-44c7-ab76-456edbe9dae9", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringString_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringString.ReplaceOrPut", t, func() {
		var k string = "27c42b9b-9040-4186-9119-e64f18fef9ca"
		var v string = "0e0cd579-a9b4-4254-9205-156bbeddbc09"
		var x string = "f29149ab-ddb9-4a16-a4a8-8518b40bccac"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("8f7d72f3-5241-4b5d-80e0-50079004a8e7", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringString_MarshalJSON(t *testing.T) {
	Convey("TestMapStringString.MarshalJSON", t, func() {
		var k string = "71c41c1b-02a3-49be-9c1b-0bc1abf5f5d7"
		var v string = "e5d2153b-03b4-4099-b782-45f087ea4b31"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"71c41c1b-02a3-49be-9c1b-0bc1abf5f5d7","value":"e5d2153b-03b4-4099-b782-45f087ea4b31"}]`)
	})
}
