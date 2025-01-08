package catena

import (
	"encoding/json"
)

func (s *Serializer) SerializeToJSON(obj any) ([]byte, error) {
	buf := s.jsonPool.Get().([]byte)
	buf = buf[:0]

	data, err := json.Marshal(obj)
	if err != nil {
		s.jsonPool.Put(buf)
		return nil, err
	}

	result := s.arena.Allocate(len(data))
	copy(result, data)

	s.jsonPool.Put(buf)
	return result, nil
}

func (s *Serializer) DeserializeFromJSON(data []byte, obj any) error {
	return json.Unmarshal(data, obj)
}
