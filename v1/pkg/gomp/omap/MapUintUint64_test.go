package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintUint64_Put(t *testing.T) {
	Convey("TestMapUintUint64.Put", t, func() {
		var k uint = 300140129
		var v uint64 = 5371093873883593936

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintUint64_Delete(t *testing.T) {
	Convey("TestMapUintUint64.Delete", t, func() {
		var k uint = 679498791
		var v uint64 = 12592007456484516739

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintUint64_Has(t *testing.T) {
	Convey("TestMapUintUint64.Has", t, func() {
		var k uint = 2983755153
		var v uint64 = 13846213669690852465

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(3144646577+4201232187), ShouldBeFalse)
	})
}


func TestMapUintUint64_Get(t *testing.T) {
	Convey("TestMapUintUint64.Get", t, func() {
		var k uint = 633477948
		var v uint64 = 7885935744185915020

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(4078972245 + 1407132682)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintUint64_GetOpt(t *testing.T) {
	Convey("TestMapUintUint64.GetOpt", t, func() {
		var k uint = 931846474
		var v uint64 = 9696192385204407560

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(3948309728 + 1528013803)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintUint64_ForEach(t *testing.T) {
	Convey("TestMapUintUint64.ForEach", t, func() {
		var k uint = 2117842866
		var v uint64 = 9396876574346753523
		hits := 0

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv uint64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintUint64_MarshalYAML(t *testing.T) {
	Convey("TestMapUintUint64.MarshalYAML", t, func() {
		var k uint = 3283671809
		var v uint64 = 11629528472208164287

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintUint64_ToYAML(t *testing.T) {
	Convey("TestMapUintUint64.ToYAML", t, func() {
		var k uint = 3053470883
		var v uint64 = 15100500901871371146

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintUint64_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintUint64.PutIfNotNil", t, func() {
		var k uint = 3329741156
		var v uint64 = 13376215520074945955

		test := omap.NewMapUintUint64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(2162035230, (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint64 = 9096088272312374785
		So(test.PutIfNotNil(3815872907, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintUint64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintUint64.ReplaceIfExists", t, func() {
		var k uint = 498108104
		var v uint64 = 11198819489269292301
		var x uint64 = 12230847645557333982

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(3300327364, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintUint64.ReplaceOrPut", t, func() {
		var k uint = 971058307
		var v uint64 = 14101170444365354593
		var x uint64 = 17081291271212899698

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(597166732, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint64_MarshalJSON(t *testing.T) {
	Convey("TestMapUintUint64.MarshalJSON", t, func() {
		var k uint = 217365958
		var v uint64 = 13789777080266667019

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":217365958,"value":13789777080266667019}]`)
	})
}

