package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapStringUint_Put(t *testing.T) {
	Convey("TestMapStringUint.Put", t, func() {
		var k string = "a479bd46-eb8e-4874-ab29-9a68304e5f10"
		var v uint = 2080604000

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapStringUint_Delete(t *testing.T) {
	Convey("TestMapStringUint.Delete", t, func() {
		var k string = "5fb507bb-2b61-4e15-a6a5-0776afe8f6eb"
		var v uint = 4187562813

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapStringUint_Has(t *testing.T) {
	Convey("TestMapStringUint.Has", t, func() {
		var k string = "72b547b6-c3b0-4334-a12d-61a0beb96740"
		var v uint = 356336148

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has("25e6eeba-c370-4275-8493-a181789cceda"+"85c5b16b-f9f4-46b6-887b-1827565fea56"), ShouldBeFalse)
	})
}

func TestMapStringUint_Get(t *testing.T) {
	Convey("TestMapStringUint.Get", t, func() {
		var k string = "cacc53fb-c46d-41eb-8a05-d481e7e4c299"
		var v uint = 2716084063

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get("3eb3d1aa-2b90-4091-ae4b-2e4372eeb5b3" + "4faade5b-63ba-47c4-863a-68419a4b65c0")
		So(b, ShouldBeFalse)
	})
}

func TestMapStringUint_GetOpt(t *testing.T) {
	Convey("TestMapStringUint.GetOpt", t, func() {
		var k string = "3a73df7f-54c0-4006-be81-ef7f4234eb07"
		var v uint = 2528057157

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt("44705c37-6d31-4730-8c42-bfd2ab673681" + "4c8bc68f-ce86-4977-bfed-cc610544f014")
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapStringUint_ForEach(t *testing.T) {
	Convey("TestMapStringUint.ForEach", t, func() {
		var k string = "8fadbb73-a944-4ae6-ad9f-f86b44a6cba3"
		var v uint = 3103106757
		hits := 0

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk string, vv uint) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapStringUint_MarshalYAML(t *testing.T) {
	Convey("TestMapStringUint.MarshalYAML", t, func() {
		var k string = "b1c2cac8-3fb0-4362-b8a3-b6cfa442d3b2"
		var v uint = 281039095

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapStringUint_ToYAML(t *testing.T) {
	Convey("TestMapStringUint.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k string = "9a47aa99-e97d-486a-94bb-70fc28351041"
			var v uint = 3257355385

			test := omap.NewMapStringUint(1)

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
			var k string = "92b3717e-c320-4dc7-b18a-e55a235a7f7f"
			var v uint = 2103133142

			test := omap.NewMapStringUint(1)
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

func TestMapStringUint_PutIfNotNil(t *testing.T) {
	Convey("TestMapStringUint.PutIfNotNil", t, func() {
		var k string = "0144e481-78ef-4a5d-a8e8-c153da45d963"
		var v uint = 598638267

		test := omap.NewMapStringUint(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil("0109dca2-70ee-4b3e-9e72-0ce7a6e43f9d", (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint = 1613552369
		So(test.PutIfNotNil("e2971d94-2638-45ab-8eb8-0f5ca30a1625", &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapStringUint_ReplaceIfExists(t *testing.T) {
	Convey("TestMapStringUint.ReplaceIfExists", t, func() {
		var k string = "a744d535-d78a-470a-acf9-9c884b03c15b"
		var v uint = 191856731
		var x uint = 3426264277

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists("72549cbb-b250-4fc4-ad94-26b49c2f8d58", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint_ReplaceOrPut(t *testing.T) {
	Convey("TestMapStringUint.ReplaceOrPut", t, func() {
		var k string = "0757e306-e945-4e1e-a0ec-6fb2eb59ca4d"
		var v uint = 3106720834
		var x uint = 4232907309

		test := omap.NewMapStringUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut("cf64e441-82c5-4642-a6d2-88034119cccc", x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapStringUint_MarshalJSON(t *testing.T) {
	Convey("TestMapStringUint.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k string = "acdff384-c14b-42b6-ac82-4614687fecdd"
			var v uint = 4247695067

			test := omap.NewMapStringUint(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":"acdff384-c14b-42b6-ac82-4614687fecdd","value":4247695067}]`)
		})

		Convey("Unordered", func() {
			var k string = "acdff384-c14b-42b6-ac82-4614687fecdd"
			var v uint = 4247695067

			test := omap.NewMapStringUint(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"acdff384-c14b-42b6-ac82-4614687fecdd":4247695067}`)
		})

	})
}
