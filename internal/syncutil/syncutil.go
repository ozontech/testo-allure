package syncutil

import (
	"sync"
	"sync/atomic"
)

type KeyedMutex[K comparable] struct {
	m sync.Map
}

func (km *KeyedMutex[K]) Lock(key K) (unlock func()) {
	value, _ := km.m.LoadOrStore(key, new(sync.Mutex))

	mu := value.(*sync.Mutex)

	mu.Lock()

	return mu.Unlock
}

type MutexGuarded[T any] struct {
	value T
	mu    sync.RWMutex
}

func (mg *MutexGuarded[T]) Load() T {
	mg.mu.RLock()
	defer mg.mu.RUnlock()

	return mg.value
}

func (mg *MutexGuarded[T]) Store(value T) {
	mg.Modify(func(v *T) {
		*v = value
	})
}

func (mg *MutexGuarded[T]) Modify(f func(value *T)) {
	mg.mu.Lock()
	defer mg.mu.Unlock()

	f(&mg.value)
}

type AtomicInt[T ~int | ~int32 | ~int64] atomic.Int64

func (a *AtomicInt[T]) Load() T {
	return T((*atomic.Int64)(a).Load())
}

func (a *AtomicInt[T]) Store(value T) {
	(*atomic.Int64)(a).Store(int64(value))
}

type AtomicValue[T any] atomic.Value

func (a *AtomicValue[T]) Load() T {
	p := (*atomic.Value)(a).Load()

	if p == nil {
		var zero T

		return zero
	}

	return p.(T)
}

func (a *AtomicValue[T]) Store(value T) {
	(*atomic.Value)(a).Store(value)
}
