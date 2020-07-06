package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapAnyUint16_Put(t *testing.T) {
	Convey("TestMapAnyUint16.Put", t, func() {
		var k interface{} = "99e5a38b-95f4-4969-ac23-1ff7d4ee8295"
		var v uint16 = 650

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapAnyUint16_Delete(t *testing.T) {
	Convey("TestMapAnyUint16.Delete", t, func() {
		var k interface{} = "577e2661-3efe-4925-836c-1be5ea396a68"
		var v uint16 = 4593

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapAnyUint16_Has(t *testing.T) {
	Convey("TestMapAnyUint16.Has", t, func() {
		var k interface{} = "d2525bd2-1bfc-4c45-b17e-2c7a13c0caf7"
		var v uint16 = 17577

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("ca72f344-3f94-4e4a-8fa4-4ffbfce92ee8"+"16a6a38d-f847-4c5f-9876-b62c13245304"), ShouldBeFalse)
	})
}

func TestMapAnyUint16_Get(t *testing.T) {
	Convey("TestMapAnyUint16.Get", t, func() {
		var k interface{} = "0db363d4-ca25-4085-859d-7ac37dc170fb"
		var v uint16 = 46060

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("9bcc0617-3698-42af-a2c6-76e485515043" + "369b5503-44e5-48e5-811f-ae42d1683f59")
		So(b, ShouldBeFalse)
	})
}

func TestMapAnyUint16_GetOpt(t *testing.T) {
	Convey("TestMapAnyUint16.GetOpt", t, func() {
		var k interface{} = "8e66771b-e492-45fd-829e-6be8759db227"
		var v uint16 = 41598

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("ee10bc01-59d1-4659-8acf-f82c68bf82a4" + "1e9c9042-b09c-40df-a2bf-ce4ffd35eae8")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapAnyUint16_ForEach(t *testing.T) {
	Convey("TestMapAnyUint16.ForEach", t, func() {
		var k interface{} = "2cd0115a-23fd-4f02-a806-87931adfa233"
		var v uint16 = 48303
		hits := 0

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk interface{}, vv uint16) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapAnyUint16_MarshalYAML(t *testing.T) {
	Convey("TestMapAnyUint16.MarshalYAML", t, func() {
		var k interface{} = "40183651-f71b-4656-b74f-92a91c8a0d06"
		var v uint16 = 51266

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapAnyUint16_ToYAML(t *testing.T) {
	Convey("TestMapAnyUint16.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "97b472b3-d327-4e73-87b5-fd0fa6c865d1"
			var v uint16 = 6315

			test := omap.NewMapAnyUint16(1)

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
			var k interface{} = "0c8db85e-5a3c-47b6-ab53-60f87711afd2"
			var v uint16 = 52935

			test := omap.NewMapAnyUint16(1)
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

func TestMapAnyUint16_PutIfNotNil(t *testing.T) {
	Convey("TestMapAnyUint16.PutIfNotNil", t, func() {
		var k interface{} = "a23ef753-eb0b-44dd-bfe3-eb3771ca7fa5"
		var v uint16 = 19700

		test := omap.NewMapAnyUint16(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("0bff99be-0bd2-48e7-9a03-95b39f649603", (*uint16)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint16 = 2929
		So(test.PutIfNotNil("25481207-697c-4676-a3c1-00ad9289eaac", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint16_ReplaceIfExists(t *testing.T) {
	Convey("TestMapAnyUint16.ReplaceIfExists", t, func() {
		var k interface{} = "96718d1d-082b-49d9-8c94-bff1404d8b8e"
		var v uint16 = 29597
		var x uint16 = 16655

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("6d51f98b-fd8d-4bec-be5f-1d2d5b2156d0", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint16_ReplaceOrPut(t *testing.T) {
	Convey("TestMapAnyUint16.ReplaceOrPut", t, func() {
		var k interface{} = "8ac63a9e-9dbe-4797-8309-9731b53fda1e"
		var v uint16 = 29027
		var x uint16 = 30604

		test := omap.NewMapAnyUint16(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("bf66b725-eb5d-47c7-998e-6af180e95084", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapAnyUint16_MarshalJSON(t *testing.T) {
	Convey("TestMapAnyUint16.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k interface{} = "06c9d67e-3b1a-4b15-ad75-41210569884c"
			var v uint16 = 26067

			test := omap.NewMapAnyUint16(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"06c9d67e-3b1a-4b15-ad75-41210569884c","value":26067}]`)
		})

		Convey("Unordered", func() {
			var k interface{} = "06c9d67e-3b1a-4b15-ad75-41210569884c"
			var v uint16 = 26067

			test := omap.NewMapAnyUint16(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"06c9d67e-3b1a-4b15-ad75-41210569884c":26067}`)
		})

	})
}
