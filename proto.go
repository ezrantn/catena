package catena

import (
	"unsafe"

	"google.golang.org/protobuf/proto"
)

// SerializeToProto serializes an object into Protobuf and stores the result in the arena.
func (s *Serializer) SerializeToProto(obj proto.Message) ([]byte, error) {
	data, err := proto.Marshal(obj)
	if err != nil {
		return nil, err
	}

	// Try to allocate memory in the arena directly for the serialized data
	result, ok := s.arena.Allocate(len(data))
	if !ok {
		return nil, err
	}

	// EXPERIMENT:
	// Use unsafe.Pointer to directly manipulate memory and avoid copying
	// Convert the result (which is a slice) into a pointer in the memory
	dst := (*[1 << 30]byte)(unsafe.Pointer(&result[0]))

	// Copy the serialized JSON data into the arena memory chunk directly
	copy((*dst)[:len(data)], data)

	// Return the memory chunk in the arena that holds the serialized data
	return result, nil
}

// DeserializeFromProto deserializes Protobuf data into an object.
func (s *Serializer) DeserializeFromProto(data []byte, obj proto.Message) error {
	return proto.Unmarshal(data, obj)
}
