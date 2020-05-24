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
		var k string = "66ebc442-5669-43ce-b4fa-4e8e2ca5611a"
		var v uint64 = 17663532307428410315

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringUint64_Delete(t *testing.T) {
	Convey("TestMapStringUint64.Delete", t, func() {
		var k string = "c45c241b-4165-46a9-a6ae-df4a2a7a541d"
		var v uint64 = 6935706149356774954

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringUint64_Has(t *testing.T) {
	Convey("TestMapStringUint64.Has", t, func() {
		var k string = "f1930be4-357f-4847-98d3-a1794b63938f"
		var v uint64 = 15569018356362562453

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("09130e5b-30b1-4aa3-8ea0-4ed07793c686"+"a73777c1-18a0-4b8d-ba6e-cb24ce315a21"), ShouldBeFalse)
	})
}


func TestMapStringUint64_Get(t *testing.T) {
	Convey("TestMapStringUint64.Get", t, func() {
		var k string = "9d96ff10-6ac0-4bf4-a03d-36ea2a0a62ea"
		var v uint64 = 8498359651613904681

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("66866304-ab74-474d-bbd3-3af71cb696cb" + "4eadee5d-21b7-4ce9-a867-90fdebce3e0a")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringUint64_GetOpt(t *testing.T) {
	Convey("TestMapStringUint64.GetOpt", t, func() {
		var k string = "3756a4b5-7b75-4039-8602-bf91173ac3b7"
		var v uint64 = 3352398008264492996

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("a417db18-f1d2-4e96-99dd-9618757c7971" + "b8e23d37-c973-4537-93d7-7d92bd0dd5c8")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringUint64_ForEach(t *testing.T) {
	Convey("TestMapStringUint64.ForEach", t, func() {
		var k string = "c182ea70-efc2-4efb-9bf9-86ff77db9f15"
		var v uint64 = 14994734888150330298
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
		var k string = "3c7c07cd-d367-4e09-bbd2-c702cde449f0"
		var v uint64 = 6381698549179781437

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
		var k string = "f0764a67-ca58-4938-b4d1-4cf8c2b5b264"
		var v uint64 = 7803154316605293321

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
		var k string = "89ee34da-7ced-44d1-ab27-1e95c8755733"
		var v uint64 = 7226509886726476608

		test := omap.NewMapStringUint64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("e8f257e5-02fa-4e22-aa79-ef23c4eb83c4", (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint64 = 6373488430964456283
		So(test.PutIfNotNil("5e3e6bd2-8a93-46d5-a897-400642344408", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringUint64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringUint64.ReplaceIfExists", t, func() {
		var k string = "1ffe5712-19ff-43be-be54-80bccbfad57c"
		var v uint64 = 10560062142390921687
		var x uint64 = 15256457522844331418

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("b9030f59-c2b4-41a1-baf2-2498ba650227", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringUint64.ReplaceOrPut", t, func() {
		var k string = "25eb27a0-0350-4ffd-8b05-64bd49c029eb"
		var v uint64 = 2558756459209986325
		var x uint64 = 13678974931633826247

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("6f0843ba-6abf-404b-b178-45bba272c3a3", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint64_MarshalJSON(t *testing.T) {
	Convey("TestMapStringUint64.MarshalJSON", t, func() {
		var k string = "4fc777aa-cdb5-4e74-ae6b-5236c2971039"
		var v uint64 = 13509666100271719395

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"4fc777aa-cdb5-4e74-ae6b-5236c2971039","value":13509666100271719395}]`)
	})
}

