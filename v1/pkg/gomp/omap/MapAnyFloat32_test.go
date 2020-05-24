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
		var k interface{} = "2de4d4e2-edf4-414b-84c5-6b84c70e7d8d"
		var v float32 = 0.893

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyFloat32_Delete(t *testing.T) {
	Convey("TestMapAnyFloat32.Delete", t, func() {
		var k interface{} = "6cd39c7d-e757-4fd9-95c0-2079f97fbd92"
		var v float32 = 0.838

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyFloat32_Has(t *testing.T) {
	Convey("TestMapAnyFloat32.Has", t, func() {
		var k interface{} = "38bb7be4-2ee3-4796-8283-96dcdc62997b"
		var v float32 = 0.242

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("2871e9e8-ba58-48f1-8b13-e0ade4f108b6"+"47a789ab-e990-4fc6-b910-8a73b02e912b"), ShouldBeFalse)
	})
}

func TestMapAnyFloat32_Get(t *testing.T) {
	Convey("TestMapAnyFloat32.Get", t, func() {
		var k interface{} = "e0480e94-4797-457b-bfcf-b7ae87ebde41"
		var v float32 = 0.486

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("d2694fa9-2ad5-4d96-a919-894883a7b002" + "99aef31c-a9a1-4733-9183-0642237aab29")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyFloat32_GetOpt(t *testing.T) {
	Convey("TestMapAnyFloat32.GetOpt", t, func() {
		var k interface{} = "9cafea26-347f-4c77-b5b7-b80de6d8aca4"
		var v float32 = 0.261

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("14168968-bccd-4e10-8f5a-68288cf970f7" + "43a36c4c-c75c-4cd8-bdec-162c83fbc192")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyFloat32_ForEach(t *testing.T) {
	Convey("TestMapAnyFloat32.ForEach", t, func() {
		var k interface{} = "4f9040ed-53d9-4ae6-8f01-b6e141900bfa"
		var v float32 = 0.383
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
		var k interface{} = "1ba2c39a-3a48-4cca-b3d0-2bb8f41d8139"
		var v float32 = 0.908

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
		var k interface{} = "ca5d0ead-e7d3-4f9a-8f80-82eb60247ec2"
		var v float32 = 0.972

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
		var k interface{} = "57292a96-e133-43ec-b8fd-fff1ee2cf91e"
		var v float32 = 0.117

		test := omap.NewMapAnyFloat32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("42cadce7-08f2-4fbf-91dd-fbf58c1f225c", (*float32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x float32 = 0.065
		So(test.PutIfNotNil("c31b3a46-3e96-472f-98ff-14e07609af66", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyFloat32.ReplaceIfExists", t, func() {
		var k interface{} = "9a7bf30f-13e2-472f-8a09-f7ec795c2715"
		var v float32 = 0.214
		var x float32 = 0.508

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("fda84feb-eb2e-4420-8471-0f97a8af8a78", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyFloat32.ReplaceOrPut", t, func() {
		var k interface{} = "15c448ab-1104-4cb8-858a-0b6dae90412f"
		var v float32 = 0.637
		var x float32 = 0.444

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("0a4bfcdf-3ed4-4f66-b71d-d882df1e6639", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyFloat32_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyFloat32.MarshalJSON", t, func() {
		var k interface{} = "5aaab412-b68d-4a1c-817f-eac228fd815d"
		var v float32 = 0.678

		test := omap.NewMapAnyFloat32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"5aaab412-b68d-4a1c-817f-eac228fd815d","value":0.678}]`)
	})
}
