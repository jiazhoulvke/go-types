package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

type Map[T1 comparable, T2 any] map[T1]T2

func (m Map[T1, T2]) URLEncode() string {
	fields := make([]string, 0, len(m))
	for k, v := range m {
		fields = append(fields, fmt.Sprintf("%v=%v", k, v))
	}
	sort.Strings(fields)
	return strings.Join(fields, "&")
}

func (m Map[T1, T2]) Json() ([]byte, error) {
	j, err := json.Marshal(m)
	return j, err
}

func (m Map[T1, T2]) JsonString() (string, error) {
	j, err := json.Marshal(m)
	return string(j), err
}

func (d *Map[T1, T2]) Scan(value any) error {
	if value == nil {
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("data is not bytes")
	}
	var data Map[T1, T2]
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}
	*d = data
	return nil
}

func (d Map[T1, T2]) Value() (driver.Value, error) {
	return json.Marshal(d)
}
