package unsafe

import (
	"fmt"
	"testing"
)

// Bechmark for Bytes2String
// go test -benchmem -bench=. ./unsafe
// Benchmark_Bytes2String100-8             30000000                48.6 ns/op           112 B/op          1 allocs/op
// Benchmark_Bytes2String100_EX-8          2000000000               0.62 ns/op            0 B/op          0 allocs/op
// Benchmark_Bytes2String1000-8             5000000               295 ns/op            1024 B/op          1 allocs/op
// Benchmark_Bytes2String1000_EX-8         2000000000               0.30 ns/op            0 B/op          0 allocs/op
// Benchmark_Bytes2String10000-8            1000000              2289 ns/op           10240 B/op          1 allocs/op
// Benchmark_Bytes2String10000_EX-8        2000000000               0.29 ns/op            0 B/op          0 allocs/op

func Benchmark_Bytes2String100(t *testing.B) {
	src := make([]byte, 100)
	for i := 0; i < t.N; i++ {
		_ = string(src)
	}
}

func Benchmark_Bytes2String100_EX(t *testing.B) {
	src := make([]byte, 100)
	for i := 0; i < t.N; i++ {
		_ = Bytes2String(src)
	}
}

func Benchmark_Bytes2String1000(t *testing.B) {
	src := make([]byte, 1000)
	for i := 0; i < t.N; i++ {
		_ = string(src)
	}
}

func Benchmark_Bytes2String1000_EX(t *testing.B) {
	src := make([]byte, 1000)
	for i := 0; i < t.N; i++ {
		_ = Bytes2String(src)
	}
}

func Benchmark_Bytes2String10000(t *testing.B) {
	src := make([]byte, 10000)
	for i := 0; i < t.N; i++ {
		_ = string(src)
	}
}

func Benchmark_Bytes2String10000_EX(t *testing.B) {
	src := make([]byte, 10000)
	for i := 0; i < t.N; i++ {
		_ = Bytes2String(src)
	}
}

// Bechmark for String2Bytes
// go test -benchmem -bench=. ./unsafe
// Benchmark_WithEscape-12                 34980004                37.2 ns/op            32 B/op          1 allocs/op
// Benchmark_NoEscape-12                   208929188                5.59 ns/op            0 B/op          0 allocs/op

func Benchmark_String2Bytes(t *testing.B) {
	src := "0123456789012345678901234567890123456789"
	for i := 0; i < t.N; i++ {
		_ = []byte(src)
	}
}

func Benchmark_String2Bytes_EX(t *testing.B) {
	src := "0123456789012345678901234567890123456789"
	for i := 0; i < t.N; i++ {
		_ = String2Bytes(src)
	}
}

// Test cases for String2Bytes
func TestString2Bytes(t *testing.T) {
	str := "1234567890"
	data := String2Bytes(str)
	fmt.Printf("%+v", data)
}