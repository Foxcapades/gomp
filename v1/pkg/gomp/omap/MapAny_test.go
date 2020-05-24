package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAny_Put(t *testing.T) {
	Convey("TestMapAny.Put", t, func() {
		var k interface{} = "7e97b713-8697-452b-a786-298dfc7e8bde"
		var v interface{} = "e252f5b9-afda-4b4a-9152-d2db2eba813c"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAny_Delete(t *testing.T) {
	Convey("TestMapAny.Delete", t, func() {
		var k interface{} = "53a73cec-27d5-434c-988d-e4d948d9e1f8"
		var v interface{} = "ecfde25a-832c-430c-9b78-1f8196dc2eba"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAny_Has(t *testing.T) {
	Convey("TestMapAny.Has", t, func() {
		var k interface{} = "140288c9-31a2-4b76-913b-9e7b3ca06ff1"
		var v interface{} = "28cc77f2-9f43-4a90-8d54-45fcf92d34e7"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("d8b970f8-1a5e-4eaa-bb50-f2e4f6200cbc"+"3a37f599-1e32-4ce5-ac2b-15f84d58e85c"), ShouldBeFalse)
	})
}


func TestMapAny_Get(t *testing.T) {
	Convey("TestMapAny.Get", t, func() {
		var k interface{} = "57fef5fc-8ded-4ed9-bea9-fb46ad74ae3d"
		var v interface{} = "8ab5e0db-e109-4899-b74e-90021631c8eb"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("66ec006c-4a1b-40e6-bcf6-7a8c3e7409bd"+"73fdaa9a-2f99-4e36-9753-b88c2c4fedd6")
		So(b, ShouldBeFalse)
	})
}

func TestMapAny_GetOpt(t *testing.T) {
	Convey("TestMapAny.GetOpt", t, func() {
		var k interface{} = "ff562aff-58e7-45c5-92b4-13845afea359"
		var v interface{} = "89b094de-2931-4d12-8b00-ae7f267347e6"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("572eed9c-f000-40aa-9a3b-fefe8706e05e"+"bda6ed8d-3d66-44e7-ba43-2ec71dedbe3d")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAny_ForEach(t *testing.T) {
	Convey("TestMapAny.ForEach", t, func() {
		var k interface{} = "74af637e-8bda-4d3b-9d5e-5a081f71566e"
		var v interface{} = "ea966907-835e-4ce5-a89f-78aa9a7cdba9"
		hits := 0

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv interface{}) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAny_MarshalYAML(t *testing.T) {
	Convey("TestMapAny.MarshalYAML", t, func() {
		var k interface{} = "cbafde03-2c04-4239-85c5-2a8f57db8ccb"
		var v interface{} = "6f865d31-0cd0-4b90-99f7-573f1f1680c7"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAny_ToYAML(t *testing.T) {
	Convey("TestMapAny.ToYAML", t, func() {
		var k interface{} = "abdce391-fc4f-4c16-b20f-c759ed90cdf0"
		var v interface{} = "4fd0f5a3-d9a3-4a61-ad3d-953f14b89d73"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapAny_PutIfNotNil(t *testing.T) {
	Convey("TestMapAny.PutIfNotNil", t, func() {
		var k interface{} = "c3671792-42ab-40ea-91af-4578c07c3fa2"
		var v interface{} = "536194dc-a194-413a-8760-0f4935859d5d"

		test := omap.NewMapAny(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("8f103e58-7815-4916-9d1e-180f3a60fe0f", (*interface{})(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x interface{} = "46e368e2-fecd-4af9-9f1a-d5d956976934"
		So(test.PutIfNotNil("50369aa0-7f5e-47d1-8fb9-022ad911d218", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAny_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAny.ReplaceIfExists", t, func() {
		var k interface{} = "48d9ba8d-3dea-4884-a9b4-7a8f3d90fefc"
		var v interface{} = "c1ef02bc-2f6a-42d8-9454-dce61443f9f0"
		var x interface{} = "53f18679-d2c9-47d7-8aca-b23d61d943bd"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("64c7d6e4-0034-4071-92d4-6137dd7be5c0", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAny_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAny.ReplaceOrPut", t, func() {
		var k interface{} = "190fae12-ff04-49a3-bd1e-b4d804c169a3"
		var v interface{} = "0d3a823f-b7d4-41a1-aea7-1c2653f86eac"
		var x interface{} = "1bd6d45a-a580-4ce7-9695-220ea262f66c"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("baf016e2-9586-4cbe-b9c2-8e876badcfc7", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAny_MarshalJSON(t *testing.T) {
	Convey("TestMapAny.MarshalJSON", t, func() {
		var k interface{} = "e52fe73d-22ce-47b7-bef6-259eb10fafc9"
		var v interface{} = "4e0d2f87-2652-49f7-b71c-6e0908b0fe23"

		test := omap.NewMapAny(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"e52fe73d-22ce-47b7-bef6-259eb10fafc9","value":"4e0d2f87-2652-49f7-b71c-6e0908b0fe23"}]`)
	})
}

