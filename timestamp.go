package types

import (
	"database/sql/driver"
	"errors"
	"strconv"
	"strings"
	"time"
)

type Timestamp int64

func TimestampOf(n int64) Timestamp {
	return Timestamp(n)
}

func (t Timestamp) Time() time.Time {
	return time.Unix(int64(t), 0)
}

func (t Timestamp) Timestamp() int64 {
	return int64(t)
}

func (t *Timestamp) SetTime(v time.Time) {
	*t = Timestamp(v.Unix())
}

func (t *Timestamp) Set(n int64) {
	*t = Timestamp(n)
}

func (t Timestamp) String() string {
	return time.Unix(int64(t), 0).Format("2006-01-02 15:04:05")
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	n, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*t = Timestamp(n)
	return nil
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(t), 10)), nil
}

func (t *Timestamp) Scan(src any) error {
	tt, ok := src.(int64)
	if !ok {
		return errors.New("data format error")
	}
	*t = Timestamp(tt)
	return nil
}

func (t Timestamp) Value() (driver.Value, error) {
	return int64(t), nil
}

type NullTimestamp struct {
	valid bool
	t     int64
}

func NullTimestampOf(n int64) NullTimestamp {
	return NullTimestamp{
		valid: true,
		t:     n,
	}
}

func (t NullTimestamp) Valid() bool {
	return t.valid
}

func (t NullTimestamp) Time() time.Time {
	return time.Unix(t.t, 0)
}

func (t NullTimestamp) Timestamp() int64 {
	return t.t
}

func (t *NullTimestamp) SetTime(v time.Time) {
	*t = NullTimestamp{
		valid: true,
		t:     v.Unix(),
	}
}

func (t *NullTimestamp) Set(n int64) {
	*t = NullTimestamp{
		valid: true,
		t:     n,
	}
}

func (t NullTimestamp) String() string {
	if !t.valid {
		return "null"
	}
	return time.Unix(t.t, 0).Format("2006-01-02 15:04:05")
}

func (t *NullTimestamp) UnmarshalJSON(data []byte) error {
	if strings.ToLower(string(data)) == "null" {
		*t = NullTimestamp{}
		return nil
	}
	n, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*t = NullTimestamp{
		valid: true,
		t:     n,
	}
	return nil
}

func (t NullTimestamp) MarshalJSON() ([]byte, error) {
	if !t.valid {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatInt(t.t, 10)), nil
}

func (t *NullTimestamp) Scan(src any) error {
	if src == nil {
		*t = NullTimestamp{}
		return nil
	}
	tt, ok := src.(int64)
	if !ok {
		return errors.New("data format error")
	}
	if tt < 0 {
		*t = NullTimestamp{}
		return nil
	}
	*t = NullTimestamp{
		valid: true,
		t:     tt,
	}
	return nil
}

func (t NullTimestamp) Value() (driver.Value, error) {
	if !t.valid {
		return nil, nil
	}
	return t.t, nil
}
