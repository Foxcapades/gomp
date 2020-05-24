package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntAny_Put(t *testing.T) {
	Convey("TestMapIntAny.Put", t, func() {
		var k int = 1840709788
		var v interface{} = "d0021003-6a10-4830-93ae-4abf1743c81e"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntAny_Delete(t *testing.T) {
	Convey("TestMapIntAny.Delete", t, func() {
		var k int = 667518929
		var v interface{} = "3379c4c7-e0ef-4b7c-8af6-72dfc5a975ab"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntAny_Has(t *testing.T) {
	Convey("TestMapIntAny.Has", t, func() {
		var k int = 1098435204
		var v interface{} = "c367045e-c7e5-4bf7-bdc8-4aff8dfd9438"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(530770805+2027084430), ShouldBeFalse)
	})
}

func TestMapIntAny_Get(t *testing.T) {
	Convey("TestMapIntAny.Get", t, func() {
		var k int = 1452863250
		var v interface{} = "ec2dc9a9-6b94-48e3-91d2-c331d53ecefb"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(420529885 + 91220924)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntAny_GetOpt(t *testing.T) {
	Convey("TestMapIntAny.GetOpt", t, func() {
		var k int = 1552805157
		var v interface{} = "f185c3e9-d97a-4457-ab12-f5ca31ef07e2"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(176687207 + 366043566)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntAny_ForEach(t *testing.T) {
	Convey("TestMapIntAny.ForEach", t, func() {
		var k int = 1757750331
		var v interface{} = "aa669186-48a8-4ce0-9280-ae960030e31c"
		hits := 0

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv interface{}) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntAny_MarshalYAML(t *testing.T) {
	Convey("TestMapIntAny.MarshalYAML", t, func() {
		var k int = 444227702
		var v interface{} = "fda9e1d8-f8d7-4c34-965e-b136057363c5"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntAny_ToYAML(t *testing.T) {
	Convey("TestMapIntAny.ToYAML", t, func() {
		var k int = 1827283102
		var v interface{} = "708937dc-78d1-44ea-9678-7ec16b700402"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntAny_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntAny.PutIfNotNil", t, func() {
		var k int = 1892035773
		var v interface{} = "e254340d-7ad9-409d-a235-28c9ec1e69c6"

		test := omap.NewMapIntAny(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1895785286, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "9f7aac86-af1b-4cbc-8b83-ea3b2013c522"
		So(test.PutIfNotNil(1588292206, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntAny_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntAny.ReplaceIfExists", t, func() {
		var k int = 1005266180
		var v interface{} = "80733d9b-bf76-4cad-951a-059ad15395ea"
		var x interface{} = "d8278841-acbd-4a31-b40e-50c130a3d219"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(542909461, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntAny_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntAny.ReplaceOrPut", t, func() {
		var k int = 906104955
		var v interface{} = "7430f8c2-785e-48c8-9ea5-cdd1f2253035"
		var x interface{} = "f49c6c83-3c6b-4b9f-bb03-9a4c33f53a3c"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1423197165, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntAny_MarshalJSON(t *testing.T) {
	Convey("TestMapIntAny.MarshalJSON", t, func() {
		var k int = 1786716199
		var v interface{} = "42a4fdbd-0646-476a-b259-1c6beaebc0c2"

		test := omap.NewMapIntAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1786716199,"value":"42a4fdbd-0646-476a-b259-1c6beaebc0c2"}]`)
	})
}
