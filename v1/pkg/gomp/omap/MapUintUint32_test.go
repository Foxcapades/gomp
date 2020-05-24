package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapUintUint32_Put(t *testing.T) {
	Convey("TestMapUintUint32.Put", t, func() {
		var k uint = 2822977630
		var v uint32 = 2557955988

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapUintUint32_Delete(t *testing.T) {
	Convey("TestMapUintUint32.Delete", t, func() {
		var k uint = 863666513
		var v uint32 = 2501434832

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapUintUint32_Has(t *testing.T) {
	Convey("TestMapUintUint32.Has", t, func() {
		var k uint = 2220674204
		var v uint32 = 1987606773

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(3943918508+1778980808), ShouldBeFalse)
	})
}


func TestMapUintUint32_Get(t *testing.T) {
	Convey("TestMapUintUint32.Get", t, func() {
		var k uint = 3498405224
		var v uint32 = 396684559

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(203463060 + 4124729119)
		So(b, ShouldBeFalse)
	})
}

func TestMapUintUint32_GetOpt(t *testing.T) {
	Convey("TestMapUintUint32.GetOpt", t, func() {
		var k uint = 1264034180
		var v uint32 = 275697827

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(2062520825 + 4196958764)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapUintUint32_ForEach(t *testing.T) {
	Convey("TestMapUintUint32.ForEach", t, func() {
		var k uint = 1930288298
		var v uint32 = 1901376613
		hits := 0

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk uint, vv uint32) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapUintUint32_MarshalYAML(t *testing.T) {
	Convey("TestMapUintUint32.MarshalYAML", t, func() {
		var k uint = 3919541651
		var v uint32 = 2963874745

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapUintUint32_ToYAML(t *testing.T) {
	Convey("TestMapUintUint32.ToYAML", t, func() {
		var k uint = 2269131907
		var v uint32 = 98749767

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapUintUint32_PutIfNotNil(t *testing.T) {
	Convey("TestMapUintUint32.PutIfNotNil", t, func() {
		var k uint = 1801286089
		var v uint32 = 3648969187

		test := omap.NewMapUintUint32(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(2333093387, (*uint32)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x uint32 = 1144873722
		So(test.PutIfNotNil(2629552536, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapUintUint32_ReplaceIfExists(t *testing.T) {
	Convey("TestMapUintUint32.ReplaceIfExists", t, func() {
		var k uint = 3456390803
		var v uint32 = 202506611
		var x uint32 = 1737594305

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(1535116534, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint32_ReplaceOrPut(t *testing.T) {
	Convey("TestMapUintUint32.ReplaceOrPut", t, func() {
		var k uint = 1574497299
		var v uint32 = 3259709620
		var x uint32 = 472151759

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(3396938747, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapUintUint32_MarshalJSON(t *testing.T) {
	Convey("TestMapUintUint32.MarshalJSON", t, func() {
		var k uint = 1856343529
		var v uint32 = 1951267038

		test := omap.NewMapUintUint32(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1856343529,"value":1951267038}]`)
	})
}
