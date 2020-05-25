package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringInt_Put(t *testing.T) {
	Convey("TestMapStringInt.Put", t, func() {
		var k string = "7e90c3ce-ae3e-45ab-b85b-0441c6fb4def"
		var v int = 1639993805

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringInt_Delete(t *testing.T) {
	Convey("TestMapStringInt.Delete", t, func() {
		var k string = "34764288-309a-46a5-8b8f-e58e5cd80d94"
		var v int = 1790571900

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringInt_Has(t *testing.T) {
	Convey("TestMapStringInt.Has", t, func() {
		var k string = "46f7fe28-42ff-41ec-b5c7-e3318f472b24"
		var v int = 1333595289

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("26a23cc5-04c8-4929-8dca-cc76a637b300"+"676822d6-4e84-410e-87f0-4d53ff3babdd"), ShouldBeFalse)
	})
}

func TestMapStringInt_Get(t *testing.T) {
	Convey("TestMapStringInt.Get", t, func() {
		var k string = "3cbc3127-472d-44a0-9261-06c77be6a358"
		var v int = 389242713

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("0b28e36d-96b6-490c-9cd7-d2af1de4838b" + "269e2bac-eacc-4e7d-891f-126a8c1fc7e3")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringInt_GetOpt(t *testing.T) {
	Convey("TestMapStringInt.GetOpt", t, func() {
		var k string = "3c9837ed-32d1-4131-a00d-f87b26b9b4ad"
		var v int = 1852095322

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("b12e90dc-69f6-436a-8d66-ade5262437de" + "a77b8a84-47e2-4492-abeb-8741a248a197")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringInt_ForEach(t *testing.T) {
	Convey("TestMapStringInt.ForEach", t, func() {
		var k string = "e3422d08-c538-46d7-b422-1ab7519d8b54"
		var v int = 1033950805
		hits := 0

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv int) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringInt_MarshalYAML(t *testing.T) {
	Convey("TestMapStringInt.MarshalYAML", t, func() {
		var k string = "af033bd4-31fd-4939-8037-2f72aa0ee592"
		var v int = 530786831

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringInt_ToYAML(t *testing.T) {
	Convey("TestMapStringInt.ToYAML", t, func() {
		var k string = "af38d384-ad0b-4711-82bb-66be5089b001"
		var v int = 1465841744

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.SequenceNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 1)
		So(xyml.IsMap(c.Content[0]), ShouldBeTrue)
	})
}

func TestMapStringInt_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringInt.PutIfNotNil", t, func() {
		var k string = "38ad784a-87e6-4bdf-9c59-68378dd39540"
		var v int = 519943356

		test := omap.NewMapStringInt(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("f4433481-f774-4b3a-85ce-2ea354ba9600", (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int = 1430095198
		So(test.PutIfNotNil("14ac935a-2e71-4964-80c7-2de6f8262ef1", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringInt_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringInt.ReplaceIfExists", t, func() {
		var k string = "0c36b76e-4eb0-44a8-8c2d-0386efc1786f"
		var v int = 51120487
		var x int = 1457618364

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("f6e07a84-38b1-46c1-8a83-1ca8e1087649", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringInt.ReplaceOrPut", t, func() {
		var k string = "499b632f-0564-4f8c-9056-4cbbabee64e8"
		var v int = 212882254
		var x int = 1596915940

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("25260120-be21-40f5-831f-a95355c49180", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringInt_MarshalJSON(t *testing.T) {
	Convey("TestMapStringInt.MarshalJSON", t, func() {
		var k string = "0bd36545-9224-4971-90ab-e48266110262"
		var v int = 1691469870

		test := omap.NewMapStringInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"0bd36545-9224-4971-90ab-e48266110262","value":1691469870}]`)
	})
}
