package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintInt64_Put(t *testing.T) {
	Convey("TestMapUintInt64.Put", t, func() {
		var k uint = 1521075727
		var v int64 = 646510827325740352

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintInt64_Delete(t *testing.T) {
	Convey("TestMapUintInt64.Delete", t, func() {
		var k uint = 777855506
		var v int64 = 8039089174950141799

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintInt64_Has(t *testing.T) {
	Convey("TestMapUintInt64.Has", t, func() {
		var k uint = 2248581630
		var v int64 = 4611221503266571624

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1109424020+2798704545), ShouldBeFalse)
	})
}

func TestMapUintInt64_Get(t *testing.T) {
	Convey("TestMapUintInt64.Get", t, func() {
		var k uint = 3058670426
		var v int64 = 5735171447896092744

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(3079481617 + 508610493)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintInt64_GetOpt(t *testing.T) {
	Convey("TestMapUintInt64.GetOpt", t, func() {
		var k uint = 1686210878
		var v int64 = 6695162742776572444

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(897492953 + 2044794664)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintInt64_ForEach(t *testing.T) {
	Convey("TestMapUintInt64.ForEach", t, func() {
		var k uint = 2345605776
		var v int64 = 9112993897755724182
		hits := 0

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv int64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintInt64_MarshalYAML(t *testing.T) {
	Convey("TestMapUintInt64.MarshalYAML", t, func() {
		var k uint = 3714059951
		var v int64 = 6725872482185823408

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintInt64_ToYAML(t *testing.T) {
	Convey("TestMapUintInt64.ToYAML", t, func() {
		var k uint = 852145451
		var v int64 = 4937302327766349760

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintInt64_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintInt64.PutIfNotNil", t, func() {
		var k uint = 920963351
		var v int64 = 3967296493878263486

		test := omap.NewMapUintInt64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(3907479792, (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int64 = 7745296167635123491
		So(test.PutIfNotNil(3682534522, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintInt64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintInt64.ReplaceIfExists", t, func() {
		var k uint = 311567733
		var v int64 = 3813076208299287592
		var x int64 = 7475172892574524276

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(774212309, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintInt64.ReplaceOrPut", t, func() {
		var k uint = 2663736617
		var v int64 = 2026083713615965069
		var x int64 = 3841469042951866003

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(826053666, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt64_MarshalJSON(t *testing.T) {
	Convey("TestMapUintInt64.MarshalJSON", t, func() {
		var k uint = 3142221108
		var v int64 = 5054049778124634890

		test := omap.NewMapUintInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":3142221108,"value":5054049778124634890}]`)
	})
}
