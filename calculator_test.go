package calculate

import (
 "testing"
)

func TestCalc(t *testing.T) {
 tests := []struct {
  expression string
  expected   float64
  shouldFail bool
 }{
  {"2 + 2", 4, false},
  {"2 * 3", 6, false},
  {"10 / 2", 5, false},
  {"5 - 3", 2, false},
  {"10 / 0", 0, true}, // Ошибка деления на ноль
  {"invalid", 0, true}, // Некорректное выражение
 }

 for _, test := range tests {
  result, err := Calc(test.expression)
  if test.shouldFail && err == nil {
   t.Errorf("expected error for expression: %s", test.expression)
  }
  if !test.shouldFail && err != nil {
   t.Errorf("unexpected error for expression: %s, error: %v", test.expression, err)
  }
  if result != test.expected {
   t.Errorf("for expression: %s, expected %v, got %v", test.expression, test.expected, result)
  }
 }
}