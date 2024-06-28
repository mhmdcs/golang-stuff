package hello

import "testing"

func TestSay(t *testing.T) {
	subtests := []struct {
		items []string
		result string
	} {
		{
		result: "Hello, world!",
		},
		{
		items: []string{"mhmd"},
		result: "Hello, mhmd!",
		},
		{
		items: []string{"saeed", "walaa"},
		result: "Hello, saeed, walaa!",
		},
	}

	for _, sb := range subtests {
		if s := Say(sb.items); s != sb.result {
			t.Errorf("wanted %s %v, got %s", sb.result, sb.items, s)
		} 
	}

	//want := "Hello, test!"
	//got := Say([]string{"test"})
	//if want != got {
	//	t.Errorf("wanted %s, got %s", want, got)
	//}
}
