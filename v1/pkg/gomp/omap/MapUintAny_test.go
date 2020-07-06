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
		var k uint = 2388441298
		var v interface{} = "1c65e23b-1fea-44c4-91cc-5d4e27dd927e"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintAny_Delete(t *testing.T) {
	Convey("TestMapUintAny.Delete", t, func() {
		var k uint = 327780646
		var v interface{} = "30bd8734-0151-4d6f-aa03-89ab5d59d953"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintAny_Has(t *testing.T) {
	Convey("TestMapUintAny.Has", t, func() {
		var k uint = 2659424036
		var v interface{} = "e8d8486d-d1bd-4e0a-aae2-a1a2e28b094d"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1597760445+2187590260), ShouldBeFalse)
	})
}

func TestMapUintAny_Get(t *testing.T) {
	Convey("TestMapUintAny.Get", t, func() {
		var k uint = 1260817416
		var v interface{} = "ce733dd8-6e6d-4611-a189-52d3c76551c5"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(1418475048 + 3849967718)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintAny_GetOpt(t *testing.T) {
	Convey("TestMapUintAny.GetOpt", t, func() {
		var k uint = 3967814468
		var v interface{} = "af939e06-7b03-45a9-8de0-dabee09f81d7"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(3608564911 + 1793966046)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintAny_ForEach(t *testing.T) {
	Convey("TestMapUintAny.ForEach", t, func() {
		var k uint = 1521171558
		var v interface{} = "8c7846c9-ca18-480b-987d-5f77af17281c"
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
		var k uint = 2902136402
		var v interface{} = "69cb49d5-a74f-4eb0-a6f1-dd7594b92019"

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
		Convey("Ordered", func() {
			var k uint = 2194440243
			var v interface{} = "205bb901-67ae-4a68-941b-efce041a57ce"

			test := omap.NewMapUintAny(1)

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
			var k uint = 475334790
			var v interface{} = "58d1587c-6c9f-40cd-8218-1154d95b089d"

			test := omap.NewMapUintAny(1)
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

func TestMapUintAny_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintAny.PutIfNotNil", t, func() {
		var k uint = 2248969623
		var v interface{} = "f8256de7-36d8-44ed-bb2f-f1ddd4d9f58d"

		test := omap.NewMapUintAny(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1890469681, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "a1bc34a5-9dfd-4ba1-b301-668caab7d64d"
		So(test.PutIfNotNil(1496765130, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintAny_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintAny.ReplaceIfExists", t, func() {
		var k uint = 2455432853
		var v interface{} = "e2a32ef2-d9c7-463b-ae23-03cd3e4e0125"
		var x interface{} = "ce240378-4c7d-4b07-9f21-1dafc38ffdee"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1144263414, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintAny_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintAny.ReplaceOrPut", t, func() {
		var k uint = 3172042759
		var v interface{} = "d37e89f6-4b0d-473f-ae4a-9dabe8f0f3db"
		var x interface{} = "02eedfdb-2839-4ccb-bbdd-ee9e9c2f8735"

		test := omap.NewMapUintAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1841059202, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintAny_MarshalJSON(t *testing.T) {
	Convey("TestMapUintAny.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k uint = 211878212
			var v interface{} = "b74520ab-f817-4534-87be-22acb702bc6d"

			test := omap.NewMapUintAny(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":211878212,"value":"b74520ab-f817-4534-87be-22acb702bc6d"}]`)
		})

		Convey("Unordered", func() {
			var k uint = 211878212
			var v interface{} = "b74520ab-f817-4534-87be-22acb702bc6d"

			test := omap.NewMapUintAny(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"211878212":"b74520ab-f817-4534-87be-22acb702bc6d"}`)
		})

	})
}
