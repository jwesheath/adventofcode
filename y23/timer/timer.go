package timer

import (
	"fmt"
	"time"
)

func Timer() func() {
	start := time.Now()
	return func() {
		fmt.Println(time.Since(start))
	}
}
