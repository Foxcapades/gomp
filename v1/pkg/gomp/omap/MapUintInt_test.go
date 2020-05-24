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
		var k uint = 1597559646
		var v int = 1932312884

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintInt_Delete(t *testing.T) {
	Convey("TestMapUintInt.Delete", t, func() {
		var k uint = 4143818140
		var v int = 501826979

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintInt_Has(t *testing.T) {
	Convey("TestMapUintInt.Has", t, func() {
		var k uint = 2226844948
		var v int = 128559010

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(3633368814+3183560535), ShouldBeFalse)
	})
}


func TestMapUintInt_Get(t *testing.T) {
	Convey("TestMapUintInt.Get", t, func() {
		var k uint = 433278482
		var v int = 970454683

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(3919117460+2127950065)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintInt_GetOpt(t *testing.T) {
	Convey("TestMapUintInt.GetOpt", t, func() {
		var k uint = 1298335765
		var v int = 655202632

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(2168934648+3440355213)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintInt_ForEach(t *testing.T) {
	Convey("TestMapUintInt.ForEach", t, func() {
		var k uint = 3056833538
		var v int = 1364809685
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
		var k uint = 3086428203
		var v int = 1003425354

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
		var k uint = 610933924
		var v int = 579001712

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
		var k uint = 3739657307
		var v int = 1583946823

		test := omap.NewMapUintInt(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(2268896797, (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int = 2144564119
		So(test.PutIfNotNil(2234739118, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintInt_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintInt.ReplaceIfExists", t, func() {
		var k uint = 713236441
		var v int = 872697163
		var x int = 1477999599

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(4229268277, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintInt.ReplaceOrPut", t, func() {
		var k uint = 1375972169
		var v int = 583261163
		var x int = 1536185790

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(2981517957, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintInt_MarshalJSON(t *testing.T) {
	Convey("TestMapUintInt.MarshalJSON", t, func() {
		var k uint = 2397626945
		var v int = 431026902

		test := omap.NewMapUintInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":2397626945,"value":431026902}]`)
	})
}

