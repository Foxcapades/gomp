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
		var k uint = 4007186405
		var v uint = 2401049689

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintUint_Delete(t *testing.T) {
	Convey("TestMapUintUint.Delete", t, func() {
		var k uint = 1885300117
		var v uint = 1271556924

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintUint_Has(t *testing.T) {
	Convey("TestMapUintUint.Has", t, func() {
		var k uint = 3095494602
		var v uint = 2156957644

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(3968175071+980491598), ShouldBeFalse)
	})
}


func TestMapUintUint_Get(t *testing.T) {
	Convey("TestMapUintUint.Get", t, func() {
		var k uint = 2954339815
		var v uint = 1855744327

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(3099079212 + 1019853387)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintUint_GetOpt(t *testing.T) {
	Convey("TestMapUintUint.GetOpt", t, func() {
		var k uint = 3620285296
		var v uint = 100144750

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(2943621831 + 4088036070)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintUint_ForEach(t *testing.T) {
	Convey("TestMapUintUint.ForEach", t, func() {
		var k uint = 1927832257
		var v uint = 112323497
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
		var k uint = 2114400318
		var v uint = 2241028232

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
		var k uint = 3339138547
		var v uint = 2769524475

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
		var k uint = 2134247852
		var v uint = 519683600

		test := omap.NewMapUintUint(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(3034937599, (*uint)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint = 4250565757
		So(test.PutIfNotNil(4106567216, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintUint_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintUint.ReplaceIfExists", t, func() {
		var k uint = 3298250556
		var v uint = 4160474270
		var x uint = 1130025621

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(3064819833, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintUint.ReplaceOrPut", t, func() {
		var k uint = 1910646104
		var v uint = 151130195
		var x uint = 474587133

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(2908832114, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint_MarshalJSON(t *testing.T) {
	Convey("TestMapUintUint.MarshalJSON", t, func() {
		var k uint = 4014887720
		var v uint = 3691998608

		test := omap.NewMapUintUint(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":4014887720,"value":3691998608}]`)
	})
}
