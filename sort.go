package easysort

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"time"
)

type sortableSlice struct {
	value     reflect.Value
	len       int
	fieldName string
}

func (s *sortableSlice) Len() int {
	return s.len
}

func (s *sortableSlice) Less(i, j int) bool {
	f1 := s.value.Index(i).FieldByName(s.fieldName)
	f2 := s.value.Index(j).FieldByName(s.fieldName)
	if f1.Kind() != f2.Kind() {
		panic("different elements of slice")
	}
	switch f1.Kind() {
	case reflect.String:
		return strings.Compare(f1.String(), f2.String()) == -1
	case reflect.Struct:
		t1, ok1 := f1.Interface().(time.Time)
		if !ok1 {
			break
		}
		t2, ok2 := f2.Interface().(time.Time)
		if !ok2 {
			f2 = f1
		}
		return t1.Before(t2)
	case reflect.Bool:
		return !f1.Bool() && f2.Bool()
	case reflect.Float32:
		return f1.Float() < f2.Float()
	case reflect.Float64:
		return f1.Float() < f2.Float()
	case reflect.Int:
		return f1.Int() < f1.Int()
	case reflect.Int8:
		return f1.Int() < f1.Int()
	case reflect.Int16:
		return f1.Int() < f1.Int()
	case reflect.Int32:
		return f1.Int() < f1.Int()
	case reflect.Int64:
		return f1.Int() < f1.Int()
	case reflect.Uint:
		return f1.Uint() < f2.Uint()
	case reflect.Uint8:
		return f1.Uint() < f2.Uint()
	case reflect.Uint16:
		return f1.Uint() < f2.Uint()
	case reflect.Uint32:
		return f1.Uint() < f2.Uint()
	case reflect.Uint64:
		return f1.Uint() < f2.Uint()
	}
	panic(fmt.Errorf("unsupported field data type: %s", f1.Kind().String()))
}

func (s *sortableSlice) Swap(i, j int) {
	v1 := s.value.Index(i)
	v2 := s.value.Index(j)
	i1 := v1.Interface()
	i2 := v2.Interface()
	v1.Set(reflect.ValueOf(i2))
	v2.Set(reflect.ValueOf(i1))
}

func ByField(v interface{}, fieldName string) {
	sliceValue := reflect.ValueOf(v)
	if sliceValue.Kind() != reflect.Slice {
		panic("value is not a slice")
	}
	len := sliceValue.Len()
	if len < 2 {
		return
	}
	field := sliceValue.Index(0).FieldByName(fieldName)
	if !field.IsValid() {
		panic(fmt.Errorf("field not exist: %s", fieldName))
	}
	metaSlice := sortableSlice{value: sliceValue, len: len, fieldName: fieldName}
	sort.Sort(&metaSlice)
}

func Reverse(v interface{}) {
	sliceValue := reflect.ValueOf(v)
	if sliceValue.Kind() != reflect.Slice {
		panic("value is not a slice")
	}
	len := sliceValue.Len()
	if len < 2 {
		return
	}
	metaSlice := sortableSlice{value: sliceValue, len: len, fieldName: fieldName}
	for i := 0; i <= len/2; i++ {
		metaSlice.Swap(i, len-i-1)
	}
}
