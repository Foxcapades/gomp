package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringInt16_Put(t *testing.T) {
	Convey("TestMapStringInt16.Put", t, func() {
		var k string = "49fdc7c5-7cd4-4048-9667-a2330ea67d9e"
		var v int16 = 26081

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringInt16_Delete(t *testing.T) {
	Convey("TestMapStringInt16.Delete", t, func() {
		var k string = "29353c1f-ad41-4a4c-89b5-38d369d91112"
		var v int16 = 30346

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringInt16_Has(t *testing.T) {
	Convey("TestMapStringInt16.Has", t, func() {
		var k string = "288a2880-de98-40e4-9fe7-fb9dbe6e4b59"
		var v int16 = 21499

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("80daa832-bba6-4072-a01f-1483601955e8"+"512ec062-ee93-4db3-8452-b704c9b580b2"), ShouldBeFalse)
	})
}

func TestMapStringInt16_Get(t *testing.T) {
	Convey("TestMapStringInt16.Get", t, func() {
		var k string = "1c161682-2f6e-4107-beb9-15b0bb0b575b"
		var v int16 = 24034

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("af1da898-f5f0-4b17-b651-dc447723d774" + "38aa9072-70c1-4f63-831a-79419e85c986")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringInt16_GetOpt(t *testing.T) {
	Convey("TestMapStringInt16.GetOpt", t, func() {
		var k string = "7269729c-a89d-4987-9198-d44e31d3e11a"
		var v int16 = 10311

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("7e1511c1-9492-479d-b367-a8e9df6d6cfc" + "371f8c4b-0453-49c4-b86c-11d0fc15b160")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringInt16_ForEach(t *testing.T) {
	Convey("TestMapStringInt16.ForEach", t, func() {
		var k string = "3b227d0e-934a-49bd-acaf-3b14f119f42e"
		var v int16 = 12710
		hits := 0

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv int16) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringInt16_MarshalYAML(t *testing.T) {
	Convey("TestMapStringInt16.MarshalYAML", t, func() {
		var k string = "4fc1dd4c-4e98-4fde-a6d5-263c6f07f453"
		var v int16 = 28175

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringInt16_ToYAML(t *testing.T) {
	Convey("TestMapStringInt16.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k string = "6d61211c-5c20-4913-8f92-71660743a429"
			var v int16 = 4714

			test := omap.NewMapStringInt16(1)

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
			var k string = "6895d09b-9975-496e-9684-98367c9cacfc"
			var v int16 = 24878

			test := omap.NewMapStringInt16(1)
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

func TestMapStringInt16_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringInt16.PutIfNotNil", t, func() {
		var k string = "755590bf-3be1-4d88-a3cf-c24f767617ac"
		var v int16 = 4112

		test := omap.NewMapStringInt16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("3d2ca0fc-34d1-4cdd-8f24-cdddff26e1f3", (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int16 = 22994
		So(test.PutIfNotNil("710a04ee-9059-4322-ae0f-5356e0c58908", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringInt16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringInt16.ReplaceIfExists", t, func() {
		var k string = "0e6bd5b8-6a48-4077-a639-7a462e7db978"
		var v int16 = 1949
		var x int16 = 25619

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("ce1f38ab-b626-4aec-9b4f-71f6633b4768", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringInt16.ReplaceOrPut", t, func() {
		var k string = "01946da4-41e4-4636-9850-13bd91f359e7"
		var v int16 = 10648
		var x int16 = 13383

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("fa96576e-10e7-4127-baa9-272e6ef7989b", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt16_MarshalJSON(t *testing.T) {
	Convey("TestMapStringInt16.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k string = "95d55759-fcad-4d42-80da-d13d8b83e243"
			var v int16 = 13621

			test := omap.NewMapStringInt16(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"95d55759-fcad-4d42-80da-d13d8b83e243","value":13621}]`)
		})

		Convey("Unordered", func() {
			var k string = "95d55759-fcad-4d42-80da-d13d8b83e243"
			var v int16 = 13621

			test := omap.NewMapStringInt16(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"95d55759-fcad-4d42-80da-d13d8b83e243":13621}`)
		})

	})
}
