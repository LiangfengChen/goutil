package goutil

import (
	"errors"
	"fmt"
	"sync"
	"testing"
)

func TestTryParallelCall(t *testing.T) {
	data := make(map[int]bool)
	mutex := sync.Mutex{}
	val, err := TryParallelCall(1000, 10, func(pos int) (interface{}, error) {
		mutex.Lock()
		defer mutex.Unlock()
		data[pos] = true
		return pos, nil
	})
	if err != nil {
		t.Errorf("TestTryParallelCall no error but return error:%v", err)
	}
	for i := 0; i < 1000; i++ {
		if _, ok := data[i]; !ok {
			t.Errorf("TestTryParallelCall pos not found: %d", i)
		}
		if val[i] != i {
			t.Errorf("TestTryParallelCall return value is not right (%d,%v)", i, val[i])
		}
	}
}

func TestParallelCall(t *testing.T) {
	format := "test:%d"
	data := make(map[int]bool)
	mutex := sync.Mutex{}
	val, err := ParallelCall(1000, 10, func(pos int) (interface{}, error) {
		mutex.Lock()
		defer mutex.Unlock()
		data[pos] = true
		return pos, errors.New(fmt.Sprintf(format, pos))
	})

	for i := 0; i < 1000; i++ {
		if _, ok := data[i]; !ok {
			t.Errorf("TestParallelCall pos not found: %d", i)
		}
		if val[i] != i {
			t.Errorf("TestParallelCall return value is not right (%d,%v)", i, val[i])
		}
		if err[i].Error() != fmt.Sprintf(format, i) {
			t.Errorf("TestParallelCall error msg is not correct (%d,%v)", i, err[i])
		}
	}
}

