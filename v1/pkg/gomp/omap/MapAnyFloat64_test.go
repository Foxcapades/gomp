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
		var k interface{} = "7377c32a-26e4-4cc6-9d1f-c8fe0545c82b"
		var v float64 = 0.118

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyFloat64_Delete(t *testing.T) {
	Convey("TestMapAnyFloat64.Delete", t, func() {
		var k interface{} = "d0598f94-4634-42ff-980d-533da7aedd13"
		var v float64 = 0.400

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyFloat64_Has(t *testing.T) {
	Convey("TestMapAnyFloat64.Has", t, func() {
		var k interface{} = "824c8734-44b9-454d-8b19-8c8d86713ced"
		var v float64 = 0.701

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("e610a7f2-0a5d-4fbd-98d0-54b26061af3d"+"345177d2-c0ad-4453-b364-dcf00c548503"), ShouldBeFalse)
	})
}

func TestMapAnyFloat64_Get(t *testing.T) {
	Convey("TestMapAnyFloat64.Get", t, func() {
		var k interface{} = "0a3524b9-297d-4380-bb23-4785c3945c3f"
		var v float64 = 0.516

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("d07444e1-b1ba-441a-926e-e5342f54f0ca" + "76d16216-30b7-43b5-9957-f6d470c2d34a")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyFloat64_GetOpt(t *testing.T) {
	Convey("TestMapAnyFloat64.GetOpt", t, func() {
		var k interface{} = "a86ae817-a85f-4d1c-bf0a-e6d504ff02a7"
		var v float64 = 0.650

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("21885ff9-dc42-4830-b3d4-23ceb3edc066" + "5075267c-5ca4-4b78-9742-b0ef2f04000b")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyFloat64_ForEach(t *testing.T) {
	Convey("TestMapAnyFloat64.ForEach", t, func() {
		var k interface{} = "d351b8de-86da-4bff-b0e1-da968570c584"
		var v float64 = 0.520
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
		var k interface{} = "39c994f8-11b3-4f1d-b9bb-8c14f42e25e8"
		var v float64 = 0.732

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
		Convey("Ordered", func() {
			var k interface{} = "c14fba86-528f-478c-a4d0-ee0e774a3534"
			var v float64 = 0.305

			test := omap.NewMapAnyFloat64(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			c, d := test.ToYAML()
			So(d, ShouldBeNil)
			So(c.Kind, ShouldEqual, yaml.SequenceNode)
			So(c.LongTag(), ShouldEqual, xyml.TagOrderedMap)
			So(len(c.Content), ShouldEqual, 1)
			So(xyml.IsMap(c.Content[0]), ShouldBeTrue)
		})

		Convey("Unordered", func() {
			var k interface{} = "62cdb7ff-037e-4427-ac38-fd6df1ed7fa1"
			var v float64 = 0.468

			test := omap.NewMapAnyFloat64(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			c, d := test.ToYAML()

			So(d, ShouldBeNil)
			So(c.Kind, ShouldEqual, yaml.MappingNode)
			So(c.LongTag(), ShouldEqual, xyml.TagMap)
			So(len(c.Content), ShouldEqual, 2)
		})
	})
}

func TestMapAnyFloat64_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyFloat64.PutIfNotNil", t, func() {
		var k interface{} = "e5361817-64f4-4d7d-ad58-e1aafc135d58"
		var v float64 = 0.572

		test := omap.NewMapAnyFloat64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("6d6bac2b-0fd5-4168-ae00-e45a1518ff75", (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float64 = 0.023
		So(test.PutIfNotNil("93e2e9cd-01e6-4f59-bb73-f800feb77373", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyFloat64.ReplaceIfExists", t, func() {
		var k interface{} = "9c77c15c-3a92-4678-8186-b544e61381b4"
		var v float64 = 0.088
		var x float64 = 0.297

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("0968cf55-e902-4571-9680-dd23b5cb0573", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyFloat64.ReplaceOrPut", t, func() {
		var k interface{} = "dc38ac5b-fbb0-4eb7-a46f-eb1c4b87c2b5"
		var v float64 = 0.007
		var x float64 = 0.476

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("ac3d0ea7-c3d4-4a73-a591-354baad1623c", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat64_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyFloat64.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "aa8d0812-c264-4c61-a0bc-7320300e7a13"
			var v float64 = 0.479

			test := omap.NewMapAnyFloat64(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"aa8d0812-c264-4c61-a0bc-7320300e7a13","value":0.479}]`)
		})

		Convey("Unordered", func() {
			var k interface{} = "aa8d0812-c264-4c61-a0bc-7320300e7a13"
			var v float64 = 0.479

			test := omap.NewMapAnyFloat64(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"aa8d0812-c264-4c61-a0bc-7320300e7a13":0.479}`)
		})

	})
}
