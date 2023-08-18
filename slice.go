package types

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

// Numeric 数字类型
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type NumericSlice[T Numeric] []T

func (ss *NumericSlice[T]) Scan(value any) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("data is not bytes")
	}
	strSlice := strings.Split(str, ",")
	slice := make([]T, 0, len(strSlice))
	for _, s := range strSlice {
		if s == "" {
			continue
		}
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		slice = append(slice, T(f))
	}
	*ss = slice
	return nil
}

func (ss NumericSlice[T]) Value() (driver.Value, error) {
	slice := make([]string, 0, len(ss))
	for _, t := range ss {
		slice = append(slice, fmt.Sprint(t))
	}
	return strings.Join(slice, ","), nil
}

type StringSlice[T ~string] []T

func (ss *StringSlice[T]) Scan(value any) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("data is not bytes")
	}
	strSlice := strings.Split(str, ",")
	slice := make([]T, 0, len(strSlice))
	for _, s := range strSlice {
		slice = append(slice, T(s))
	}
	*ss = slice
	return nil
}

func (ss StringSlice[T]) Value() (driver.Value, error) {
	slice := make([]string, 0, len(ss))
	for _, t := range ss {
		slice = append(slice, fmt.Sprint(t))
	}
	return strings.Join(slice, ","), nil
}
