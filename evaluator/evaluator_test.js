func TestBangOperator(t *testing.T) {
  tests := []struct {
    input string
    expected bool
  }{
    {"!true", false},
    {"!false", true},
    {"!5", false},
    {"!!true", true},
    {"!!false", false},
    {"!!5", true},
  }

  for _, tt := range tests {
    evaluated := testEval(tt.input)
    TestBooleanObject(t, evaluated, tt.expected)
  }
}

