package catena

import (
	"encoding/json"
	"testing"
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

func BenchmarkCustomSerialization(b *testing.B) {
	user := createTestUser()
	arena := NewArena(1024 * 1024)
	jsonSerializer := &JSONSerializer{}
	for i := 0; i < b.N; i++ {
		arena.position = 0
		_, err := SerializeObject(jsonSerializer, user, arena)
		if err != nil {
			b.Fatalf("Custom serialization failed: %v", err)
		}
	}
}

func BenchmarkCustomDeserialization(b *testing.B) {
	user := createTestUser()
	arena := NewArena(1024 * 1024)
	jsonSerializer := &JSONSerializer{}
	serializedData, _ := SerializeObject(jsonSerializer, user, arena)
	b.ResetTimer() // Reset the timer to exclude setup time
	for i := 0; i < b.N; i++ {
		var deserializedUser User
		err := DeserializeObject(jsonSerializer, serializedData, &deserializedUser)
		if err != nil {
			b.Fatalf("Custom deserialization failed: %v", err)
		}
	}
}
