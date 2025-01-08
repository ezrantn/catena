# catena

**catena** is a custom serialization library that provides high-performance memory management using an arena-based allocator. The library is designed to optimize memory usage during serialization and deserialization processes by allocating memory in a single, pre-allocated chunk (arena). This can be beneficial for applications where frequent object serialization occurs and memory management needs to be efficient.

## Installation

```bash
go get github.com/ezrantn/catena
```

## Usage

### **Arena**

The core component of this library is the Arena, which is responsible for managing memory. An arena allocates a block of memory once, and then multiple objects can be serialized into that memory without requiring frequent allocations.

```go
package main

import "github.com/ezrantn/catena"

func main() {
    // Create a new arena with 1MB of memory
    arena := catena.NewArena(1024 * 1024)
}
```

### **Serialize an Object**

To serialize an object using a `JSONSerializer`:

```go
type User struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
}

func main() {
    user := User{Name: "Alice", Age: 30}
    arena := catena.NewArena(1024 * 1024)
    jsonSerializer := &catena.JSONSerializer{}
    
    // Serialize the object
    serializedData, err := jsonSerializer.Serialize(user, arena)
    if err != nil {
        panic(err)
    }
    
    fmt.Println("Serialized Data:", string(serializedData))
}
```

### **Deserialize an Object**

To deserialize an object from the serialized data:

```go
// Deserialize the object
var deserializedUser User
err = jsonSerializer.Deserialize(serializedData, &deserializedUser)
if err != nil {
    panic(err)
}
    
fmt.Println("Deserialized User:", deserializedUser)
```

## Benchmark

Here are the benchmark results comparing standard JSON serialization/deserialization against the custom serialization/deserialization implemented in the Catena library. The benchmark was run on an Intel i5-11400H CPU with 8GB RAM.

```bash
goos: linux
goarch: amd64
pkg: github.com/ezrantn/catena
cpu: 11th Gen Intel(R) Core(TM) i5-11400H @ 2.70GHz
BenchmarkJSONSerialization-12            7569145               161.3 ns/op            48 B/op          1 allocs/op
BenchmarkJSONDeserialization-12          1887522               637.7 ns/op           280 B/op          7 allocs/op
BenchmarkCustomSerialization-12          7210110               170.7 ns/op            48 B/op          1 allocs/op
BenchmarkCustomDeserialization-12        1820454               673.4 ns/op           280 B/op          7 allocs/op
PASS
ok      github.com/ezrantn/catena       6.527s
```

### Key Takeaways

1. The custom serialization is slightly slower than the standard JSON serialization, with a minimal performance overhead of about 8.6 ns per operation.
2. Memory allocation for both custom and standard JSON serialization is the same (48 bytes), showing that the custom implementation does not introduce unnecessary overhead.
3. The custom deserialization time is close to the standard JSON deserialization, with only a small difference of around 11 ns per operation.

## License

This tool is open-source and available under the MIT License.

## Contributing

Contributions are welcome! Please feel free to submit a pull request. For major changes, please open an issue first to discuss what you would like to change.
