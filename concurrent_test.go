package goutil

import (
	"sync"
	"testing"
)

func TestParallelCall(t *testing.T) {
	data := make(map[int]bool)
	mutex := sync.Mutex{}
	val, err := ParallelCall(1000, func(pos int) (interface{}, error) {
		mutex.Lock()
		defer mutex.Unlock()
		data[pos] = true
		return pos, nil
	})
	if err != nil {
		t.Errorf("TestParallelCall no error but return error:%v", err)
	}
	for i := 0; i < 1000; i++ {
		if _, ok := data[i]; !ok {
			t.Errorf("TestParallelCall pos not found: %d", i)
		}
		if val[i] != i {
			t.Errorf("TestParallelCall return value is not right (%d,%v)", i, val[i])
		}
	}
}

