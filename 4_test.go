package tooffer

import (
	"testing"
)

func TestEscapeSpace(t *testing.T) {
	cases := []struct {
		in   []byte
		want []byte
	}{
		{
			[]byte("We are happy."),
			[]byte("We%20are%20happy."),
		},
	}

	for _, c := range cases {
		got := EscapeSpace(c.in)
		if string(got) != string(c.want) {
			t.Errorf("SpaceEscape failed, want: %s, got: %s.", c.want, got)
		}
	}
}
