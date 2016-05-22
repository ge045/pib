package main

import (
	"fmt"
	"path"
	"testing"
)

func TestScan4Pictures(t *testing.T) {
	cases := []struct {
		in   string
		want []string
	}{
		{
			path.Join("test", "d1"),
			[]string{
				path.Join("test", "d1", "d11", "1.jpg"),
				path.Join("test", "d1", "d11", "1.png"),
			},
		},
		{
			path.Join("test", "d2"),
			[]string{
				path.Join("test", "d2", "d21", "2.png"),
				path.Join("test", "d2", "d21", "3.png"),
				path.Join("test", "d2", "d22", "1.png"),
				path.Join("test", "d2", "d22", "2.png"),
				path.Join("test", "d2", "d22", "3.png"),
			},
		},
	}
	for _, c := range cases {
		got := PictureScan(c.in)

		if got == nil && c.want == nil {
			fmt.Printf("NIL") //return 1
		} else if len(got) != len(c.want) {
			fmt.Printf("len mismatch") //return 1
			t.Errorf("%q == %q, want %q", c.in, got, c.want)
		} else {
			for i, v := range got {
				if v != c.want[i] {
					t.Errorf("%q == %q, want %q", c.in, v, c.want[i])
				}
			}
		}
	}
}
