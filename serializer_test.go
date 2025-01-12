package catena

import (
	"encoding/json"
	"testing"

	pb "github.com/ezrantn/catena/catena_proto"
)

func TestSerializerSerializeToJSON(t *testing.T) {
	// Initialize the serializer with a specific arena size (e.g., 1MB)
	serializer := NewSerializer(1024 * 1024)

	// Create a user object for testing
	user := &User{
		Name:  "Alice",
		Email: "alice@example.com",
	}

	// Serialize the User object to JSON
	data, err := serializer.SerializeToJSON(user)
	if err != nil {
		t.Fatalf("failed to serialize user to JSON: %v", err)
	}

	// Check if the serialized data is valid JSON
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("serialized data is not valid JSON: %v", err)
	}

	// Optionally, check if the serialized data contains the expected fields
	if result["name"] != "Alice" || result["email"] != "alice@example.com" {
		t.Errorf("expected name=Alice and email=alice@example.com, but got name=%v and email=%v", result["name"], result["email"])
	}

}

func TestSerializerDeserializeToJSON(t *testing.T) {
	// Initialize the serializer with a specific arena size (e.g., 1MB)
	serializer := NewSerializer(1024 * 1024)

	// Create a user object for testing
	user := &User{
		Name:  "Alice",
		Email: "alice@example.com",
	}

	// Serialize the User object to JSON
	serializedData, err := serializer.SerializeToJSON(user)
	if err != nil {
		t.Fatalf("failed to serialize user to JSON: %v", err)
	}

	// Deserialized the JSON data back into a User object
	var deserializedUser User
	err = serializer.DeserializeFromJSON(serializedData, &deserializedUser)
	if err != nil {
		t.Fatalf("failed to deserialize JSON data: %v", err)
	}

	// Verify that the deserialized data matches the original data
	if deserializedUser.Name != "Alice" || deserializedUser.Email != "alice@example.com" {
		t.Errorf("expected name=Alice and email=alice@example.com, but got name=%v and email=%v", deserializedUser.Name, deserializedUser.Email)
	}
}

func TestSerializerSerializeToProto(t *testing.T) {
	// Initialize the serializer with a specific arena size (e.g., 1MB)
	serializer := NewSerializer(1024 * 1024) // 1MB memory

	// Create a User object for testing (using proto.Message interface)
	user := &pb.ProtoUser{
		Name:  "John",
		Email: "john@mail.com",
	}

	// Serialize the User object to Proto
	data, err := serializer.SerializeToProto(user)
	if err != nil {
		t.Fatalf("failed to serialize user to Proto: %v", err)
	}

	// Verify that the serialized data is valid Proto (attempt to unmarshal it back)
	var protoUser pb.ProtoUser
	err = serializer.DeserializeFromProto(data, &protoUser)
	if err != nil {
		t.Fatalf("failed to deserialize Proto data: %v", err)
	}

	// Verify that the deserialized data matches the original data
	if protoUser.Name != "John" || protoUser.Email != "john@mail.com" {
		t.Errorf("expected name=John and email=john@mail.com, but got name=%v and email=%v", protoUser.Name, protoUser.Email)
	}
}

func TestSerializerDeserializeFromProto(t *testing.T) {
	// Initialize the serializer with a specific arena size (e.g., 1MB)
	serializer := NewSerializer(1024 * 1024) // 1MB memory

	// Create a User object for testing (using proto.Message interface)
	user := &pb.ProtoUser{
		Name:  "John",
		Email: "john@mail.com",
	}

	// Serialize the User object to Proto
	serializedData, err := serializer.SerializeToProto(user)
	if err != nil {
		t.Fatalf("failed to serialize user to Proto: %v", err)
	}

	// Deserialize the Proto data back into a User object
	var deserializedUser pb.ProtoUser
	err = serializer.DeserializeFromProto(serializedData, &deserializedUser)
	if err != nil {
		t.Fatalf("failed to deserialize Proto data: %v", err)
	}

	// Verify that the deserialized data matches the original data
	if deserializedUser.Name != "John" || deserializedUser.Email != "john@mail.com" {
		t.Errorf("expected name=John and email=john@mail.com, but got name=%v and email=%v", deserializedUser.Name, deserializedUser.Email)
	}
}
