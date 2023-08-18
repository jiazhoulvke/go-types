package types

import (
	"encoding/json"
	"testing"

	"github.com/glebarez/sqlite"
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/gorm"
)

func TestMap(t *testing.T) {
	Convey("Map", t, func() {
		type mm struct {
			Id int64            `json:"id"`
			M  Map[string, any] `json:"m"`
		}
		m := mm{
			Id: 1,
			M:  Map[string, any]{"foo": 1},
		}
		b, err := json.Marshal(m)
		So(err, ShouldBeNil)
		So(string(b), ShouldEqual, `{"id":1,"m":{"foo":1}}`)
		db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		So(err, ShouldBeNil)
		err = db.AutoMigrate(&m)
		So(err, ShouldBeNil)
		err = db.Create(&m).Error
		So(err, ShouldBeNil)
		var m1 mm
		err = db.Take(&m1, 1).Error
		So(err, ShouldBeNil)
		So(m1.M["foo"], ShouldEqual, 1)
	})
}
