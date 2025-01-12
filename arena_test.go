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
		_, err := proto.Marshal(user) // Google's Protobuf serialization
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

func TestArenaAllocate(t *testing.T) {
	size := 1024
	arena := NewArena(size)

	// Allocate within available memory
	data, ok := arena.Allocate(512)
	if !ok || len(data) != 512 {
		t.Fatalf("expected allocation of 512 bytes to succeed, got ok=%v, len(data)=%d", ok, len(data))
	}

	// Allocate exceeding available memory
	data, ok = arena.Allocate(600)
	if ok || data != nil {
		t.Fatalf("expected allocation of 600 bytes to fail, got ok=%v, data=%v", ok, data)
	}
}

func TestArenaReset(t *testing.T) {
	size := 1024
	arena := NewArena(size)

	// Allocate some memory
	_, ok := arena.Allocate(512)
	if !ok {
		t.Fatalf("allocation of 512 bytes failed")
	}

	// Reset the arena
	arena.Reset()

	// Allocate again, should reuse the memory
	data, ok := arena.Allocate(1024)
	if !ok || len(data) != 1024 {
		t.Fatalf("expected allocation of 1024 bytes to succeed after reset, got ok=%v, len(data)=%d", ok, len(data))
	}
}

func TestArenaManagerGetPut(t *testing.T) {
	size := 1024
	manager := NewArenaManager(size)

	// Get an arena from the manager
	arena := manager.Get()
	if arena == nil {
		t.Fatalf("expected arena to be non-nil")
	}

	// Allocate new memory
	_, ok := arena.Allocate(512)
	if !ok {
		t.Fatalf("allocation of 512 bytes failed")
	}

	// Return the arena to the pool
	manager.Put(arena)

	// Get the arena again and ensure it's reset
	arena = manager.Get()
	data, ok := arena.Allocate(1024)
	if !ok || len(data) != 1024 {
		t.Fatalf("expected allocation of 1024 bytes to succeed after returning to pool, ok=%v, len(data)=%d", ok, len(data))
	}
}

func TestArenaManagerMultipleArenas(t *testing.T) {
	size := 1024
	manager := NewArenaManager(size)

	// Get multiple arenas from the manager
	arena1 := manager.Get()
	arena2 := manager.Get()

	if arena1 == arena2 {
		t.Fatalf("expected different arenas, but got the same instance")
	}

	// Return both arenas to the pool
	manager.Put(arena1)
	manager.Put(arena2)

	// Get arenas again and ensure they are reset
	arena1 = manager.Get()
	arena2 = manager.Get()

	_, ok1 := arena1.Allocate(512)
	_, ok2 := arena2.Allocate(512)

	if !ok1 || !ok2 {
		t.Fatalf("expected allocations in both arenas to succeed, got ok1=%v, ok=%v", ok1, ok2)
	}
}
