package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntInt8_Put(t *testing.T) {
	Convey("TestMapIntInt8.Put", t, func() {
		var k int = 127907959
		var v int8 = 68

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntInt8_Delete(t *testing.T) {
	Convey("TestMapIntInt8.Delete", t, func() {
		var k int = 1182000149
		var v int8 = 4

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntInt8_Has(t *testing.T) {
	Convey("TestMapIntInt8.Has", t, func() {
		var k int = 59081892
		var v int8 = 125

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(989722213+1808733745), ShouldBeFalse)
	})
}

func TestMapIntInt8_Get(t *testing.T) {
	Convey("TestMapIntInt8.Get", t, func() {
		var k int = 1287496631
		var v int8 = 78

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(696481174 + 2066681235)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntInt8_GetOpt(t *testing.T) {
	Convey("TestMapIntInt8.GetOpt", t, func() {
		var k int = 1739880030
		var v int8 = 23

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(710842142 + 1752325169)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntInt8_ForEach(t *testing.T) {
	Convey("TestMapIntInt8.ForEach", t, func() {
		var k int = 841867331
		var v int8 = 4
		hits := 0

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv int8) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntInt8_MarshalYAML(t *testing.T) {
	Convey("TestMapIntInt8.MarshalYAML", t, func() {
		var k int = 297036478
		var v int8 = 92

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntInt8_ToYAML(t *testing.T) {
	Convey("TestMapIntInt8.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k int = 1430266816
			var v int8 = 13

			test := omap.NewMapIntInt8(1)

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
			var k int = 1466190448
			var v int8 = 52

			test := omap.NewMapIntInt8(1)
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

func TestMapIntInt8_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntInt8.PutIfNotNil", t, func() {
		var k int = 1665343564
		var v int8 = 52

		test := omap.NewMapIntInt8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1683871896, (*int8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int8 = 98
		So(test.PutIfNotNil(1274015086, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntInt8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntInt8.ReplaceIfExists", t, func() {
		var k int = 1162140558
		var v int8 = 65
		var x int8 = 58

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(879667099, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntInt8.ReplaceOrPut", t, func() {
		var k int = 70075112
		var v int8 = 58
		var x int8 = 2

		test := omap.NewMapIntInt8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1969297309, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt8_MarshalJSON(t *testing.T) {
	Convey("TestMapIntInt8.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k int = 198025260
			var v int8 = 64

			test := omap.NewMapIntInt8(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":198025260,"value":64}]`)
		})

		Convey("Unordered", func() {
			var k int = 198025260
			var v int8 = 64

			test := omap.NewMapIntInt8(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"198025260":64}`)
		})

	})
}
