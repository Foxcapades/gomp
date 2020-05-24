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
		var k uint = 3763222118
		var v float64 = 0.313

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintFloat64_Delete(t *testing.T) {
	Convey("TestMapUintFloat64.Delete", t, func() {
		var k uint = 2845774518
		var v float64 = 0.126

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintFloat64_Has(t *testing.T) {
	Convey("TestMapUintFloat64.Has", t, func() {
		var k uint = 3181327064
		var v float64 = 0.624

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(814026972+199367856), ShouldBeFalse)
	})
}

func TestMapUintFloat64_Get(t *testing.T) {
	Convey("TestMapUintFloat64.Get", t, func() {
		var k uint = 3173737859
		var v float64 = 0.772

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(2043507941 + 330629406)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintFloat64_GetOpt(t *testing.T) {
	Convey("TestMapUintFloat64.GetOpt", t, func() {
		var k uint = 3086704277
		var v float64 = 0.870

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(268226988 + 882364038)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintFloat64_ForEach(t *testing.T) {
	Convey("TestMapUintFloat64.ForEach", t, func() {
		var k uint = 1078848317
		var v float64 = 0.566
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
		var k uint = 3197811395
		var v float64 = 0.841

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
		var k uint = 2907256667
		var v float64 = 0.123

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintFloat64_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintFloat64.PutIfNotNil", t, func() {
		var k uint = 1475535589
		var v float64 = 0.415

		test := omap.NewMapUintFloat64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(2593585747, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float64 = 0.403
		So(test.PutIfNotNil(1632978391, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintFloat64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintFloat64.ReplaceIfExists", t, func() {
		var k uint = 697221335
		var v float64 = 0.861
		var x float64 = 0.090

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(2561022206, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintFloat64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintFloat64.ReplaceOrPut", t, func() {
		var k uint = 3280269601
		var v float64 = 0.138
		var x float64 = 0.193

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(3485524978, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintFloat64_MarshalJSON(t *testing.T) {
	Convey("TestMapUintFloat64.MarshalJSON", t, func() {
		var k uint = 366404833
		var v float64 = 0.555

		test := omap.NewMapUintFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":366404833,"value":0.555}]`)
	})
}
