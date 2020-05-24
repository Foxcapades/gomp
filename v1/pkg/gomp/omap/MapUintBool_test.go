package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintBool_Put(t *testing.T) {
	Convey("TestMapUintBool.Put", t, func() {
		var k uint = 3318554097
		var v bool = false

		test := omap.NewMapUintBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintBool_Delete(t *testing.T) {
	Convey("TestMapUintBool.Delete", t, func() {
		var k uint = 894698308
		var v bool = false

		test := omap.NewMapUintBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintBool_Has(t *testing.T) {
	Convey("TestMapUintBool.Has", t, func() {
		var k uint = 2457971398
		var v bool = false

		test := omap.NewMapUintBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(2396078780+3544947539), ShouldBeFalse)
	})
}

func TestMapUintBool_Get(t *testing.T) {
	Convey("TestMapUintBool.Get", t, func() {
		var k uint = 2751430474
		var v bool = false

		test := omap.NewMapUintBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1331440829 + 592019664)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintBool_GetOpt(t *testing.T) {
	Convey("TestMapUintBool.GetOpt", t, func() {
		var k uint = 1169799061
		var v bool = false

		test := omap.NewMapUintBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(2434837342 + 2904789990)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintBool_ForEach(t *testing.T) {
	Convey("TestMapUintBool.ForEach", t, func() {
		var k uint = 1662238368
		var v bool = false
		hits := 0

		test := omap.NewMapUintBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv bool) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintBool_MarshalYAML(t *testing.T) {
	Convey("TestMapUintBool.MarshalYAML", t, func() {
		var k uint = 412178453
		var v bool = false

		test := omap.NewMapUintBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintBool_ToYAML(t *testing.T) {
	Convey("TestMapUintBool.ToYAML", t, func() {
		var k uint = 2714636781
		var v bool = false

		test := omap.NewMapUintBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintBool_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintBool.PutIfNotNil", t, func() {
		var k uint = 4049896568
		var v bool = false

		test := omap.NewMapUintBool(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*bool)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(3753818389, (*bool)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x bool = false
		So(test.PutIfNotNil(3146231441, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintBool_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintBool.ReplaceIfExists", t, func() {
		var k uint = 3665810487
		var v bool = false
		var x bool = false

		test := omap.NewMapUintBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(774555407, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintBool_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintBool.ReplaceOrPut", t, func() {
		var k uint = 287285739
		var v bool = false
		var x bool = false

		test := omap.NewMapUintBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(2980216582, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintBool_MarshalJSON(t *testing.T) {
	Convey("TestMapUintBool.MarshalJSON", t, func() {
		var k uint = 3322603370
		var v bool = false

		test := omap.NewMapUintBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":3322603370,"value":false}]`)
	})
}
