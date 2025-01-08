package catena

// Serializer is the interface that all serialization formats must implement
type Serializer interface {
	Serialize(object interface{}, arena *Arena) ([]byte, error)
	Deserialize(data []byte, object interface{}) error
}

// SerializeObject serializes an object using the provided serializer format (JSON, Proto, etc.).
func SerializeObject(serializer Serializer, object interface{}, arena *Arena) ([]byte, error) {
	return serializer.Serialize(object, arena)
}

// DeserializeObject deserializes data into an object using the provided serializer format.
func DeserializeObject(serializer Serializer, data []byte, object interface{}) error {
	return serializer.Deserialize(data, object)
}

type JSONSerializer struct{}
type ProtoSerializer struct{}
