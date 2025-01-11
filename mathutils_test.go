package mathutils

import "testing"

// Тестирование функции Add
func TestAdd(t *testing.T) {
    tests := []struct {
        a, b, expected int
    }{
        {1, 2, 3},
        {-1, -1, -2},
        {0, 0, 0},
        {1000000, 1000000, 2000000},
        {-1, 2, 1},
    }

    for _, test := range tests {
        t.Run("", func(t *testing.T) {
            result := Add(test.a, test.b)
            if result != test.expected {
                t.Errorf("Add(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
            }
        })
    }
}

// Тестирование функции Sub
func TestSub(t *testing.T) {
    tests := []struct {
        a, b, expected int
    }{
        {3, 2, 1},
        {-1, -1, 0},
        {0, 0, 0},
        {1000000, 999999, 1},
        {5, 7, -2},
    }

    for _, test := range tests {
        t.Run("", func(t *testing.T) {
            result := Sub(test.a, test.b)
            if result != test.expected {
                t.Errorf("Sub(%d, %d) = %d; expected %d", test.a, test.b, result, test.expected)
            }
        })
    }
}
