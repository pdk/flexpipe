package flexpipe

import (
	"fmt"
	"strconv"
)

// Record is a flexible container of values, ala a dynamic struct.
type Record struct {
	names  map[string]int
	values []interface{}
}

// NewRecord constructs a new Record.
func NewRecord(names map[string]int, values []interface{}) Record {
	return Record{
		names:  names,
		values: values,
	}
}

// Names constructs a name index, suitable for passing to NewRecord.
func Names(names ...string) map[string]int {

	m := map[string]int{}

	for i, name := range names {
		m[name] = i
	}

	return m
}

// Values constructs a series of values, suitable for passing to NewRecord.
func Values(values ...interface{}) []interface{} {
	return values
}

// StringOf accesses the Record by name and returns a string value. If there is
// any problem, the second return value (ok) will be false.
func (r Record) StringOf(name string) (stringValue string, ok bool) {
	p, ok := r.names[name]
	if !ok {
		return "", false
	}

	return r.StringAt(p)
}

// StringAt returns the value at position p (0-based). Return value ok indicates
// if value could successfuly converted to a string.
func (r Record) StringAt(p int) (stringValue string, ok bool) {
	if p >= len(r.values) {
		return "", false
	}

	v := r.values[p]

	if s, ok := v.(string); ok {
		return s, true
	}

	if s, ok := v.(fmt.Stringer); ok {
		return s.String(), true
	}

	if i, ok := v.(int64); ok {
		return strconv.FormatInt(i, 10), true
	}

	return "", false
}

// Int64Of attempts to get the value by name as an int64. ok true/false
// indicates success/failure.
func (r Record) Int64Of(name string) (int64Value int64, ok bool) {
	p, ok := r.names[name]
	if !ok {
		return 0, false
	}

	return r.Int64At(p)
}

// Int64At attempts to get the value by position as an int64. ok true/false
// indicates success/failure.
func (r Record) Int64At(p int) (int64Value int64, ok bool) {
	if p >= len(r.values) {
		return 0, false
	}

	v := r.values[p]

	if i, ok := v.(int64); ok {
		return i, true
	}

	if i, ok := v.(int); ok {
		return int64(i), true
	}

	if s, ok := v.(string); ok {
		return int64ValueOf(s)
	}

	if s, ok := v.(fmt.Stringer); ok {
		return int64ValueOf(s.String())
	}

	return 0, false
}

// int64ValueOf attempts to convert a string to an int64. If there is any
// problem, we return ok value false. (We don't care exactly what the error
// was.)
func int64ValueOf(v string) (int64Value int64, ok bool) {
	i, err := strconv.ParseInt(v, 10, 64)
	return i, err == nil
}
