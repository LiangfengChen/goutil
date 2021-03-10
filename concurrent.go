package goutil


// ParallelCall will call the callback function in parallel.
// if we met error will stop and return the error, if there are several errors, will return one randomly
func ParallelCall(count int, callback func(pos int) (interface{}, error)) ([]interface{}, error) {
	const MaxGoRoutineSize = 4
	result := make([]interface{}, count)
	var resultError error = nil
	parallel := make(chan bool, MaxGoRoutineSize)
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
	for i := 0; i < MaxGoRoutineSize; i++ {
		parallel <- true
	}
	return result, resultError
}
