package operator

import "testing"

func TestCompareList(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 4}
	c := [...]int{1, 4, 2, 3}

	t.Log(a == b)
	t.Log(a == c)
}
