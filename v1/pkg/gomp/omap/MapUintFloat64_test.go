package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintFloat64_Put(t *testing.T) {
	Convey("TestMapUintFloat64.Put", t, func() {
		var k uint = 2323341487
		var v float64 = 0.238

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintFloat64_Delete(t *testing.T) {
	Convey("TestMapUintFloat64.Delete", t, func() {
		var k uint = 2859055562
		var v float64 = 0.458

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintFloat64_Has(t *testing.T) {
	Convey("TestMapUintFloat64.Has", t, func() {
		var k uint = 2759729940
		var v float64 = 0.786

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1479370707+298265664), ShouldBeFalse)
	})
}

func TestMapUintFloat64_Get(t *testing.T) {
	Convey("TestMapUintFloat64.Get", t, func() {
		var k uint = 2866262573
		var v float64 = 0.400

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(2519684331 + 1032865909)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintFloat64_GetOpt(t *testing.T) {
	Convey("TestMapUintFloat64.GetOpt", t, func() {
		var k uint = 1737505599
		var v float64 = 0.538

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(2764885414 + 3536745725)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintFloat64_ForEach(t *testing.T) {
	Convey("TestMapUintFloat64.ForEach", t, func() {
		var k uint = 3927750508
		var v float64 = 0.292
		hits := 0

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv float64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintFloat64_MarshalYAML(t *testing.T) {
	Convey("TestMapUintFloat64.MarshalYAML", t, func() {
		var k uint = 435158843
		var v float64 = 0.891

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintFloat64_ToYAML(t *testing.T) {
	Convey("TestMapUintFloat64.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k uint = 2477525175
			var v float64 = 0.933

			test := omap.NewMapUintFloat64(1)

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
			var k uint = 3633339639
			var v float64 = 0.706

			test := omap.NewMapUintFloat64(1)
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

func TestMapUintFloat64_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintFloat64.PutIfNotNil", t, func() {
		var k uint = 3123843816
		var v float64 = 0.516

		test := omap.NewMapUintFloat64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(3024444362, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float64 = 0.320
		So(test.PutIfNotNil(231162788, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintFloat64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintFloat64.ReplaceIfExists", t, func() {
		var k uint = 1353456300
		var v float64 = 0.459
		var x float64 = 0.217

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(3028489571, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintFloat64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintFloat64.ReplaceOrPut", t, func() {
		var k uint = 3469293234
		var v float64 = 0.211
		var x float64 = 0.901

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(404991487, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintFloat64_MarshalJSON(t *testing.T) {
	Convey("TestMapUintFloat64.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k uint = 2432302163
			var v float64 = 0.507

			test := omap.NewMapUintFloat64(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":2432302163,"value":0.507}]`)
		})

		Convey("Unordered", func() {
			var k uint = 2432302163
			var v float64 = 0.507

			test := omap.NewMapUintFloat64(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"2432302163":0.507}`)
		})

	})
}
