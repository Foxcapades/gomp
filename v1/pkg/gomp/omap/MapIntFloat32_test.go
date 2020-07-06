package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntFloat32_Put(t *testing.T) {
	Convey("TestMapIntFloat32.Put", t, func() {
		var k int = 2125005779
		var v float32 = 0.978

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntFloat32_Delete(t *testing.T) {
	Convey("TestMapIntFloat32.Delete", t, func() {
		var k int = 340073423
		var v float32 = 0.501

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntFloat32_Has(t *testing.T) {
	Convey("TestMapIntFloat32.Has", t, func() {
		var k int = 1069766276
		var v float32 = 0.425

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1798550724+601174171), ShouldBeFalse)
	})
}

func TestMapIntFloat32_Get(t *testing.T) {
	Convey("TestMapIntFloat32.Get", t, func() {
		var k int = 116278600
		var v float32 = 0.200

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(967072994 + 110219588)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntFloat32_GetOpt(t *testing.T) {
	Convey("TestMapIntFloat32.GetOpt", t, func() {
		var k int = 105506518
		var v float32 = 0.330

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1165044941 + 415189566)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntFloat32_ForEach(t *testing.T) {
	Convey("TestMapIntFloat32.ForEach", t, func() {
		var k int = 24709330
		var v float32 = 0.209
		hits := 0

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv float32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntFloat32_MarshalYAML(t *testing.T) {
	Convey("TestMapIntFloat32.MarshalYAML", t, func() {
		var k int = 1624658777
		var v float32 = 0.212

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntFloat32_ToYAML(t *testing.T) {
	Convey("TestMapIntFloat32.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k int = 1148897637
			var v float32 = 0.696

			test := omap.NewMapIntFloat32(1)

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
			var k int = 422710357
			var v float32 = 0.391

			test := omap.NewMapIntFloat32(1)
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

func TestMapIntFloat32_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntFloat32.PutIfNotNil", t, func() {
		var k int = 1584579463
		var v float32 = 0.222

		test := omap.NewMapIntFloat32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1722213713, (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float32 = 0.425
		So(test.PutIfNotNil(21975443, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntFloat32.ReplaceIfExists", t, func() {
		var k int = 117254569
		var v float32 = 0.821
		var x float32 = 0.853

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1129748768, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntFloat32.ReplaceOrPut", t, func() {
		var k int = 202399855
		var v float32 = 0.008
		var x float32 = 0.379

		test := omap.NewMapIntFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1660377840, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat32_MarshalJSON(t *testing.T) {
	Convey("TestMapIntFloat32.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k int = 1643002107
			var v float32 = 0.070

			test := omap.NewMapIntFloat32(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":1643002107,"value":0.07}]`)
		})

		Convey("Unordered", func() {
			var k int = 1643002107
			var v float32 = 0.070

			test := omap.NewMapIntFloat32(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"1643002107":0.07}`)
		})

	})
}
