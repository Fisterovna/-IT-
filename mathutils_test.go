package mathutils

import "testing"

func TestAdd(t *testing.T) {
    result := Add(3, 5)
    if result != 8 {
        t.Errorf("Expected 8, got %f", result)
    }
}

func TestSub(t *testing.T) {
    result := Sub(10, 4)
    if result != 6 {
        t.Errorf("Expected 6, got %f", result)
    }
}
