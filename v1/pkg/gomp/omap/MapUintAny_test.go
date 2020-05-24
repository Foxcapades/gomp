package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintAny_Put(t *testing.T) {
	Convey("TestMapUintAny.Put", t, func() {
		var k uint = 151746108
		var v interface{} = "79a798c5-ef27-40cc-a0ba-872758ff5eb7"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintAny_Delete(t *testing.T) {
	Convey("TestMapUintAny.Delete", t, func() {
		var k uint = 3931238617
		var v interface{} = "232aa85f-a818-4669-b4ea-a748aecc78cf"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintAny_Has(t *testing.T) {
	Convey("TestMapUintAny.Has", t, func() {
		var k uint = 370170447
		var v interface{} = "ef4602a6-6de9-451c-96a2-42a0e3cebe86"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(3257495364+1903072760), ShouldBeFalse)
	})
}

func TestMapUintAny_Get(t *testing.T) {
	Convey("TestMapUintAny.Get", t, func() {
		var k uint = 818947736
		var v interface{} = "3845af22-4642-47ff-8e67-ef392d30dbb7"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(3449447239 + 2663380057)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintAny_GetOpt(t *testing.T) {
	Convey("TestMapUintAny.GetOpt", t, func() {
		var k uint = 1230941462
		var v interface{} = "fa44ae8e-f41d-4b89-af85-d30339173e91"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(3435842879 + 1971203250)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintAny_ForEach(t *testing.T) {
	Convey("TestMapUintAny.ForEach", t, func() {
		var k uint = 3402666314
		var v interface{} = "bff3c711-1acb-406d-a37a-4a89b9020724"
		hits := 0

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv interface{}) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintAny_MarshalYAML(t *testing.T) {
	Convey("TestMapUintAny.MarshalYAML", t, func() {
		var k uint = 3037588673
		var v interface{} = "2f330a74-b5fc-4cae-ad20-7d6c0c570db5"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintAny_ToYAML(t *testing.T) {
	Convey("TestMapUintAny.ToYAML", t, func() {
		var k uint = 1361531495
		var v interface{} = "b1d8534b-bb27-4392-a234-09f9a3b4d949"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintAny_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintAny.PutIfNotNil", t, func() {
		var k uint = 3619196893
		var v interface{} = "aeea46ce-a3cd-4709-9def-7351ca1c3038"

		test := omap.NewMapUintAny(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(2097723908, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "a729692e-80ff-48d7-bca2-cc7c8119a2be"
		So(test.PutIfNotNil(3049561793, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintAny_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintAny.ReplaceIfExists", t, func() {
		var k uint = 2209404738
		var v interface{} = "bca7ac04-263a-4b1f-927a-99f46956ad88"
		var x interface{} = "da3c3330-a554-422f-9798-a09438369365"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(4292657844, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintAny_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintAny.ReplaceOrPut", t, func() {
		var k uint = 119577303
		var v interface{} = "c53280a2-9b84-4a5c-b777-8b9855a9288c"
		var x interface{} = "d9ffdc95-06dd-45c7-b7bc-a3d499aa163d"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(3183805263, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintAny_MarshalJSON(t *testing.T) {
	Convey("TestMapUintAny.MarshalJSON", t, func() {
		var k uint = 3701789742
		var v interface{} = "5cbc0d01-21b5-4982-9551-2dbfb3d96ded"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":3701789742,"value":"5cbc0d01-21b5-4982-9551-2dbfb3d96ded"}]`)
	})
}
