package unsafe

import "testing"

//go test -benchmem -bench=. ./unsafe
//Benchmark_Bytes2String100-8             30000000                48.6 ns/op           112 B/op          1 allocs/op
//Benchmark_Bytes2String100_EX-8          2000000000               0.62 ns/op            0 B/op          0 allocs/op
//Benchmark_Bytes2String1000-8             5000000               295 ns/op            1024 B/op          1 allocs/op
//Benchmark_Bytes2String1000_EX-8         2000000000               0.30 ns/op            0 B/op          0 allocs/op
//Benchmark_Bytes2String10000-8            1000000              2289 ns/op           10240 B/op          1 allocs/op
//Benchmark_Bytes2String10000_EX-8        2000000000               0.29 ns/op            0 B/op          0 allocs/op

func Benchmark_Bytes2String100(t *testing.B) {
	src := make([]byte, 100)
	for i := 0; i < t.N; i++ {
		_ = string(src)
	}
}

func Benchmark_Bytes2String100_EX(t *testing.B) {
	src := make([]byte, 100)
	for i := 0; i < t.N; i++ {
		_ = Bytes2str(src)
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
		_ = Bytes2str(src)
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
		_ = Bytes2str(src)
	}
}
