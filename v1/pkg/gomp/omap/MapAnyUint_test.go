package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyUint_Put(t *testing.T) {
	Convey("TestMapAnyUint.Put", t, func() {
		var k interface{} = "56486312-0a2e-49c0-8cc5-99a8cc508222"
		var v uint = 3708950243

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyUint_Delete(t *testing.T) {
	Convey("TestMapAnyUint.Delete", t, func() {
		var k interface{} = "577f824a-89e2-41b1-9753-6d25ce4178e8"
		var v uint = 1818549593

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyUint_Has(t *testing.T) {
	Convey("TestMapAnyUint.Has", t, func() {
		var k interface{} = "0f27959b-d19b-407f-8ae7-e630212b0488"
		var v uint = 1785006685

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("b79f144c-c173-431b-acea-011de84efb99"+"24fb18d2-c655-4558-b0c5-595bfdb4e067"), ShouldBeFalse)
	})
}


func TestMapAnyUint_Get(t *testing.T) {
	Convey("TestMapAnyUint.Get", t, func() {
		var k interface{} = "76f5027c-5296-41e7-a6a5-e31a03898c41"
		var v uint = 2444392625

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("a713a51d-b39f-4326-9746-5504dd5e6a9c" + "d457fe4a-e4bc-4883-910a-8b82771f52a2")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyUint_GetOpt(t *testing.T) {
	Convey("TestMapAnyUint.GetOpt", t, func() {
		var k interface{} = "d0f6c4a1-c503-4f43-b4fd-cdbed07e7849"
		var v uint = 2743892666

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("c7236412-6309-48d1-acdc-a6dbed00f1e5" + "dfc781d0-eaae-48f3-ad9d-c79e33c41e32")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyUint_ForEach(t *testing.T) {
	Convey("TestMapAnyUint.ForEach", t, func() {
		var k interface{} = "b040d2f5-e37e-4406-a687-4ac6c033e158"
		var v uint = 179520163
		hits := 0

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv uint) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyUint_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyUint.MarshalYAML", t, func() {
		var k interface{} = "0710d5bf-3eb1-4b60-bc0d-17251da2c04f"
		var v uint = 3509751768

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyUint_ToYAML(t *testing.T) {
	Convey("TestMapAnyUint.ToYAML", t, func() {
		var k interface{} = "2b05cc3d-05d5-4714-8eff-bb7b00308676"
		var v uint = 1743606297

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyUint_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyUint.PutIfNotNil", t, func() {
		var k interface{} = "ef47d774-9479-428f-a72b-24a771f6557c"
		var v uint = 4095746407

		test := omap.NewMapAnyUint(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("b8f0d14d-cbe2-49fb-95e9-881bc356003c", (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint = 3962025320
		So(test.PutIfNotNil("62c00f27-bd1f-4108-9fe2-40caa61b07a2", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyUint.ReplaceIfExists", t, func() {
		var k interface{} = "e9b9fc36-e98c-40a8-aa8b-d28d3b85e081"
		var v uint = 415131427
		var x uint = 3314737456

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("e3a60a4a-2340-4b0e-85cd-93cf506f350b", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyUint.ReplaceOrPut", t, func() {
		var k interface{} = "f4b5d285-b50a-486f-a5bb-110e85e97872"
		var v uint = 3304023291
		var x uint = 2920691060

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("2cab0de3-cd49-48f0-8de6-5eaf67abd2c6", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyUint.MarshalJSON", t, func() {
		var k interface{} = "c590147b-f4fc-4dbb-8b39-31111ea67930"
		var v uint = 3975892170

		test := omap.NewMapAnyUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"c590147b-f4fc-4dbb-8b39-31111ea67930","value":3975892170}]`)
	})
}
