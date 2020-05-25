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
		var k int = 475612719
		var v rune = 724479785

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntRune_Delete(t *testing.T) {
	Convey("TestMapIntRune.Delete", t, func() {
		var k int = 1924598750
		var v rune = 1444917013

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntRune_Has(t *testing.T) {
	Convey("TestMapIntRune.Has", t, func() {
		var k int = 1260047181
		var v rune = 236064769

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(212778955+898487068), ShouldBeFalse)
	})
}

func TestMapIntRune_Get(t *testing.T) {
	Convey("TestMapIntRune.Get", t, func() {
		var k int = 555823231
		var v rune = 1773695913

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(1007730461 + 783093618)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntRune_GetOpt(t *testing.T) {
	Convey("TestMapIntRune.GetOpt", t, func() {
		var k int = 392129278
		var v rune = 1290935656

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1543945639 + 252207744)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntRune_ForEach(t *testing.T) {
	Convey("TestMapIntRune.ForEach", t, func() {
		var k int = 2045881338
		var v rune = 16021497
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
		var k int = 79054550
		var v rune = 2072282596

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
		var k int = 1188756185
		var v rune = 457714541

		test := omap.NewMapIntRune(1)

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

func TestMapIntRune_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntRune.PutIfNotNil", t, func() {
		var k int = 1210827967
		var v rune = 1827082878

		test := omap.NewMapIntRune(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1986314726, (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x rune = 1225689614
		So(test.PutIfNotNil(671656956, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntRune_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntRune.ReplaceIfExists", t, func() {
		var k int = 1408072371
		var v rune = 915188223
		var x rune = 921295076

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1246280523, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntRune_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntRune.ReplaceOrPut", t, func() {
		var k int = 1878630615
		var v rune = 1914549241
		var x rune = 1502051531

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1449613410, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntRune_MarshalJSON(t *testing.T) {
	Convey("TestMapIntRune.MarshalJSON", t, func() {
		var k int = 1798770900
		var v rune = 746890536

		test := omap.NewMapIntRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1798770900,"value":746890536}]`)
	})
}
