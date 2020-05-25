package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintRune_Put(t *testing.T) {
	Convey("TestMapUintRune.Put", t, func() {
		var k uint = 1695110748
		var v rune = 1521637524

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintRune_Delete(t *testing.T) {
	Convey("TestMapUintRune.Delete", t, func() {
		var k uint = 1450494209
		var v rune = 16011128

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintRune_Has(t *testing.T) {
	Convey("TestMapUintRune.Has", t, func() {
		var k uint = 1295458503
		var v rune = 465983648

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(2567198735+659194763), ShouldBeFalse)
	})
}

func TestMapUintRune_Get(t *testing.T) {
	Convey("TestMapUintRune.Get", t, func() {
		var k uint = 1077341569
		var v rune = 1892284059

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(1945109427 + 931146386)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintRune_GetOpt(t *testing.T) {
	Convey("TestMapUintRune.GetOpt", t, func() {
		var k uint = 3596990169
		var v rune = 1576681560

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(3258287079 + 1492961151)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintRune_ForEach(t *testing.T) {
	Convey("TestMapUintRune.ForEach", t, func() {
		var k uint = 984384265
		var v rune = 451519044
		hits := 0

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv rune) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintRune_MarshalYAML(t *testing.T) {
	Convey("TestMapUintRune.MarshalYAML", t, func() {
		var k uint = 1959180485
		var v rune = 479365953

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintRune_ToYAML(t *testing.T) {
	Convey("TestMapUintRune.ToYAML", t, func() {
		var k uint = 2994217473
		var v rune = 5692739

		test := omap.NewMapUintRune(1)

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

func TestMapUintRune_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintRune.PutIfNotNil", t, func() {
		var k uint = 2398580335
		var v rune = 1560787546

		test := omap.NewMapUintRune(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(2118991998, (*rune)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x rune = 1531977981
		So(test.PutIfNotNil(3768894561, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintRune_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintRune.ReplaceIfExists", t, func() {
		var k uint = 1997595459
		var v rune = 401490240
		var x rune = 220824240

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(2354190044, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintRune_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintRune.ReplaceOrPut", t, func() {
		var k uint = 298011763
		var v rune = 784094500
		var x rune = 1074106136

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(3515415647, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintRune_MarshalJSON(t *testing.T) {
	Convey("TestMapUintRune.MarshalJSON", t, func() {
		var k uint = 2679471825
		var v rune = 1608040931

		test := omap.NewMapUintRune(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":2679471825,"value":1608040931}]`)
	})
}
