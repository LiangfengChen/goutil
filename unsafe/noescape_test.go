package unsafe

import (
	"testing"
)

//go test -benchmem -bench=. ./unsafe
//Benchmark_WithEscape-12         35929988                33.2 ns/op            32 B/op          1 allocs/op
//Benchmark_NoEscape-12           217909807                5.49 ns/op            0 B/op          0 allocs/op

type dataChecker struct {
	value1 string
	value2 int64
	value3 *dataChecker
}

//go:noinline
func assignValue(data *dataChecker)  {
	data.value1 = "value2"
	data.value2 = 2
	data.value3 = &dataChecker{}
}

//go:noinline
func assignValue2(data *dataChecker)  {
	data.value1 = "value2"
	data.value2 = 2
	// WARN: it is just used to check noescape, pls don't use it like this.
	// as the value3 is already invalid outside the function
	data.value3 = NoEscape(&dataChecker{}).(*dataChecker)
}

func Benchmark_WithEscape(b *testing.B) {
	for i := 0; i < b.N; i++{
		data := &dataChecker{}
		assignValue(data)
		data.value1 = "value"
	}
}

func Benchmark_NoEscape(b *testing.B) {
	for i := 0; i < b.N; i++{
		data := &dataChecker{}
		assignValue2(data)
		data.value1 = "value"
	}
}