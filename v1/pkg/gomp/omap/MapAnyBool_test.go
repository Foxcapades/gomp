package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyBool_Put(t *testing.T) {
	Convey("TestMapAnyBool.Put", t, func() {
		var k interface{} = "d5b08f1c-e20c-4a35-bcbc-bd22e56d1cea"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyBool_Delete(t *testing.T) {
	Convey("TestMapAnyBool.Delete", t, func() {
		var k interface{} = "2909c89e-0e20-4349-bca6-e9dc65ffd8c0"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyBool_Has(t *testing.T) {
	Convey("TestMapAnyBool.Has", t, func() {
		var k interface{} = "ea57bb0b-b973-4251-8596-47f3cdbaf17d"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("03d4530e-8b88-49e8-bb21-e966d81ead70"+"e98fb5eb-3092-4411-aca3-08242e6ef7b1"), ShouldBeFalse)
	})
}

func TestMapAnyBool_Get(t *testing.T) {
	Convey("TestMapAnyBool.Get", t, func() {
		var k interface{} = "95763157-d99d-4e6d-822b-029ca079deb5"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("6397677c-6944-4e34-b6c5-8df758dd28c0" + "21d928b1-d3bb-4822-9495-2aca0c7e21f5")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyBool_GetOpt(t *testing.T) {
	Convey("TestMapAnyBool.GetOpt", t, func() {
		var k interface{} = "d8484321-d34d-4d61-96b6-2a290a6cf819"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("ccf66dad-55c4-4c81-817c-abbf8c01214a" + "eac7315f-8a16-4d84-a4a5-46005754ac31")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyBool_ForEach(t *testing.T) {
	Convey("TestMapAnyBool.ForEach", t, func() {
		var k interface{} = "c524edf2-712f-460e-8d8e-42f62987d8ac"
		var v bool = false
		hits := 0

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv bool) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyBool_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyBool.MarshalYAML", t, func() {
		var k interface{} = "28281c49-d956-4529-831d-d5dc65b34d64"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyBool_ToYAML(t *testing.T) {
	Convey("TestMapAnyBool.ToYAML", t, func() {
		var k interface{} = "5a76408b-5d72-45ff-a7df-76e065d4d92c"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyBool_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyBool.PutIfNotNil", t, func() {
		var k interface{} = "391a4f56-5dc2-4a41-a782-431018924d72"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*bool)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("01faa60c-c3eb-45f7-bd56-056118ac57c8", (*bool)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x bool = false
		So(test.PutIfNotNil("57742357-f1e0-4a9a-941f-83cbb11e1a6c", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyBool_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyBool.ReplaceIfExists", t, func() {
		var k interface{} = "6379cbde-1ffc-4762-af9f-09cebf00f3f6"
		var v bool = false
		var x bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("ced0287f-f950-4437-950d-00990a561a4e", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyBool_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyBool.ReplaceOrPut", t, func() {
		var k interface{} = "0ab513ff-fb6e-4b93-aea1-616e8cf894f7"
		var v bool = false
		var x bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("3717c7a2-4075-44b6-8551-1decac539f2c", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyBool_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyBool.MarshalJSON", t, func() {
		var k interface{} = "3ace1c86-f778-4777-9b9b-c87e72d6b2f5"
		var v bool = false

		test := omap.NewMapAnyBool(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"3ace1c86-f778-4777-9b9b-c87e72d6b2f5","value":false}]`)
	})
}
