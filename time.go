package types

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"strconv"
	"time"
)

type Time time.Time

func TimeOf(v time.Time) Time {
	return Time(v)
}

func (t Time) Time() time.Time {
	return time.Time(t)
}

func (t Time) Timestamp() int64 {
	return t.Time().Unix()
}

func (t *Time) Set(v time.Time) {
	*t = Time(v)
}

func (t Time) String() string {
	return t.Time().Format("2006-01-02 15:04:05")
}

func (t *Time) UnmarshalJSON(data []byte) error {
	n, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*t = Time(time.Unix(n, 0))
	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(t.Time().Unix(), 10)), nil
}

func (t *Time) Scan(src any) error {
	tt, ok := src.(time.Time)
	if !ok {
		return errors.New("data format error")
	}
	*t = Time(tt)
	return nil
}

func (t Time) Value() (driver.Value, error) {
	tt := time.Time(t)
	return tt, nil
}

type NullTime struct {
	t time.Time
	v bool
}

func NullTimeOf(v time.Time) NullTime {
	var t NullTime
	t.Set(v)
	return t
}

func (t *NullTime) Set(v time.Time) {
	t.t = v
	t.v = true
}

func (t *NullTime) Reset() {
	t.t = time.Unix(0, 0)
	t.v = false
}

func (t NullTime) Time() time.Time {
	return t.t
}

func (t NullTime) Timestamp() int64 {
	return t.Time().Unix()
}

func (t *NullTime) Valid() bool {
	return t.v
}

var timestampFormats = []string{
	"2006-01-02 15:04:05.999999999-07:00",
	"2006-01-02T15:04:05.999999999-07:00",
	"2006-01-02 15:04:05.999999999",
	"2006-01-02T15:04:05.999999999",
	"2006-01-02 15:04:05",
	"2006-01-02T15:04:05",
	"2006-01-02 15:04",
	"2006-01-02T15:04",
	"2006-01-02",
	"2006/01/02 15:04:05",
}

func (t *NullTime) Scan(value any) error {
	t.t, t.v = value.(time.Time)
	if t.v {
		return nil
	}
	var ns sql.NullString
	if err := ns.Scan(value); err != nil {
		return err
	}
	if !ns.Valid {
		return nil
	}
	for _, tf := range timestampFormats {
		if tt, err := time.Parse(tf, ns.String); err == nil {
			t.t = tt
			t.v = true
			return nil
		}
	}
	return nil
}

func (t NullTime) Value() (driver.Value, error) {
	if !t.Valid() {
		return nil, nil
	}
	return t.t, nil
}

func (t NullTime) String() string {
	if !t.v {
		return ""
	}
	return t.t.Format("2006-01-02 15:04:05")
}

func (t NullTime) MarshalJSON() ([]byte, error) {
	if !t.v {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatInt(time.Time(t.t).Unix(), 10)), nil
}

func (t *NullTime) UnmarshalJSON(data []byte) error {
	n, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	a := time.Unix(int64(n), 0)
	tt := NullTime{
		t: a,
		v: n > 0,
	}
	*t = tt
	return nil
}
