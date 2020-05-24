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
		var k int = 1775993622
		var v uint = 616726289

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntUint_Delete(t *testing.T) {
	Convey("TestMapIntUint.Delete", t, func() {
		var k int = 725392006
		var v uint = 3431801820

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntUint_Has(t *testing.T) {
	Convey("TestMapIntUint.Has", t, func() {
		var k int = 1152425457
		var v uint = 1052129802

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1290281624+1041432950), ShouldBeFalse)
	})
}

func TestMapIntUint_Get(t *testing.T) {
	Convey("TestMapIntUint.Get", t, func() {
		var k int = 1588421212
		var v uint = 716403925

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(138405513 + 1346033051)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntUint_GetOpt(t *testing.T) {
	Convey("TestMapIntUint.GetOpt", t, func() {
		var k int = 826417461
		var v uint = 1563927694

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(359275309 + 1402369215)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntUint_ForEach(t *testing.T) {
	Convey("TestMapIntUint.ForEach", t, func() {
		var k int = 1930305116
		var v uint = 3262178325
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
		var k int = 1084563052
		var v uint = 888619397

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
		var k int = 2103064784
		var v uint = 1965357404

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntUint_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntUint.PutIfNotNil", t, func() {
		var k int = 2005251484
		var v uint = 1716303691

		test := omap.NewMapIntUint(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1988993034, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint = 1494012642
		So(test.PutIfNotNil(130776568, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntUint_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntUint.ReplaceIfExists", t, func() {
		var k int = 1209440733
		var v uint = 2630478861
		var x uint = 2942159598

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1757526878, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntUint.ReplaceOrPut", t, func() {
		var k int = 1438900658
		var v uint = 1423692167
		var x uint = 903030732

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(517369117, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint_MarshalJSON(t *testing.T) {
	Convey("TestMapIntUint.MarshalJSON", t, func() {
		var k int = 24977686
		var v uint = 3569082655

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":24977686,"value":3569082655}]`)
	})
}
