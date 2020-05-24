package omap_test

import (
	"testing"

	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/yaml.v3"

	"github.com/Foxcapades/gomp/v1/pkg/gomp/omap"
)

func TestMapIntInt_Put(t *testing.T) {
	Convey("TestMapIntInt.Put", t, func() {
		var k int = 287704894
		var v int = 328884768

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Key, ShouldEqual, k)
		So(test.At(0).Val, ShouldEqual, v)
	})
}

func TestMapIntInt_Delete(t *testing.T) {
	Convey("TestMapIntInt.Delete", t, func() {
		var k int = 61260044
		var v int = 847764718

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Delete(k), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 0)
	})
}

func TestMapIntInt_Has(t *testing.T) {
	Convey("TestMapIntInt.Has", t, func() {
		var k int = 1509513363
		var v int = 1882953600

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.Has(k), ShouldBeTrue)
		So(test.Has(229773555+951604515), ShouldBeFalse)
	})
}


func TestMapIntInt_Get(t *testing.T) {
	Convey("TestMapIntInt.Get", t, func() {
		var k int = 1431139235
		var v int = 1759153471

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.Get(k)
		So(b, ShouldBeTrue)
		So(a, ShouldEqual, v)

		a, b = test.Get(692215161 + 1008046602)
		So(b, ShouldBeFalse)
	})
}

func TestMapIntInt_GetOpt(t *testing.T) {
	Convey("TestMapIntInt.GetOpt", t, func() {
		var k int = 327223700
		var v int = 1658636838

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a := test.GetOpt(k)
		So(a.IsPresent(), ShouldBeTrue)
		So(a.Get(), ShouldEqual, v)

		a = test.GetOpt(873573441 + 2097487185)
		So(a.IsNil(), ShouldBeTrue)
	})
}

func TestMapIntInt_ForEach(t *testing.T) {
	Convey("TestMapIntInt.ForEach", t, func() {
		var k int = 116431822
		var v int = 932227246
		hits := 0

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ForEach(func(kk int, vv int) {
			So(kk, ShouldEqual, k)
			So(vv, ShouldEqual, v)
			hits++
		}), ShouldPointTo, test)
		So(hits, ShouldEqual, 1)
	})
}

func TestMapIntInt_MarshalYAML(t *testing.T) {
	Convey("TestMapIntInt.MarshalYAML", t, func() {
		var k int = 172493801
		var v int = 159282929

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalYAML()
		So(b, ShouldBeNil)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)

		So(a, ShouldResemble, c)
	})
}

func TestMapIntInt_ToYAML(t *testing.T) {
	Convey("TestMapIntInt.ToYAML", t, func() {
		var k int = 2080025841
		var v int = 1128452475

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		c, d := test.ToYAML()
		So(d, ShouldBeNil)
		So(c.Kind, ShouldEqual, yaml.MappingNode)
		So(c.Tag, ShouldEqual, xyml.TagOrderedMap)
		So(len(c.Content), ShouldEqual, 2)
	})
}

func TestMapIntInt_PutIfNotNil(t *testing.T) {
	Convey("TestMapIntInt.PutIfNotNil", t, func() {
		var k int = 500584896
		var v int = 380225100

		test := omap.NewMapIntInt(1)

		So(test.PutIfNotNil(k, &v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.PutIfNotNil(k, (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.PutIfNotNil(1850847291, (*int)(nil)), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		var x int = 454203990
		So(test.PutIfNotNil(1132252158, &x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)
	})
}

func TestMapIntInt_ReplaceIfExists(t *testing.T) {
	Convey("TestMapIntInt.ReplaceIfExists", t, func() {
		var k int = 1398623570
		var v int = 941761902
		var x int = 467466823

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceIfExists(668436384, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, v)

		So(test.ReplaceIfExists(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt_ReplaceOrPut(t *testing.T) {
	Convey("TestMapIntInt.ReplaceOrPut", t, func() {
		var k int = 433377888
		var v int = 1408442547
		var x int = 411044851

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		So(test.ReplaceOrPut(1713388809, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(1).Val, ShouldEqual, x)

		So(test.ReplaceOrPut(k, x), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 2)
		So(test.At(0).Val, ShouldEqual, x)
	})
}

func TestMapIntInt_MarshalJSON(t *testing.T) {
	Convey("TestMapIntInt.MarshalJSON", t, func() {
		var k int = 1488703509
		var v int = 523995787

		test := omap.NewMapIntInt(1)

		So(test.Put(k, v), ShouldPointTo, test)
		So(test.Len(), ShouldEqual, 1)

		a, b := test.MarshalJSON()
		So(b, ShouldBeNil)
		So(string(a), ShouldEqual, `[{"key":1488703509,"value":523995787}]`)
	})
}
