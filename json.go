package catena

import (
	"encoding/json"
	"fmt"
)

// Serialize serializes an object into JSON format
func (j *JSONSerializer) Serialize(object interface{}, arena *Arena) ([]byte, error) {
	data, err := json.Marshal(object)
	if err != nil {
		return nil, err
	}

	// Check if there is enough space in the arena
	if len(data) > len(arena.memory)-arena.position {
		return nil, fmt.Errorf("insufficient space in arena: need %d bytes, but only %d bytes are available", len(data), len(arena.memory)-arena.position)
	}

	// Allocate space in arena for the serialized data
	arenaData := arena.Allocate(len(data))
	copy(arenaData, data)
	return arenaData, nil
}

// Deserialize deserializes JSON data into an object
func (j *JSONSerializer) Deserialize(data []byte, object interface{}) error {
	return json.Unmarshal(data, object)
}
