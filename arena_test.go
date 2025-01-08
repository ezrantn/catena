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
