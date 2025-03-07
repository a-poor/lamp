package lamp_test

import (
	"fmt"
	"testing"

	"github.com/a-poor/lamp"
)

func TestLampMergeRemote(t *testing.T) {
  // The remote value should win
  t.Run("remote ts > local ts", func(t *testing.T) {
    l := lamp.New("a", 0)
    l.MergeRemote(1, "b", 1)
    if l.Get() != 1 {
      t.Errorf("expected 1, got %v", l.Get())
    }
  })
  // The local value should win
  t.Run("remote ts < local ts", func(t *testing.T) {
    l := lamp.New("a", 0)
    l.Set(0)
    l.MergeRemote(1, "b", 0)
    if l.Get() != 0 {
      t.Errorf("expected 0, got %v", l.Get())
    }
  })
  // Tie version. Remote id greater. Remote wins.
  t.Run("remote ts == local ts, remote cid > local cid", func(t *testing.T) {
    l := lamp.New("a", 0)
    l.MergeRemote(1, "b", 0)
    if l.Get() != 1 {
      t.Errorf("expected 1, got %v", l.Get())
    }
  })
  // Tie version. Local id greater. Local wins.
  t.Run("remote ts == local ts, remote cid < local cid", func(t *testing.T) {
    l := lamp.New("b", 0)
    l.MergeRemote(1, "a", 0)
    if l.Get() != 0 {
      t.Errorf("expected 0, got %v", l.Get())
    }
  })
  // Tie version. Same id. Shouldn't happen? No-op.
  t.Run("remote ts == local ts, remote cid == local cid", func(t *testing.T) {
    l := lamp.New("a", 0)
    l.MergeRemote(1, "a", 0)
    if l.Get() != 0 {
      t.Errorf("expected 0, got %v", l.Get())
    }
  })
}

func ExampleLamp_MergeRemote() {
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
}

