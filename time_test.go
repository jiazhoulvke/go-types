package types

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/glebarez/sqlite"
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/gorm"
)

func TestTime(t *testing.T) {
	Convey("Time", t, func() {
		t1 := Time(time.Unix(1234, 0))
		b, err := json.Marshal(t1)
		So(err, ShouldBeNil)
		So(string(b), ShouldEqual, "1234")
		type ss struct {
			Id int64
			T  Time `json:"t"`
		}
		var s1 ss
		err = json.Unmarshal([]byte(`{"t":1234}`), &s1)
		So(err, ShouldBeNil)
		So(s1.T.Timestamp(), ShouldEqual, 1234)
		db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		So(err, ShouldBeNil)
		err = db.AutoMigrate(&s1)
		So(err, ShouldBeNil)
		So(db.Create(&s1).Error, ShouldBeNil)
		var s2 ss
		So(db.Take(&s2, 1).Error, ShouldBeNil)
		So(s2.T.Timestamp(), ShouldEqual, 1234)
	})
}

func TestNullTime(t *testing.T) {
	Convey("NullTime", t, func() {
		t1 := NullTimeOf(time.Unix(1234, 0))
		b, err := json.Marshal(t1)
		So(err, ShouldBeNil)
		So(string(b), ShouldEqual, "1234")
		type ss struct {
			Id int64
			T  NullTime `json:"t"`
		}
		var s1 ss
		err = json.Unmarshal([]byte(`{"t":null}`), &s1)
		So(err, ShouldBeNil)
		err = json.Unmarshal([]byte("{}"), &s1)
		So(err, ShouldBeNil)
		err = json.Unmarshal([]byte(`{"t":1234}`), &s1)
		So(err, ShouldBeNil)
		So(s1.T.Timestamp(), ShouldEqual, 1234)
		db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		So(err, ShouldBeNil)
		err = db.AutoMigrate(&s1)
		So(err, ShouldBeNil)
		So(db.Create(&s1).Error, ShouldBeNil)
		var s2 ss
		So(db.Take(&s2, 1).Error, ShouldBeNil)
		So(s2.T.Timestamp(), ShouldEqual, 1234)
	})
}
