package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringUint16_Put(t *testing.T) {
	Convey("TestMapStringUint16.Put", t, func() {
		var k string = "5f378015-ba31-481d-bbda-e2605229c91f"
		var v uint16 = 14346

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringUint16_Delete(t *testing.T) {
	Convey("TestMapStringUint16.Delete", t, func() {
		var k string = "2855f60e-92fb-4d9e-ad7e-900c648bf3b6"
		var v uint16 = 48798

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringUint16_Has(t *testing.T) {
	Convey("TestMapStringUint16.Has", t, func() {
		var k string = "0fa8204a-3d72-4601-8f1d-209bb360f74a"
		var v uint16 = 32626

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("8c0dad8f-2414-4988-8b86-4e030fe0c832"+"ccdbdf16-58a6-4e5b-b135-223dedb1046b"), ShouldBeFalse)
	})
}


func TestMapStringUint16_Get(t *testing.T) {
	Convey("TestMapStringUint16.Get", t, func() {
		var k string = "dad7a953-f555-4349-9d58-022dbb80b321"
		var v uint16 = 57221

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("cd1d4810-0adf-491e-a593-d38efc489c1e" + "8fb47fd7-a4ae-408e-9959-0ae7c0789b88")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringUint16_GetOpt(t *testing.T) {
	Convey("TestMapStringUint16.GetOpt", t, func() {
		var k string = "18861bb5-1fa1-405a-87c9-9cd6d2750b2a"
		var v uint16 = 37990

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("191b3d16-0696-43cf-9c2a-a16b73db080d" + "a8b0326f-7352-41c7-b6fe-11cb42ac7c06")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringUint16_ForEach(t *testing.T) {
	Convey("TestMapStringUint16.ForEach", t, func() {
		var k string = "d2c4c855-9ebd-4d41-8b6e-efeb0d898a96"
		var v uint16 = 20037
		hits := 0

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv uint16) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringUint16_MarshalYAML(t *testing.T) {
	Convey("TestMapStringUint16.MarshalYAML", t, func() {
		var k string = "27343f21-b421-40ac-9f0b-c7c6e834cfbf"
		var v uint16 = 53593

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringUint16_ToYAML(t *testing.T) {
	Convey("TestMapStringUint16.ToYAML", t, func() {
		var k string = "5a73408e-9308-4738-b453-5d5edceb3dd6"
		var v uint16 = 26794

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringUint16_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringUint16.PutIfNotNil", t, func() {
		var k string = "1a2b7dd4-c87c-4f34-b241-a50e6e82616a"
		var v uint16 = 55311

		test := omap.NewMapStringUint16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("ed0ec7a4-02d9-4290-b9ae-a746af112752", (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint16 = 15158
		So(test.PutIfNotNil("514f4967-f1de-4d90-bbf4-38f23f06ac80", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringUint16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringUint16.ReplaceIfExists", t, func() {
		var k string = "a5a92480-d4c9-4123-987a-b7c7e98b8337"
		var v uint16 = 11720
		var x uint16 = 41296

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("e78dfabb-5611-4268-9637-1ae75a55b23b", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringUint16.ReplaceOrPut", t, func() {
		var k string = "3951e1c0-885b-4403-b66c-e38c269e9f62"
		var v uint16 = 60152
		var x uint16 = 48969

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("79f48e10-b117-488f-9fba-61439f755aac", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint16_MarshalJSON(t *testing.T) {
	Convey("TestMapStringUint16.MarshalJSON", t, func() {
		var k string = "6e43b772-e847-47ff-94b1-12e5b895bb84"
		var v uint16 = 42818

		test := omap.NewMapStringUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"6e43b772-e847-47ff-94b1-12e5b895bb84","value":42818}]`)
	})
}
