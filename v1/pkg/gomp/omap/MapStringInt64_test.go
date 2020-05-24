package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringInt64_Put(t *testing.T) {
	Convey("TestMapStringInt64.Put", t, func() {
		var k string = "48c5f44a-db6d-479d-91eb-743bcb091b47"
		var v int64 = 8111399539087912765

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringInt64_Delete(t *testing.T) {
	Convey("TestMapStringInt64.Delete", t, func() {
		var k string = "f070a0cf-d551-4711-92dd-36da3eeae1c8"
		var v int64 = 8415372234761462824

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringInt64_Has(t *testing.T) {
	Convey("TestMapStringInt64.Has", t, func() {
		var k string = "6c67546d-26c0-44b9-829f-1470b269fa5b"
		var v int64 = 8231093891108184357

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("abd3baa2-d2a6-43aa-85f2-5447902d5011"+"06a6a1e7-02ac-4071-8551-653a29211afc"), ShouldBeFalse)
	})
}

func TestMapStringInt64_Get(t *testing.T) {
	Convey("TestMapStringInt64.Get", t, func() {
		var k string = "0c2feec8-3cc7-4c3c-a74c-5fc4ffd68af7"
		var v int64 = 1310761637193358581

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("dc9f7148-fd32-4e07-895c-9bfcd6d35b94" + "1c838f97-dde0-4d4d-8426-a9cfc7ea77d1")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringInt64_GetOpt(t *testing.T) {
	Convey("TestMapStringInt64.GetOpt", t, func() {
		var k string = "5857fd80-c303-4712-8bdf-edcbb103f755"
		var v int64 = 7719540285690734212

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("5165e83f-f5ee-4bb4-9309-bae0d67adcd2" + "503c8ec5-f5af-40e9-838b-48083755150a")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringInt64_ForEach(t *testing.T) {
	Convey("TestMapStringInt64.ForEach", t, func() {
		var k string = "a52336bd-3f66-4491-b694-e28a578a08f3"
		var v int64 = 6055119758438726088
		hits := 0

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv int64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringInt64_MarshalYAML(t *testing.T) {
	Convey("TestMapStringInt64.MarshalYAML", t, func() {
		var k string = "9900c2fc-c107-4607-b152-16abb9287b27"
		var v int64 = 7612181820988171234

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringInt64_ToYAML(t *testing.T) {
	Convey("TestMapStringInt64.ToYAML", t, func() {
		var k string = "28405bd8-72ee-42a9-90c9-a15701702eaf"
		var v int64 = 4287131316171532298

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringInt64_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringInt64.PutIfNotNil", t, func() {
		var k string = "b2ca4913-1ab3-4d46-8460-970ba0233333"
		var v int64 = 6129061306608743087

		test := omap.NewMapStringInt64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("67a9aacc-a2b7-4ce5-b0c8-69f547574dbf", (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int64 = 4738440880806391729
		So(test.PutIfNotNil("8e5b6255-1601-43a9-8d62-fbd7df37e75c", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringInt64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringInt64.ReplaceIfExists", t, func() {
		var k string = "b4ce0515-f46e-42f2-9453-f5b193ddbd34"
		var v int64 = 4664305699890220898
		var x int64 = 5025671435372331710

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("9c13ece8-ea8f-43a2-b378-829aae51a1f9", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringInt64.ReplaceOrPut", t, func() {
		var k string = "924931ac-a441-4df9-b3ee-c7505e30edda"
		var v int64 = 3721467147521410659
		var x int64 = 3597763883305162638

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("5ceafb9d-41fd-458f-a83a-234c0078d6c1", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt64_MarshalJSON(t *testing.T) {
	Convey("TestMapStringInt64.MarshalJSON", t, func() {
		var k string = "e067be49-8637-4f8c-a733-8904fb78d239"
		var v int64 = 7124147097868215171

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"e067be49-8637-4f8c-a733-8904fb78d239","value":7124147097868215171}]`)
	})
}
