package catena

import (
	"encoding/json"
	"testing"

	pb "github.com/ezrantn/catena/catena_proto"
	"google.golang.org/protobuf/proto"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func createTestUser() *User {
	return &User{
		Name:  "Alice",
		Email: "alice@example.com",
	}
}

func createTestProtoUser() *pb.ProtoUser {
	return &pb.ProtoUser{
		Name:  "Alice",
		Email: "alice@example.com",
	}
}

func BenchmarkJSONSerialization(b *testing.B) {
	user := createTestUser()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(user)
		if err != nil {
			b.Fatalf("JSON serialization failed: %v", err)
		}
	}
}

func BenchmarkJSONDeserialization(b *testing.B) {
	user := createTestUser()
	serializedData, _ := json.Marshal(user)
	b.ResetTimer() // Reset the timer to exclude setup time
	for i := 0; i < b.N; i++ {
		var deserializedUser User
		err := json.Unmarshal(serializedData, &deserializedUser)
		if err != nil {
			b.Fatalf("JSON deserialization failed: %v", err)
		}
	}
}

func BenchmarkCatenaJSONSerialization(b *testing.B) {
	user := &User{
		Name:  "Alice",
		Email: "alice@example.com",
	}
	serializer := NewSerializer(1024 * 1024)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := serializer.SerializeToJSON(user)
			if err != nil {
				b.Fatalf("JSON serialization failed: %v", err)
			}
		}
	})
}

func BenchmarkCatenaJSONDeserialization(b *testing.B) {
	user := &User{
		Name:  "Alice",
		Email: "alice@example.com",
	}
	serializer := NewSerializer(1024 * 1024)
	data, _ := serializer.SerializeToJSON(user)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var deserializedUser User
			err := serializer.DeserializeFromJSON(data, &deserializedUser)
			if err != nil {
				b.Fatalf("JSON deserialization failed: %v", err)
			}
		}
	})
}

func BenchmarkGoogleProtoSerialization(b *testing.B) {
	user := &pb.ProtoUser{
		Name:  "Alice",
		Email: "alice@example.com",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(user)
		if err != nil {
			b.Fatalf("Google Protobuf serialization failed: %v", err)
		}
	}
}

func BenchmarkGoogleProtoDeserialization(b *testing.B) {
	user := &pb.ProtoUser{
		Name:  "Alice",
		Email: "alice@example.com",
	}
	data, err := proto.Marshal(user) // Pre-serialize the ProtoUser
	if err != nil {
		b.Fatalf("Pre-serialization failed: %v", err)
	}

	var deserializedUser pb.ProtoUser

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := proto.Unmarshal(data, &deserializedUser) // Google's Protobuf deserialization
		if err != nil {
			b.Fatalf("Google Protobuf deserialization failed: %v", err)
		}
	}
}

func BenchmarkCatenaProtoSerialization(b *testing.B) {
	user := createTestProtoUser()
	serializer := NewSerializer(1024 * 1024)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := serializer.SerializeToProto(user)
		if err != nil {
			b.Fatalf("Proto serialization failed: %v", err)
		}
	}
}

func BenchmarkCatenaProtoDeserialization(b *testing.B) {
	user := createTestProtoUser()
	serializer := NewSerializer(1024 * 1024)
	data, _ := serializer.SerializeToProto(user)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var deserializedUser pb.ProtoUser
		err := serializer.DeserializeFromProto(data, &deserializedUser)
		if err != nil {
			b.Fatalf("Proto deserialization failed: %v", err)
		}
	}
}
