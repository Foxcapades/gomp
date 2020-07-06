package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringInt32_Put(t *testing.T) {
	Convey("TestMapStringInt32.Put", t, func() {
		var k string = "03c7a91f-b089-4be8-b609-d52e4a2903eb"
		var v int32 = 2055696339

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringInt32_Delete(t *testing.T) {
	Convey("TestMapStringInt32.Delete", t, func() {
		var k string = "7a7694fa-438d-425c-8e41-00e4008b7602"
		var v int32 = 214546178

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringInt32_Has(t *testing.T) {
	Convey("TestMapStringInt32.Has", t, func() {
		var k string = "ba5bfc6e-01b0-45cb-9546-040d1107577d"
		var v int32 = 2088085965

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("ee5ad873-4f48-442c-90ed-b983e97e75d3"+"0f34aedd-c22a-4ebd-8133-3a099c716d90"), ShouldBeFalse)
	})
}

func TestMapStringInt32_Get(t *testing.T) {
	Convey("TestMapStringInt32.Get", t, func() {
		var k string = "a3118682-fe66-4302-a265-1f485568402a"
		var v int32 = 1451158052

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("f0cda0d8-766a-44ac-b4e1-0f8ce05466ff" + "f410a6bc-b89c-4921-adb1-f083fce2ef3f")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringInt32_GetOpt(t *testing.T) {
	Convey("TestMapStringInt32.GetOpt", t, func() {
		var k string = "1313cb7d-b5e7-480f-a919-edad3ad0b3f2"
		var v int32 = 955751141

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("730ae327-05bb-4560-9716-dbe346f10310" + "eb63b93f-0a36-469c-9ed4-5c5a8305c47d")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringInt32_ForEach(t *testing.T) {
	Convey("TestMapStringInt32.ForEach", t, func() {
		var k string = "807ab8d6-cd3a-43a7-be3f-2bf74742cf6a"
		var v int32 = 918676985
		hits := 0

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv int32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringInt32_MarshalYAML(t *testing.T) {
	Convey("TestMapStringInt32.MarshalYAML", t, func() {
		var k string = "7874d159-7598-4480-9484-3b01b932c91d"
		var v int32 = 345056021

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringInt32_ToYAML(t *testing.T) {
	Convey("TestMapStringInt32.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k string = "ac0d9f55-f6f8-4a34-a59b-094ec79e024c"
			var v int32 = 939929820

			test := omap.NewMapStringInt32(1)

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
			var k string = "467f5575-d566-434d-a250-c7388d4528dc"
			var v int32 = 1949921439

			test := omap.NewMapStringInt32(1)
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

func TestMapStringInt32_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringInt32.PutIfNotNil", t, func() {
		var k string = "0aee316d-7710-49a6-8726-4aa4696a52c9"
		var v int32 = 1166247463

		test := omap.NewMapStringInt32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("890aeee1-e220-437b-a661-1deeb9c4faaa", (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int32 = 1145443706
		So(test.PutIfNotNil("ae9b94cb-5626-4d6d-91e0-6587c6452ccf", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringInt32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringInt32.ReplaceIfExists", t, func() {
		var k string = "bf676e78-c057-4dd7-96e4-f0e85214505c"
		var v int32 = 801185404
		var x int32 = 1502120580

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("31f3ffcc-c0ac-4855-8d8f-1c8ced9b109c", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringInt32.ReplaceOrPut", t, func() {
		var k string = "42711a9e-3fbd-4a41-9ff4-9e1e18e21ba8"
		var v int32 = 1243219376
		var x int32 = 377468902

		test := omap.NewMapStringInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("3417edf8-08d0-491f-a981-7a7422b7a32e", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt32_MarshalJSON(t *testing.T) {
	Convey("TestMapStringInt32.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k string = "f4069649-75a8-4665-9154-c63d93f88daa"
			var v int32 = 152265266

			test := omap.NewMapStringInt32(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"f4069649-75a8-4665-9154-c63d93f88daa","value":152265266}]`)
		})

		Convey("Unordered", func() {
			var k string = "f4069649-75a8-4665-9154-c63d93f88daa"
			var v int32 = 152265266

			test := omap.NewMapStringInt32(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"f4069649-75a8-4665-9154-c63d93f88daa":152265266}`)
		})

	})
}
