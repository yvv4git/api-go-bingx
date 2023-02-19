package utils

import "time"

// CurrentTimestamp - used for get current timestamp.
func CurrentTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
