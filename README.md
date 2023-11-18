## Gofx

A simple go library for utility functions related to map access.

## Features

- Map
    - Get value for a given key
    - Set the value for a given key
    - Get or Set value for a given key
    - Delete the pair for the given key
    - Get all keys
    - Get all values
    - Get the length of the map


## Example

```go
mymap := gofx.Map[string, int]{
    "S": 2,
    "M": 4,
    "L": 6,
}
// Get the value associated with a key
value, exists := mymap.Get("M")
if exists {
    fmt.Printf("M value is: %d\n", value)
}
```

```go
// Get or set a value with a default if the key doesn't exist
key := "XL"
// when the key is not present, the default value is based on the data type
value, exists = n.Get(key)
fmt.Printf("%s value is: %d\n", key, value)
```

```go
// Set a new key-value pair
n.Set("XS", 0)

value, exists = n.GetSet("XL", 8)
fmt.Printf("%s value is: %d\n", key, value)
```

```go
// Delete a key-value pair
deletedValue, deleted := n.Delete(key)
if deleted {
    fmt.Printf("%s value deleted with value: %d", key, deletedValue)
}
```

```go
// Retrieve all keys and values
keys := n.Keys()
fmt.Println("Keys:", keys)
values := n.Values()
fmt.Println("Values:", values)

// Get the length of the nap
length := n.Len()
fmt.Println("Map Length:", length)
```


