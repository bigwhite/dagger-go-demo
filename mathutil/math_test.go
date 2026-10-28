package mathutil

import "testing"

func TestAdd(t *testing.T) {
    got := Add(2, 3)
    if got != 5 {
        t.Fatalf("Add(2, 3) = %d, want 5", got)
    }
}

func TestAddNegative(t *testing.T) {
    got := Add(-2, -3)
    if got != -5 {
        t.Fatalf("Add(-2, -3) = %d, want -5", got)
    }
}
