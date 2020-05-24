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
		var k interface{} = "a4ad41fe-6664-44ae-9fa7-932f076a4c78"
		var v float64 = 0.277

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyFloat64_Delete(t *testing.T) {
	Convey("TestMapAnyFloat64.Delete", t, func() {
		var k interface{} = "a02b3ba2-2466-442b-ba3c-2df00a371616"
		var v float64 = 0.903

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyFloat64_Has(t *testing.T) {
	Convey("TestMapAnyFloat64.Has", t, func() {
		var k interface{} = "bd554668-916c-4ea4-aa77-a761fd54c4bb"
		var v float64 = 0.004

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("9a6d9335-afbf-48a0-9054-7c3be9026bc3"+"43d4b10e-5dd6-4ac2-bae3-aece6ab14fcc"), ShouldBeFalse)
	})
}


func TestMapAnyFloat64_Get(t *testing.T) {
	Convey("TestMapAnyFloat64.Get", t, func() {
		var k interface{} = "6c184f94-890f-43d8-a4de-cbfa976159a9"
		var v float64 = 0.211

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("dd5e87b3-625d-4f22-aef3-3d485d8fdbff" + "1e15e2f0-4bb2-4c82-9053-f957cd1afd20")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyFloat64_GetOpt(t *testing.T) {
	Convey("TestMapAnyFloat64.GetOpt", t, func() {
		var k interface{} = "117fcce3-1fb0-496f-b7e9-5d99dcc26835"
		var v float64 = 0.048

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("1fdfcb48-24ba-4813-888d-830f5b01e222" + "a20655ea-42a7-4439-a0bc-44af05342461")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyFloat64_ForEach(t *testing.T) {
	Convey("TestMapAnyFloat64.ForEach", t, func() {
		var k interface{} = "e7d09d78-5a08-4c49-9ae0-f84bc69ed3b3"
		var v float64 = 0.159
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
		var k interface{} = "abf48d74-a8dc-43e0-bd25-4cb2308c4bf5"
		var v float64 = 0.381

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
		var k interface{} = "0ca89739-7265-4e01-9629-5e7ecae97711"
		var v float64 = 0.266

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
		var k interface{} = "2e6d0e58-569d-4214-af03-55b9720f1b31"
		var v float64 = 0.842

		test := omap.NewMapAnyFloat64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("f8b91387-ab29-4186-a73c-93565cdfb79e", (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float64 = 0.705
		So(test.PutIfNotNil("9bdb3b7e-ab37-44f5-a06d-38fdf54ba719", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyFloat64.ReplaceIfExists", t, func() {
		var k interface{} = "af7ecd84-f5d9-4058-9f14-8f06e26309a3"
		var v float64 = 0.953
		var x float64 = 0.224

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("1099dc70-d10e-4214-9e6b-3d60bb7f18dc", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyFloat64.ReplaceOrPut", t, func() {
		var k interface{} = "e4d06635-6d11-40dd-9974-91df0330fea4"
		var v float64 = 0.526
		var x float64 = 0.725

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("52df5b2a-a9fb-4983-bd7c-0dfd17f15d70", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat64_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyFloat64.MarshalJSON", t, func() {
		var k interface{} = "03bb2696-4639-4cc5-8012-919bcefad5da"
		var v float64 = 0.942

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"03bb2696-4639-4cc5-8012-919bcefad5da","value":0.942}]`)
	})
}
