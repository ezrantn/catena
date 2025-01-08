package catena

import "sync"

type Arena struct {
	memory   []byte     // Arena memory chunk
	position int        // Current position in arena
	mu       sync.Mutex // Thread-safe operation
}

func NewArena(size int) *Arena {
	return &Arena{
		memory:   make([]byte, size),
		position: 0,
	}
}

// Allocates memory from the arena for serialized objects
func (a *Arena) Allocate(size int) []byte {
	a.mu.Lock()
	defer a.mu.Unlock()

	// Double the arena size if needed
	if a.position+size > len(a.memory) {
		newSize := max(len(a.memory)*2, a.position+size)
		newMemory := make([]byte, newSize)
		copy(newMemory, a.memory)
		a.memory = newMemory
	}

	start := a.position
	a.position += size
	return a.memory[start:a.position]
}
