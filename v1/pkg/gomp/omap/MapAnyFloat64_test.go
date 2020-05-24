package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyFloat64_Put(t *testing.T) {
	Convey("TestMapAnyFloat64.Put", t, func() {
		var k interface{} = "e8cfab99-3e9f-49db-be97-a55f33effbda"
		var v float64 = 0.591

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyFloat64_Delete(t *testing.T) {
	Convey("TestMapAnyFloat64.Delete", t, func() {
		var k interface{} = "cf2313b9-22d4-45e2-8a09-b593008b36a0"
		var v float64 = 0.995

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyFloat64_Has(t *testing.T) {
	Convey("TestMapAnyFloat64.Has", t, func() {
		var k interface{} = "9c0cae1f-22aa-46c9-90f7-23cdc416c03a"
		var v float64 = 0.398

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("581a3f3f-c675-4b85-9e3c-9988e6ec7eca"+"7978670e-29a5-4458-a77c-ef0d9cd6d92e"), ShouldBeFalse)
	})
}

func TestMapAnyFloat64_Get(t *testing.T) {
	Convey("TestMapAnyFloat64.Get", t, func() {
		var k interface{} = "673defe7-eac6-417d-bab8-2ebe9a7a3203"
		var v float64 = 0.645

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("88ddee4e-a94a-44bd-ab1b-360b9c1e76d1" + "63338a8c-7531-4062-a0ea-ca43633ab705")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyFloat64_GetOpt(t *testing.T) {
	Convey("TestMapAnyFloat64.GetOpt", t, func() {
		var k interface{} = "a5ee2bcb-d3fe-4458-92b4-43c9b0507da3"
		var v float64 = 0.650

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("7e38af8d-05a1-4307-97ae-ca57b1d7f275" + "d7bce286-2e29-422c-a9ce-c66af1e364c1")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyFloat64_ForEach(t *testing.T) {
	Convey("TestMapAnyFloat64.ForEach", t, func() {
		var k interface{} = "f2eb9f25-7a0c-41cb-b602-716578d4c80a"
		var v float64 = 0.538
		hits := 0

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv float64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyFloat64_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyFloat64.MarshalYAML", t, func() {
		var k interface{} = "657e76fc-7ca3-4a4c-8713-dd36c1d7700b"
		var v float64 = 0.373

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyFloat64_ToYAML(t *testing.T) {
	Convey("TestMapAnyFloat64.ToYAML", t, func() {
		var k interface{} = "f083ecf4-3cbb-4f87-a90d-f5eb2aebb487"
		var v float64 = 0.249

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyFloat64_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyFloat64.PutIfNotNil", t, func() {
		var k interface{} = "24072981-0e3e-4f1b-8e7e-fbb172970729"
		var v float64 = 0.274

		test := omap.NewMapAnyFloat64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("edf95c8e-72f6-4bf6-9ffa-e1e005be001b", (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float64 = 0.697
		So(test.PutIfNotNil("022be7e3-fddc-4421-aaa4-2d8c08dc5b1b", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyFloat64.ReplaceIfExists", t, func() {
		var k interface{} = "94a454bc-bbf9-49c6-89d8-ad6abc6367e0"
		var v float64 = 0.558
		var x float64 = 0.331

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("a19c7bee-a295-456d-b4d4-4d1f12d599b3", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyFloat64.ReplaceOrPut", t, func() {
		var k interface{} = "9f1429fc-3f47-473d-9e47-ae709e365f56"
		var v float64 = 0.649
		var x float64 = 0.056

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("e9ffe3ad-f805-4541-96b7-9eada7e01c5e", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat64_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyFloat64.MarshalJSON", t, func() {
		var k interface{} = "2c01c22a-040c-4ce1-87a2-2ca2ef9d685d"
		var v float64 = 0.477

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"2c01c22a-040c-4ce1-87a2-2ca2ef9d685d","value":0.477}]`)
	})
}
