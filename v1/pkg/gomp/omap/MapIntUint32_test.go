package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntUint32_Put(t *testing.T) {
	Convey("TestMapIntUint32.Put", t, func() {
		var k int = 1060041166
		var v uint32 = 2620680614

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntUint32_Delete(t *testing.T) {
	Convey("TestMapIntUint32.Delete", t, func() {
		var k int = 1140789278
		var v uint32 = 2121201325

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntUint32_Has(t *testing.T) {
	Convey("TestMapIntUint32.Has", t, func() {
		var k int = 1702294465
		var v uint32 = 357508670

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(2125524606+219298375), ShouldBeFalse)
	})
}


func TestMapIntUint32_Get(t *testing.T) {
	Convey("TestMapIntUint32.Get", t, func() {
		var k int = 1173555791
		var v uint32 = 1609870889

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(89849910 + 1144918131)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntUint32_GetOpt(t *testing.T) {
	Convey("TestMapIntUint32.GetOpt", t, func() {
		var k int = 1744745364
		var v uint32 = 2101136069

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1090917313 + 1408967612)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntUint32_ForEach(t *testing.T) {
	Convey("TestMapIntUint32.ForEach", t, func() {
		var k int = 245803399
		var v uint32 = 2102525737
		hits := 0

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv uint32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntUint32_MarshalYAML(t *testing.T) {
	Convey("TestMapIntUint32.MarshalYAML", t, func() {
		var k int = 123134062
		var v uint32 = 3347820224

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntUint32_ToYAML(t *testing.T) {
	Convey("TestMapIntUint32.ToYAML", t, func() {
		var k int = 1381806980
		var v uint32 = 2183660464

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntUint32_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntUint32.PutIfNotNil", t, func() {
		var k int = 1275243333
		var v uint32 = 3822511747

		test := omap.NewMapIntUint32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1522658897, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint32 = 2789861503
		So(test.PutIfNotNil(1672632951, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntUint32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntUint32.ReplaceIfExists", t, func() {
		var k int = 2136200667
		var v uint32 = 3122162072
		var x uint32 = 1622410704

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(308555685, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntUint32.ReplaceOrPut", t, func() {
		var k int = 1320062723
		var v uint32 = 3344290128
		var x uint32 = 1453115209

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1610564448, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint32_MarshalJSON(t *testing.T) {
	Convey("TestMapIntUint32.MarshalJSON", t, func() {
		var k int = 150880476
		var v uint32 = 1555985531

		test := omap.NewMapIntUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":150880476,"value":1555985531}]`)
	})
}
