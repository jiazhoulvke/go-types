package types

import (
	"encoding/json"
	"testing"

	"github.com/glebarez/sqlite"
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/gorm"
)

func TestNumericSlice(t *testing.T) {
	Convey("NumericSlice", t, func() {
		type ss struct {
			Id int64                 `json:"id"`
			Ns NumericSlice[float64] `json:"ns"`
		}
		s1 := ss{
			Id: 1,
			Ns: []float64{1.23, 3.14},
		}
		b, err := json.Marshal(s1)
		So(err, ShouldBeNil)
		So(string(b), ShouldEqual, `{"id":1,"ns":[1.23,3.14]}`)
		db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		So(err, ShouldBeNil)
		err = db.AutoMigrate(&s1)
		So(err, ShouldBeNil)
		err = db.Create(&s1).Error
		So(err, ShouldBeNil)
		var s2 ss
		err = db.Take(&s2, 1).Error
		So(err, ShouldBeNil)
		So(s2.Ns, ShouldResemble, NumericSlice[float64]{1.23, 3.14})
	})
}
