package easysort

import (
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

var testA = TestStruct{"a", time.Unix(0, 0), 1.11, 1.11, false, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, []byte("testA")}
var testB = TestStruct{"b", time.Unix(100000, 0), 2.22, 2.22, false, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, []byte("testB")}
var testC = TestStruct{"c", time.Unix(100000000, 0), 3.33, 3.33, true, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, []byte("testC")}

func TestSortByField(t *testing.T) {
	slice := []TestStruct{testA, testC, testB}
	ByField(slice, "Byte")
	t.Logf("result: %v", slice)
}

func TestReverse(t *testing.T) {
	slice := []TestStruct{testA, testC, testB}
	Reverse(slice)
	t.Logf("reverse: %s", slice)
}
