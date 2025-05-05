package main

import "testing"



func TestCleanInput(t *Testing.T) {
  cases := []struct {
    input string
    expected []string
  }{

    input: "  hello world ",
    expected: []string {"hello", "world"},
  },
  //Add more test cases just setting up for now

  for _, c := range cases {
    actual := cleanInput(c.input)
    //Need to check the length of the actual silce against the expected silce
    //if they don't match, I will use t.Errorf to print an error messagge
    //and fail the test
    for i := range actual {
      word := actual[i]
      expectedWord := c.expected[i]
    }
  }
}
