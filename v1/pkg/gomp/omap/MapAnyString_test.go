package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyString_Put(t *testing.T) {
	Convey("TestMapAnyString.Put", t, func() {
		var k interface{} = "c709a4e3-1da8-4288-9e97-b2f96c76aef9"
		var v string = "7c8e30a6-060a-4949-85d3-7a684e52bae6"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyString_Delete(t *testing.T) {
	Convey("TestMapAnyString.Delete", t, func() {
		var k interface{} = "e2024c93-54cc-455b-a355-259a664a2f00"
		var v string = "41fc166a-c9c4-404f-b4ca-71620da2a7bc"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyString_Has(t *testing.T) {
	Convey("TestMapAnyString.Has", t, func() {
		var k interface{} = "bd12e45d-c82b-4561-9036-0d96cf8310dd"
		var v string = "d7f5c59a-48c4-4252-9fb1-0ca0f8a6152f"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("ffa22b31-80d1-4ca2-8c05-0fba0e3b6e5c"+"d9163411-98e0-40de-9d24-7cf56b2a46ea"), ShouldBeFalse)
	})
}

func TestMapAnyString_Get(t *testing.T) {
	Convey("TestMapAnyString.Get", t, func() {
		var k interface{} = "2374ce64-5be6-4806-aef8-786037ecf278"
		var v string = "854b9bb3-8938-4c4e-bd7a-5559aeffa2be"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("ecdafdf1-0f2b-42c6-b147-7cdba5af616e" + "5c7e5b7c-1bb5-4eef-9cb5-818d7a7b5fb3")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyString_GetOpt(t *testing.T) {
	Convey("TestMapAnyString.GetOpt", t, func() {
		var k interface{} = "47057148-b4a5-4ede-8ed6-99ddefeaf032"
		var v string = "417f8053-a9e6-4fc7-8918-e18c7c116aa3"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("31f418fc-33d2-4f65-b841-36333e7bb993" + "97eb5248-37c6-49aa-9162-b4a30386c2d9")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyString_ForEach(t *testing.T) {
	Convey("TestMapAnyString.ForEach", t, func() {
		var k interface{} = "603228c2-e932-4a4a-b49b-73542a2df027"
		var v string = "eb79f4de-5c56-44bb-9818-73b4228891cb"
		hits := 0

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv string) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyString_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyString.MarshalYAML", t, func() {
		var k interface{} = "57fdfd4a-8829-48ed-b5bc-42ba6b5a66e7"
		var v string = "33a8efe5-7f0b-4e95-8393-48f3ede0642a"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyString_ToYAML(t *testing.T) {
	Convey("TestMapAnyString.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "e70ca00c-0285-427d-a58f-88ed39b0ae42"
			var v string = "9f3afbc6-a357-4d50-99ff-f02fa15d5fbb"

			test := omap.NewMapAnyString(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			c, d := test.ToYAML()
			So(d, ShouldBeNil)
			So(c.Kind, ShouldEqual, yaml.SequenceNode)
			So(c.LongTag(), ShouldEqual, xyml.TagOrderedMap)
			So(len(c.Content), ShouldEqual, 1)
			So(xyml.IsMap(c.Content[0]), ShouldBeTrue)
		})

		Convey("Unordered", func() {
			var k interface{} = "18f4be88-248d-49de-ae84-ad5b5bfa2ee3"
			var v string = "a624f7f2-6efa-4587-ad57-7abcb7bfaef2"

			test := omap.NewMapAnyString(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			c, d := test.ToYAML()

			So(d, ShouldBeNil)
			So(c.Kind, ShouldEqual, yaml.MappingNode)
			So(c.LongTag(), ShouldEqual, xyml.TagMap)
			So(len(c.Content), ShouldEqual, 2)
		})
	})
}

func TestMapAnyString_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyString.PutIfNotNil", t, func() {
		var k interface{} = "90b33a60-bb27-4d91-9a12-075adb0718ca"
		var v string = "cfbfc2d0-0989-4345-a3d6-cf95ca1117ea"

		test := omap.NewMapAnyString(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("b999f7d8-1c80-4840-9b6f-a3314f1f7233", (*string)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x string = "9e8d4911-4a21-4114-9da9-1226b9894f22"
		So(test.PutIfNotNil("b20071d5-f04f-4a44-9033-5d1bad130fb8", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyString_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyString.ReplaceIfExists", t, func() {
		var k interface{} = "7a1589be-9ee8-4d3b-8ba6-9e5c293a7bdc"
		var v string = "ba3edb0d-61fb-405b-8616-c277c4ef3435"
		var x string = "597c08be-83af-4ada-8f8e-830ac2fb18c2"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("2b577521-d6c7-47ea-a050-eb8831d6ee57", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyString_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyString.ReplaceOrPut", t, func() {
		var k interface{} = "0c1bebbf-a658-457f-83c2-d89cf59516ca"
		var v string = "7e71dea3-a430-4c35-928f-64b97caac542"
		var x string = "1fce46ad-ad7d-4278-b092-08194252300d"

		test := omap.NewMapAnyString(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("7496c66d-4524-4ea9-83e3-70982b0fafad", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyString_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyString.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "0a3339b9-294b-4237-85b3-2a2be7ccabdd"
			var v string = "a509fab4-4212-45b5-aaaa-ad1849eba9fa"

			test := omap.NewMapAnyString(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"0a3339b9-294b-4237-85b3-2a2be7ccabdd","value":"a509fab4-4212-45b5-aaaa-ad1849eba9fa"}]`)
		})

		Convey("Unordered", func() {
			var k interface{} = "0a3339b9-294b-4237-85b3-2a2be7ccabdd"
			var v string = "a509fab4-4212-45b5-aaaa-ad1849eba9fa"

			test := omap.NewMapAnyString(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"0a3339b9-294b-4237-85b3-2a2be7ccabdd":"a509fab4-4212-45b5-aaaa-ad1849eba9fa"}`)
		})

	})
}
