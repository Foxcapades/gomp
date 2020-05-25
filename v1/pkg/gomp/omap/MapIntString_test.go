package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntString_Put(t *testing.T) {
	Convey("TestMapIntString.Put", t, func() {
		var k int = 1030866802
		var v string = "3cf692f1-4249-43a9-8602-84701dcffc84"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntString_Delete(t *testing.T) {
	Convey("TestMapIntString.Delete", t, func() {
		var k int = 332486340
		var v string = "8970c6ae-1be1-4673-b921-edd9676a3885"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntString_Has(t *testing.T) {
	Convey("TestMapIntString.Has", t, func() {
		var k int = 903876024
		var v string = "8dff0c29-ec46-4817-aa94-8742ef626dc0"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(2008339820+363457496), ShouldBeFalse)
	})
}

func TestMapIntString_Get(t *testing.T) {
	Convey("TestMapIntString.Get", t, func() {
		var k int = 1534446063
		var v string = "fb24c3eb-4451-432b-beb4-f0538ca2cd2d"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(1845523950 + 587995302)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntString_GetOpt(t *testing.T) {
	Convey("TestMapIntString.GetOpt", t, func() {
		var k int = 1166250999
		var v string = "01506878-54a4-438c-ac9d-7d8c02e72623"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(791225701 + 3769085)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntString_ForEach(t *testing.T) {
	Convey("TestMapIntString.ForEach", t, func() {
		var k int = 1648970418
		var v string = "a7ee0982-7368-4ce0-a801-1163ffb9bf6f"
		hits := 0

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv string) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntString_MarshalYAML(t *testing.T) {
	Convey("TestMapIntString.MarshalYAML", t, func() {
		var k int = 322651319
		var v string = "cf3f2562-9199-455b-bbbd-4ccdc72eb59f"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntString_ToYAML(t *testing.T) {
	Convey("TestMapIntString.ToYAML", t, func() {
		var k int = 1467550005
		var v string = "6cc150a0-5586-44df-92e6-bad98956f22e"

		test := omap.NewMapIntString(1)

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

func TestMapIntString_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntString.PutIfNotNil", t, func() {
		var k int = 942225852
		var v string = "d137cf17-bb4b-4151-a4d0-430feb529c0c"

		test := omap.NewMapIntString(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1237189664, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "61849deb-2052-4b9b-be08-38d680706f7d"
		So(test.PutIfNotNil(1379283362, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntString_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntString.ReplaceIfExists", t, func() {
		var k int = 1479765705
		var v string = "852fcbb8-0adf-42f4-8729-c2dd154596f4"
		var x string = "15b3d3b7-9f09-4082-bb09-fdb98240005f"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1531642089, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntString_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntString.ReplaceOrPut", t, func() {
		var k int = 1267867475
		var v string = "33d5c29c-e0f0-4e41-badf-654eceb64008"
		var x string = "464b24ab-76c8-4127-8756-e4b18192ab3c"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1685310176, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntString_MarshalJSON(t *testing.T) {
	Convey("TestMapIntString.MarshalJSON", t, func() {
		var k int = 1130578852
		var v string = "d8aae3f1-cb21-4851-afd8-84dd1d942952"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1130578852,"value":"d8aae3f1-cb21-4851-afd8-84dd1d942952"}]`)
	})
}
