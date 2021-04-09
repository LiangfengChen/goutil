package goutil

import (
	"strconv"
	"strings"
)

// JoinUint32Slice concatenates the uint32 elements of its first argument to create a single string. The separator
// string sep is placed between elements in the resulting string.
func JoinUint32Slice(slice []uint32, sep string) string {
	strList := make([]string, 0, len(slice))
	for _, meta := range slice {
		strList = append(strList, strconv.Itoa(int(meta)))
	}
	return strings.Join(strList, sep)
}

// JoinInt32Slice concatenates the int32 elements of its first argument to create a single string. The separator
// string sep is placed between elements in the resulting string.
func JoinInt32Slice(slice []int32, sep string) string {
	strList := make([]string, 0, len(slice))
	for _, meta := range slice {
		strList = append(strList, strconv.Itoa(int(meta)))
	}
	return strings.Join(strList, sep)
}
