package safesum

import (
	"fmt"
	"sync"
)

type safesum struct {
	mu  sync.Mutex
	sum int
}

func NewSafeSum() safesum {
	return safesum{sum: 0}
}

func (sum *safesum) Add(n int) {
	sum.mu.Lock()
	sum.sum += n
	sum.mu.Unlock()
}

func (sum *safesum) Print() {
	fmt.Println(sum.sum)
}
