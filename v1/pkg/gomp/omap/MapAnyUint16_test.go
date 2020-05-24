package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyUint16_Put(t *testing.T) {
	Convey("TestMapAnyUint16.Put", t, func() {
		var k interface{} = "8f6db7e3-c64e-490c-8db8-9813c432d9a9"
		var v uint16 = 10630

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyUint16_Delete(t *testing.T) {
	Convey("TestMapAnyUint16.Delete", t, func() {
		var k interface{} = "b4e54dec-81cc-49c8-bcd8-91e4c9ec0282"
		var v uint16 = 60737

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyUint16_Has(t *testing.T) {
	Convey("TestMapAnyUint16.Has", t, func() {
		var k interface{} = "3a8a5c27-2635-435b-b554-44a156701543"
		var v uint16 = 38522

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("7f581903-3850-46ff-b0b2-c2ef08e4b34d"+"cf9c6a30-b17c-455c-ae70-1e4a7c0a30fd"), ShouldBeFalse)
	})
}

func TestMapAnyUint16_Get(t *testing.T) {
	Convey("TestMapAnyUint16.Get", t, func() {
		var k interface{} = "3e0bb4bc-6b32-4b51-862f-17de7fec968c"
		var v uint16 = 47313

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("2afa4e72-4dd6-4277-ab05-614a6f3c3f77" + "6de1294f-b700-4868-8ccd-a3f57f1b7bc5")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyUint16_GetOpt(t *testing.T) {
	Convey("TestMapAnyUint16.GetOpt", t, func() {
		var k interface{} = "a30e2912-ea41-4e9c-ba46-5f8eeb2a7a5c"
		var v uint16 = 22481

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("5c8885fc-cfb1-474c-8149-92d0bfbe7618" + "454cebe6-e512-4e8e-bfbe-ceb36b3a8c5c")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyUint16_ForEach(t *testing.T) {
	Convey("TestMapAnyUint16.ForEach", t, func() {
		var k interface{} = "299a0beb-aab3-4483-9843-7c7fc401ee21"
		var v uint16 = 58189
		hits := 0

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv uint16) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyUint16_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyUint16.MarshalYAML", t, func() {
		var k interface{} = "cc208421-6ffc-481a-8acf-fdad9e7ea412"
		var v uint16 = 58977

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyUint16_ToYAML(t *testing.T) {
	Convey("TestMapAnyUint16.ToYAML", t, func() {
		var k interface{} = "fc84cd2e-963f-4cad-a5eb-08cac673f395"
		var v uint16 = 40865

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyUint16_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyUint16.PutIfNotNil", t, func() {
		var k interface{} = "7aeac365-e85c-44ce-9344-2a91df22d631"
		var v uint16 = 27914

		test := omap.NewMapAnyUint16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("bf4e22f6-d561-45f5-bf18-07ec883870aa", (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint16 = 43429
		So(test.PutIfNotNil("bd9f13a5-bb2f-4dd5-b0a9-c334492c47b9", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyUint16.ReplaceIfExists", t, func() {
		var k interface{} = "d5b026b7-8417-43b4-9dfb-a3e842d9e8e3"
		var v uint16 = 8048
		var x uint16 = 50987

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("96be0bf0-3862-4c51-8a92-7b5957d5d705", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyUint16.ReplaceOrPut", t, func() {
		var k interface{} = "94eef0e3-04d1-463e-a7a6-cfe6b0833eb7"
		var v uint16 = 29248
		var x uint16 = 14632

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("f05a0edd-cd63-4cb1-b4a9-eb194c8894d5", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint16_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyUint16.MarshalJSON", t, func() {
		var k interface{} = "89304e44-6b7e-46a5-b463-0825795a305e"
		var v uint16 = 38651

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"89304e44-6b7e-46a5-b463-0825795a305e","value":38651}]`)
	})
}
