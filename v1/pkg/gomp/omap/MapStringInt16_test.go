package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringInt16_Put(t *testing.T) {
	Convey("TestMapStringInt16.Put", t, func() {
		var k string = "052af1a2-e414-48e3-a489-04217b8dc589"
		var v int16 = 25615

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringInt16_Delete(t *testing.T) {
	Convey("TestMapStringInt16.Delete", t, func() {
		var k string = "5f788d80-dacf-4330-a35a-fbf798b116df"
		var v int16 = 13451

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringInt16_Has(t *testing.T) {
	Convey("TestMapStringInt16.Has", t, func() {
		var k string = "2db903be-01be-4cb9-87cd-2aa9fb6a72a3"
		var v int16 = 25520

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("d042bbdb-867d-445b-9fec-2894f49904dc"+"45cef269-8641-455c-95b6-5259ebb187bd"), ShouldBeFalse)
	})
}

func TestMapStringInt16_Get(t *testing.T) {
	Convey("TestMapStringInt16.Get", t, func() {
		var k string = "fdc221e3-d040-43df-9658-a0d57a56b3eb"
		var v int16 = 10971

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("cfb50052-3977-4cd3-aa28-8ea2b2b0f2d7" + "bb61f958-f7cf-4382-9fcc-e59325553cf5")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringInt16_GetOpt(t *testing.T) {
	Convey("TestMapStringInt16.GetOpt", t, func() {
		var k string = "db6e1c20-e44b-45bb-bd88-83c746648716"
		var v int16 = 11340

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("7c86a226-efbe-455f-9bab-d93ad3730e07" + "61076998-1a28-4ce5-a416-c08a42c68b02")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringInt16_ForEach(t *testing.T) {
	Convey("TestMapStringInt16.ForEach", t, func() {
		var k string = "1a9e4287-d0f0-4719-a7fd-eebea9963bf2"
		var v int16 = 31642
		hits := 0

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv int16) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringInt16_MarshalYAML(t *testing.T) {
	Convey("TestMapStringInt16.MarshalYAML", t, func() {
		var k string = "535862c9-099f-4baa-8bc7-14b793f87f91"
		var v int16 = 28647

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringInt16_ToYAML(t *testing.T) {
	Convey("TestMapStringInt16.ToYAML", t, func() {
		var k string = "2dcc68f9-a265-4205-a6a0-7bcc097e9049"
		var v int16 = 5703

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringInt16_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringInt16.PutIfNotNil", t, func() {
		var k string = "49e1e88b-00b9-40f9-817d-c4840b869337"
		var v int16 = 10009

		test := omap.NewMapStringInt16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("acf2a911-36c5-4486-a516-76791c0a12cc", (*int16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int16 = 31665
		So(test.PutIfNotNil("a58cb5d6-4d3f-419c-b76c-79b90f69c6dc", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringInt16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringInt16.ReplaceIfExists", t, func() {
		var k string = "45c83861-dd90-47c3-b302-19cd48e6f967"
		var v int16 = 30134
		var x int16 = 28385

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("1cbabec6-6af2-47fe-b351-50358b8bdf86", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringInt16.ReplaceOrPut", t, func() {
		var k string = "b3349f66-cd0a-4126-9fd9-0aaa08672979"
		var v int16 = 4364
		var x int16 = 849

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("4be1cd61-d964-4b92-ba08-fc54efccc391", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt16_MarshalJSON(t *testing.T) {
	Convey("TestMapStringInt16.MarshalJSON", t, func() {
		var k string = "3d4a9f09-9b5a-4d27-b3ac-1e133e329b05"
		var v int16 = 32342

		test := omap.NewMapStringInt16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"3d4a9f09-9b5a-4d27-b3ac-1e133e329b05","value":32342}]`)
	})
}
