package main

import "testing"



func TestCleanInput(t *testing.T) {
  cases := []struct {
	input    string
	expected []string
}{
	{
		input:    "  hello  world  ",
		expected: []string{"hello", "world"},
	},
  {
    input:    "   game    over  ",
    expected: []string{"game", "over"},
  },
  {
    input:    "   please   try    again",
    expected: []string{"please", "try", "again"},
  },
  {
    input:    "   you   shall   not   pass",
    expected: []string{"you", "shall", "not", "pass"},
  },
  {
    input:    "",
    expected: []string{},
  },
}

  for _, c := range cases {
	actual := cleanInput(c.input)
  if len(actual) != len(c.expected){
    t.Errorf("Expected Length %d, got %d", len(c.expected), len(actual))
    continue
  }
	  for i := range actual {
		  word := actual[i]
		  expectedWord := c.expected[i]
      if word != expectedWord {
        t.Errorf("At index %d: expected word %s, got %s", i, expectedWord, word)
     }
	  }
  }

}
