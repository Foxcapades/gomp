package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyByte_Put(t *testing.T) {
	Convey("TestMapAnyByte.Put", t, func() {
		var k interface{} = "03fc81f9-6bfe-4693-9278-970afea044e6"
		var v byte = 11

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyByte_Delete(t *testing.T) {
	Convey("TestMapAnyByte.Delete", t, func() {
		var k interface{} = "55f3b07f-2a1c-4e58-bf0e-2440eacb7550"
		var v byte = 170

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyByte_Has(t *testing.T) {
	Convey("TestMapAnyByte.Has", t, func() {
		var k interface{} = "ca7f87b8-1273-4fae-81f4-1bc88010487b"
		var v byte = 171

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("24484720-dd2a-4973-b4a5-a978fbefdbe7"+"e16565ce-f1e3-479c-b570-bcb0335f7605"), ShouldBeFalse)
	})
}


func TestMapAnyByte_Get(t *testing.T) {
	Convey("TestMapAnyByte.Get", t, func() {
		var k interface{} = "6d4a9a39-7886-4690-a5c4-3efea077e527"
		var v byte = 224

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("f89a627f-47f7-4b2b-8e69-44d162a3d755" + "da19feb5-1a62-4980-b150-9202fe7b7419")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyByte_GetOpt(t *testing.T) {
	Convey("TestMapAnyByte.GetOpt", t, func() {
		var k interface{} = "65b1a3ff-9045-4ef1-ac2e-1747b518d6d2"
		var v byte = 56

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("c9726e8c-a8e0-41e9-8ae3-d2d939d256ab" + "40df39c2-08bb-42d8-bc83-f12c484635c4")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyByte_ForEach(t *testing.T) {
	Convey("TestMapAnyByte.ForEach", t, func() {
		var k interface{} = "811bb8b6-626a-448e-97a6-471733a03411"
		var v byte = 252
		hits := 0

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv byte) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyByte_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyByte.MarshalYAML", t, func() {
		var k interface{} = "62a1e388-aa4a-4452-8358-175ac156396b"
		var v byte = 204

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyByte_ToYAML(t *testing.T) {
	Convey("TestMapAnyByte.ToYAML", t, func() {
		var k interface{} = "d18d06dd-21fe-4adb-8f67-fcf60211a987"
		var v byte = 123

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyByte_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyByte.PutIfNotNil", t, func() {
		var k interface{} = "d0d20042-60dd-4410-a204-e07b01b4619e"
		var v byte = 220

		test := omap.NewMapAnyByte(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("d3612337-07e4-4f86-9c3b-0f5e58834cf7", (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x byte = 243
		So(test.PutIfNotNil("3f9e1e85-8c51-40b9-be6d-2d56cb8906d0", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyByte_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyByte.ReplaceIfExists", t, func() {
		var k interface{} = "fb1c5d15-9ca0-49a0-9e8b-4fe95c99a44d"
		var v byte = 6
		var x byte = 29

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("04678e9e-8416-416f-b3da-395c38fb9656", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyByte_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyByte.ReplaceOrPut", t, func() {
		var k interface{} = "9014da1e-6816-44f5-b63e-20b10dfdd346"
		var v byte = 177
		var x byte = 77

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("bba0b255-3926-460c-9cee-810f80c68736", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyByte_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyByte.MarshalJSON", t, func() {
		var k interface{} = "f7637337-3928-487e-a7c5-9c1678732ea0"
		var v byte = 158

		test := omap.NewMapAnyByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"f7637337-3928-487e-a7c5-9c1678732ea0","value":158}]`)
	})
}
