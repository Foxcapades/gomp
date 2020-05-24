package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntInt32_Put(t *testing.T) {
	Convey("TestMapIntInt32.Put", t, func() {
		var k int = 245632976
		var v int32 = 1715819466

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntInt32_Delete(t *testing.T) {
	Convey("TestMapIntInt32.Delete", t, func() {
		var k int = 820501641
		var v int32 = 1095591454

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntInt32_Has(t *testing.T) {
	Convey("TestMapIntInt32.Has", t, func() {
		var k int = 246853569
		var v int32 = 1613031933

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(443991004+1955820965), ShouldBeFalse)
	})
}


func TestMapIntInt32_Get(t *testing.T) {
	Convey("TestMapIntInt32.Get", t, func() {
		var k int = 36083099
		var v int32 = 1554335723

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(2063123621 + 1508851752)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntInt32_GetOpt(t *testing.T) {
	Convey("TestMapIntInt32.GetOpt", t, func() {
		var k int = 344508060
		var v int32 = 636473250

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(36859793 + 254519676)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntInt32_ForEach(t *testing.T) {
	Convey("TestMapIntInt32.ForEach", t, func() {
		var k int = 512013681
		var v int32 = 617965838
		hits := 0

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv int32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntInt32_MarshalYAML(t *testing.T) {
	Convey("TestMapIntInt32.MarshalYAML", t, func() {
		var k int = 1669949648
		var v int32 = 2106623058

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntInt32_ToYAML(t *testing.T) {
	Convey("TestMapIntInt32.ToYAML", t, func() {
		var k int = 1609318951
		var v int32 = 188841267

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntInt32_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntInt32.PutIfNotNil", t, func() {
		var k int = 1516480693
		var v int32 = 388529047

		test := omap.NewMapIntInt32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(16519848, (*int32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int32 = 1536206172
		So(test.PutIfNotNil(1135351481, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntInt32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntInt32.ReplaceIfExists", t, func() {
		var k int = 38103086
		var v int32 = 1997778166
		var x int32 = 1919510068

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(29005737, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntInt32.ReplaceOrPut", t, func() {
		var k int = 1838478869
		var v int32 = 2058241901
		var x int32 = 381696373

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(466817316, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt32_MarshalJSON(t *testing.T) {
	Convey("TestMapIntInt32.MarshalJSON", t, func() {
		var k int = 889223529
		var v int32 = 1069995254

		test := omap.NewMapIntInt32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":889223529,"value":1069995254}]`)
	})
}

