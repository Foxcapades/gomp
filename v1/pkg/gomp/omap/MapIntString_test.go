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
		var k int = 750336693
		var v string = "995df7cc-b6a6-4b1c-904c-2b98d57c6100"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntString_Delete(t *testing.T) {
	Convey("TestMapIntString.Delete", t, func() {
		var k int = 29439602
		var v string = "562d9c1f-e60a-45b2-88bd-74b339e3bffe"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntString_Has(t *testing.T) {
	Convey("TestMapIntString.Has", t, func() {
		var k int = 1325854673
		var v string = "a000a9d3-6792-4f6a-b8f8-f748daec58ed"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(596643659+448655125), ShouldBeFalse)
	})
}


func TestMapIntString_Get(t *testing.T) {
	Convey("TestMapIntString.Get", t, func() {
		var k int = 1401814555
		var v string = "e8745d9f-0c24-4bef-bf46-3cb91612312a"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1053037711+1062463378)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntString_GetOpt(t *testing.T) {
	Convey("TestMapIntString.GetOpt", t, func() {
		var k int = 143323058
		var v string = "a04c4584-add0-4eb6-b2e9-07932ee10f4a"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(282362111+932200096)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntString_ForEach(t *testing.T) {
	Convey("TestMapIntString.ForEach", t, func() {
		var k int = 1939606774
		var v string = "234f01b8-9cac-4cbe-a9f0-7276abbe18d2"
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
		var k int = 1176139219
		var v string = "b34b6f30-2d29-4c09-a2da-e2d9ae2960db"

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
		var k int = 485189586
		var v string = "5f1c4fe5-2cf9-4a2d-b673-aa0bcc5760e4"

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
		var k int = 1557131364
		var v string = "c2ef2987-5bbf-4aa0-a5f7-b39083591da1"

		test := omap.NewMapIntString(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1697657887, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "6028e1db-85c6-4074-b197-5882cfd51b87"
		So(test.PutIfNotNil(700513044, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntString_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntString.ReplaceIfExists", t, func() {
		var k int = 1833811875
		var v string = "cad6c32f-261e-410f-ab92-1d18e20ba52e"
		var x string = "e0af6e38-73bb-4687-8ed5-3cbaed6c3676"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1219851442, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntString_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntString.ReplaceOrPut", t, func() {
		var k int = 422149440
		var v string = "11ae5e35-6a11-4ef5-8fe1-9b648fe348d7"
		var x string = "fc40ba92-6f30-44fa-956a-293902528581"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(582919337, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntString_MarshalJSON(t *testing.T) {
	Convey("TestMapIntString.MarshalJSON", t, func() {
		var k int = 1886053744
		var v string = "c59a2faf-61fc-45f7-9eb7-b0945ac3f013"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1886053744,"value":"c59a2faf-61fc-45f7-9eb7-b0945ac3f013"}]`)
	})
}

