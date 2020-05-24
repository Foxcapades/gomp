package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintUint32_Put(t *testing.T) {
	Convey("TestMapUintUint32.Put", t, func() {
		var k uint = 4027655770
		var v uint32 = 529440278

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintUint32_Delete(t *testing.T) {
	Convey("TestMapUintUint32.Delete", t, func() {
		var k uint = 3393360237
		var v uint32 = 1985621270

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintUint32_Has(t *testing.T) {
	Convey("TestMapUintUint32.Has", t, func() {
		var k uint = 2839522708
		var v uint32 = 2053729237

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1569378621+1564551874), ShouldBeFalse)
	})
}

func TestMapUintUint32_Get(t *testing.T) {
	Convey("TestMapUintUint32.Get", t, func() {
		var k uint = 1966905074
		var v uint32 = 1591890253

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(4072286991 + 1792924963)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintUint32_GetOpt(t *testing.T) {
	Convey("TestMapUintUint32.GetOpt", t, func() {
		var k uint = 2408674986
		var v uint32 = 1042718913

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(948213447 + 3995558886)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintUint32_ForEach(t *testing.T) {
	Convey("TestMapUintUint32.ForEach", t, func() {
		var k uint = 851530776
		var v uint32 = 2763701458
		hits := 0

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv uint32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintUint32_MarshalYAML(t *testing.T) {
	Convey("TestMapUintUint32.MarshalYAML", t, func() {
		var k uint = 2902678438
		var v uint32 = 2878145478

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintUint32_ToYAML(t *testing.T) {
	Convey("TestMapUintUint32.ToYAML", t, func() {
		var k uint = 1621245106
		var v uint32 = 3038339052

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintUint32_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintUint32.PutIfNotNil", t, func() {
		var k uint = 2192920645
		var v uint32 = 2634499786

		test := omap.NewMapUintUint32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1259487490, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint32 = 2493442572
		So(test.PutIfNotNil(2729722727, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintUint32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintUint32.ReplaceIfExists", t, func() {
		var k uint = 3014828734
		var v uint32 = 1687717317
		var x uint32 = 3489860568

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1084313760, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintUint32.ReplaceOrPut", t, func() {
		var k uint = 997711782
		var v uint32 = 1051191202
		var x uint32 = 1729264129

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1321653920, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint32_MarshalJSON(t *testing.T) {
	Convey("TestMapUintUint32.MarshalJSON", t, func() {
		var k uint = 3541667415
		var v uint32 = 3584702580

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":3541667415,"value":3584702580}]`)
	})
}
