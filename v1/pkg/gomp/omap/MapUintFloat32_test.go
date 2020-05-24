package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintFloat32_Put(t *testing.T) {
	Convey("TestMapUintFloat32.Put", t, func() {
		var k uint = 2247039020
		var v float32 = 0.661

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintFloat32_Delete(t *testing.T) {
	Convey("TestMapUintFloat32.Delete", t, func() {
		var k uint = 1010954877
		var v float32 = 0.363

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintFloat32_Has(t *testing.T) {
	Convey("TestMapUintFloat32.Has", t, func() {
		var k uint = 570911090
		var v float32 = 0.830

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(2351150567+1268789495), ShouldBeFalse)
	})
}

func TestMapUintFloat32_Get(t *testing.T) {
	Convey("TestMapUintFloat32.Get", t, func() {
		var k uint = 1024548918
		var v float32 = 0.935

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(518215597 + 493320466)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintFloat32_GetOpt(t *testing.T) {
	Convey("TestMapUintFloat32.GetOpt", t, func() {
		var k uint = 2216244281
		var v float32 = 0.595

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(2352246899 + 2834445622)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintFloat32_ForEach(t *testing.T) {
	Convey("TestMapUintFloat32.ForEach", t, func() {
		var k uint = 2888327318
		var v float32 = 0.966
		hits := 0

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv float32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintFloat32_MarshalYAML(t *testing.T) {
	Convey("TestMapUintFloat32.MarshalYAML", t, func() {
		var k uint = 11741271
		var v float32 = 0.370

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintFloat32_ToYAML(t *testing.T) {
	Convey("TestMapUintFloat32.ToYAML", t, func() {
		var k uint = 3300785251
		var v float32 = 0.670

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintFloat32_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintFloat32.PutIfNotNil", t, func() {
		var k uint = 2709965072
		var v float32 = 0.281

		test := omap.NewMapUintFloat32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(2240507926, (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float32 = 0.832
		So(test.PutIfNotNil(2681446701, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintFloat32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintFloat32.ReplaceIfExists", t, func() {
		var k uint = 3054550330
		var v float32 = 0.014
		var x float32 = 0.727

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(475576912, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintFloat32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintFloat32.ReplaceOrPut", t, func() {
		var k uint = 2732827270
		var v float32 = 0.515
		var x float32 = 0.679

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1847967889, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintFloat32_MarshalJSON(t *testing.T) {
	Convey("TestMapUintFloat32.MarshalJSON", t, func() {
		var k uint = 1204094430
		var v float32 = 0.028

		test := omap.NewMapUintFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1204094430,"value":0.028}]`)
	})
}
