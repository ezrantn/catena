package catena

import (
	"encoding/json"
	"unsafe"
)

// SerializeToJSON serializes an object into JSON and stores the result in the arena.
func (s *Serializer) SerializeToJSON(obj any) ([]byte, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	// Try to allocate memory in the arena directly for the serialized data
	result, ok := s.arena.Allocate(len(data))
	if !ok {
		return nil, err // Allocation failed, return error
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

// DeserializeFromJSON deserializes JSON data into an object.
func (s *Serializer) DeserializeFromJSON(data []byte, obj any) error {
	return json.Unmarshal(data, obj)
}
