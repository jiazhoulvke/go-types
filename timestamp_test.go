package types

import (
	"encoding/json"
	"testing"

	"github.com/glebarez/sqlite"
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/gorm"
)

func TestTimestamp(t *testing.T) {
	Convey("Timestamp", t, func() {
		type ts struct {
			T Timestamp `json:"t" gorm:"column:t"`
		}

		db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		So(err, ShouldBeNil)
		err = db.AutoMigrate(&ts{})
		So(err, ShouldBeNil)

		t1 := Timestamp(1234)
		b, err := json.Marshal(t1)
		So(err, ShouldBeNil)
		So(string(b), ShouldEqual, "1234")
		var s1 ts
		err = json.Unmarshal([]byte(`{"t":1234}`), &s1)
		So(err, ShouldBeNil)
		So(s1.T.Timestamp(), ShouldEqual, 1234)
		So(db.Create(&s1).Error, ShouldBeNil)
		var s2 ts
		So(db.Take(&s2).Error, ShouldBeNil)
		So(s2.T.Timestamp(), ShouldEqual, 1234)
	})
}

func TestNullTimestamp(t *testing.T) {
	Convey("NullTimestamp", t, func() {
		type tsn struct {
			Id int64         `json:"id" gorm:"column:id"`
			T  NullTimestamp `json:"t" gorm:"column:t;type:int"`
		}

		db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		So(err, ShouldBeNil)
		err = db.AutoMigrate(&tsn{})
		So(err, ShouldBeNil)

		var t1 NullTimestamp
		b, err := json.Marshal(t1)
		So(err, ShouldBeNil)
		So(string(b), ShouldEqual, "null")

		var s1 tsn
		err = json.Unmarshal([]byte(`{"t":1234}`), &s1)
		So(err, ShouldBeNil)
		So(s1.T.Timestamp(), ShouldEqual, 1234)
		So(db.Create(&s1).Error, ShouldBeNil)
		So(s1.Id, ShouldEqual, 1)
		var s2 tsn
		So(db.Take(&s2).Error, ShouldBeNil)
		So(s2.T.Timestamp(), ShouldEqual, 1234)

		var v1 tsn
		err = json.Unmarshal([]byte(`{"t":null}`), &v1)
		So(err, ShouldBeNil)
		So(v1.T.Valid(), ShouldBeFalse)
		So(db.Create(&v1).Error, ShouldBeNil)
		So(v1.Id, ShouldEqual, 2)
		var v2 tsn
		So(db.Take(&v2, 2).Error, ShouldBeNil)
		So(v2.T.Valid(), ShouldBeFalse)
	})
}
