package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringUint64_Put(t *testing.T) {
	Convey("TestMapStringUint64.Put", t, func() {
		var k string = "821b02f6-af8c-439e-9d62-8c13eec03d40"
		var v uint64 = 2350022639414248035

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringUint64_Delete(t *testing.T) {
	Convey("TestMapStringUint64.Delete", t, func() {
		var k string = "947c4a65-ee0a-476c-8c65-fb8ea85c66bc"
		var v uint64 = 152415794905590416

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringUint64_Has(t *testing.T) {
	Convey("TestMapStringUint64.Has", t, func() {
		var k string = "c8cd3145-44d7-4d7a-91e3-75f050fc637c"
		var v uint64 = 7286235230430855642

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("3a89c3ce-00a5-45a0-9423-a3b0a3af1a39"+"3545024f-f50c-4fad-b934-b87608aebd5f"), ShouldBeFalse)
	})
}


func TestMapStringUint64_Get(t *testing.T) {
	Convey("TestMapStringUint64.Get", t, func() {
		var k string = "b8c9b433-13fa-438c-b224-615073060eb1"
		var v uint64 = 15157452511365941799

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get("495004fc-12b1-4713-98f5-dc45f2aaf2c2" + "73dd2418-e58d-4ed0-87b1-b218ffada925")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringUint64_GetOpt(t *testing.T) {
	Convey("TestMapStringUint64.GetOpt", t, func() {
		var k string = "198c4005-1e0a-4c50-84ee-3aa8a0421efe"
		var v uint64 = 2919787773188631391

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("8af159ed-3459-4ae1-a91e-cc3d5b1136fc" + "7b3d4819-ce71-405e-a656-d38709d7812d")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringUint64_ForEach(t *testing.T) {
	Convey("TestMapStringUint64.ForEach", t, func() {
		var k string = "0b8f30de-be4a-4f5f-a4c9-97980c7bfe81"
		var v uint64 = 6901912218310014329
		hits := 0

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv uint64) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringUint64_MarshalYAML(t *testing.T) {
	Convey("TestMapStringUint64.MarshalYAML", t, func() {
		var k string = "095242a5-38e4-474e-a775-4b45730115e1"
		var v uint64 = 2948861355782621774

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringUint64_ToYAML(t *testing.T) {
	Convey("TestMapStringUint64.ToYAML", t, func() {
		var k string = "6d7c8421-b3dc-493b-b6b3-f44f80c2aa6b"
		var v uint64 = 8915631205399650371

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapStringUint64_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringUint64.PutIfNotNil", t, func() {
		var k string = "bef5ec50-8cd4-43a0-bda3-ba957d9ed07d"
		var v uint64 = 10527628601282100838

		test := omap.NewMapStringUint64(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("4b95207e-beaf-4249-8fc5-ea2b75b093bd", (*uint64)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint64 = 5946279488135478239
		So(test.PutIfNotNil("c65eab2a-4e07-484a-ac69-55d3062bc880", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringUint64_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringUint64.ReplaceIfExists", t, func() {
		var k string = "27b9e9c1-fd95-4d63-b4c2-d58717dd073f"
		var v uint64 = 16344675991005283317
		var x uint64 = 7294338973595902937

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("43f1f781-64a5-4519-8645-cadd723b9aa3", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint64_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringUint64.ReplaceOrPut", t, func() {
		var k string = "0d2c7752-bfe4-4530-8723-754de27b85b9"
		var v uint64 = 12402411923700131427
		var x uint64 = 6631328768793095454

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("b1b57cd7-050a-40b7-9c11-7acdaf35ed80", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint64_MarshalJSON(t *testing.T) {
	Convey("TestMapStringUint64.MarshalJSON", t, func() {
		var k string = "aa588666-48c8-4157-aeea-3da6e07ac485"
		var v uint64 = 17067222309253608432

		test := omap.NewMapStringUint64(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":"aa588666-48c8-4157-aeea-3da6e07ac485","value":17067222309253608432}]`)
	})
}
