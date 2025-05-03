package approx

import (
	"testing"
)

func TestCompactApproximatorFNV(t *testing.T) {
	approximator, err := NewCompactApproximator(512, 3, FNV)
	if err != nil {
		t.Fatal(err)
	}

	key := []byte("user123")
	approximator.Insert(key, 10)
	value := approximator.Get(key)
	if value < 10 {
		t.Errorf("Expected at least 10, got %d", value)
	}

	approximator.Insert(key, 20)
	value = approximator.Get(key)
	if value < 20 {
		t.Errorf("Expected at least 20, got %d", value)
	}
}

func TestCompactApproximatorSHA256(t *testing.T) {
	approximator, err := NewCompactApproximator(512, 4, SHA256)
	if err != nil {
		t.Fatal(err)
	}

	key := []byte("session456")
	approximator.Insert(key, 42)
	value := approximator.Get(key)
	if value < 42 {
		t.Errorf("Expected at least 42, got %d", value)
	}
}

func TestMultipleKeys(t *testing.T) {
	tracker, err := NewCompactApproximator(1024, 3, FNV)
	if err != nil {
		t.Fatal(err)
	}

	keys := [][]byte{[]byte("k1"), []byte("k2"), []byte("k3")}
	values := []uint64{5, 10, 15}

	for i, key := range keys {
		tracker.Insert(key, values[i])
	}

	for i, key := range keys {
		got := tracker.Get(key)
		if got < values[i] {
			t.Errorf("Expected at least %d for key %s, got %d", values[i], key, got)
		}
	}
}
