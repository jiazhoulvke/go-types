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

func TimestampNow() Timestamp {
	return Timestamp(time.Now().Unix())
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
	switch v := src.(type) {
	case []byte:
		tt, err := strToTime(string(v))
		if err != nil {
			return err
		}
		*t = Timestamp(tt.Unix())
		return nil
	case string:
		tt, err := strToTime(v)
		if err != nil {
			return err
		}
		*t = Timestamp(tt.Unix())
		return nil
	case time.Time:
		*t = Timestamp(v.Unix())
		return nil
	case int64:
		*t = Timestamp(v)
		return nil
	}
	return errors.New("data format error")
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

func NullTimestampNow() NullTimestamp {
	return NullTimestamp{
		valid: true,
		t:     time.Now().Unix(),
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
	switch v := src.(type) {
	case []byte:
		tt, err := strToTime(string(v))
		if err != nil {
			return err
		}
		*t = NullTimestampOf(tt.Unix())
		return nil
	case string:
		tt, err := strToTime(v)
		if err != nil {
			return err
		}
		*t = NullTimestampOf(tt.Unix())
		return nil
	case time.Time:
		*t = NullTimestampOf(v.Unix())
		return nil
	case int64:
		*t = NullTimestampOf(v)
		return nil
	}
	return errors.New("data format error")
}

func (t NullTimestamp) Value() (driver.Value, error) {
	if !t.valid {
		return nil, nil
	}
	return t.t, nil
}
