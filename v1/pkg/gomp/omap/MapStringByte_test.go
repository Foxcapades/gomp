package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringByte_Put(t *testing.T) {
	Convey("TestMapStringByte.Put", t, func() {
		var k string = "bc2215b0-83d4-41f8-b62c-31316812a4ac"
		var v byte = 105

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringByte_Delete(t *testing.T) {
	Convey("TestMapStringByte.Delete", t, func() {
		var k string = "2e6e0284-0aad-4a6f-a701-47223beeb6ad"
		var v byte = 128

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringByte_Has(t *testing.T) {
	Convey("TestMapStringByte.Has", t, func() {
		var k string = "ec8cfda5-1269-4a8a-a92c-1f892221b88a"
		var v byte = 74

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("abb90929-909b-496e-a384-f6e5ec20b877"+"b23f7a53-7e69-4987-be3a-843c96e1156a"), ShouldBeFalse)
	})
}

func TestMapStringByte_Get(t *testing.T) {
	Convey("TestMapStringByte.Get", t, func() {
		var k string = "d9e7a89e-7787-4808-a559-5336ec9355ef"
		var v byte = 218

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("7fb27e10-4873-4d5f-8b39-5e84aa0cc7b0" + "19925c26-568b-4ca5-8536-5eb044091ac2")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringByte_GetOpt(t *testing.T) {
	Convey("TestMapStringByte.GetOpt", t, func() {
		var k string = "010831f3-be1a-4b25-84ae-617bf1231886"
		var v byte = 111

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("b4e5cf1d-0699-4189-84d0-85d9be4b9af7" + "dc56db73-31a5-43b3-95f2-04fce68c0caa")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringByte_ForEach(t *testing.T) {
	Convey("TestMapStringByte.ForEach", t, func() {
		var k string = "b6993a57-1529-4fe1-bd6b-7fe97290933d"
		var v byte = 254
		hits := 0

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv byte) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringByte_MarshalYAML(t *testing.T) {
	Convey("TestMapStringByte.MarshalYAML", t, func() {
		var k string = "1d6a7a4c-858a-4bfe-9ad0-87d504f30d9d"
		var v byte = 170

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringByte_ToYAML(t *testing.T) {
	Convey("TestMapStringByte.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k string = "4a7a68b8-5250-4e55-b6ab-cb4167edeb99"
			var v byte = 231

			test := omap.NewMapStringByte(1)

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
			var k string = "565f3ef6-63ca-4b2a-89ff-3980c52d2d15"
			var v byte = 128

			test := omap.NewMapStringByte(1)
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

func TestMapStringByte_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringByte.PutIfNotNil", t, func() {
		var k string = "8cf67236-dbb1-4a04-9bd4-261ec4ef62d1"
		var v byte = 10

		test := omap.NewMapStringByte(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("efe57c64-47ed-4605-ba79-2e5f8b3c5355", (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x byte = 102
		So(test.PutIfNotNil("634ef0ac-16d5-4a31-ada0-0d79edfff6ce", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringByte_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringByte.ReplaceIfExists", t, func() {
		var k string = "0580f468-1ff3-4dfd-8104-378deaa3c06f"
		var v byte = 14
		var x byte = 216

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("cce0b1d7-a584-4011-af34-431cde61fcae", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringByte_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringByte.ReplaceOrPut", t, func() {
		var k string = "b11d488e-b42f-406d-ba07-5b7561e89b10"
		var v byte = 211
		var x byte = 0

		test := omap.NewMapStringByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("d4022b33-9fdb-4ca5-8e7b-0c154e24cfa8", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringByte_MarshalJSON(t *testing.T) {
	Convey("TestMapStringByte.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k string = "179fdb12-2a8e-45c7-8fdc-4b3000a8483e"
			var v byte = 124

			test := omap.NewMapStringByte(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"179fdb12-2a8e-45c7-8fdc-4b3000a8483e","value":124}]`)
		})

		Convey("Unordered", func() {
			var k string = "179fdb12-2a8e-45c7-8fdc-4b3000a8483e"
			var v byte = 124

			test := omap.NewMapStringByte(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"179fdb12-2a8e-45c7-8fdc-4b3000a8483e":124}`)
		})

	})
}
