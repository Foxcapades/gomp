package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntInt64_Put(t *testing.T) {
	Convey("TestMapIntInt64.Put", t, func() {
		var k int = 850180116
		var v int64 = 3412735533454777820

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntInt64_Delete(t *testing.T) {
	Convey("TestMapIntInt64.Delete", t, func() {
		var k int = 991942433
		var v int64 = 401220724821667065

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntInt64_Has(t *testing.T) {
	Convey("TestMapIntInt64.Has", t, func() {
		var k int = 1999144250
		var v int64 = 7322743954280573338

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(34351926+636939728), ShouldBeFalse)
	})
}

func TestMapIntInt64_Get(t *testing.T) {
	Convey("TestMapIntInt64.Get", t, func() {
		var k int = 1747644693
		var v int64 = 8896954659552692760

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(1202349099 + 1946264638)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntInt64_GetOpt(t *testing.T) {
	Convey("TestMapIntInt64.GetOpt", t, func() {
		var k int = 896370897
		var v int64 = 8007567424646371633

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(2141043937 + 906293871)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntInt64_ForEach(t *testing.T) {
	Convey("TestMapIntInt64.ForEach", t, func() {
		var k int = 2056944463
		var v int64 = 3037516473968670926
		hits := 0

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv int64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntInt64_MarshalYAML(t *testing.T) {
	Convey("TestMapIntInt64.MarshalYAML", t, func() {
		var k int = 774455767
		var v int64 = 2395929978196370676

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntInt64_ToYAML(t *testing.T) {
	Convey("TestMapIntInt64.ToYAML", t, func() {
		var k int = 824302521
		var v int64 = 7894079724697301638

		test := omap.NewMapIntInt64(1)

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

func TestMapIntInt64_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntInt64.PutIfNotNil", t, func() {
		var k int = 96469019
		var v int64 = 7665314068164947885

		test := omap.NewMapIntInt64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(441939075, (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int64 = 4114282557146531995
		So(test.PutIfNotNil(445797185, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntInt64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntInt64.ReplaceIfExists", t, func() {
		var k int = 950891983
		var v int64 = 2056935792651446362
		var x int64 = 167273866149821514

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1024443931, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntInt64.ReplaceOrPut", t, func() {
		var k int = 1561852673
		var v int64 = 4684745435096929055
		var x int64 = 3581396157271868203

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1819106381, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt64_MarshalJSON(t *testing.T) {
	Convey("TestMapIntInt64.MarshalJSON", t, func() {
		var k int = 2070118336
		var v int64 = 5853004860950808376

		test := omap.NewMapIntInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":2070118336,"value":5853004860950808376}]`)
	})
}
