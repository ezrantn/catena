# Catena

**Catena** is a custom serialization library that provides high-performance memory management using an arena-based allocator. The library is designed to optimize memory usage during serialization and deserialization processes by allocating memory in a single, pre-allocated chunk (arena). This can be beneficial for applications where frequent object serialization occurs and memory management needs to be efficient.

> [!IMPORTANT]
> This is an experimental project and is not intended for production use. I created this project while exploring Go's memory arena. Thank you!

## Installation

```bash
go get github.com/ezrantn/catena
```

## Usage

### JSON Serialization Example:

This example demonstrates how to serialize and deserialize a `User` object using the Catena serializer with JSON format.

```go
// ...
type User struct {
   Name  string `json:"name"`
   Email string `json:"email"`
}

// Create a User object
user := &User{
   Name:  "John",
   Email: "john@mail.com",
}

// Initialize the serializer with a specific memory arena size (e.g., 1 MB)
serializer := catena.NewSerializer(1024 * 1024)

// Serialize the User object to JSON
data, err := serializer.SerializeToJSON(user)
if err != nil {
   log.Fatalf("failed to serialize user to JSON: %v", err)
}

// Print the serialized JSON data
fmt.Printf("Serialized JSON: %s\n", string(data))
```

```go
var deserializedUser User
err = serializer.DeserializeFromJSON(data, &deserializedUser)
if err != nil {
   log.Fatalf("failed to deserialize JSON data: %v", err)
}

// Verify that the deserialized data matches the original data
if deserializedUser.Name != user.Name || deserializedUser.Email != user.Email {
   log.Fatalf("deserialization failed: expected %v, got %v", user, deserializedUser)
}

// Print the deserialized User object
fmt.Printf("Deserialized User: %+v\n", deserializedUser)
```

### Proto Serialization Example:

For serializing and deserializing using Protocol Buffers (Proto), use the following approach:

```go
// Create a ProtoUser object (generated from a .proto file)
user := &pb.ProtoUser{
   Name:  "John",
   Email: "john@mail.com",
}

// Initialize the serializer with a specific memory arena size (e.g., 1 MB)
serializer := catena.NewSerializer(1024 * 1024)

// Serialize the ProtoUser object to Proto format
data, err := serializer.SerializeToProto(user)
if err != nil {
   log.Fatalf("failed to serialize user to Proto: %v", err)
}

// Print the serialized Proto data (in bytes)
fmt.Printf("Serialized Proto: %v\n", data)
```

```go
// Deserialize the Proto data back into a ProtoUser object
var deserializedUser pb.ProtoUser
err = serializer.DeserializeFromProto(data, &deserializedUser)
if err != nil {
   log.Fatalf("failed to deserialize Proto data: %v", err)
}

// Verify that the deserialized data matches the original data
if deserializedUser.Name != user.Name || deserializedUser.Email != user.Email {
   log.Fatalf("deserialization failed: expected %v, got %v", user, deserializedUser)
}

// Print the deserialized ProtoUser object
fmt.Printf("Deserialized ProtoUser: %+v\n", deserializedUser)
```

## Benchmark

Here are the benchmark results comparing standard JSON serialization / deserialization against the custom serialization/deserialization implemented in the Catena library. The benchmark was run on an Intel i5-11400H CPU with 8GB RAM.

```bash
goarch: amd64
pkg: github.com/ezrantn/catena
cpu: 11th Gen Intel(R) Core(TM) i5-11400H @ 2.70GHz
BenchmarkJSONSerialization-12                    7626043               154.9 ns/op            48 B/op          1 allocs/op
BenchmarkJSONDeserialization-12                  1939724               620.1 ns/op           280 B/op          7 allocs/op
BenchmarkCatenaJSONSerialization-12             29596701                45.91 ns/op           48 B/op          1 allocs/op
BenchmarkCatenaJSONDeserialization-12            7336686               156.7 ns/op           280 B/op          7 allocs/op
BenchmarkGoogleProtoSerialization-12            14637853                80.19 ns/op           32 B/op          1 allocs/op
BenchmarkGoogleProtoDeserialization-12          10426597               108.2 ns/op            29 B/op          2 allocs/op
BenchmarkCatenaProtoSerialization-12            13550955                77.88 ns/op           32 B/op          1 allocs/op
BenchmarkCatenaProtoDeserialization-12           7948450               147.1 ns/op           109 B/op          3 allocs/op
PASS
ok      github.com/ezrantn/catena       10.868s
```

### Takeaways

1. **JSON Serialization and Deserialization:**

   - Catena JSON Serialization achieves a significant performance boost (45.91 ns/op vs. 154.9 ns/op), providing a 3.38x improvement over the standard JSON serialization. This speedup comes with a slight increase in memory usage (72 bytes vs. 48 bytes) and a minor increase in allocations (2 allocations vs. 1). While memory usage is a bit higher, the increased speed can be very beneficial for applications that require faster serialization without sacrificing much on memory efficiency.

   - Catena JSON Deserialization also demonstrates a substantial speedup (156.7 ns/op vs. 620.1 ns/op), achieving a 3.96x improvement. However, it retains the same memory usage (280 bytes) and allocations (7). The performance increase is especially valuable in high-throughput systems that rely on quick deserialization, though the trade-off in memory and allocations should be considered in environments where minimizing allocations is a priority.
2. **Proto Serialization and Deserialization:**

   - Catena Proto Serialization shows a slight performance loss (77.88 ns/op vs. 80.19 ns/op) compared to the Google Proto library. The custom implementation is still quite close in speed but offers added flexibility through the memory arena and object pooling. This flexibility introduces additional memory usage (56 bytes vs. 32 bytes) and allocations (2 vs. 1), but allows for fine-tuned control over memory management, which can be more beneficial for systems where memory reuse and scalability are crucial.
  
   - Catena Proto Deserialization is slower than Google's implementation (147.1 ns/op vs. 108.2 ns/op), showing a 1.36x performance difference. However, the added memory usage (109 bytes vs. 29 bytes) and allocations (3 vs. 2) come with the advantage of more scalable memory handling, which might be desirable in systems that need to efficiently manage memory at scale, even at the cost of a slight increase in deserialization time.

### Conclusion

The Catena library offers substantial improvements in performance (especially for JSON serialization/deserialization) when using unsafe operations and direct memory manipulation. The trade-offs are increased memory usage and allocations, but these are outweighed by the enhanced speed, particularly in high-concurrency systems or those processing large volumes of data.

While the standard libraries (JSON and Proto) outperform Catena in terms of raw speed and memory efficiency in simpler use cases, Catena excels in scenarios where custom memory management and scalability are more important. This is particularly useful for applications where large datasets need to be processed efficiently, and memory reuse is critical for performance.

Ultimately, the decision to use Catena over standard libraries depends on the application's requirements. For performance-sensitive applications dealing with large, frequent data serialization and deserialization tasks, Catena's approach provides better control, though with slight trade-offs in memory usage and allocation overhead.

## Testing

Here’s how to run the tests locally:

To execute unit tests, use the following command:

```bash
make test
```

To run tests with code coverage, use:

```bash
make cov
```

To perform benchmarks, execute:

```bash
make bench
```

## License

This tool is open-source and available under the [MIT License](https://github.com/ezrantn/catena/blob/main/LICENSE).

## Contributing

Contributions are welcome! Please feel free to submit a pull request. For major changes, please open an issue first to discuss what you would like to change.
