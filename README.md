# lamp

[![Go Reference](https://pkg.go.dev/badge/github.com/a-poor/lamp.svg)](https://pkg.go.dev/github.com/a-poor/lamp)

`lamp` is a [Lamport timestamp](https://en.wikipedia.org/wiki/Lamport_timestamp)
implementation in Go.

## Docs

```
type Lamp[T any] struct {
	// Has unexported fields.
}
    Lamp is a Lamport timestamped value.

func New[T any](cid string, value T) *Lamp[T]
    New creates a new Lamp instance.

func (l *Lamp[T]) Get() T
    Get returns the value of the Lamp.

func (l *Lamp[T]) MergeRemote(v T, cid string, ts uint)
    MergeRemote merges a remote change into the Lamp.

func (l *Lamp[T]) Set(v T)
    Set sets the value of the Lamp (locally).
```

## Example

```go
// Create a new value
l := lamp.New("a", 0)
fmt.Println(l.Get())

// Merge a remote value (1) with a higher
// timestamp (1)
l.MergeRemote(1, "b", 1)
fmt.Println(l.Get())

// Set the local value to 2
l.Set(2)
fmt.Println(l.Get())

// Set the local value to 3
l.Set(3)
fmt.Println(l.Get())

// Merge a remote value (1) with a lower
// timestamp (0) and nothing happens
l.MergeRemote(1, "b", 0)
fmt.Println(l.Get())

// Output: 
// 0
// 1
// 2
// 3
// 3
```

