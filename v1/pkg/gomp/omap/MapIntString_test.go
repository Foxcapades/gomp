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
		var k int = 271196503
		var v string = "87881034-f217-4766-b576-065e2852b075"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntString_Delete(t *testing.T) {
	Convey("TestMapIntString.Delete", t, func() {
		var k int = 606331788
		var v string = "5b17edf0-b798-4790-a2ae-91028aad65e8"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntString_Has(t *testing.T) {
	Convey("TestMapIntString.Has", t, func() {
		var k int = 2000560583
		var v string = "8bcd7f5e-3830-4036-b779-98c848b476b1"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(2031852723+1189233742), ShouldBeFalse)
	})
}


func TestMapIntString_Get(t *testing.T) {
	Convey("TestMapIntString.Get", t, func() {
		var k int = 320294848
		var v string = "17aa8e26-c328-48d4-95a7-75a2d07f75e1"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1734236398 + 824001108)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntString_GetOpt(t *testing.T) {
	Convey("TestMapIntString.GetOpt", t, func() {
		var k int = 513202734
		var v string = "4b435298-5cfb-471b-9f03-b72ac12fab1d"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(910886606 + 332565253)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntString_ForEach(t *testing.T) {
	Convey("TestMapIntString.ForEach", t, func() {
		var k int = 1475729251
		var v string = "1eb4732b-5c89-4b38-95ba-2d8fb4c47baa"
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
		var k int = 844313309
		var v string = "5098113d-ed6c-4711-9c60-1c1ada57552b"

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
		var k int = 904226590
		var v string = "2b9b5b93-c5fc-4419-8ebc-71ccb140faf3"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntString_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntString.PutIfNotNil", t, func() {
		var k int = 1092456597
		var v string = "1db13aba-8f21-4b18-ae13-408175ae12f6"

		test := omap.NewMapIntString(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1454849785, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "5af01a5b-7724-47bc-8b34-2f2cba4614c5"
		So(test.PutIfNotNil(1285917542, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntString_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntString.ReplaceIfExists", t, func() {
		var k int = 1886462377
		var v string = "9e89f3af-afca-4b9d-859c-ffcdd7aaff27"
		var x string = "39db36dc-51f7-416a-a43d-43e7abe4c565"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(255571706, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntString_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntString.ReplaceOrPut", t, func() {
		var k int = 491078762
		var v string = "129c0bb8-af56-49fe-bdc6-c675573abdd0"
		var x string = "d5ef231c-eb88-4437-8155-96d0d531ef85"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(2060046268, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntString_MarshalJSON(t *testing.T) {
	Convey("TestMapIntString.MarshalJSON", t, func() {
		var k int = 1534788711
		var v string = "4747c6bf-4523-476b-8e67-47399d05bc01"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1534788711,"value":"4747c6bf-4523-476b-8e67-47399d05bc01"}]`)
	})
}

