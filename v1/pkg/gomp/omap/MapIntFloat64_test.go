package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntFloat64_Put(t *testing.T) {
	Convey("TestMapIntFloat64.Put", t, func() {
		var k int = 1659842122
		var v float64 = 0.041

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntFloat64_Delete(t *testing.T) {
	Convey("TestMapIntFloat64.Delete", t, func() {
		var k int = 164202907
		var v float64 = 0.187

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntFloat64_Has(t *testing.T) {
	Convey("TestMapIntFloat64.Has", t, func() {
		var k int = 1436238993
		var v float64 = 0.596

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1900446977+972700989), ShouldBeFalse)
	})
}

func TestMapIntFloat64_Get(t *testing.T) {
	Convey("TestMapIntFloat64.Get", t, func() {
		var k int = 1058938835
		var v float64 = 0.650

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(1773333429 + 166507222)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntFloat64_GetOpt(t *testing.T) {
	Convey("TestMapIntFloat64.GetOpt", t, func() {
		var k int = 643946416
		var v float64 = 0.539

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1409157945 + 240498055)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntFloat64_ForEach(t *testing.T) {
	Convey("TestMapIntFloat64.ForEach", t, func() {
		var k int = 632236992
		var v float64 = 0.864
		hits := 0

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv float64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntFloat64_MarshalYAML(t *testing.T) {
	Convey("TestMapIntFloat64.MarshalYAML", t, func() {
		var k int = 1689376112
		var v float64 = 0.271

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntFloat64_ToYAML(t *testing.T) {
	Convey("TestMapIntFloat64.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k int = 2046463278
			var v float64 = 0.103

			test := omap.NewMapIntFloat64(1)

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
			var k int = 1408332047
			var v float64 = 0.136

			test := omap.NewMapIntFloat64(1)
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

func TestMapIntFloat64_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntFloat64.PutIfNotNil", t, func() {
		var k int = 1499176672
		var v float64 = 0.651

		test := omap.NewMapIntFloat64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1180920487, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float64 = 0.481
		So(test.PutIfNotNil(737486622, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntFloat64.ReplaceIfExists", t, func() {
		var k int = 577440171
		var v float64 = 0.322
		var x float64 = 0.882

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(661448516, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntFloat64.ReplaceOrPut", t, func() {
		var k int = 1648678581
		var v float64 = 0.691
		var x float64 = 0.194

		test := omap.NewMapIntFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(501304628, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntFloat64_MarshalJSON(t *testing.T) {
	Convey("TestMapIntFloat64.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k int = 1567142070
			var v float64 = 0.515

			test := omap.NewMapIntFloat64(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":1567142070,"value":0.515}]`)
		})

		Convey("Unordered", func() {
			var k int = 1567142070
			var v float64 = 0.515

			test := omap.NewMapIntFloat64(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"1567142070":0.515}`)
		})

	})
}
