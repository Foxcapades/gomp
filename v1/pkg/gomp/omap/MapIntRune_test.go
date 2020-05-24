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
		var k int = 460109856
		var v rune = 1424629575

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntRune_Delete(t *testing.T) {
	Convey("TestMapIntRune.Delete", t, func() {
		var k int = 621803344
		var v rune = 109593868

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntRune_Has(t *testing.T) {
	Convey("TestMapIntRune.Has", t, func() {
		var k int = 2126004469
		var v rune = 441888826

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(39571396+660812291), ShouldBeFalse)
	})
}

func TestMapIntRune_Get(t *testing.T) {
	Convey("TestMapIntRune.Get", t, func() {
		var k int = 1777621445
		var v rune = 1365811289

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(552875319 + 2129951120)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntRune_GetOpt(t *testing.T) {
	Convey("TestMapIntRune.GetOpt", t, func() {
		var k int = 100249928
		var v rune = 1868995725

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(278705934 + 1580771416)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntRune_ForEach(t *testing.T) {
	Convey("TestMapIntRune.ForEach", t, func() {
		var k int = 640924757
		var v rune = 2109701261
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
		var k int = 2093047564
		var v rune = 1476360598

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
		var k int = 2065453708
		var v rune = 1633539456

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
		var k int = 1559223453
		var v rune = 312726043

		test := omap.NewMapIntRune(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1305000734, (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x rune = 1834122265
		So(test.PutIfNotNil(212712158, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntRune_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntRune.ReplaceIfExists", t, func() {
		var k int = 61210836
		var v rune = 1424769722
		var x rune = 2136494981

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(98030003, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntRune_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntRune.ReplaceOrPut", t, func() {
		var k int = 1256414833
		var v rune = 863364425
		var x rune = 1796088315

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(756157397, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntRune_MarshalJSON(t *testing.T) {
	Convey("TestMapIntRune.MarshalJSON", t, func() {
		var k int = 777045248
		var v rune = 2010086074

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":777045248,"value":2010086074}]`)
	})
}
