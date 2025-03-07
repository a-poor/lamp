package lamp

// Lamp is a Lamport timestamped value.
type Lamp[T any] struct {
  cid string
  ts uint
  value T
  // onlocal func(cid string, ts uint, value T)
  // onremote 
}

// New creates a new Lamp instance.
func New[T any](cid string, value T) *Lamp[T] {
  return &Lamp[T]{cid, 0, value}
}

// Get returns the value of the Lamp.
func (l *Lamp[T]) Get() T {
  return l.value
}

// Set sets the value of the Lamp (locally).
func (l *Lamp[T]) Set(v T) {
  l.value = v
  l.ts++
}

// MergeRemote merges a remote change into the Lamp.
func (l *Lamp[T]) MergeRemote(v T, cid string, ts uint) {
  // If the remote timestamp is greater than
  // the local timestamp, update the local
  // value (and timestamp).
  if l.ts < ts {
    l.value = v
    l.ts = ts + 1
    return
  }

  // If the remote timestamp is equal to the
  // local timestamp, and the remote client
  // id is greater than the local client id,
  // update the local value (and timestamp).
  if l.ts == ts && l.cid < cid {
    l.value = v
    l.ts = ts + 1
  }

  // Otherwise, do nothing.
  // ...
}

