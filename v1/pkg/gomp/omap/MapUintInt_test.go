package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintInt_Put(t *testing.T) {
	Convey("TestMapUintInt.Put", t, func() {
		var k uint = 3248818075
		var v int = 1318286908

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintInt_Delete(t *testing.T) {
	Convey("TestMapUintInt.Delete", t, func() {
		var k uint = 4084095277
		var v int = 977207923

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintInt_Has(t *testing.T) {
	Convey("TestMapUintInt.Has", t, func() {
		var k uint = 1261188576
		var v int = 1405240990

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(2073694353+1356136993), ShouldBeFalse)
	})
}

func TestMapUintInt_Get(t *testing.T) {
	Convey("TestMapUintInt.Get", t, func() {
		var k uint = 1946488083
		var v int = 70349737

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(1378209325 + 1329822314)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintInt_GetOpt(t *testing.T) {
	Convey("TestMapUintInt.GetOpt", t, func() {
		var k uint = 147630073
		var v int = 544804047

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(849596670 + 4284715951)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintInt_ForEach(t *testing.T) {
	Convey("TestMapUintInt.ForEach", t, func() {
		var k uint = 1010545967
		var v int = 1076267976
		hits := 0

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv int) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintInt_MarshalYAML(t *testing.T) {
	Convey("TestMapUintInt.MarshalYAML", t, func() {
		var k uint = 971308278
		var v int = 1082917969

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintInt_ToYAML(t *testing.T) {
	Convey("TestMapUintInt.ToYAML", t, func() {
		var k uint = 2132619147
		var v int = 1478715757

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintInt_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintInt.PutIfNotNil", t, func() {
		var k uint = 4176992843
		var v int = 433341168

		test := omap.NewMapUintInt(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(2239300238, (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int = 1939013664
		So(test.PutIfNotNil(2642179852, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintInt_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintInt.ReplaceIfExists", t, func() {
		var k uint = 1317765241
		var v int = 579168723
		var x int = 1791838055

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(388129252, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintInt.ReplaceOrPut", t, func() {
		var k uint = 936174979
		var v int = 2058729380
		var x int = 99266701

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(3782831504, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt_MarshalJSON(t *testing.T) {
	Convey("TestMapUintInt.MarshalJSON", t, func() {
		var k uint = 1137415377
		var v int = 271910662

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1137415377,"value":271910662}]`)
	})
}
