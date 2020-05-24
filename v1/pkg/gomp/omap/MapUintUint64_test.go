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
		var k uint = 401073296
		var v uint64 = 14241965901004183428

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintUint64_Delete(t *testing.T) {
	Convey("TestMapUintUint64.Delete", t, func() {
		var k uint = 3614298117
		var v uint64 = 3199546508052667233

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintUint64_Has(t *testing.T) {
	Convey("TestMapUintUint64.Has", t, func() {
		var k uint = 3764498302
		var v uint64 = 12369901253421588522

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(3877571309+1749534206), ShouldBeFalse)
	})
}


func TestMapUintUint64_Get(t *testing.T) {
	Convey("TestMapUintUint64.Get", t, func() {
		var k uint = 3631849055
		var v uint64 = 5085685987775272736

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(3704035988 + 431609023)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintUint64_GetOpt(t *testing.T) {
	Convey("TestMapUintUint64.GetOpt", t, func() {
		var k uint = 645624430
		var v uint64 = 3445191104954471595

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(3335481991 + 26910770)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintUint64_ForEach(t *testing.T) {
	Convey("TestMapUintUint64.ForEach", t, func() {
		var k uint = 3561222906
		var v uint64 = 9488581547869744665
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
		var k uint = 3447752248
		var v uint64 = 3385609147004464541

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
		var k uint = 3786302112
		var v uint64 = 1868980048074670534

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
		var k uint = 3922286644
		var v uint64 = 3297466213237054303

		test := omap.NewMapUintUint64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1626210465, (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint64 = 15647843301317855333
		So(test.PutIfNotNil(1958188569, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintUint64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintUint64.ReplaceIfExists", t, func() {
		var k uint = 4284706707
		var v uint64 = 16046660520637596882
		var x uint64 = 10935010837922465011

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(2094038570, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintUint64.ReplaceOrPut", t, func() {
		var k uint = 3036648625
		var v uint64 = 3284364105808057494
		var x uint64 = 9379814765346700894

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(3999323616, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint64_MarshalJSON(t *testing.T) {
	Convey("TestMapUintUint64.MarshalJSON", t, func() {
		var k uint = 4165425127
		var v uint64 = 11039243419998285807

		test := omap.NewMapUintUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":4165425127,"value":11039243419998285807}]`)
	})
}
