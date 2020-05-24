package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringUint64_Put(t *testing.T) {
	Convey("TestMapStringUint64.Put", t, func() {
		var k string = "d67041d7-a992-455b-8303-2e7db0b45784"
		var v uint64 = 13667506668107291369

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringUint64_Delete(t *testing.T) {
	Convey("TestMapStringUint64.Delete", t, func() {
		var k string = "8dce66f0-6b0e-4335-99fb-0d286966a39e"
		var v uint64 = 961013324129767155

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringUint64_Has(t *testing.T) {
	Convey("TestMapStringUint64.Has", t, func() {
		var k string = "1c3cda66-b3a5-4ce7-b5d4-80b2617ed679"
		var v uint64 = 10962715121876111267

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("ec1303fc-744b-4415-815c-c8570eaa63f4"+"61fedf78-1eec-4662-aca9-364a5dfa6c51"), ShouldBeFalse)
	})
}


func TestMapStringUint64_Get(t *testing.T) {
	Convey("TestMapStringUint64.Get", t, func() {
		var k string = "e5e8aeef-0b74-44fb-8c07-93be644876b6"
		var v uint64 = 6860521981105206200

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("6df25d5e-f537-4eea-96c1-9fcff34468c8"+"9839d8d7-f581-4c63-b2b2-9e1f0aa9611d")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringUint64_GetOpt(t *testing.T) {
	Convey("TestMapStringUint64.GetOpt", t, func() {
		var k string = "f520a063-a7b7-419d-a5f8-553bf2fdbd89"
		var v uint64 = 11036125744844063142

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("ace79ef0-5fca-4104-866c-ef9ee7d1b8e6"+"fad8505c-42ad-4e2c-bea7-8c2ee771a1e4")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringUint64_ForEach(t *testing.T) {
	Convey("TestMapStringUint64.ForEach", t, func() {
		var k string = "a6564ea3-769d-4d94-b508-2db0eab6f6f6"
		var v uint64 = 1540237488085844762
		hits := 0

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv uint64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringUint64_MarshalYAML(t *testing.T) {
	Convey("TestMapStringUint64.MarshalYAML", t, func() {
		var k string = "898358b5-e5b0-49c6-8701-26787e7b0d93"
		var v uint64 = 11837867407160495306

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringUint64_ToYAML(t *testing.T) {
	Convey("TestMapStringUint64.ToYAML", t, func() {
		var k string = "a2a4cd0e-0d17-4125-a5ed-4d10e3652f77"
		var v uint64 = 17165788438577579054

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringUint64_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringUint64.PutIfNotNil", t, func() {
		var k string = "52608d09-a804-4ee6-bd8c-bb28d78b5970"
		var v uint64 = 13663077014701149945

		test := omap.NewMapStringUint64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("65fca94f-27e5-46fe-a5b1-4aa6fb2984ca", (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint64 = 1707991068423035663
		So(test.PutIfNotNil("6dc70d5a-16c0-4f75-8d26-e25b0bda873e", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringUint64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringUint64.ReplaceIfExists", t, func() {
		var k string = "6998fbe2-465b-4aca-a9c5-8ec692f27814"
		var v uint64 = 2824121101518164065
		var x uint64 = 4940910970242948622

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("b4579aa5-78a6-4e59-9fb8-b0f383f4cefc", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringUint64.ReplaceOrPut", t, func() {
		var k string = "173c1c0e-43be-4493-a4be-446be108c94e"
		var v uint64 = 12890824664756983546
		var x uint64 = 2541856675805714302

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("1e703ca6-86cc-4852-8784-c3bce7e88180", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint64_MarshalJSON(t *testing.T) {
	Convey("TestMapStringUint64.MarshalJSON", t, func() {
		var k string = "0eeaf21b-94b4-4d58-9717-0b93645ad6ee"
		var v uint64 = 16808662775594644164

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"0eeaf21b-94b4-4d58-9717-0b93645ad6ee","value":16808662775594644164}]`)
	})
}

