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
		var k int = 742694901
		var v float32 = 0.147

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntFloat32_Delete(t *testing.T) {
	Convey("TestMapIntFloat32.Delete", t, func() {
		var k int = 1624480180
		var v float32 = 0.254

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntFloat32_Has(t *testing.T) {
	Convey("TestMapIntFloat32.Has", t, func() {
		var k int = 1196486033
		var v float32 = 0.831

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(138584955+1859063876), ShouldBeFalse)
	})
}

func TestMapIntFloat32_Get(t *testing.T) {
	Convey("TestMapIntFloat32.Get", t, func() {
		var k int = 1729059421
		var v float32 = 0.625

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(121640006 + 133068468)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntFloat32_GetOpt(t *testing.T) {
	Convey("TestMapIntFloat32.GetOpt", t, func() {
		var k int = 888372039
		var v float32 = 0.183

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1162923158 + 626121946)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntFloat32_ForEach(t *testing.T) {
	Convey("TestMapIntFloat32.ForEach", t, func() {
		var k int = 1378592786
		var v float32 = 0.580
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
		var k int = 825997729
		var v float32 = 0.893

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
		var k int = 2085121491
		var v float32 = 0.761

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
		var k int = 1001989572
		var v float32 = 0.460

		test := omap.NewMapIntFloat32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(366467913, (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float32 = 0.891
		So(test.PutIfNotNil(665001883, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntFloat32.ReplaceIfExists", t, func() {
		var k int = 2004080112
		var v float32 = 0.198
		var x float32 = 0.663

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(2057054773, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntFloat32.ReplaceOrPut", t, func() {
		var k int = 1324452433
		var v float32 = 0.617
		var x float32 = 0.082

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1995035771, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat32_MarshalJSON(t *testing.T) {
	Convey("TestMapIntFloat32.MarshalJSON", t, func() {
		var k int = 998905396
		var v float32 = 0.142

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":998905396,"value":0.142}]`)
	})
}
