package devide

import (
	"testing"
)

func TestD(t *testing.T) {
	test1, _ := D(10000, 3, true)
	test2, _ := D(10000, 3, false)
	test3, _ := D(5, 3, false)
	_, ok := D(1999, 1, true)
	if (test1[0] != 4000) && (test1[1] != 3000) && (test1[2] != 3000) {
		t.Error("Loi")
	}

	if (test2[0] != 3334) && (test2[1] != 3333) && (test2[2] != 3333) {
		t.Error("Loi")
	}

	if (test3[0] != 2) && (test3[1] != 2) && (test3[2] != 1) {
		t.Error("Loi")
	}

	if ok == true {
		t.Error("Loi")
	}

}
