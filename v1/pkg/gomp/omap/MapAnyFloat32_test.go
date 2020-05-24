package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyFloat32_Put(t *testing.T) {
	Convey("TestMapAnyFloat32.Put", t, func() {
		var k interface{} = "64ff1fd6-9d24-4461-82fd-70d5b99cd493"
		var v float32 = 0.435

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyFloat32_Delete(t *testing.T) {
	Convey("TestMapAnyFloat32.Delete", t, func() {
		var k interface{} = "eaa583f0-9398-46d8-88ac-9234ff49a98f"
		var v float32 = 0.144

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyFloat32_Has(t *testing.T) {
	Convey("TestMapAnyFloat32.Has", t, func() {
		var k interface{} = "b5762003-f60e-4e9e-adec-332b1dc8456c"
		var v float32 = 0.838

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("6ad3f78e-6be5-4b2f-ae9f-c2901c8b04c5"+"5e1170c8-811c-4c39-a6e5-61828d7129a9"), ShouldBeFalse)
	})
}


func TestMapAnyFloat32_Get(t *testing.T) {
	Convey("TestMapAnyFloat32.Get", t, func() {
		var k interface{} = "76041844-9ad0-4ea7-9ab0-d0e922edc50a"
		var v float32 = 0.330

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("b5896ae0-27df-48d8-a1a9-ad8bfdc78ea7"+"224a77ed-cfd4-4747-8556-591d63d10f7c")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyFloat32_GetOpt(t *testing.T) {
	Convey("TestMapAnyFloat32.GetOpt", t, func() {
		var k interface{} = "92d8e267-a0cd-4794-b227-347082cf61f0"
		var v float32 = 0.829

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("21ed3122-4aeb-4b54-b1f4-8e3fa9c44da1"+"937f3882-948a-48d5-a59d-9e443271c9d5")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyFloat32_ForEach(t *testing.T) {
	Convey("TestMapAnyFloat32.ForEach", t, func() {
		var k interface{} = "a40eaedb-4027-4dac-95a7-e43e037ab0cf"
		var v float32 = 0.062
		hits := 0

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv float32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyFloat32_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyFloat32.MarshalYAML", t, func() {
		var k interface{} = "e44b021c-c9ca-448c-aaea-4ce80898ecaf"
		var v float32 = 0.593

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyFloat32_ToYAML(t *testing.T) {
	Convey("TestMapAnyFloat32.ToYAML", t, func() {
		var k interface{} = "3dff5a66-5ba2-4f25-a4c3-a2e9e9428a4c"
		var v float32 = 0.852

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAnyFloat32_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyFloat32.PutIfNotNil", t, func() {
		var k interface{} = "8be463fd-44eb-470d-9565-69891f000c80"
		var v float32 = 0.531

		test := omap.NewMapAnyFloat32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("6160f4cc-f4fe-4378-9baf-bcf07221fa89", (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float32 = 0.939
		So(test.PutIfNotNil("d0717148-ac23-4794-acc7-c9951a95db98", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyFloat32.ReplaceIfExists", t, func() {
		var k interface{} = "a8158422-6a44-4147-ae08-d72c7294d745"
		var v float32 = 0.461
		var x float32 = 0.280

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("517267eb-196d-44b7-a67b-ecd637cb7a37", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyFloat32.ReplaceOrPut", t, func() {
		var k interface{} = "7418b3fb-5f13-4e20-8629-ca7b94bdf216"
		var v float32 = 0.561
		var x float32 = 0.435

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("901ae8e5-33c9-4631-82e2-112ad13a4b6c", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat32_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyFloat32.MarshalJSON", t, func() {
		var k interface{} = "d386c259-3867-4d10-850b-eec0734bb3fc"
		var v float32 = 0.297

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"d386c259-3867-4d10-850b-eec0734bb3fc","value":0.297}]`)
	})
}

