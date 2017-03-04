package common

import "time"

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}
func MakeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}