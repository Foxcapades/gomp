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
		var k string = "921d454b-509b-4201-ad43-9685a5e7d3de"
		var v float64 = 0.502

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringFloat64_Delete(t *testing.T) {
	Convey("TestMapStringFloat64.Delete", t, func() {
		var k string = "9572047b-13a6-4d22-a707-d89c4e7804f3"
		var v float64 = 0.925

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringFloat64_Has(t *testing.T) {
	Convey("TestMapStringFloat64.Has", t, func() {
		var k string = "98ccebe8-fb3f-4352-9db6-b765213d4c60"
		var v float64 = 0.328

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("84c3b30d-3613-4f57-895d-487d0ffaaaa7"+"59882884-9511-4845-8d68-0b0a12db39c0"), ShouldBeFalse)
	})
}

func TestMapStringFloat64_Get(t *testing.T) {
	Convey("TestMapStringFloat64.Get", t, func() {
		var k string = "1f73e898-108b-4ac3-bea2-9e0483e444d5"
		var v float64 = 0.999

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("b43378b0-7dfd-4bc5-bdcb-5fd69dd8e40c" + "91265241-5b40-4d24-b0b8-bbd93950d60a")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringFloat64_GetOpt(t *testing.T) {
	Convey("TestMapStringFloat64.GetOpt", t, func() {
		var k string = "f3ba6532-a340-453c-903c-704f879a7e97"
		var v float64 = 0.263

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("82130870-80a9-475b-b1c1-bf216ad9c39e" + "e7e07464-b6de-464f-b086-e2b6c3580cf3")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringFloat64_ForEach(t *testing.T) {
	Convey("TestMapStringFloat64.ForEach", t, func() {
		var k string = "c5c947db-2fc7-4928-bf5b-2b154b3e6fe5"
		var v float64 = 0.292
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
		var k string = "a2f371c9-26d5-469a-a716-a99380ee0229"
		var v float64 = 0.961

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
		var k string = "8cfe153e-27fe-483d-a46d-232844c02e75"
		var v float64 = 0.605

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
		var k string = "1d08afae-a63b-424b-84a1-427cd92d5e66"
		var v float64 = 0.439

		test := omap.NewMapStringFloat64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("81533d78-08e0-4199-aa7b-b64da1ef7e88", (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float64 = 0.233
		So(test.PutIfNotNil("155c9099-6d20-44d2-b9d5-e7575cfd97dd", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringFloat64.ReplaceIfExists", t, func() {
		var k string = "208d460a-9d5a-4b73-891b-b6703b9baa90"
		var v float64 = 0.090
		var x float64 = 0.741

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("a178a987-0d7f-4417-ba73-122990303c80", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringFloat64.ReplaceOrPut", t, func() {
		var k string = "0d9558b6-f9fa-48b1-b47d-15cae0407c70"
		var v float64 = 0.051
		var x float64 = 0.692

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("35a67402-2f35-48dd-aa99-524c68370ad3", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat64_MarshalJSON(t *testing.T) {
	Convey("TestMapStringFloat64.MarshalJSON", t, func() {
		var k string = "2df85949-d969-4945-9822-f788e3822ffc"
		var v float64 = 0.517

		test := omap.NewMapStringFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"2df85949-d969-4945-9822-f788e3822ffc","value":0.517}]`)
	})
}
