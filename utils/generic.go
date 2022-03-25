package utils

import "time"

func TimeNow() string {
	return time.Now().Format(time.RFC3339)
}
