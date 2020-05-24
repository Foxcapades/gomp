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
		var k uint = 3458269577
		var v int = 1438965648

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintInt_Delete(t *testing.T) {
	Convey("TestMapUintInt.Delete", t, func() {
		var k uint = 719556398
		var v int = 326697044

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintInt_Has(t *testing.T) {
	Convey("TestMapUintInt.Has", t, func() {
		var k uint = 1864665455
		var v int = 2126622836

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(1871311408+896630589), ShouldBeFalse)
	})
}


func TestMapUintInt_Get(t *testing.T) {
	Convey("TestMapUintInt.Get", t, func() {
		var k uint = 3565681569
		var v int = 2114892975

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1268050383 + 2251956968)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintInt_GetOpt(t *testing.T) {
	Convey("TestMapUintInt.GetOpt", t, func() {
		var k uint = 439056639
		var v int = 1977741920

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(2522188711 + 2024225111)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintInt_ForEach(t *testing.T) {
	Convey("TestMapUintInt.ForEach", t, func() {
		var k uint = 194001260
		var v int = 737599383
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
		var k uint = 1196970658
		var v int = 1123500461

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
		var k uint = 637770918
		var v int = 271226049

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
		var k uint = 819653411
		var v int = 573365776

		test := omap.NewMapUintInt(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1478436174, (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int = 320956971
		So(test.PutIfNotNil(2436522464, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintInt_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintInt.ReplaceIfExists", t, func() {
		var k uint = 807793777
		var v int = 1025774125
		var x int = 1294531588

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(2309506758, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintInt.ReplaceOrPut", t, func() {
		var k uint = 1378322890
		var v int = 1035485620
		var x int = 1108425899

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(2003527793, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt_MarshalJSON(t *testing.T) {
	Convey("TestMapUintInt.MarshalJSON", t, func() {
		var k uint = 2154412544
		var v int = 2106578226

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":2154412544,"value":2106578226}]`)
	})
}
