package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintInt_Put(t *testing.T) {
	Convey("TestMapUintInt.Put", t, func() {
		var k uint = 2549709675
		var v int = 1810706985

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintInt_Delete(t *testing.T) {
	Convey("TestMapUintInt.Delete", t, func() {
		var k uint = 3374927430
		var v int = 621842298

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintInt_Has(t *testing.T) {
	Convey("TestMapUintInt.Has", t, func() {
		var k uint = 2195876457
		var v int = 301180880

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1749466400+368508652), ShouldBeFalse)
	})
}

func TestMapUintInt_Get(t *testing.T) {
	Convey("TestMapUintInt.Get", t, func() {
		var k uint = 1019397751
		var v int = 791056606

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(4212630400 + 3984098689)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintInt_GetOpt(t *testing.T) {
	Convey("TestMapUintInt.GetOpt", t, func() {
		var k uint = 264088188
		var v int = 1177792518

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(2043308116 + 39060934)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintInt_ForEach(t *testing.T) {
	Convey("TestMapUintInt.ForEach", t, func() {
		var k uint = 3145432718
		var v int = 39462293
		hits := 0

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv int) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintInt_MarshalYAML(t *testing.T) {
	Convey("TestMapUintInt.MarshalYAML", t, func() {
		var k uint = 1357871080
		var v int = 473752834

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintInt_ToYAML(t *testing.T) {
	Convey("TestMapUintInt.ToYAML", t, func() {
		var k uint = 3434357569
		var v int = 539759328

		test := omap.NewMapUintInt(1)

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

func TestMapUintInt_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintInt.PutIfNotNil", t, func() {
		var k uint = 2803511613
		var v int = 1421100473

		test := omap.NewMapUintInt(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(2037965675, (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int = 822639804
		So(test.PutIfNotNil(3978906146, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintInt_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintInt.ReplaceIfExists", t, func() {
		var k uint = 2108909713
		var v int = 44324211
		var x int = 2110384823

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(817217831, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintInt.ReplaceOrPut", t, func() {
		var k uint = 3694143755
		var v int = 274121040
		var x int = 784249140

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1185012364, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt_MarshalJSON(t *testing.T) {
	Convey("TestMapUintInt.MarshalJSON", t, func() {
		var k uint = 1378781493
		var v int = 1776557083

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1378781493,"value":1776557083}]`)
	})
}
