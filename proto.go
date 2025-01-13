package catena

import (
	"google.golang.org/protobuf/proto"
)

// SerializeToProto serializes an object into Protobuf and stores the result in the arena.
func (s *Serializer) SerializeToProto(obj proto.Message) ([]byte, error) {
	buf := s.protoPool.Get().([]byte)
	defer s.protoPool.Put(buf[:0])

	data, err := proto.Marshal(obj)
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

// DeserializeFromProto deserializes Protobuf data into an object.
func (s *Serializer) DeserializeFromProto(data []byte, obj proto.Message) error {
	return proto.Unmarshal(data, obj)
}
