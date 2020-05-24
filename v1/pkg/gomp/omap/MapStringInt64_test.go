package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringInt64_Put(t *testing.T) {
	Convey("TestMapStringInt64.Put", t, func() {
		var k string = "afc56730-ee88-4cf6-a7f1-ac6bb452fa6a"
		var v int64 = 2777896761030433531

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringInt64_Delete(t *testing.T) {
	Convey("TestMapStringInt64.Delete", t, func() {
		var k string = "5a7ff205-bd08-48d3-b0cb-55790bff2890"
		var v int64 = 3754381534386467736

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringInt64_Has(t *testing.T) {
	Convey("TestMapStringInt64.Has", t, func() {
		var k string = "cba8e7b9-dcaa-4108-a6a2-3983de557d55"
		var v int64 = 184160961946060757

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("fd02243e-a38f-4ab3-9070-9a44797d380d"+"c7a5f469-97a1-43ff-bb73-bbd02b96255c"), ShouldBeFalse)
	})
}


func TestMapStringInt64_Get(t *testing.T) {
	Convey("TestMapStringInt64.Get", t, func() {
		var k string = "7de19f27-5dd1-41fc-8050-92ccad30157b"
		var v int64 = 6771727014215989688

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("4772a985-a204-496d-aa86-9f0f4d1cc6e2" + "a3d240a6-4ed3-4e25-a097-f54357388f0d")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringInt64_GetOpt(t *testing.T) {
	Convey("TestMapStringInt64.GetOpt", t, func() {
		var k string = "195778f4-2a8e-401d-a9eb-6a580131794d"
		var v int64 = 5683187985490565009

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("ad70dbed-dd3f-4393-8589-4ea11dd37dd0" + "687dff72-fee5-4c3d-90d5-c5a5cc689a2f")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringInt64_ForEach(t *testing.T) {
	Convey("TestMapStringInt64.ForEach", t, func() {
		var k string = "52706137-9fb1-45a2-8033-fde8745c9432"
		var v int64 = 7832104391819701985
		hits := 0

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv int64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringInt64_MarshalYAML(t *testing.T) {
	Convey("TestMapStringInt64.MarshalYAML", t, func() {
		var k string = "897f876a-d1a2-4a30-8c06-8d153abda488"
		var v int64 = 4105433550886927146

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringInt64_ToYAML(t *testing.T) {
	Convey("TestMapStringInt64.ToYAML", t, func() {
		var k string = "2ce31952-5bf2-4c68-960c-61278ea268ad"
		var v int64 = 8532456874297088754

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringInt64_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringInt64.PutIfNotNil", t, func() {
		var k string = "7e7906b1-7dbb-4966-a1b2-7d749d18dbb5"
		var v int64 = 394664991344316168

		test := omap.NewMapStringInt64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("85b6428f-f759-4e9a-9143-3bea533c36a0", (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int64 = 8938657874279386289
		So(test.PutIfNotNil("ec3fe249-83af-4ed6-9d8e-01fa9710c8a5", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringInt64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringInt64.ReplaceIfExists", t, func() {
		var k string = "ee1b0e3a-30c1-4be2-af67-9436ba02a540"
		var v int64 = 8614742169396859656
		var x int64 = 1192632364069905525

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("df94e903-7566-42a0-990c-7e22ea4ef95b", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringInt64.ReplaceOrPut", t, func() {
		var k string = "421a0257-8520-49d8-bf89-0d755a14c92d"
		var v int64 = 2646364984317989043
		var x int64 = 7974644206437311822

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("a2a5398e-02c0-4b75-9659-dc0727d06da8", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt64_MarshalJSON(t *testing.T) {
	Convey("TestMapStringInt64.MarshalJSON", t, func() {
		var k string = "4c33c6e5-e59f-4ec2-b541-63c416129b39"
		var v int64 = 8710798189780972324

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"4c33c6e5-e59f-4ec2-b541-63c416129b39","value":8710798189780972324}]`)
	})
}

