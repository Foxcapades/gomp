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
		var k uint = 1609239643
		var v string = "9908b5bc-dcf9-42d0-831e-1784fed6801c"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintString_Delete(t *testing.T) {
	Convey("TestMapUintString.Delete", t, func() {
		var k uint = 3038956583
		var v string = "63710f3a-9028-46c2-96e9-ba11d0e3adf1"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintString_Has(t *testing.T) {
	Convey("TestMapUintString.Has", t, func() {
		var k uint = 1983753827
		var v string = "5db2b3f8-f084-4dd2-8b05-b02dae45c643"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(4201412355+1444303754), ShouldBeFalse)
	})
}

func TestMapUintString_Get(t *testing.T) {
	Convey("TestMapUintString.Get", t, func() {
		var k uint = 2556884762
		var v string = "58f0f373-427d-4d74-84c3-4fc83daf4727"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(1730755124 + 1095977028)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintString_GetOpt(t *testing.T) {
	Convey("TestMapUintString.GetOpt", t, func() {
		var k uint = 712873347
		var v string = "d7956a96-4672-4fb7-a7d0-a0056e2f7ada"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1570047713 + 2670037207)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintString_ForEach(t *testing.T) {
	Convey("TestMapUintString.ForEach", t, func() {
		var k uint = 3650554542
		var v string = "eaab958f-e70d-4920-a22f-527b4aed6417"
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
		var k uint = 3011343461
		var v string = "0dba4983-8d22-496f-8b5a-9b26691f3f27"

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
		var k uint = 2056416736
		var v string = "7f8ebf87-3d24-4099-9437-111bcf9a7ba8"

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
		var k uint = 2067078857
		var v string = "4532617f-349e-448c-aebc-60c650addf1f"

		test := omap.NewMapUintString(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(2151311592, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "f6f21e65-9171-426b-8713-dc4451f9954d"
		So(test.PutIfNotNil(4057893365, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintString_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintString.ReplaceIfExists", t, func() {
		var k uint = 3860517454
		var v string = "0e988990-d7c1-4ca5-9b04-23557e606b31"
		var x string = "84bf8941-0335-44fc-a0a2-ebfd99509265"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1826958951, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintString_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintString.ReplaceOrPut", t, func() {
		var k uint = 249676545
		var v string = "e6675783-8e50-43d8-a977-84fc546543c9"
		var x string = "0bd8a7dc-aa33-4c50-bdeb-213dc49b1b87"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(2592420637, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintString_MarshalJSON(t *testing.T) {
	Convey("TestMapUintString.MarshalJSON", t, func() {
		var k uint = 3656614618
		var v string = "9427cb7b-df0c-4d32-9e06-c962f2e751b7"

		test := omap.NewMapUintString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":3656614618,"value":"9427cb7b-df0c-4d32-9e06-c962f2e751b7"}]`)
	})
}
