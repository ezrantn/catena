package catena

import "google.golang.org/protobuf/proto"

// Serialize serializes an object into Protocol Buffers format.
func (p *ProtoSerializer) Serialize(object interface{}, arena *Arena) ([]byte, error) {
	data, err := proto.Marshal(object.(proto.Message))
	if err != nil {
		return nil, err
	}

	// Allocate space in the arena for the serialized data.
	arenaData := arena.Allocate(len(data))
	copy(arenaData, data)
	return arenaData, nil
}

// Deserialize deserializes Protocol Buffers data into an object.
func (p *ProtoSerializer) Deserialize(data []byte, object interface{}) error {
	return proto.Unmarshal(data, object.(proto.Message))
}
