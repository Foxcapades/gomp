package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringString_Put(t *testing.T) {
	Convey("TestMapStringString.Put", t, func() {
		var k string = "bf1da2f7-dae8-4c3f-9431-9fa7524cf44e"
		var v string = "eb0c0524-aada-481d-a732-49cc3832b353"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringString_Delete(t *testing.T) {
	Convey("TestMapStringString.Delete", t, func() {
		var k string = "e23ae186-5d52-42d6-a778-5e4f23be613c"
		var v string = "a7830927-f2e4-4e7b-8c61-63f5d0a499d5"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringString_Has(t *testing.T) {
	Convey("TestMapStringString.Has", t, func() {
		var k string = "be57a7aa-6281-456c-a7d8-c31268150ada"
		var v string = "53e72af5-b011-4d2e-a3f4-b6d18415f9a0"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("c12fc602-70c8-4f62-b21c-4a4830a4890a"+"fdb15dbb-ee0f-4d8f-a024-92ff3776837c"), ShouldBeFalse)
	})
}


func TestMapStringString_Get(t *testing.T) {
	Convey("TestMapStringString.Get", t, func() {
		var k string = "a4823a1a-8dcf-40ec-ae53-664ff5049d6b"
		var v string = "3a11750c-53e7-4b01-8a35-cb0e79d16aa1"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("7b699701-4bd1-4221-90d6-2fce5a08b2d7" + "1f849803-e708-4a9e-afb8-8362c2384ba9")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringString_GetOpt(t *testing.T) {
	Convey("TestMapStringString.GetOpt", t, func() {
		var k string = "9c488887-35bb-49c6-9862-2912565eef54"
		var v string = "148c2cd3-f721-4173-83d4-4670ccffe414"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("b72aed31-fe4a-407d-91c2-e4303285f04f" + "d49d7273-9853-4afa-a66b-284697424371")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringString_ForEach(t *testing.T) {
	Convey("TestMapStringString.ForEach", t, func() {
		var k string = "04c0203b-24a5-4e72-be6f-07cc5724af53"
		var v string = "ffc7fd5c-75f6-4c83-89ce-9ffc2ae78284"
		hits := 0

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv string) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringString_MarshalYAML(t *testing.T) {
	Convey("TestMapStringString.MarshalYAML", t, func() {
		var k string = "4e5a5b2a-ac49-4ea8-bf99-46be87eb61ce"
		var v string = "db522a6d-96b4-4eb7-a7fd-04fe94a1b566"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringString_ToYAML(t *testing.T) {
	Convey("TestMapStringString.ToYAML", t, func() {
		var k string = "da27f885-658b-46f3-950e-0d3118b05ef0"
		var v string = "cf410008-10dd-49d6-82c9-4882b84363cf"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringString_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringString.PutIfNotNil", t, func() {
		var k string = "ca7d0f02-14eb-466f-9631-7e92373e1119"
		var v string = "5873fe98-f9ea-42d8-aadd-23ce6fee213d"

		test := omap.NewMapStringString(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("5940c9d8-6550-4026-9cf9-a3994ea46b2e", (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "401a9f47-1733-408e-a71c-6df3fbf55fca"
		So(test.PutIfNotNil("200a5c0d-0354-4c5d-b325-14fdc67efc89", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringString_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringString.ReplaceIfExists", t, func() {
		var k string = "6a6fc075-da2a-4c97-bff3-11d4c54aacbe"
		var v string = "edfc1c6a-2fc2-428e-8bbd-7588cf67180e"
		var x string = "81e01997-f923-4897-acd0-d9aac04a9ab6"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("1a80a07d-ec5b-4711-89b8-5148635b82cb", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringString_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringString.ReplaceOrPut", t, func() {
		var k string = "f6cb554c-2128-44a5-8b9e-e8230eccfd68"
		var v string = "331a187c-7301-4928-97ff-7562444ad248"
		var x string = "f4dcb971-05c4-4798-a969-816463eaf313"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("255ed251-a668-4abf-a7dd-04ba665ca5c7", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringString_MarshalJSON(t *testing.T) {
	Convey("TestMapStringString.MarshalJSON", t, func() {
		var k string = "bd677d34-4073-4c8a-acb2-057faefc2680"
		var v string = "592865fb-66a2-4945-81a9-dc10961e5196"

		test := omap.NewMapStringString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"bd677d34-4073-4c8a-acb2-057faefc2680","value":"592865fb-66a2-4945-81a9-dc10961e5196"}]`)
	})
}

