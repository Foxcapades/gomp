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
		var k string = "81787846-ff8a-45fa-a270-8daa7b1b7fe9"
		var v int64 = 8148548612212397501

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringInt64_Delete(t *testing.T) {
	Convey("TestMapStringInt64.Delete", t, func() {
		var k string = "d8503a80-9bb5-4259-bc8c-1dcec89a3d28"
		var v int64 = 2767245423226527670

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringInt64_Has(t *testing.T) {
	Convey("TestMapStringInt64.Has", t, func() {
		var k string = "cdb3e63e-6343-4759-90db-10c3f08e820f"
		var v int64 = 804547478610168050

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("cf41047a-6343-47b2-b3bd-a464d898b76d"+"7196e7a5-bf7a-4831-8e08-b8222bdfc741"), ShouldBeFalse)
	})
}


func TestMapStringInt64_Get(t *testing.T) {
	Convey("TestMapStringInt64.Get", t, func() {
		var k string = "216b7d5f-1ef6-4df7-a8f9-8920112a767b"
		var v int64 = 4731839807591617202

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("34b96fb5-1a08-4838-b8fe-32b54e885662" + "e0e7ca91-159d-4048-aaf9-5051e63a2717")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringInt64_GetOpt(t *testing.T) {
	Convey("TestMapStringInt64.GetOpt", t, func() {
		var k string = "11be1026-a541-4f4d-b1b0-574dd71b3e28"
		var v int64 = 391527534374499213

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("d6bee7ea-ff50-4901-81e5-4ae14a8acd8b" + "4b064f15-a9b8-4a7e-bb60-6d2bd083c1e0")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringInt64_ForEach(t *testing.T) {
	Convey("TestMapStringInt64.ForEach", t, func() {
		var k string = "d0233889-4b89-4ad3-9161-b0edcfdffb6f"
		var v int64 = 7050934425361495188
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
		var k string = "cd5989a2-4e00-45a3-89fb-1250d9a5009f"
		var v int64 = 4672340503510267074

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
		var k string = "4c5aad6f-0583-4aea-8fb5-c4d2c2b41ea1"
		var v int64 = 8605879846971181869

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
		var k string = "7ac30459-1c68-467c-bc5f-cc551ef84cc3"
		var v int64 = 4783088298666092113

		test := omap.NewMapStringInt64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("25ba845f-2dcc-44a1-acba-788c0fdb8798", (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int64 = 7638135071870216587
		So(test.PutIfNotNil("a584fb8a-653a-44b6-8d30-f672b637eec8", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringInt64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringInt64.ReplaceIfExists", t, func() {
		var k string = "307da318-b7d5-4f75-ad6e-6b6dfd28d8c8"
		var v int64 = 8020840926263143584
		var x int64 = 944480521919664070

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("c96661df-b9fe-4d1c-9db5-2127cc743efc", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringInt64.ReplaceOrPut", t, func() {
		var k string = "bb535865-dae2-433c-8da8-62b69ba5e749"
		var v int64 = 8049960607921514600
		var x int64 = 7977955135774187311

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("439d29ef-4657-4683-89ac-04fef0683b97", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt64_MarshalJSON(t *testing.T) {
	Convey("TestMapStringInt64.MarshalJSON", t, func() {
		var k string = "d434cb61-9956-4526-909f-6a01d889f34f"
		var v int64 = 6204637147112348171

		test := omap.NewMapStringInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"d434cb61-9956-4526-909f-6a01d889f34f","value":6204637147112348171}]`)
	})
}
