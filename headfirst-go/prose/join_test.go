package prose

import (
	"fmt"
	"testing"
)

type testData struct {
	list []string
	want string
}

func TestJoinWithSeparator(t *testing.T) {
	tests := []testData{
		testData{list: []string{}, want: ""},
		testData{list: []string{"apple"}, want: "apple"},
		testData{list: []string{"apple", "orange"}, want: "apple and orange"},
		testData{list: []string{"apple", "orange", "pear"}, want: "apple, orange, and pear"},
	}
	for _, test := range tests {
		got := JoinWithSeparator(test.list, ",")
		if got != test.want {
			t.Errorf(errorString(test.list, got, test.want))
		}
	}
}

func TestOneElement(t *testing.T) {
	list := []string{"apple"}
	want := "apple"
	got := JoinWithSeparator(list, ",")
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	want := "apple and orange"
	got := JoinWithSeparator(list, ",")
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	want := "apple, orange, and pear"
	got := JoinWithSeparator(list, ",")
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithSeparator(%#v) = \"%s\", want \"%s\"", list, got, want)
}
