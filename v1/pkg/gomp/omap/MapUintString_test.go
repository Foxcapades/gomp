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
		var k uint = 1854934854
		var v string = "a0e4a442-d44f-435b-bffe-81bca42584ff"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintString_Delete(t *testing.T) {
	Convey("TestMapUintString.Delete", t, func() {
		var k uint = 1791462742
		var v string = "48186b2f-2901-4c3a-b927-63d2d75a8f07"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintString_Has(t *testing.T) {
	Convey("TestMapUintString.Has", t, func() {
		var k uint = 350677186
		var v string = "e656d649-5c0d-4f78-a644-1ed30689d6cd"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(830566946+1188440254), ShouldBeFalse)
	})
}

func TestMapUintString_Get(t *testing.T) {
	Convey("TestMapUintString.Get", t, func() {
		var k uint = 1563139179
		var v string = "b69f2716-3074-4c96-ac25-90ffa8513450"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(1497610564 + 3688507506)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintString_GetOpt(t *testing.T) {
	Convey("TestMapUintString.GetOpt", t, func() {
		var k uint = 3371702041
		var v string = "89e3840d-1970-4a44-9af6-40ff2aba8db4"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(596321497 + 560824660)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintString_ForEach(t *testing.T) {
	Convey("TestMapUintString.ForEach", t, func() {
		var k uint = 1287829580
		var v string = "07ef3fd1-cf4e-47b2-98b7-227fa4e2735a"
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
		var k uint = 1524675227
		var v string = "2c3319ef-4dbf-4968-9d7e-c105bcf12d1a"

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
		var k uint = 1357245802
		var v string = "226f3ee8-4f9a-4f53-a59d-7f7c1918fec9"

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
		var k uint = 649407013
		var v string = "6a68afca-b122-47c9-8190-15c28cc83dbc"

		test := omap.NewMapUintString(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1498027288, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "10b40cbd-c6dd-40a0-baa6-7f0807c18e90"
		So(test.PutIfNotNil(3263488426, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintString_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintString.ReplaceIfExists", t, func() {
		var k uint = 1425498997
		var v string = "0dd2dd57-0c80-40d0-a809-249aa5193690"
		var x string = "c4b4e1dd-11c5-4ac2-9321-994da09e17f3"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(647015573, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintString_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintString.ReplaceOrPut", t, func() {
		var k uint = 3592793358
		var v string = "1049a256-8faf-4156-9c35-face5494b533"
		var x string = "ea79aedc-fdfc-414e-9099-015458db5c1d"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(2360359677, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintString_MarshalJSON(t *testing.T) {
	Convey("TestMapUintString.MarshalJSON", t, func() {
		var k uint = 2992954310
		var v string = "54786027-81eb-48ae-9db8-5693e21b0aca"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":2992954310,"value":"54786027-81eb-48ae-9db8-5693e21b0aca"}]`)
	})
}
