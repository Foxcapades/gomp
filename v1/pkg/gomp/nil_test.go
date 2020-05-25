package gomp_test

import (
	"github.com/Foxcapades/gomp/v1/pkg/gomp"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDeref(t *testing.T) {
	Convey("Deref", t, func() {
		So(gomp.Deref(3), ShouldEqual, 3)
		So(gomp.Deref([]string{"hi"}), ShouldResemble, []string{"hi"})
		So(gomp.Deref(map[string]string{"hi": "bye"}), ShouldResemble,
			map[string]string{"hi": "bye"})
		foo := 666
		So(gomp.Deref(&foo), ShouldEqual, 666)
	})
}
