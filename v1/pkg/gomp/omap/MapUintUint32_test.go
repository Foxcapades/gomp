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
		var k uint = 804387621
		var v uint32 = 2322774889

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintUint32_Delete(t *testing.T) {
	Convey("TestMapUintUint32.Delete", t, func() {
		var k uint = 2546951691
		var v uint32 = 2757919272

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintUint32_Has(t *testing.T) {
	Convey("TestMapUintUint32.Has", t, func() {
		var k uint = 4292962799
		var v uint32 = 985676696

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1259736284+4245896495), ShouldBeFalse)
	})
}

func TestMapUintUint32_Get(t *testing.T) {
	Convey("TestMapUintUint32.Get", t, func() {
		var k uint = 1812420167
		var v uint32 = 1522030917

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(4246000510 + 3223922041)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintUint32_GetOpt(t *testing.T) {
	Convey("TestMapUintUint32.GetOpt", t, func() {
		var k uint = 3920273532
		var v uint32 = 4108292519

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(2284369679 + 3557714748)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintUint32_ForEach(t *testing.T) {
	Convey("TestMapUintUint32.ForEach", t, func() {
		var k uint = 1709853510
		var v uint32 = 1406175794
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
		var k uint = 2941613996
		var v uint32 = 2844487259

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
		var k uint = 1663704659
		var v uint32 = 1717880188

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
		var k uint = 1873286049
		var v uint32 = 3119088696

		test := omap.NewMapUintUint32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(3501309939, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint32 = 1152607515
		So(test.PutIfNotNil(2928699158, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintUint32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintUint32.ReplaceIfExists", t, func() {
		var k uint = 1451699220
		var v uint32 = 2701942887
		var x uint32 = 2880075964

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(3263119622, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintUint32.ReplaceOrPut", t, func() {
		var k uint = 514229479
		var v uint32 = 3526069884
		var x uint32 = 1076251374

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(3074036712, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint32_MarshalJSON(t *testing.T) {
	Convey("TestMapUintUint32.MarshalJSON", t, func() {
		var k uint = 4149664777
		var v uint32 = 909984975

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":4149664777,"value":909984975}]`)
	})
}
