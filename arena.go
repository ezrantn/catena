package catena

import "sync"

// arena represents a memory arena that holds a large block of memory for efficient allocation.
// memory 	= arena memory chunk (a pre-allocated block of memory)
// position = current position in the arena for allocation
type arena struct {
	memory   []byte
	position int
}

// arenaManager manages a pool of arena objects for efficient reuse.
// pool = pool of reusable arena objects
// size = size of each arena managed by the pool
type arenaManager struct {
	pool sync.Pool
	size int
}

// NewArenaManager creates and returns a new arenaManager with a given arena size.
// It initialize the pool with a new arena object when needed and
// creates a new arena with the specified size
func NewArenaManager(size int) *arenaManager {
	return &arenaManager{
		size: size,
		pool: sync.Pool{
			New: func() any {
				return NewArena(size)
			},
		},
	}
}

// Get retrieves an arena from the pool (or creates a new one if empty).
func (am *arenaManager) Get() *arena {
	return am.pool.Get().(*arena)
}

// Put returns an arena back to the pool after resetting it for reuse.
// Reset the arena before putting it back in the pool to clear it's state
func (am *arenaManager) Put(arena *arena) {
	arena.Reset()
	am.pool.Put(arena)
}

// NewArena creates and returns a new Arena with a given size.
func NewArena(size int) *arena {
	return &arena{
		memory:   make([]byte, size),
		position: 0,
	}
}

// Allocates memory from the arena for serialized objects
func (a *arena) Allocate(size int) ([]byte, bool) {
	if a.position+size > len(a.memory) {
		// Allocation failed due to insufficient space
		return nil, false
	}

	start := a.position
	a.position += size
	return a.memory[start:a.position], true
}

// Reset resets the arena's position to the beginning to reuse the memory.
func (a *arena) Reset() {
	a.position = 0
}
