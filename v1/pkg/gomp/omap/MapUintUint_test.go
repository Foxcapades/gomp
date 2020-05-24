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
		var k uint = 2037986963
		var v uint = 846468835

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintUint_Delete(t *testing.T) {
	Convey("TestMapUintUint.Delete", t, func() {
		var k uint = 4106566399
		var v uint = 1521021553

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintUint_Has(t *testing.T) {
	Convey("TestMapUintUint.Has", t, func() {
		var k uint = 631511834
		var v uint = 1405022633

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1745691025+826287658), ShouldBeFalse)
	})
}

func TestMapUintUint_Get(t *testing.T) {
	Convey("TestMapUintUint.Get", t, func() {
		var k uint = 4222551692
		var v uint = 3038764912

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(2917600120 + 2525296420)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintUint_GetOpt(t *testing.T) {
	Convey("TestMapUintUint.GetOpt", t, func() {
		var k uint = 3097308597
		var v uint = 767666310

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(3779294268 + 1566940566)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintUint_ForEach(t *testing.T) {
	Convey("TestMapUintUint.ForEach", t, func() {
		var k uint = 2088706659
		var v uint = 2530563251
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
		var k uint = 1116784310
		var v uint = 3537225124

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
		var k uint = 2769339389
		var v uint = 1179144841

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
		var k uint = 1804925799
		var v uint = 3913440555

		test := omap.NewMapUintUint(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1177994037, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint = 2430445687
		So(test.PutIfNotNil(1818722851, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintUint_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintUint.ReplaceIfExists", t, func() {
		var k uint = 162755311
		var v uint = 1221468297
		var x uint = 718742132

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1223155925, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintUint.ReplaceOrPut", t, func() {
		var k uint = 1831845704
		var v uint = 3877242222
		var x uint = 3259045408

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(2771192883, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint_MarshalJSON(t *testing.T) {
	Convey("TestMapUintUint.MarshalJSON", t, func() {
		var k uint = 2409171353
		var v uint = 897164643

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":2409171353,"value":897164643}]`)
	})
}
