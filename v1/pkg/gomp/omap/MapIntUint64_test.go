package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntUint64_Put(t *testing.T) {
	Convey("TestMapIntUint64.Put", t, func() {
		var k int = 224528821
		var v uint64 = 7026551793418468102

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntUint64_Delete(t *testing.T) {
	Convey("TestMapIntUint64.Delete", t, func() {
		var k int = 1033733650
		var v uint64 = 4724491074647007096

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntUint64_Has(t *testing.T) {
	Convey("TestMapIntUint64.Has", t, func() {
		var k int = 1989216099
		var v uint64 = 12967328373775260787

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1768429276+652093724), ShouldBeFalse)
	})
}

func TestMapIntUint64_Get(t *testing.T) {
	Convey("TestMapIntUint64.Get", t, func() {
		var k int = 897526297
		var v uint64 = 1730099981071377426

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(1706736850 + 1036942674)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntUint64_GetOpt(t *testing.T) {
	Convey("TestMapIntUint64.GetOpt", t, func() {
		var k int = 1673784437
		var v uint64 = 12702475403523338068

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(108726162 + 540956739)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntUint64_ForEach(t *testing.T) {
	Convey("TestMapIntUint64.ForEach", t, func() {
		var k int = 1188267967
		var v uint64 = 16406025550121680011
		hits := 0

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv uint64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntUint64_MarshalYAML(t *testing.T) {
	Convey("TestMapIntUint64.MarshalYAML", t, func() {
		var k int = 526343258
		var v uint64 = 13303551199148085711

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntUint64_ToYAML(t *testing.T) {
	Convey("TestMapIntUint64.ToYAML", t, func() {
		var k int = 72451473
		var v uint64 = 8701497950073253868

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntUint64_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntUint64.PutIfNotNil", t, func() {
		var k int = 946205235
		var v uint64 = 16115420760079533951

		test := omap.NewMapIntUint64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(811533446, (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint64 = 3951189552369881134
		So(test.PutIfNotNil(1065524146, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntUint64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntUint64.ReplaceIfExists", t, func() {
		var k int = 613298069
		var v uint64 = 793293855622280624
		var x uint64 = 5233310757055673633

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(748932212, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntUint64.ReplaceOrPut", t, func() {
		var k int = 1697732272
		var v uint64 = 6099948485243198864
		var x uint64 = 2353224418260255631

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1774647717, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint64_MarshalJSON(t *testing.T) {
	Convey("TestMapIntUint64.MarshalJSON", t, func() {
		var k int = 1069140245
		var v uint64 = 10751223762647717374

		test := omap.NewMapIntUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1069140245,"value":10751223762647717374}]`)
	})
}
