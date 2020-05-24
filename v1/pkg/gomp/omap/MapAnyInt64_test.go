package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyInt64_Put(t *testing.T) {
	Convey("TestMapAnyInt64.Put", t, func() {
		var k interface{} = "b59769bc-c933-42bf-8dbe-b20c304c4123"
		var v int64 = 5944352518196675106

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyInt64_Delete(t *testing.T) {
	Convey("TestMapAnyInt64.Delete", t, func() {
		var k interface{} = "609473e3-5c80-4528-b91e-8fc78c82abc6"
		var v int64 = 8396043208160535266

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyInt64_Has(t *testing.T) {
	Convey("TestMapAnyInt64.Has", t, func() {
		var k interface{} = "e5490f70-8938-4fd2-bde7-f7b418520100"
		var v int64 = 9005921525178920262

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("63e3849d-401d-4acd-b423-b6bc8cbc4733"+"43504878-1e6b-4ef4-b537-160779cc6c0b"), ShouldBeFalse)
	})
}

func TestMapAnyInt64_Get(t *testing.T) {
	Convey("TestMapAnyInt64.Get", t, func() {
		var k interface{} = "1c78d0cf-c06a-4070-aa94-91064882e486"
		var v int64 = 2122429176122880998

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("45b0a79d-74fe-4160-9312-1fcfdf63f167" + "57df5589-b7bd-4f5e-bb94-69695aa50874")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyInt64_GetOpt(t *testing.T) {
	Convey("TestMapAnyInt64.GetOpt", t, func() {
		var k interface{} = "c2915008-42b6-4399-bb91-9cc9e89f18f2"
		var v int64 = 4013047287309121981

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("fa69e210-2b85-4aa7-a1cf-f80976331f40" + "eb6f79e6-fe73-4d76-a425-60563ea4199f")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyInt64_ForEach(t *testing.T) {
	Convey("TestMapAnyInt64.ForEach", t, func() {
		var k interface{} = "81aa366a-90e9-48fa-aba1-46f5367003bf"
		var v int64 = 366387212216142011
		hits := 0

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv int64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyInt64_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyInt64.MarshalYAML", t, func() {
		var k interface{} = "f04ab244-49d3-48ae-a372-937659837b86"
		var v int64 = 6798890954782261955

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyInt64_ToYAML(t *testing.T) {
	Convey("TestMapAnyInt64.ToYAML", t, func() {
		var k interface{} = "5f51577e-3c0e-449b-9ede-fedf07fec3a9"
		var v int64 = 4347996720186710695

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyInt64_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyInt64.PutIfNotNil", t, func() {
		var k interface{} = "dc285cd4-4d64-4b72-8275-95f2af93ae17"
		var v int64 = 157950995485560561

		test := omap.NewMapAnyInt64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("4a6773cd-32ba-447f-a7d4-ade4e4c010a2", (*int64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int64 = 8790986734363850803
		So(test.PutIfNotNil("c665322b-eb2a-401d-b9b4-246a96f08100", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyInt64.ReplaceIfExists", t, func() {
		var k interface{} = "7cb5ad85-0291-4211-84b4-54a6c1bfde8c"
		var v int64 = 2855284181919604339
		var x int64 = 5750115586826296369

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("bf8860ff-8cc6-4065-99f9-746fa5c8d940", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyInt64.ReplaceOrPut", t, func() {
		var k interface{} = "fc5ed89b-968c-4057-b46c-0685550cfc5c"
		var v int64 = 6667973405933083816
		var x int64 = 3583272721950849632

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("30263544-c441-4984-a16a-cbce013bd7cd", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyInt64_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyInt64.MarshalJSON", t, func() {
		var k interface{} = "70829c9f-bb89-47d8-ab97-5378f8936639"
		var v int64 = 8531031015402636699

		test := omap.NewMapAnyInt64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"70829c9f-bb89-47d8-ab97-5378f8936639","value":8531031015402636699}]`)
	})
}
