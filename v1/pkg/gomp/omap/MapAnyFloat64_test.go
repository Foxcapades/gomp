package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyFloat64_Put(t *testing.T) {
	Convey("TestMapAnyFloat64.Put", t, func() {
		var k interface{} = "9d5c027f-9fc0-44cc-b990-29b8b394a7dd"
		var v float64 = 0.585

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyFloat64_Delete(t *testing.T) {
	Convey("TestMapAnyFloat64.Delete", t, func() {
		var k interface{} = "2435b781-57b0-4ad0-82e4-653cdaf3c1be"
		var v float64 = 0.671

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyFloat64_Has(t *testing.T) {
	Convey("TestMapAnyFloat64.Has", t, func() {
		var k interface{} = "d235f2b1-94c8-45cd-a7a4-1358e26f7acc"
		var v float64 = 0.120

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("3790eed0-dc6e-4709-bdd9-a3c2bff96bf4"+"eaf1aa6a-a86f-4854-ba6a-70219de28867"), ShouldBeFalse)
	})
}

func TestMapAnyFloat64_Get(t *testing.T) {
	Convey("TestMapAnyFloat64.Get", t, func() {
		var k interface{} = "2c7e5b3a-af29-4bc7-b0d0-fe139685948c"
		var v float64 = 0.547

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("eddd715e-2aba-4fe3-991d-9fc89afbc218" + "31f37c36-30e8-4c7e-8c29-f8d07bf904ec")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyFloat64_GetOpt(t *testing.T) {
	Convey("TestMapAnyFloat64.GetOpt", t, func() {
		var k interface{} = "4a739376-aab4-4aca-a414-7b89b3a5e090"
		var v float64 = 0.800

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("cd0f24ed-4591-491d-84d9-f40a237385d9" + "8bc5c8de-a16f-4f6c-8147-dc99dbdaeffd")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyFloat64_ForEach(t *testing.T) {
	Convey("TestMapAnyFloat64.ForEach", t, func() {
		var k interface{} = "ff970c35-1e3d-4718-aa45-189eceeabd17"
		var v float64 = 0.637
		hits := 0

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv float64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyFloat64_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyFloat64.MarshalYAML", t, func() {
		var k interface{} = "a28459de-ebb3-40d6-9941-5f7743ad96b0"
		var v float64 = 0.812

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyFloat64_ToYAML(t *testing.T) {
	Convey("TestMapAnyFloat64.ToYAML", t, func() {
		var k interface{} = "d183d2f9-38bf-4f15-a67c-1ac508cb1fa4"
		var v float64 = 0.376

		test := omap.NewMapAnyFloat64(1)

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

func TestMapAnyFloat64_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyFloat64.PutIfNotNil", t, func() {
		var k interface{} = "09a30d22-5e0e-4774-8c72-ba356be653d1"
		var v float64 = 0.718

		test := omap.NewMapAnyFloat64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("40539aee-3dd0-4773-a017-0bd21ea4dc7b", (*float64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float64 = 0.048
		So(test.PutIfNotNil("c9639c81-3af0-4600-8762-eb5a7364d574", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyFloat64.ReplaceIfExists", t, func() {
		var k interface{} = "7746d298-9ad6-48a4-894b-41a4b1215aa2"
		var v float64 = 0.416
		var x float64 = 0.325

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("81cdfd5c-3784-432e-bf9b-1518b4508c17", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyFloat64.ReplaceOrPut", t, func() {
		var k interface{} = "ac596301-58f2-4c24-86a8-b51a9fd68263"
		var v float64 = 0.531
		var x float64 = 0.071

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("0c0bd389-cab8-4ef1-a66a-d5885b16f77d", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat64_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyFloat64.MarshalJSON", t, func() {
		var k interface{} = "8b4597f3-1a2f-4c91-b5d8-9b2eaef9aca1"
		var v float64 = 0.113

		test := omap.NewMapAnyFloat64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"8b4597f3-1a2f-4c91-b5d8-9b2eaef9aca1","value":0.113}]`)
	})
}
