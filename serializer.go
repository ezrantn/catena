package catena

import "sync"

type Serializer struct {
	arena     *Arena
	jsonPool  sync.Pool
	protoPool sync.Pool
}

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
