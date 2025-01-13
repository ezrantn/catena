package catena

import "sync"

// Serializer is a struct that manages serialization, memory allocation, and memory pooling.
type Serializer struct {
	arena     *arena
	jsonPool  sync.Pool
	protoPool sync.Pool
}

// NewSerializer creates and initializes a new Serializer with a specified arena size.
// It also initializes pools for JSON and Protobuf serialization buffers.
func NewSerializer(arenaSize int) Serializer {
	return Serializer{
		arena: NewArena(arenaSize),
		jsonPool: sync.Pool{
			New: func() any {
				return make([]byte, 0, 1024)
			},
		},
		protoPool: sync.Pool{
			New: func() any {
				return make([]byte, 0, 1024)
			},
		},
	}
}
