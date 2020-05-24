package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintString_Put(t *testing.T) {
	Convey("TestMapUintString.Put", t, func() {
		var k uint = 4056414347
		var v string = "b6c24600-1267-4ee8-a3a0-779585832223"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintString_Delete(t *testing.T) {
	Convey("TestMapUintString.Delete", t, func() {
		var k uint = 1783602139
		var v string = "69c32bcf-b2f9-4d5d-b88f-daa50611325a"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintString_Has(t *testing.T) {
	Convey("TestMapUintString.Has", t, func() {
		var k uint = 2577732997
		var v string = "dc3060e8-3ba2-4e29-855b-2b60aa1be0e9"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(3000662617+1514839095), ShouldBeFalse)
	})
}


func TestMapUintString_Get(t *testing.T) {
	Convey("TestMapUintString.Get", t, func() {
		var k uint = 1724234298
		var v string = "a598b4c6-4e9a-478f-b535-672d8cc9a11d"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(461681877 + 1471064723)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintString_GetOpt(t *testing.T) {
	Convey("TestMapUintString.GetOpt", t, func() {
		var k uint = 3567788637
		var v string = "2ac44c1f-b7ba-4314-bfe1-4ad65f652a34"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(4008398967 + 468310645)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintString_ForEach(t *testing.T) {
	Convey("TestMapUintString.ForEach", t, func() {
		var k uint = 429279060
		var v string = "a1ca647f-bd2b-4ce7-abd8-39b30f872d30"
		hits := 0

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv string) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintString_MarshalYAML(t *testing.T) {
	Convey("TestMapUintString.MarshalYAML", t, func() {
		var k uint = 1254099318
		var v string = "1f0b9191-a7e5-4845-b9a8-2b6d1ad14b8f"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintString_ToYAML(t *testing.T) {
	Convey("TestMapUintString.ToYAML", t, func() {
		var k uint = 435049921
		var v string = "cd5c8f3a-9246-4a3f-8379-41c07c99a992"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintString_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintString.PutIfNotNil", t, func() {
		var k uint = 4249054960
		var v string = "1fac1f12-3b84-4438-807d-931d7f11e429"

		test := omap.NewMapUintString(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(384013914, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "213aba5a-c7be-4be8-bf77-d1fa4c1c869d"
		So(test.PutIfNotNil(1402201836, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintString_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintString.ReplaceIfExists", t, func() {
		var k uint = 3099265160
		var v string = "493568a8-fce9-487f-a629-3028697f5043"
		var x string = "7a478463-68d8-4dee-9e9c-03295b08631a"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1372383518, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintString_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintString.ReplaceOrPut", t, func() {
		var k uint = 2668771057
		var v string = "d91fef40-f9dc-4780-8df1-1cb6b2be71af"
		var x string = "978f3aeb-9f32-4a10-887b-d6dc731c3dcf"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(404151916, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintString_MarshalJSON(t *testing.T) {
	Convey("TestMapUintString.MarshalJSON", t, func() {
		var k uint = 2490115165
		var v string = "d859bc89-6a3f-42c7-835b-308ba0356d79"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":2490115165,"value":"d859bc89-6a3f-42c7-835b-308ba0356d79"}]`)
	})
}
