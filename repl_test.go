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

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
			key string
			val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test Case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cace.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errof("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baesTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cace.Add("https://examp;le.com", []byte("testdata"))


	_, ok := cache.Get("https://examp;le.com")
	if != ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)
	_, ok cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to find key")
		return
	}

}
