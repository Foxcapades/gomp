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
		var k int = 930956904
		var v uint = 1647541763

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntUint_Delete(t *testing.T) {
	Convey("TestMapIntUint.Delete", t, func() {
		var k int = 1663930550
		var v uint = 916717484

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntUint_Has(t *testing.T) {
	Convey("TestMapIntUint.Has", t, func() {
		var k int = 149653793
		var v uint = 565950260

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(782630291+340838023), ShouldBeFalse)
	})
}

func TestMapIntUint_Get(t *testing.T) {
	Convey("TestMapIntUint.Get", t, func() {
		var k int = 1544549245
		var v uint = 3340226745

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(2106122857 + 1755776901)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntUint_GetOpt(t *testing.T) {
	Convey("TestMapIntUint.GetOpt", t, func() {
		var k int = 328610713
		var v uint = 223659681

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(431065225 + 68926317)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntUint_ForEach(t *testing.T) {
	Convey("TestMapIntUint.ForEach", t, func() {
		var k int = 621620808
		var v uint = 1023752601
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
		var k int = 1536777058
		var v uint = 1196140758

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
		Convey("Ordered", func() {
			var k int = 475607135
			var v uint = 1910457360

			test := omap.NewMapIntUint(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			c, d := test.ToYAML()
			So(d, ShouldBeNil)
			So(c.Kind, ShouldEqual, yaml.SequenceNode)
			So(c.LongTag(), ShouldEqual, xyml.TagOrderedMap)
			So(len(c.Content), ShouldEqual, 1)
			So(xyml.IsMap(c.Content[0]), ShouldBeTrue)
		})

		Convey("Unordered", func() {
			var k int = 1810829071
			var v uint = 2650102610

			test := omap.NewMapIntUint(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			c, d := test.ToYAML()

			So(d, ShouldBeNil)
			So(c.Kind, ShouldEqual, yaml.MappingNode)
			So(c.LongTag(), ShouldEqual, xyml.TagMap)
			So(len(c.Content), ShouldEqual, 2)
		})
	})
}

func TestMapIntUint_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntUint.PutIfNotNil", t, func() {
		var k int = 1763388937
		var v uint = 3486467599

		test := omap.NewMapIntUint(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(605939597, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint = 4279551363
		So(test.PutIfNotNil(1731169632, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntUint_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntUint.ReplaceIfExists", t, func() {
		var k int = 1580714873
		var v uint = 490938822
		var x uint = 4224111294

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1368987300, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntUint.ReplaceOrPut", t, func() {
		var k int = 204192471
		var v uint = 2859279086
		var x uint = 2928712884

		test := omap.NewMapIntUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(469475928, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint_MarshalJSON(t *testing.T) {
	Convey("TestMapIntUint.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k int = 1465625947
			var v uint = 317313

			test := omap.NewMapIntUint(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":1465625947,"value":317313}]`)
		})

		Convey("Unordered", func() {
			var k int = 1465625947
			var v uint = 317313

			test := omap.NewMapIntUint(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"1465625947":317313}`)
		})

	})
}
