package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntRune_Put(t *testing.T) {
	Convey("TestMapIntRune.Put", t, func() {
		var k int = 1333949346
		var v rune = 220399048

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntRune_Delete(t *testing.T) {
	Convey("TestMapIntRune.Delete", t, func() {
		var k int = 1407651016
		var v rune = 561757521

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntRune_Has(t *testing.T) {
	Convey("TestMapIntRune.Has", t, func() {
		var k int = 782821688
		var v rune = 1433878035

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1149499002+443913921), ShouldBeFalse)
	})
}

func TestMapIntRune_Get(t *testing.T) {
	Convey("TestMapIntRune.Get", t, func() {
		var k int = 302923101
		var v rune = 164209742

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(1444884229 + 292247848)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntRune_GetOpt(t *testing.T) {
	Convey("TestMapIntRune.GetOpt", t, func() {
		var k int = 1814773714
		var v rune = 2009020821

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(2142340031 + 909709894)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntRune_ForEach(t *testing.T) {
	Convey("TestMapIntRune.ForEach", t, func() {
		var k int = 1396456536
		var v rune = 1557193987
		hits := 0

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv rune) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntRune_MarshalYAML(t *testing.T) {
	Convey("TestMapIntRune.MarshalYAML", t, func() {
		var k int = 2126931915
		var v rune = 1486780623

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntRune_ToYAML(t *testing.T) {
	Convey("TestMapIntRune.ToYAML", t, func() {
		var k int = 1997449569
		var v rune = 409112030

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntRune_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntRune.PutIfNotNil", t, func() {
		var k int = 263337551
		var v rune = 1966150409

		test := omap.NewMapIntRune(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1604161290, (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x rune = 1920951008
		So(test.PutIfNotNil(1447215369, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntRune_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntRune.ReplaceIfExists", t, func() {
		var k int = 645702847
		var v rune = 692606526
		var x rune = 950252411

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(728165548, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntRune_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntRune.ReplaceOrPut", t, func() {
		var k int = 34762380
		var v rune = 1547597476
		var x rune = 426044190

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1778180340, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntRune_MarshalJSON(t *testing.T) {
	Convey("TestMapIntRune.MarshalJSON", t, func() {
		var k int = 790980913
		var v rune = 1182590916

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":790980913,"value":1182590916}]`)
	})
}
