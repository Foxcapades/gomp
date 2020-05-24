package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntInt_Put(t *testing.T) {
	Convey("TestMapIntInt.Put", t, func() {
		var k int = 731700009
		var v int = 68769341

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntInt_Delete(t *testing.T) {
	Convey("TestMapIntInt.Delete", t, func() {
		var k int = 2061727911
		var v int = 98305988

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntInt_Has(t *testing.T) {
	Convey("TestMapIntInt.Has", t, func() {
		var k int = 1149636917
		var v int = 1610748341

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1328830671+2051273981), ShouldBeFalse)
	})
}

func TestMapIntInt_Get(t *testing.T) {
	Convey("TestMapIntInt.Get", t, func() {
		var k int = 700927103
		var v int = 1594387273

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(1218554589 + 426038943)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntInt_GetOpt(t *testing.T) {
	Convey("TestMapIntInt.GetOpt", t, func() {
		var k int = 473763786
		var v int = 1427497969

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1687026700 + 397143562)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntInt_ForEach(t *testing.T) {
	Convey("TestMapIntInt.ForEach", t, func() {
		var k int = 577835790
		var v int = 962237204
		hits := 0

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv int) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntInt_MarshalYAML(t *testing.T) {
	Convey("TestMapIntInt.MarshalYAML", t, func() {
		var k int = 995838708
		var v int = 1398479153

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntInt_ToYAML(t *testing.T) {
	Convey("TestMapIntInt.ToYAML", t, func() {
		var k int = 1115795077
		var v int = 182775176

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntInt_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntInt.PutIfNotNil", t, func() {
		var k int = 997834585
		var v int = 1771407564

		test := omap.NewMapIntInt(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1960558818, (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int = 167814552
		So(test.PutIfNotNil(1182768914, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntInt_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntInt.ReplaceIfExists", t, func() {
		var k int = 1933410176
		var v int = 850409958
		var x int = 707132253

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1297756705, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntInt.ReplaceOrPut", t, func() {
		var k int = 1074992660
		var v int = 854170439
		var x int = 607885257

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1827211461, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt_MarshalJSON(t *testing.T) {
	Convey("TestMapIntInt.MarshalJSON", t, func() {
		var k int = 727057543
		var v int = 235340622

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":727057543,"value":235340622}]`)
	})
}
