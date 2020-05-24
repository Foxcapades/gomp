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
		var k uint = 3195164876
		var v uint64 = 14943015831524131118

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintUint64_Delete(t *testing.T) {
	Convey("TestMapUintUint64.Delete", t, func() {
		var k uint = 3824609091
		var v uint64 = 15056373514664388886

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintUint64_Has(t *testing.T) {
	Convey("TestMapUintUint64.Has", t, func() {
		var k uint = 2195095246
		var v uint64 = 17318055835187908965

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1141763642+57042133), ShouldBeFalse)
	})
}


func TestMapUintUint64_Get(t *testing.T) {
	Convey("TestMapUintUint64.Get", t, func() {
		var k uint = 2506323977
		var v uint64 = 4844531649570386764

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(73986321+1925120676)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintUint64_GetOpt(t *testing.T) {
	Convey("TestMapUintUint64.GetOpt", t, func() {
		var k uint = 4193000543
		var v uint64 = 7258649412438209559

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(3684870492+3135122109)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintUint64_ForEach(t *testing.T) {
	Convey("TestMapUintUint64.ForEach", t, func() {
		var k uint = 2360813432
		var v uint64 = 12804441906295073762
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
		var k uint = 2403444366
		var v uint64 = 7496626535180669951

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
		var k uint = 1672019475
		var v uint64 = 2697482299904269207

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
		var k uint = 1672189925
		var v uint64 = 4414407328072527814

		test := omap.NewMapUintUint64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(471067701, (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint64 = 4599242050934324803
		So(test.PutIfNotNil(5123943, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintUint64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintUint64.ReplaceIfExists", t, func() {
		var k uint = 3524362603
		var v uint64 = 7270273366926959294
		var x uint64 = 3623947803923801473

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(3162365561, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintUint64.ReplaceOrPut", t, func() {
		var k uint = 3508919764
		var v uint64 = 16565547506189355814
		var x uint64 = 8846290979205175743

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1833722599, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint64_MarshalJSON(t *testing.T) {
	Convey("TestMapUintUint64.MarshalJSON", t, func() {
		var k uint = 2057631784
		var v uint64 = 1872242443227096458

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":2057631784,"value":1872242443227096458}]`)
	})
}

