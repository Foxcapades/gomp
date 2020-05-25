package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntFloat64_Put(t *testing.T) {
	Convey("TestMapIntFloat64.Put", t, func() {
		var k int = 317132069
		var v float64 = 0.946

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntFloat64_Delete(t *testing.T) {
	Convey("TestMapIntFloat64.Delete", t, func() {
		var k int = 649115469
		var v float64 = 0.169

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntFloat64_Has(t *testing.T) {
	Convey("TestMapIntFloat64.Has", t, func() {
		var k int = 1657436770
		var v float64 = 0.663

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1959572014+1602525503), ShouldBeFalse)
	})
}

func TestMapIntFloat64_Get(t *testing.T) {
	Convey("TestMapIntFloat64.Get", t, func() {
		var k int = 259031675
		var v float64 = 0.732

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(1897187953 + 1809874984)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntFloat64_GetOpt(t *testing.T) {
	Convey("TestMapIntFloat64.GetOpt", t, func() {
		var k int = 396329855
		var v float64 = 0.593

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1118647868 + 314921125)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntFloat64_ForEach(t *testing.T) {
	Convey("TestMapIntFloat64.ForEach", t, func() {
		var k int = 85833081
		var v float64 = 0.544
		hits := 0

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv float64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntFloat64_MarshalYAML(t *testing.T) {
	Convey("TestMapIntFloat64.MarshalYAML", t, func() {
		var k int = 1128194736
		var v float64 = 0.105

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntFloat64_ToYAML(t *testing.T) {
	Convey("TestMapIntFloat64.ToYAML", t, func() {
		var k int = 455631748
		var v float64 = 0.879

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.SequenceNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 1)
		So(xyml.IsMap(c.Content[0]), ShouldBeTrue)
	})
}

func TestMapIntFloat64_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntFloat64.PutIfNotNil", t, func() {
		var k int = 734748165
		var v float64 = 0.564

		test := omap.NewMapIntFloat64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(2136268133, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float64 = 0.015
		So(test.PutIfNotNil(1078466750, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntFloat64.ReplaceIfExists", t, func() {
		var k int = 265688216
		var v float64 = 0.602
		var x float64 = 0.413

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(657781498, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntFloat64.ReplaceOrPut", t, func() {
		var k int = 1023568961
		var v float64 = 0.746
		var x float64 = 0.486

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1264104950, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat64_MarshalJSON(t *testing.T) {
	Convey("TestMapIntFloat64.MarshalJSON", t, func() {
		var k int = 1840325498
		var v float64 = 0.129

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1840325498,"value":0.129}]`)
	})
}
