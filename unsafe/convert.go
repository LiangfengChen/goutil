package unsafe

import "unsafe"

// Bytes2str is a faster way to convert bytes to string
// there is no allocations in this function
// BE CAREFUL to use as the input and output share the same memory
// benchmark with 144 bytes (more bytes more benefit you get)
//Benchmark_Bytes2String100-8             30000000                48.6 ns/op           112 B/op          1 allocs/op
//Benchmark_Bytes2String100_EX-8          2000000000               0.62 ns/op            0 B/op          0 allocs/op
//Benchmark_Bytes2String1000-8             5000000               295 ns/op            1024 B/op          1 allocs/op
//Benchmark_Bytes2String1000_EX-8         2000000000               0.30 ns/op            0 B/op          0 allocs/op
//Benchmark_Bytes2String10000-8            1000000              2289 ns/op           10240 B/op          1 allocs/op
//Benchmark_Bytes2String10000_EX-8        2000000000               0.29 ns/op            0 B/op          0 allocs/op
func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

