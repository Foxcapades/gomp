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
		var k int = 265411658
		var v string = "84b2fadd-cd18-419e-9e39-a270ed74c94d"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntString_Delete(t *testing.T) {
	Convey("TestMapIntString.Delete", t, func() {
		var k int = 1226376816
		var v string = "4f5ebe9d-d8f4-47b4-bd41-e52dd89a9e9a"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntString_Has(t *testing.T) {
	Convey("TestMapIntString.Has", t, func() {
		var k int = 2142323398
		var v string = "36bc0b8a-499f-47a6-8acf-794273404237"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(723039746+1895125681), ShouldBeFalse)
	})
}

func TestMapIntString_Get(t *testing.T) {
	Convey("TestMapIntString.Get", t, func() {
		var k int = 2015788780
		var v string = "b78a05bd-5380-4611-b9be-f42a70c3a972"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1170344831 + 1900183492)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntString_GetOpt(t *testing.T) {
	Convey("TestMapIntString.GetOpt", t, func() {
		var k int = 826118200
		var v string = "c6ce44b3-55f7-435a-b089-8efea7b8a806"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(84664697 + 985730605)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntString_ForEach(t *testing.T) {
	Convey("TestMapIntString.ForEach", t, func() {
		var k int = 1382986722
		var v string = "b432f65a-945a-48b0-8687-41d9d21adf8e"
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
		var k int = 2117098144
		var v string = "bce6d0e1-db1f-4bee-b80b-67a78ec719d3"

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
		var k int = 1545668144
		var v string = "dc22bfce-dff2-429b-b498-fb69e369fd1e"

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
		var k int = 1052433817
		var v string = "afaa19ce-086f-4765-abca-a8a886ffb648"

		test := omap.NewMapIntString(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1319565887, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "e9de443e-4dd3-4f1e-a09d-51b47ff07bd8"
		So(test.PutIfNotNil(1251075549, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntString_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntString.ReplaceIfExists", t, func() {
		var k int = 1293564298
		var v string = "f61dba94-ca15-402a-8c2a-54febf9d1f95"
		var x string = "059b6f51-4040-4e22-88b7-2a6dc83e28f7"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(196233999, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntString_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntString.ReplaceOrPut", t, func() {
		var k int = 971202424
		var v string = "12710ff1-c3d2-449e-9307-a2f882d4a69c"
		var x string = "7e8f344f-74b2-44d2-9b6f-f97375afa976"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1770426312, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntString_MarshalJSON(t *testing.T) {
	Convey("TestMapIntString.MarshalJSON", t, func() {
		var k int = 373304162
		var v string = "5b83eefb-c667-4ed3-821f-3a590dfe62ae"

		test := omap.NewMapIntString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":373304162,"value":"5b83eefb-c667-4ed3-821f-3a590dfe62ae"}]`)
	})
}
