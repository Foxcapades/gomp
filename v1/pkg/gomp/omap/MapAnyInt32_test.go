package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyInt32_Put(t *testing.T) {
	Convey("TestMapAnyInt32.Put", t, func() {
		var k interface{} = "272bbf7b-861f-444b-a4ea-9f36b3f8ecdd"
		var v int32 = 960363711

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyInt32_Delete(t *testing.T) {
	Convey("TestMapAnyInt32.Delete", t, func() {
		var k interface{} = "8a1f8198-017b-4d63-95fd-0cc94cd1dc4d"
		var v int32 = 1565259724

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyInt32_Has(t *testing.T) {
	Convey("TestMapAnyInt32.Has", t, func() {
		var k interface{} = "04830855-5771-4867-a55c-a12e3f98821e"
		var v int32 = 297619107

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("f061f106-7c45-4ae9-8928-d74786c2ae68"+"870ce880-190d-4bcf-b6b1-cdbffceeefdc"), ShouldBeFalse)
	})
}

func TestMapAnyInt32_Get(t *testing.T) {
	Convey("TestMapAnyInt32.Get", t, func() {
		var k interface{} = "11333969-a066-43f2-8c0e-2152a209686c"
		var v int32 = 1619777612

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("56633abd-38be-447c-9da5-a0e39c7f325c" + "ce0cda67-a107-4bc1-a392-e406aafca7a3")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyInt32_GetOpt(t *testing.T) {
	Convey("TestMapAnyInt32.GetOpt", t, func() {
		var k interface{} = "202f5ca4-d4c3-4fec-bc61-fe906d9ac545"
		var v int32 = 71290522

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("02baec03-8653-4179-9c64-cf9bd3256199" + "6f559210-41f9-43c5-841d-dda5707434bd")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyInt32_ForEach(t *testing.T) {
	Convey("TestMapAnyInt32.ForEach", t, func() {
		var k interface{} = "8eb73b67-df4e-439b-820f-b558afe17593"
		var v int32 = 1880953958
		hits := 0

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv int32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyInt32_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyInt32.MarshalYAML", t, func() {
		var k interface{} = "95419882-36b6-49f1-bcbc-961839f46164"
		var v int32 = 309488539

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyInt32_ToYAML(t *testing.T) {
	Convey("TestMapAnyInt32.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "69d73e73-742c-488d-834e-ef1f89fbfee6"
			var v int32 = 783139143

			test := omap.NewMapAnyInt32(1)

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
			var k interface{} = "21fba9b8-b845-41f0-81a7-11eec7e7b04f"
			var v int32 = 1712027292

			test := omap.NewMapAnyInt32(1)
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

func TestMapAnyInt32_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyInt32.PutIfNotNil", t, func() {
		var k interface{} = "43536f98-2518-463c-8fb4-2ee4f240a2c3"
		var v int32 = 979096235

		test := omap.NewMapAnyInt32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("ff25e3ef-d60c-4b8c-b6b4-74eaedf68459", (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int32 = 1815728017
		So(test.PutIfNotNil("a8003f11-b03a-4e43-8f9b-ea982d8b688b", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyInt32.ReplaceIfExists", t, func() {
		var k interface{} = "da8e4d45-a9bc-44c3-949d-c055bc04b4cb"
		var v int32 = 1872449060
		var x int32 = 563599842

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("1aee499b-8235-4210-b0f9-f32b5de25aa8", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyInt32.ReplaceOrPut", t, func() {
		var k interface{} = "ad038969-2dd6-4ec8-9c1f-32158bd82501"
		var v int32 = 1624338967
		var x int32 = 1226689057

		test := omap.NewMapAnyInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("d90e013b-e577-4370-9d49-6f9ef0ca0473", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt32_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyInt32.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "9186beef-65d3-4133-b834-8ab1cef48015"
			var v int32 = 3281612

			test := omap.NewMapAnyInt32(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"9186beef-65d3-4133-b834-8ab1cef48015","value":3281612}]`)
		})

		Convey("Unordered", func() {
			var k interface{} = "9186beef-65d3-4133-b834-8ab1cef48015"
			var v int32 = 3281612

			test := omap.NewMapAnyInt32(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"9186beef-65d3-4133-b834-8ab1cef48015":3281612}`)
		})

	})
}
