package spinlock

import (
	"runtime"
	"sync"
	"sync/atomic"
)

// Locker is a spinlock implementation.
//
// A Locker must not be copied after first use.
type Locker struct {
	_    sync.Mutex // for copy protection compiler warning
	lock uintptr
}

// Lock locks l.
// If the lock is already in use, the calling goroutine
// blocks until the locker is available.
func (l *Locker) Lock() {
loop:
	if !atomic.CompareAndSwapUintptr(&l.lock, 0, 1) {
		runtime.Gosched()
		goto loop
	}
}

// Unlock unlocks l.
func (l *Locker) Unlock() {
	atomic.StoreUintptr(&l.lock, 0)
}
