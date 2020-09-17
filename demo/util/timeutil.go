package util

import (
	"fmt"
	"time"
)

// GetTime : 获取时间
func GetTime() {
	now := time.Now()
	fmt.Printf("%v", now)
}
