package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntByte_Put(t *testing.T) {
	Convey("TestMapIntByte.Put", t, func() {
		var k int = 1667158397
		var v byte = 245

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntByte_Delete(t *testing.T) {
	Convey("TestMapIntByte.Delete", t, func() {
		var k int = 1501360596
		var v byte = 139

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntByte_Has(t *testing.T) {
	Convey("TestMapIntByte.Has", t, func() {
		var k int = 892919563
		var v byte = 182

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(512263112+256342522), ShouldBeFalse)
	})
}


func TestMapIntByte_Get(t *testing.T) {
	Convey("TestMapIntByte.Get", t, func() {
		var k int = 818396241
		var v byte = 187

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(1583207789+1856878969)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntByte_GetOpt(t *testing.T) {
	Convey("TestMapIntByte.GetOpt", t, func() {
		var k int = 1868841418
		var v byte = 133

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(1548082856+1450508305)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntByte_ForEach(t *testing.T) {
	Convey("TestMapIntByte.ForEach", t, func() {
		var k int = 1520847786
		var v byte = 121
		hits := 0

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv byte) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntByte_MarshalYAML(t *testing.T) {
	Convey("TestMapIntByte.MarshalYAML", t, func() {
		var k int = 405336843
		var v byte = 208

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntByte_ToYAML(t *testing.T) {
	Convey("TestMapIntByte.ToYAML", t, func() {
		var k int = 1484612887
		var v byte = 168

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntByte_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntByte.PutIfNotNil", t, func() {
		var k int = 1450562656
		var v byte = 30

		test := omap.NewMapIntByte(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(169543620, (*byte)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x byte = 145
		So(test.PutIfNotNil(2139147250, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntByte_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntByte.ReplaceIfExists", t, func() {
		var k int = 2082416836
		var v byte = 42
		var x byte = 84

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(688836600, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntByte_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntByte.ReplaceOrPut", t, func() {
		var k int = 302870501
		var v byte = 217
		var x byte = 173

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(902906577, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntByte_MarshalJSON(t *testing.T) {
	Convey("TestMapIntByte.MarshalJSON", t, func() {
		var k int = 1886987058
		var v byte = 121

		test := omap.NewMapIntByte(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1886987058,"value":121}]`)
	})
}

