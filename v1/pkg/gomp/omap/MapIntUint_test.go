package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntUint_Put(t *testing.T) {
	Convey("TestMapIntUint.Put", t, func() {
		var k int = 1635107658
		var v uint = 2186139297

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntUint_Delete(t *testing.T) {
	Convey("TestMapIntUint.Delete", t, func() {
		var k int = 305565093
		var v uint = 4218210736

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntUint_Has(t *testing.T) {
	Convey("TestMapIntUint.Has", t, func() {
		var k int = 1348956462
		var v uint = 3656245963

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(779966835+1073508089), ShouldBeFalse)
	})
}

func TestMapIntUint_Get(t *testing.T) {
	Convey("TestMapIntUint.Get", t, func() {
		var k int = 1372376885
		var v uint = 2741534292

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(269613520 + 1597648792)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntUint_GetOpt(t *testing.T) {
	Convey("TestMapIntUint.GetOpt", t, func() {
		var k int = 402736263
		var v uint = 1131431436

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1800581026 + 1429699656)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntUint_ForEach(t *testing.T) {
	Convey("TestMapIntUint.ForEach", t, func() {
		var k int = 797820817
		var v uint = 3551542197
		hits := 0

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv uint) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntUint_MarshalYAML(t *testing.T) {
	Convey("TestMapIntUint.MarshalYAML", t, func() {
		var k int = 1266953310
		var v uint = 1181138404

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntUint_ToYAML(t *testing.T) {
	Convey("TestMapIntUint.ToYAML", t, func() {
		var k int = 1182910486
		var v uint = 757526674

		test := omap.NewMapIntUint(1)

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

func TestMapIntUint_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntUint.PutIfNotNil", t, func() {
		var k int = 1777824793
		var v uint = 1130305917

		test := omap.NewMapIntUint(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(6523476, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint = 1981634468
		So(test.PutIfNotNil(1694353679, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntUint_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntUint.ReplaceIfExists", t, func() {
		var k int = 1976061866
		var v uint = 512973796
		var x uint = 27597848

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1310516259, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntUint.ReplaceOrPut", t, func() {
		var k int = 1579706853
		var v uint = 3506233727
		var x uint = 1213755016

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1763939036, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint_MarshalJSON(t *testing.T) {
	Convey("TestMapIntUint.MarshalJSON", t, func() {
		var k int = 1180674490
		var v uint = 148889116

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1180674490,"value":148889116}]`)
	})
}
