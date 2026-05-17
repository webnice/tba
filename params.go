package tgbotapi

import (
	"encoding/json"
	"reflect"
	"strconv"
)

// Params represents a set of parameters that gets passed to a request.
type Params map[string]string

// AddNonEmpty adds a value if it not an empty string.
func (p *Params) AddNonEmpty(key, value string) {
	if value != "" {
		(*p)[key] = value
	}
}

// AddNonZeroInteger is the same as AddNonZero except uses an int64.
func (p *Params) AddNonZeroInteger(key string, value int64) {
	if value != 0 {
		(*p)[key] = strconv.FormatInt(value, 10)
	}
}

// AddBool adds a value of a bool if it is true.
func (p *Params) AddBool(key string, value bool) {
	if value {
		(*p)[key] = strconv.FormatBool(value)
	}
}

// AddNonZeroFloat adds a floating point value that is not zero.
func (p *Params) AddNonZeroFloat(key string, value float64) {
	if value != 0 {
		(*p)[key] = strconv.FormatFloat(value, 'f', 6, 64)
	}
}

// AddInterface adds an interface if it is not nil and can be JSON marshaled.
func (p *Params) AddInterface(key string, value interface{}) error {
	if value == nil || (reflect.ValueOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil()) {
		return nil
	}

	b, err := json.Marshal(value)
	if err != nil {
		return err
	}

	(*p)[key] = string(b)

	return nil
}

// AddFirstValid attempts to add the first item that is not a default value.
//
// For example, AddFirstValid(0, "", "test") would add "test".
func (p *Params) AddFirstValid(key string, args ...any) (err error) {
	var (
		buf []byte
		arg any
	)
	for _, arg = range args {
		switch v := arg.(type) {
		case int:
			if v != 0 {
				(*p)[key] = strconv.Itoa(v)
				return
			}
		case int64:
			if v != 0 {
				(*p)[key] = strconv.FormatInt(v, 10)
				return
			}
		case string:
			if v != "" {
				(*p)[key] = v
				return
			}
		case nil:
		default:
			if buf, err = json.Marshal(arg); err != nil {
				return
			}
			(*p)[key] = string(buf)
			return
		}
	}

	return
}

// AddNonZeroSliceString Добавление не пустых срезов строк.
func (p *Params) AddNonZeroSliceString(key string, args ...[]string) (err error) {
	var (
		buf []string
		out []byte
		n   int
	)

	for n = range args {
		buf = append(buf, args[n]...)
	}
	if len(buf) <= 0 {
		return
	}
	if out, err = json.Marshal(buf); err != nil {
		return
	}
	(*p)[key] = string(out)

	return
}

// AddNonZeroSliceInt64 Добавление не пустых срезов чисел.
func (p *Params) AddNonZeroSliceInt64(key string, args ...[]int64) (err error) {
	var (
		buf []int64
		out []byte
		n   int
	)

	for n = range args {
		buf = append(buf, args[n]...)
	}
	if len(buf) <= 0 {
		return
	}
	if out, err = json.Marshal(buf); err != nil {
		return
	}
	(*p)[key] = string(out)

	return
}

// Merge merges two sets of parameters. Overwrites old fields if present
func (p *Params) Merge(p1 Params) {
	for k, v := range p1 {
		(*p)[k] = v
	}
}
