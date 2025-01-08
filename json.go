package catena

import (
	"encoding/json"
)

func (s *Serializer) SerializeToJSON(obj any) ([]byte, error) {
	buf := s.jsonPool.Get().([]byte)
	defer s.jsonPool.Put(buf[:0])

	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	result, ok := s.arena.Allocate(len(data))
	if !ok {
		return nil, err
	}

	copy(result, data)
	return result, nil
}

func (s *Serializer) DeserializeFromJSON(data []byte, obj any) error {
	return json.Unmarshal(data, obj)
}
