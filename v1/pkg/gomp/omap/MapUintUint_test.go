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
		var k uint = 4288782820
		var v uint = 2097607349

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintUint_Delete(t *testing.T) {
	Convey("TestMapUintUint.Delete", t, func() {
		var k uint = 896843533
		var v uint = 3880177886

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintUint_Has(t *testing.T) {
	Convey("TestMapUintUint.Has", t, func() {
		var k uint = 1181908411
		var v uint = 3804899197

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1633253665+2263024289), ShouldBeFalse)
	})
}

func TestMapUintUint_Get(t *testing.T) {
	Convey("TestMapUintUint.Get", t, func() {
		var k uint = 4021568913
		var v uint = 1149082510

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(1214515042 + 3505409533)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintUint_GetOpt(t *testing.T) {
	Convey("TestMapUintUint.GetOpt", t, func() {
		var k uint = 2628322762
		var v uint = 2843824436

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1694563923 + 881560593)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintUint_ForEach(t *testing.T) {
	Convey("TestMapUintUint.ForEach", t, func() {
		var k uint = 2563958437
		var v uint = 1283651226
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
		var k uint = 232458985
		var v uint = 2152310472

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
		var k uint = 1685224503
		var v uint = 1665428142

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
		var k uint = 4113686793
		var v uint = 421846641

		test := omap.NewMapUintUint(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(706425458, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint = 3282910013
		So(test.PutIfNotNil(1597293067, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintUint_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintUint.ReplaceIfExists", t, func() {
		var k uint = 981784256
		var v uint = 1394013605
		var x uint = 2867266338

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1307403160, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintUint.ReplaceOrPut", t, func() {
		var k uint = 2964397826
		var v uint = 2036593349
		var x uint = 1948033605

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(3178762812, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint_MarshalJSON(t *testing.T) {
	Convey("TestMapUintUint.MarshalJSON", t, func() {
		var k uint = 335659181
		var v uint = 1354714519

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":335659181,"value":1354714519}]`)
	})
}
