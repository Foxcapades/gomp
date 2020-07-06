package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyUint32_Put(t *testing.T) {
	Convey("TestMapAnyUint32.Put", t, func() {
		var k interface{} = "b9e89582-2695-4af0-ab5e-0fc6c3bee6a4"
		var v uint32 = 4080365327

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyUint32_Delete(t *testing.T) {
	Convey("TestMapAnyUint32.Delete", t, func() {
		var k interface{} = "fcc8e0b3-5008-427d-88c3-dffd35068d69"
		var v uint32 = 935364274

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyUint32_Has(t *testing.T) {
	Convey("TestMapAnyUint32.Has", t, func() {
		var k interface{} = "cedbe9aa-7ac5-468c-afb9-57fe64f07a74"
		var v uint32 = 1082442335

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("f980ecba-6c1a-4679-a57f-25c713a04f19"+"e5c1f1eb-ab7d-4f11-8f0f-63c71fa7dac8"), ShouldBeFalse)
	})
}

func TestMapAnyUint32_Get(t *testing.T) {
	Convey("TestMapAnyUint32.Get", t, func() {
		var k interface{} = "fcc90d02-b243-4a05-88d6-eac33cbe35e8"
		var v uint32 = 2537457170

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("48bebd0f-e84c-4b95-a9f1-106b67bea046" + "5b78767f-ba21-45d5-b615-8da649eafbc4")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyUint32_GetOpt(t *testing.T) {
	Convey("TestMapAnyUint32.GetOpt", t, func() {
		var k interface{} = "2d9b8d32-2309-47f8-ac94-c9e6fcf6d9bd"
		var v uint32 = 589372693

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("de15dbc8-2b6d-4ff0-afee-eed12b16bee1" + "e17e667c-dd53-4f9b-8d41-576dda0aa2cd")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyUint32_ForEach(t *testing.T) {
	Convey("TestMapAnyUint32.ForEach", t, func() {
		var k interface{} = "a94a1fc7-0d3e-48e6-8bf3-08fb0f1df583"
		var v uint32 = 1797159878
		hits := 0

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv uint32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyUint32_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyUint32.MarshalYAML", t, func() {
		var k interface{} = "41b6e6bb-4c32-4e46-8a34-8f46adde0a9a"
		var v uint32 = 985449296

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyUint32_ToYAML(t *testing.T) {
	Convey("TestMapAnyUint32.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "ac60a54c-6723-41c5-866f-3ac074063bd6"
			var v uint32 = 4102305042

			test := omap.NewMapAnyUint32(1)

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
			var k interface{} = "30c94c9f-c8dc-40fe-bb7a-406c481f3764"
			var v uint32 = 990732455

			test := omap.NewMapAnyUint32(1)
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

func TestMapAnyUint32_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyUint32.PutIfNotNil", t, func() {
		var k interface{} = "e7e626ae-aaa1-491d-a034-dd850053721c"
		var v uint32 = 2904895041

		test := omap.NewMapAnyUint32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("a824fa6b-a16f-4fbf-8e8d-bd055002f41a", (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint32 = 2386167410
		So(test.PutIfNotNil("0eaecb53-e249-4e7b-8df2-e34a3c0439eb", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyUint32.ReplaceIfExists", t, func() {
		var k interface{} = "8276eadd-9219-4ce4-9640-654c0ea0fcf0"
		var v uint32 = 3598390695
		var x uint32 = 4149613080

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("6a021e1c-17c7-49aa-a3c0-1997c9e3648e", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyUint32.ReplaceOrPut", t, func() {
		var k interface{} = "541290f9-cca6-41f8-bb7a-4ab769c639a4"
		var v uint32 = 1139522731
		var x uint32 = 4140906017

		test := omap.NewMapAnyUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("ad5ef9b5-e66e-4b08-a596-e8a8981435cd", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint32_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyUint32.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "49f96e20-bf16-491a-bcbb-76282ae6f05c"
			var v uint32 = 198718819

			test := omap.NewMapAnyUint32(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"49f96e20-bf16-491a-bcbb-76282ae6f05c","value":198718819}]`)
		})

		Convey("Unordered", func() {
			var k interface{} = "49f96e20-bf16-491a-bcbb-76282ae6f05c"
			var v uint32 = 198718819

			test := omap.NewMapAnyUint32(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"49f96e20-bf16-491a-bcbb-76282ae6f05c":198718819}`)
		})

	})
}
