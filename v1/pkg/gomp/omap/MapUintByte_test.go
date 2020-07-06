package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintByte_Put(t *testing.T) {
	Convey("TestMapUintByte.Put", t, func() {
		var k uint = 3730392604
		var v byte = 29

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintByte_Delete(t *testing.T) {
	Convey("TestMapUintByte.Delete", t, func() {
		var k uint = 461868688
		var v byte = 69

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintByte_Has(t *testing.T) {
	Convey("TestMapUintByte.Has", t, func() {
		var k uint = 1723464370
		var v byte = 35

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(3989314449+4280962888), ShouldBeFalse)
	})
}

func TestMapUintByte_Get(t *testing.T) {
	Convey("TestMapUintByte.Get", t, func() {
		var k uint = 2758799903
		var v byte = 87

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		_, b = test.Get(2823221282 + 99809012)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintByte_GetOpt(t *testing.T) {
	Convey("TestMapUintByte.GetOpt", t, func() {
		var k uint = 2619882204
		var v byte = 26

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(2430807918 + 1475282605)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintByte_ForEach(t *testing.T) {
	Convey("TestMapUintByte.ForEach", t, func() {
		var k uint = 3025137739
		var v byte = 63
		hits := 0

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv byte) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintByte_MarshalYAML(t *testing.T) {
	Convey("TestMapUintByte.MarshalYAML", t, func() {
		var k uint = 1670481506
		var v byte = 12

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintByte_ToYAML(t *testing.T) {
	Convey("TestMapUintByte.ToYAML", t, func() {
		Convey("Ordered", func() {
			var k uint = 674268476
			var v byte = 245

			test := omap.NewMapUintByte(1)

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
			var k uint = 474385857
			var v byte = 66

			test := omap.NewMapUintByte(1)
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

func TestMapUintByte_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintByte.PutIfNotNil", t, func() {
		var k uint = 2914726553
		var v byte = 130

		test := omap.NewMapUintByte(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(3835267113, (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x byte = 233
		So(test.PutIfNotNil(1513073580, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintByte_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintByte.ReplaceIfExists", t, func() {
		var k uint = 4044804554
		var v byte = 226
		var x byte = 103

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(3237315667, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintByte_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintByte.ReplaceOrPut", t, func() {
		var k uint = 2425489249
		var v byte = 196
		var x byte = 130

		test := omap.NewMapUintByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(4216928327, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintByte_MarshalJSON(t *testing.T) {
	Convey("TestMapUintByte.MarshalJSON", t, func() {
		Convey("Ordered", func() {
			var k uint = 1183922559
			var v byte = 123

			test := omap.NewMapUintByte(1)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `[{"key":1183922559,"value":123}]`)
		})

		Convey("Unordered", func() {
			var k uint = 1183922559
			var v byte = 123

			test := omap.NewMapUintByte(1)
			test.SerializeOrdered(false)

			So(test.Put(k, v), ShouldPointTo, test)
			So(test.Len(), ShouldEqual, 1)

			a, b := test.MarshalJSON()
			So(b, ShouldBeNil)
			So(string(a), ShouldEqual, `{"1183922559":123}`)
		})

	})
}
