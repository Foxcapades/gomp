package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyInt16_Put(t *testing.T) {
	Convey("TestMapAnyInt16.Put", t, func() {
		var k interface{} = "7ba86072-29e5-4d14-99d5-577c73bc0fb8"
		var v int16 = 4429

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyInt16_Delete(t *testing.T) {
	Convey("TestMapAnyInt16.Delete", t, func() {
		var k interface{} = "694a91cd-6725-4664-8345-3b082b7cba44"
		var v int16 = 20513

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyInt16_Has(t *testing.T) {
	Convey("TestMapAnyInt16.Has", t, func() {
		var k interface{} = "c237c2cd-01e5-4f38-8502-bdebb6a26af8"
		var v int16 = 9317

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("cf3fd607-3f36-464c-9573-726b57fad08c"+"9308fc63-454f-4140-be36-38dc1600cd30"), ShouldBeFalse)
	})
}

func TestMapAnyInt16_Get(t *testing.T) {
	Convey("TestMapAnyInt16.Get", t, func() {
		var k interface{} = "aac2c567-8949-4904-af07-83a85d4c2574"
		var v int16 = 5306

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("69b034fc-de6d-4513-a736-ffba5a36ac4f" + "956824e7-fb14-466a-b462-40fcbbf886ba")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyInt16_GetOpt(t *testing.T) {
	Convey("TestMapAnyInt16.GetOpt", t, func() {
		var k interface{} = "a24a2d0c-908e-4491-9693-fbddf1ba2055"
		var v int16 = 12526

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("557344eb-96a7-4b9d-986c-7226636e2ed6" + "8bf89325-d7a5-4c95-8671-dce7696fab98")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyInt16_ForEach(t *testing.T) {
	Convey("TestMapAnyInt16.ForEach", t, func() {
		var k interface{} = "bb9d943f-1e60-4751-ad4b-1ff5917aca97"
		var v int16 = 6483
		hits := 0

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv int16) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyInt16_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyInt16.MarshalYAML", t, func() {
		var k interface{} = "c4246fc2-ce99-4967-be2c-f8992fd9d5f9"
		var v int16 = 18830

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyInt16_ToYAML(t *testing.T) {
	Convey("TestMapAnyInt16.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "d976dadb-ee50-4dd0-8df0-b2b4d5bd00db"
			var v int16 = 27379

			test := omap.NewMapAnyInt16(1)

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
			var k interface{} = "4e6e9566-dc8b-4ddd-a790-1321f18d06e8"
			var v int16 = 32571

			test := omap.NewMapAnyInt16(1)
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

func TestMapAnyInt16_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyInt16.PutIfNotNil", t, func() {
		var k interface{} = "98b6f903-32e2-4c0a-b233-c64e643b4a9b"
		var v int16 = 14231

		test := omap.NewMapAnyInt16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("e6e98161-93d9-4621-9d2b-424a56fabfb8", (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int16 = 4833
		So(test.PutIfNotNil("eccd8e90-b23e-4520-b52d-580e264441a7", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyInt16.ReplaceIfExists", t, func() {
		var k interface{} = "96c8745a-d23f-4a2a-91e1-054cfe7a7681"
		var v int16 = 13508
		var x int16 = 9823

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("761cc6ac-3159-405e-9636-fadad7f2af69", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyInt16.ReplaceOrPut", t, func() {
		var k interface{} = "aa1d0133-3fe3-4a68-93e1-510df05b1dd6"
		var v int16 = 3259
		var x int16 = 22479

		test := omap.NewMapAnyInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("511ce0b5-ca0a-4a59-be9b-1acca62848f7", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt16_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyInt16.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "34e5c1ae-d2ed-4e69-b4ea-c5cbc0c58e70"
			var v int16 = 32333

			test := omap.NewMapAnyInt16(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"34e5c1ae-d2ed-4e69-b4ea-c5cbc0c58e70","value":32333}]`)
		})

		Convey("Unordered", func() {
			var k interface{} = "34e5c1ae-d2ed-4e69-b4ea-c5cbc0c58e70"
			var v int16 = 32333

			test := omap.NewMapAnyInt16(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"34e5c1ae-d2ed-4e69-b4ea-c5cbc0c58e70":32333}`)
		})

	})
}
