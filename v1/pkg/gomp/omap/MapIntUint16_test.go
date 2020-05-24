package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntUint16_Put(t *testing.T) {
	Convey("TestMapIntUint16.Put", t, func() {
		var k int = 509434607
		var v uint16 = 48356

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntUint16_Delete(t *testing.T) {
	Convey("TestMapIntUint16.Delete", t, func() {
		var k int = 121954712
		var v uint16 = 12890

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntUint16_Has(t *testing.T) {
	Convey("TestMapIntUint16.Has", t, func() {
		var k int = 1533604689
		var v uint16 = 17507

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(115200836+1694252616), ShouldBeFalse)
	})
}

func TestMapIntUint16_Get(t *testing.T) {
	Convey("TestMapIntUint16.Get", t, func() {
		var k int = 1307579961
		var v uint16 = 37619

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(844434428 + 742734289)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntUint16_GetOpt(t *testing.T) {
	Convey("TestMapIntUint16.GetOpt", t, func() {
		var k int = 992318495
		var v uint16 = 57195

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(2116572813 + 868106356)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntUint16_ForEach(t *testing.T) {
	Convey("TestMapIntUint16.ForEach", t, func() {
		var k int = 108596705
		var v uint16 = 28840
		hits := 0

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv uint16) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntUint16_MarshalYAML(t *testing.T) {
	Convey("TestMapIntUint16.MarshalYAML", t, func() {
		var k int = 1625587658
		var v uint16 = 64886

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntUint16_ToYAML(t *testing.T) {
	Convey("TestMapIntUint16.ToYAML", t, func() {
		var k int = 1748097505
		var v uint16 = 52328

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntUint16_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntUint16.PutIfNotNil", t, func() {
		var k int = 1476378342
		var v uint16 = 18811

		test := omap.NewMapIntUint16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1305811042, (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint16 = 4728
		So(test.PutIfNotNil(507175231, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntUint16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntUint16.ReplaceIfExists", t, func() {
		var k int = 1092208413
		var v uint16 = 59714
		var x uint16 = 44656

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1292553667, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntUint16.ReplaceOrPut", t, func() {
		var k int = 1905150102
		var v uint16 = 18292
		var x uint16 = 17383

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(447695895, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint16_MarshalJSON(t *testing.T) {
	Convey("TestMapIntUint16.MarshalJSON", t, func() {
		var k int = 1140334902
		var v uint16 = 34685

		test := omap.NewMapIntUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1140334902,"value":34685}]`)
	})
}
