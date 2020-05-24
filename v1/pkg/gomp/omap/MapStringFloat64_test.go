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
		var k string = "370352f0-224e-461f-8cd2-02d423fc6e9b"
		var v float64 = 0.629

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringFloat64_Delete(t *testing.T) {
	Convey("TestMapStringFloat64.Delete", t, func() {
		var k string = "1c2e6fe9-039d-4ef3-8075-1be06afee5f5"
		var v float64 = 0.809

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringFloat64_Has(t *testing.T) {
	Convey("TestMapStringFloat64.Has", t, func() {
		var k string = "740e0c93-0102-44ef-a82b-b7db67af6bdc"
		var v float64 = 0.084

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("4de3209e-0f87-4d24-b70b-7b058cd2ccd9"+"6f0d2046-dee6-440c-a426-9c12d19ddd0c"), ShouldBeFalse)
	})
}

func TestMapStringFloat64_Get(t *testing.T) {
	Convey("TestMapStringFloat64.Get", t, func() {
		var k string = "defdcc36-b24c-406d-ac9e-3a3ab16066f4"
		var v float64 = 0.402

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("823ae762-f03d-4741-ba29-264bd0f51f79" + "ba3e80f3-2474-4d28-bb60-a895085bd1d9")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringFloat64_GetOpt(t *testing.T) {
	Convey("TestMapStringFloat64.GetOpt", t, func() {
		var k string = "17b730df-4f9d-42c2-94db-dc99288147ca"
		var v float64 = 0.740

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("75a88db1-2993-46c2-87b7-d7d4bd158a39" + "c64748a6-e111-4fc4-9c2e-4d7b1e958491")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringFloat64_ForEach(t *testing.T) {
	Convey("TestMapStringFloat64.ForEach", t, func() {
		var k string = "08eede03-266d-4f9c-aae7-7dbbf97ba768"
		var v float64 = 0.469
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
		var k string = "202e5c33-9a9d-476a-93fb-8647d7aa7d72"
		var v float64 = 0.306

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
		var k string = "ea656aa3-d851-4333-909e-fdd38d8723d5"
		var v float64 = 0.895

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
		var k string = "c912d011-bc2f-4f2c-8a8c-621cde327acb"
		var v float64 = 0.999

		test := omap.NewMapStringFloat64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("b3c9ae94-fc5b-47fb-90ae-e423d815ca73", (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float64 = 0.750
		So(test.PutIfNotNil("ac0d08df-e220-4543-b2ce-de454df1502d", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringFloat64.ReplaceIfExists", t, func() {
		var k string = "697d9836-dda2-49a0-8660-ed40a401fa01"
		var v float64 = 0.317
		var x float64 = 0.036

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("ea860261-e90d-4ae4-ba7a-e2f6d27ba21b", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringFloat64.ReplaceOrPut", t, func() {
		var k string = "0a225e83-c4bf-4509-910f-d034bb154a6d"
		var v float64 = 0.957
		var x float64 = 0.466

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("f1a35a90-da12-47e4-b06d-2119ae1d360b", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat64_MarshalJSON(t *testing.T) {
	Convey("TestMapStringFloat64.MarshalJSON", t, func() {
		var k string = "a4342e37-f65c-48ef-976f-00f64a5e479c"
		var v float64 = 0.412

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"a4342e37-f65c-48ef-976f-00f64a5e479c","value":0.412}]`)
	})
}
