package easysort

import (
	"sort"
	"testing"
	"time"
)

type TestStruct struct {
	String  string
	Time    time.Time
	Float32 float64
	Float64 float32
	Bool    bool
	Int     int
	Int8    int8
	Int16   int16
	Int32   int32
	Int64   int64
	Uint    uint
	Uint16  uint16
	Uint32  uint32
	Uint64  uint64
	Byte    byte
	Bytes   []byte
}

type TestStructs []TestStruct

func (ts TestStructs) Len() int {
	return len(ts)
}

func (ts TestStructs) Less(i, j int) bool {
	return ts[i].Int < ts[j].Int
}

func (ts TestStructs) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

var testA = TestStruct{"a", time.Unix(0, 0), 1.11, 1.11, false, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, []byte("testA")}
var testB = TestStruct{"b", time.Unix(100000, 0), 2.22, 2.22, false, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, []byte("testB")}
var testC = TestStruct{"c", time.Unix(100000000, 0), 3.33, 3.33, true, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, []byte("testC")}

func TestSortByField(t *testing.T) {
	slice := []TestStruct{testA, testC, testB}
	ByField(slice, "Bytes")
	t.Logf("result: %v", slice)
}

/*
func TestSortByFieldMapStringInterface(t *testing.T) {
	slice := []map[string]interface{}{
		map[string]interface{}{"user":"testb", "age": 45},
		map[string]interface{}{"user":"testa", "age": 35},
	}
	ByField(slice, "age")
	t.Logf("result: %v", slice)
}
*/
func TestReverse(t *testing.T) {
	slice := []TestStruct{testA, testC, testB}
	Reverse(slice)
	t.Logf("reverse: %v", slice)
}

func BenchmarkByFieldInt(b *testing.B) {
	len := 100000
	slice := make([]TestStruct, len)
	for i := 0; i < len-2; i++ {
		slice[i] = testB
	}
	slice[len-2] = testC
	slice[len-1] = testA
	for i := 0; i < b.N; i++ {
		ByField(slice, "Int")
	}
}

func BenchmarkNativeByInt(b *testing.B) {
	len := 100000
	slice := make(TestStructs, len)
	for i := 0; i < len-2; i++ {
		slice[i] = testB
	}
	slice[len-2] = testC
	slice[len-1] = testA
	for i := 0; i < b.N; i++ {
		sort.Sort(slice)
	}
}
