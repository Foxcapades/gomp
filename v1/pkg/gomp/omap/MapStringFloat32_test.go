package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringFloat32_Put(t *testing.T) {
	Convey("TestMapStringFloat32.Put", t, func() {
		var k string = "4e6e3f3b-3112-4704-87ea-5dc4c73b517e"
		var v float32 = 0.591

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringFloat32_Delete(t *testing.T) {
	Convey("TestMapStringFloat32.Delete", t, func() {
		var k string = "3888cf58-b435-4b4e-ac60-535ec1e5023e"
		var v float32 = 0.669

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringFloat32_Has(t *testing.T) {
	Convey("TestMapStringFloat32.Has", t, func() {
		var k string = "29e62720-bad0-4c74-b5c3-6d0ce3c8b3f7"
		var v float32 = 0.027

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("7a596c70-12a3-407c-beec-6c2105e96600"+"dc1f7981-4e26-4fc5-8ee4-604fc8ab3988"), ShouldBeFalse)
	})
}


func TestMapStringFloat32_Get(t *testing.T) {
	Convey("TestMapStringFloat32.Get", t, func() {
		var k string = "9c38eb5a-8852-4487-a712-51363ad0d0d1"
		var v float32 = 0.737

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("5bb100ff-775a-4817-82b6-d4ffb6675bdc" + "7e220fdb-a3c2-4ebe-a1ea-92ce5b4d13ed")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringFloat32_GetOpt(t *testing.T) {
	Convey("TestMapStringFloat32.GetOpt", t, func() {
		var k string = "90dedd59-5681-4466-bec9-afbd16536595"
		var v float32 = 0.167

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("62cff689-24e7-4195-b99d-c425b3d381ff" + "13596bf3-ec34-4ee1-bc4e-a8a0b6def53f")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringFloat32_ForEach(t *testing.T) {
	Convey("TestMapStringFloat32.ForEach", t, func() {
		var k string = "f19a64b3-b56b-4844-854c-1ca024f89278"
		var v float32 = 0.292
		hits := 0

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv float32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringFloat32_MarshalYAML(t *testing.T) {
	Convey("TestMapStringFloat32.MarshalYAML", t, func() {
		var k string = "1a7bfc81-ecf1-44ea-a79d-712291b52574"
		var v float32 = 0.534

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringFloat32_ToYAML(t *testing.T) {
	Convey("TestMapStringFloat32.ToYAML", t, func() {
		var k string = "b49eda73-a768-4331-a2a6-cd5d6f0ab1f6"
		var v float32 = 0.512

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringFloat32_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringFloat32.PutIfNotNil", t, func() {
		var k string = "b53eb89d-6229-4d8f-900e-7125282538f7"
		var v float32 = 0.179

		test := omap.NewMapStringFloat32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("39fe27c9-e647-4e77-9877-9620ab966bbe", (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float32 = 0.905
		So(test.PutIfNotNil("02197612-64d3-4970-91f5-ff9206fda429", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringFloat32.ReplaceIfExists", t, func() {
		var k string = "5bb7e4a9-460d-41f9-820c-74e9550853e0"
		var v float32 = 0.451
		var x float32 = 0.260

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("7ce851a7-2cd5-4e7b-bbc2-edf96b5ea1a9", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringFloat32.ReplaceOrPut", t, func() {
		var k string = "e5d976d7-0ae1-4398-8826-e8b4c6c943df"
		var v float32 = 0.329
		var x float32 = 0.274

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("4d6e704b-0fa1-4bec-8ce6-d7e48391731d", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringFloat32_MarshalJSON(t *testing.T) {
	Convey("TestMapStringFloat32.MarshalJSON", t, func() {
		var k string = "c698e45f-e076-4b89-bdb5-2394da88e0a0"
		var v float32 = 0.655

		test := omap.NewMapStringFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"c698e45f-e076-4b89-bdb5-2394da88e0a0","value":0.655}]`)
	})
}

