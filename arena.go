package catena

import "fmt"

type Arena struct {
	memory   []byte // Arena memory chunk
	position int    // Current position in arena
}

func NewArena(size int) *Arena {
	return &Arena{
		memory:   make([]byte, size),
		position: 0,
	}
}

// Allocates memory from the arena for serialized objects
func (a *Arena) Allocate(size int) []byte {
	if a.position+size > len(a.memory) {
		fmt.Printf("Requested size: %d, Available space: %d\n", size, len(a.memory)-a.position)
		panic("Arena memory exhausted")
	}

	start := a.position
	a.position += size
	return a.memory[start:a.position]
}
