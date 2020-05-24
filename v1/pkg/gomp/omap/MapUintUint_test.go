package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintUint_Put(t *testing.T) {
	Convey("TestMapUintUint.Put", t, func() {
		var k uint = 2691231150
		var v uint = 1705811775

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintUint_Delete(t *testing.T) {
	Convey("TestMapUintUint.Delete", t, func() {
		var k uint = 300404557
		var v uint = 1084272573

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintUint_Has(t *testing.T) {
	Convey("TestMapUintUint.Has", t, func() {
		var k uint = 2278519221
		var v uint = 4094020200

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(838308395+4149398527), ShouldBeFalse)
	})
}


func TestMapUintUint_Get(t *testing.T) {
	Convey("TestMapUintUint.Get", t, func() {
		var k uint = 3029561488
		var v uint = 4136646913

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(3950870811+3792043745)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintUint_GetOpt(t *testing.T) {
	Convey("TestMapUintUint.GetOpt", t, func() {
		var k uint = 4281709939
		var v uint = 2106934652

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(2627202181+3006487686)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintUint_ForEach(t *testing.T) {
	Convey("TestMapUintUint.ForEach", t, func() {
		var k uint = 492146059
		var v uint = 1380871328
		hits := 0

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv uint) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintUint_MarshalYAML(t *testing.T) {
	Convey("TestMapUintUint.MarshalYAML", t, func() {
		var k uint = 1321828555
		var v uint = 2374220982

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintUint_ToYAML(t *testing.T) {
	Convey("TestMapUintUint.ToYAML", t, func() {
		var k uint = 4084362382
		var v uint = 1065522991

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintUint_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintUint.PutIfNotNil", t, func() {
		var k uint = 3993326573
		var v uint = 1011478546

		test := omap.NewMapUintUint(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1585558984, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint = 3605826262
		So(test.PutIfNotNil(2457225146, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintUint_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintUint.ReplaceIfExists", t, func() {
		var k uint = 1477105224
		var v uint = 1340913351
		var x uint = 3642381606

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(270509632, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintUint.ReplaceOrPut", t, func() {
		var k uint = 162523868
		var v uint = 516896398
		var x uint = 4007389509

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1154063565, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint_MarshalJSON(t *testing.T) {
	Convey("TestMapUintUint.MarshalJSON", t, func() {
		var k uint = 1365632013
		var v uint = 4194407668

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1365632013,"value":4194407668}]`)
	})
}

