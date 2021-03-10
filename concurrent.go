package goutil


// TryParallelCall will try to call the callback function in parallel. If error occurs, stop and return
// if we get error will stop and return the error, if there are several errors, will only return one randomly
func TryParallelCall(count int, maxGoRoutine int, callback func(pos int) (interface{}, error)) ([]interface{}, error) {
	result := make([]interface{}, count)
	var resultError error = nil
	parallel := make(chan bool, maxGoRoutine)
	for i := 0; i < count && nil == resultError; i++ {
		parallel <- true
		go func(pos int) {
			val, err := callback(pos)
			if err != nil {
				resultError = err
			}
			result[pos] = val
			<-parallel
		}(i)
	}
	for i := 0; i < maxGoRoutine; i++ {
		parallel <- true
	}
	return result, resultError
}

// ParallelCall will call the callback function in parallel. If error occurs, will continue to finish the calling
// it will return all the result into the array
func ParallelCall(count int, maxGoRoutine int, callback func(pos int) (interface{}, error)) ([]interface{}, []error) {
	result := make([]interface{}, count)
	resultError := make([]error, count)
	parallel := make(chan bool, maxGoRoutine)
	for i := 0; i < count; i++ {
		parallel <- true
		go func(pos int) {
			val, err := callback(pos)
			resultError[pos] = err
			result[pos] = val
			<-parallel
		}(i)
	}
	for i := 0; i < maxGoRoutine; i++ {
		parallel <- true
	}
	return result, resultError
}
