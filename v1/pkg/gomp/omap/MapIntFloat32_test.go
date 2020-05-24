package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntFloat32_Put(t *testing.T) {
	Convey("TestMapIntFloat32.Put", t, func() {
		var k int = 1760574509
		var v float32 = 0.433

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntFloat32_Delete(t *testing.T) {
	Convey("TestMapIntFloat32.Delete", t, func() {
		var k int = 386897538
		var v float32 = 0.232

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntFloat32_Has(t *testing.T) {
	Convey("TestMapIntFloat32.Has", t, func() {
		var k int = 1237444107
		var v float32 = 0.631

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(953765715+1000224453), ShouldBeFalse)
	})
}

func TestMapIntFloat32_Get(t *testing.T) {
	Convey("TestMapIntFloat32.Get", t, func() {
		var k int = 266098633
		var v float32 = 0.525

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1210481763 + 1683454495)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntFloat32_GetOpt(t *testing.T) {
	Convey("TestMapIntFloat32.GetOpt", t, func() {
		var k int = 1631006805
		var v float32 = 0.110

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1773277055 + 891159687)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntFloat32_ForEach(t *testing.T) {
	Convey("TestMapIntFloat32.ForEach", t, func() {
		var k int = 1389364564
		var v float32 = 0.534
		hits := 0

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv float32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntFloat32_MarshalYAML(t *testing.T) {
	Convey("TestMapIntFloat32.MarshalYAML", t, func() {
		var k int = 1326953255
		var v float32 = 0.971

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntFloat32_ToYAML(t *testing.T) {
	Convey("TestMapIntFloat32.ToYAML", t, func() {
		var k int = 756225089
		var v float32 = 0.177

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntFloat32_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntFloat32.PutIfNotNil", t, func() {
		var k int = 1826820807
		var v float32 = 0.811

		test := omap.NewMapIntFloat32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1878775988, (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float32 = 0.053
		So(test.PutIfNotNil(1595580153, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntFloat32.ReplaceIfExists", t, func() {
		var k int = 927188660
		var v float32 = 0.036
		var x float32 = 0.827

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(663966268, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntFloat32.ReplaceOrPut", t, func() {
		var k int = 1371233728
		var v float32 = 0.612
		var x float32 = 0.907

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(98457472, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat32_MarshalJSON(t *testing.T) {
	Convey("TestMapIntFloat32.MarshalJSON", t, func() {
		var k int = 1135137209
		var v float32 = 0.597

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1135137209,"value":0.597}]`)
	})
}
