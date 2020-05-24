package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringUint_Put(t *testing.T) {
	Convey("TestMapStringUint.Put", t, func() {
		var k string = "24494f07-7f36-44b7-8075-5fafff3bd33f"
		var v uint = 1867489244

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringUint_Delete(t *testing.T) {
	Convey("TestMapStringUint.Delete", t, func() {
		var k string = "63d0907e-c187-4bf2-9c98-395b533590b9"
		var v uint = 2920541394

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringUint_Has(t *testing.T) {
	Convey("TestMapStringUint.Has", t, func() {
		var k string = "998de4a3-4039-46f5-82d8-73352ed8017f"
		var v uint = 1643646862

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("8ffda8cd-ee11-401c-9365-ed1eef0cac27"+"9d8e5462-f025-4905-abc0-353a3e142fa9"), ShouldBeFalse)
	})
}


func TestMapStringUint_Get(t *testing.T) {
	Convey("TestMapStringUint.Get", t, func() {
		var k string = "95893d4e-2c04-4b07-b86e-14649a8a5451"
		var v uint = 2560308187

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("484ad4c6-5a74-45ef-8efb-492ee7e6766f" + "6cedac08-8ae5-47b7-88c0-a26a6b1cc79d")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringUint_GetOpt(t *testing.T) {
	Convey("TestMapStringUint.GetOpt", t, func() {
		var k string = "0ebe3c99-5d0c-4a86-9cde-6bbed07170e6"
		var v uint = 364322436

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("0b6e5f8c-8a43-4233-b2e4-08d0cdbf2c56" + "2b8b2d89-480f-4f4f-90a8-03b545106bb3")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringUint_ForEach(t *testing.T) {
	Convey("TestMapStringUint.ForEach", t, func() {
		var k string = "2d6d94b2-d9b6-49b9-be2e-d6790ea53637"
		var v uint = 3484804331
		hits := 0

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv uint) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringUint_MarshalYAML(t *testing.T) {
	Convey("TestMapStringUint.MarshalYAML", t, func() {
		var k string = "8e2c1218-2499-4384-898a-6629dc36575f"
		var v uint = 3080369506

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringUint_ToYAML(t *testing.T) {
	Convey("TestMapStringUint.ToYAML", t, func() {
		var k string = "b0bec76c-c83f-4846-b1cc-253f597806a2"
		var v uint = 4199054346

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringUint_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringUint.PutIfNotNil", t, func() {
		var k string = "2e49fe15-3414-452b-8d97-b25a1dda61f7"
		var v uint = 2850272465

		test := omap.NewMapStringUint(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("96b500da-f0d7-41ba-b07d-a77bb63c4203", (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint = 1025019933
		So(test.PutIfNotNil("af239746-9caa-4709-9fb9-f0b7f7f8e4e1", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringUint_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringUint.ReplaceIfExists", t, func() {
		var k string = "c12a3b38-b7d8-470a-b8ae-60885accad86"
		var v uint = 1055354881
		var x uint = 2984320434

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("9bd147b9-0c0b-41f3-a8ae-d9a36c9df7e8", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringUint.ReplaceOrPut", t, func() {
		var k string = "0228b32b-e6f9-4e30-974d-6014d2350a27"
		var v uint = 3317131420
		var x uint = 268805136

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("7fd31817-e4c5-42db-bd24-b24c4bcc98fb", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint_MarshalJSON(t *testing.T) {
	Convey("TestMapStringUint.MarshalJSON", t, func() {
		var k string = "a37a830c-6275-4041-8bd8-20b23a437f3a"
		var v uint = 529724927

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"a37a830c-6275-4041-8bd8-20b23a437f3a","value":529724927}]`)
	})
}
