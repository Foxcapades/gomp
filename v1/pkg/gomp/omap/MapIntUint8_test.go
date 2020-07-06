package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntUint8_Put(t *testing.T) {
	Convey("TestMapIntUint8.Put", t, func() {
		var k int = 520235989
		var v uint8 = 216

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntUint8_Delete(t *testing.T) {
	Convey("TestMapIntUint8.Delete", t, func() {
		var k int = 1235939988
		var v uint8 = 100

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntUint8_Has(t *testing.T) {
	Convey("TestMapIntUint8.Has", t, func() {
		var k int = 183394750
		var v uint8 = 37

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1386875197+92759709), ShouldBeFalse)
	})
}

func TestMapIntUint8_Get(t *testing.T) {
	Convey("TestMapIntUint8.Get", t, func() {
		var k int = 937106504
		var v uint8 = 237

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(2146647952 + 1548663687)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntUint8_GetOpt(t *testing.T) {
	Convey("TestMapIntUint8.GetOpt", t, func() {
		var k int = 1311646820
		var v uint8 = 221

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(914925856 + 529174130)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntUint8_ForEach(t *testing.T) {
	Convey("TestMapIntUint8.ForEach", t, func() {
		var k int = 563221157
		var v uint8 = 210
		hits := 0

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv uint8) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntUint8_MarshalYAML(t *testing.T) {
	Convey("TestMapIntUint8.MarshalYAML", t, func() {
		var k int = 2033265206
		var v uint8 = 141

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntUint8_ToYAML(t *testing.T) {
	Convey("TestMapIntUint8.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k int = 1255448012
			var v uint8 = 132

			test := omap.NewMapIntUint8(1)

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
			var k int = 182557919
			var v uint8 = 42

			test := omap.NewMapIntUint8(1)
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

func TestMapIntUint8_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntUint8.PutIfNotNil", t, func() {
		var k int = 154353390
		var v uint8 = 162

		test := omap.NewMapIntUint8(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(65287999, (*uint8)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint8 = 122
		So(test.PutIfNotNil(2035499104, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntUint8_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntUint8.ReplaceIfExists", t, func() {
		var k int = 1372112658
		var v uint8 = 27
		var x uint8 = 82

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1037830202, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint8_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntUint8.ReplaceOrPut", t, func() {
		var k int = 1236331679
		var v uint8 = 221
		var x uint8 = 186

		test := omap.NewMapIntUint8(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(905941176, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntUint8_MarshalJSON(t *testing.T) {
	Convey("TestMapIntUint8.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k int = 1090926295
			var v uint8 = 134

			test := omap.NewMapIntUint8(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":1090926295,"value":134}]`)
		})

		Convey("Unordered", func() {
			var k int = 1090926295
			var v uint8 = 134

			test := omap.NewMapIntUint8(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"1090926295":134}`)
		})

	})
}
