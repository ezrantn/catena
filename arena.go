package catena

import "sync"

type Arena struct {
	memory   []byte // Arena memory chunk
	position int    // Current position in arena
}

type ArenaManager struct {
	pool sync.Pool
	size int
}

func NewArenaManager(size int) *ArenaManager {
	return &ArenaManager{
		size: size,
		pool: sync.Pool{
			New: func() any {
				return NewArena(size)
			},
		},
	}
}

func (am *ArenaManager) Get() *Arena {
	return am.pool.Get().(*Arena)
}

func (am *ArenaManager) Put(arena *Arena) {
	arena.Reset()
	am.pool.Put(arena)
}

func NewArena(size int) *Arena {
	return &Arena{
		memory:   make([]byte, size),
		position: 0,
	}
}

// Allocates memory from the arena for serialized objects
func (a *Arena) Allocate(size int) ([]byte, bool) {
	if a.position+size > len(a.memory) {
		// Allocation failed due to insufficient space
		return nil, false
	}

	start := a.position
	a.position += size
	return a.memory[start:a.position], true
}

func (a *Arena) Reset() {
	a.position = 0
}
