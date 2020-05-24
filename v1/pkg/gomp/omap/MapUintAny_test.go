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
		var k uint = 3917287802
		var v interface{} = "efa607f8-426e-44c9-8fbf-3d72844835b9"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintAny_Delete(t *testing.T) {
	Convey("TestMapUintAny.Delete", t, func() {
		var k uint = 4161162692
		var v interface{} = "caf36f04-517b-4e6e-a20d-808073d5dfd9"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintAny_Has(t *testing.T) {
	Convey("TestMapUintAny.Has", t, func() {
		var k uint = 2453184314
		var v interface{} = "8f93d60a-2792-48f0-bc14-6ee09551782b"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1516083343+17774235), ShouldBeFalse)
	})
}

func TestMapUintAny_Get(t *testing.T) {
	Convey("TestMapUintAny.Get", t, func() {
		var k uint = 1755730072
		var v interface{} = "7832a093-d806-4fed-a250-11e04c920582"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1530255009 + 9355777)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintAny_GetOpt(t *testing.T) {
	Convey("TestMapUintAny.GetOpt", t, func() {
		var k uint = 1248437310
		var v interface{} = "81f3fb68-ad2d-48f1-b749-3e6b42f5142f"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1066832964 + 404055303)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintAny_ForEach(t *testing.T) {
	Convey("TestMapUintAny.ForEach", t, func() {
		var k uint = 1170646042
		var v interface{} = "6a01d363-97e2-4e1a-b7f6-ec24e385ae39"
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
		var k uint = 381534437
		var v interface{} = "08441384-8b77-41c2-afd8-b255731addc2"

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
		var k uint = 668340015
		var v interface{} = "7168be5a-c47c-4a0a-b26d-fc59be151cbb"

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
		var k uint = 1939383837
		var v interface{} = "01ee8687-0a4e-4f7d-b1c3-79833b04ac4d"

		test := omap.NewMapUintAny(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(4108562536, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "46924646-416c-4f00-a2b2-43510baf3071"
		So(test.PutIfNotNil(4013707384, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintAny_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintAny.ReplaceIfExists", t, func() {
		var k uint = 3290799376
		var v interface{} = "1cf78b67-9c03-4ce4-a903-fc8bb8c128f3"
		var x interface{} = "1c478167-86c3-4f8d-8485-7b2129aa7968"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(3379131346, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintAny_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintAny.ReplaceOrPut", t, func() {
		var k uint = 2088417659
		var v interface{} = "cd2e0b1a-3c32-4323-84bb-a137ffec4ff1"
		var x interface{} = "d8496485-2498-4588-9f9b-db12bc5cbd90"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(3127416620, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintAny_MarshalJSON(t *testing.T) {
	Convey("TestMapUintAny.MarshalJSON", t, func() {
		var k uint = 3230668779
		var v interface{} = "fd08e10d-7615-4d59-aeb4-de2658b33b58"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":3230668779,"value":"fd08e10d-7615-4d59-aeb4-de2658b33b58"}]`)
	})
}
