package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringUint_Put(t *testing.T) {
	Convey("TestMapStringUint.Put", t, func() {
		var k string = "c15e9d8e-a414-4e36-8bf4-2b2d369b1d9a"
		var v uint = 1434471365

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringUint_Delete(t *testing.T) {
	Convey("TestMapStringUint.Delete", t, func() {
		var k string = "f3e6db20-72f9-4351-a513-98e16ef94f9a"
		var v uint = 1519087335

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringUint_Has(t *testing.T) {
	Convey("TestMapStringUint.Has", t, func() {
		var k string = "3bcdcb2c-eff7-4961-9d9d-0f80b6346a69"
		var v uint = 2111118454

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("0382854b-d3fa-4a20-a66b-ddbea5da8a7e"+"4e6e24a6-9bec-4ad1-ab80-dcada8dca68e"), ShouldBeFalse)
	})
}

func TestMapStringUint_Get(t *testing.T) {
	Convey("TestMapStringUint.Get", t, func() {
		var k string = "a6e05e7b-4dc9-4e64-850b-65fc06fb47b3"
		var v uint = 3468501533

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("7674d40c-6bf9-4be1-b2ee-b925db8762e3" + "b3f093ba-f66c-445c-9404-9f688ace8b23")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringUint_GetOpt(t *testing.T) {
	Convey("TestMapStringUint.GetOpt", t, func() {
		var k string = "9d2d0ddc-2fcc-414f-bd73-014c7c716133"
		var v uint = 8607859

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("8585c351-c659-43bc-ad27-b33f7be6eee5" + "af4fce83-c1ec-4dc2-93d4-06a4c8a89a04")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringUint_ForEach(t *testing.T) {
	Convey("TestMapStringUint.ForEach", t, func() {
		var k string = "98b55f3c-010b-44f9-a9b0-07ed9f1bc21a"
		var v uint = 1211140054
		hits := 0

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv uint) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringUint_MarshalYAML(t *testing.T) {
	Convey("TestMapStringUint.MarshalYAML", t, func() {
		var k string = "7dd12bc7-351f-403c-9f00-19cdb04eb1e0"
		var v uint = 2393387145

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringUint_ToYAML(t *testing.T) {
	Convey("TestMapStringUint.ToYAML", t, func() {
		var k string = "5095bac0-b233-48ad-89fc-b9f617745641"
		var v uint = 3958186716

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringUint_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringUint.PutIfNotNil", t, func() {
		var k string = "a5f4e64e-fc37-43f1-815b-2af2a5fb1b71"
		var v uint = 2243623855

		test := omap.NewMapStringUint(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("c8434621-16ef-43d7-bc76-91e457bf830c", (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint = 3780618646
		So(test.PutIfNotNil("c5c039fa-d81a-411e-8e7f-3090f27c5475", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringUint_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringUint.ReplaceIfExists", t, func() {
		var k string = "5c5e28e9-3bf8-4599-b4c0-3c1485b73d3f"
		var v uint = 1360005017
		var x uint = 720360342

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("886fe888-e521-4f96-9510-1f774e2f8565", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringUint.ReplaceOrPut", t, func() {
		var k string = "917a1f31-ea55-4a9e-85c4-a7665ffa7834"
		var v uint = 3142735583
		var x uint = 1668005435

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("4fe35ec9-d25b-4c04-81e2-564523cbedc0", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint_MarshalJSON(t *testing.T) {
	Convey("TestMapStringUint.MarshalJSON", t, func() {
		var k string = "7db9cb60-b788-495e-8079-66b57b9e207e"
		var v uint = 1176757059

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"7db9cb60-b788-495e-8079-66b57b9e207e","value":1176757059}]`)
	})
}
