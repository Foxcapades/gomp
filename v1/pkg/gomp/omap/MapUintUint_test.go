package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintUint_Put(t *testing.T) {
	Convey("TestMapUintUint.Put", t, func() {
		var k uint = 88514979
		var v uint = 560137662

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintUint_Delete(t *testing.T) {
	Convey("TestMapUintUint.Delete", t, func() {
		var k uint = 1755913377
		var v uint = 1727432710

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintUint_Has(t *testing.T) {
	Convey("TestMapUintUint.Has", t, func() {
		var k uint = 295641132
		var v uint = 59892679

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(2700926086+2595055410), ShouldBeFalse)
	})
}

func TestMapUintUint_Get(t *testing.T) {
	Convey("TestMapUintUint.Get", t, func() {
		var k uint = 2106279098
		var v uint = 1860332979

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(1669971750 + 3388843463)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintUint_GetOpt(t *testing.T) {
	Convey("TestMapUintUint.GetOpt", t, func() {
		var k uint = 3188628568
		var v uint = 1131272595

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(3902187832 + 1343141496)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintUint_ForEach(t *testing.T) {
	Convey("TestMapUintUint.ForEach", t, func() {
		var k uint = 1613497708
		var v uint = 1822154564
		hits := 0

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv uint) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintUint_MarshalYAML(t *testing.T) {
	Convey("TestMapUintUint.MarshalYAML", t, func() {
		var k uint = 241745971
		var v uint = 2892565960

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintUint_ToYAML(t *testing.T) {
	Convey("TestMapUintUint.ToYAML", t, func() {
		var k uint = 1095859078
		var v uint = 2121339562

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintUint_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintUint.PutIfNotNil", t, func() {
		var k uint = 3880093453
		var v uint = 4086587947

		test := omap.NewMapUintUint(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(3785045458, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint = 3153306157
		So(test.PutIfNotNil(826880005, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintUint_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintUint.ReplaceIfExists", t, func() {
		var k uint = 724559517
		var v uint = 3178241185
		var x uint = 788837016

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(340987760, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintUint.ReplaceOrPut", t, func() {
		var k uint = 31337864
		var v uint = 3351263303
		var x uint = 1919436958

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(383974811, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint_MarshalJSON(t *testing.T) {
	Convey("TestMapUintUint.MarshalJSON", t, func() {
		var k uint = 973453616
		var v uint = 2608501528

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":973453616,"value":2608501528}]`)
	})
}
