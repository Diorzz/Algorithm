package list

import (
	"testing"
)

func TestList(t *testing.T) {
	l := NewList()

	l.TailAdd(1)
	l.TailAdd(2)
	if l.String() != "1,2" {
		t.Errorf("Expect 1,2 actual %s", l)
	}

	l.HeadAdd(3)
	if l.String() != "3,1,2" {
		t.Errorf("Expect 3,1,2 actual %s", l)
	}

	l.Reverse()
	if l.String() != "2,1,3" {
		t.Errorf("Expect 2,1,3 actual %s", l)
	}

	r := l.HasCircle()
	if r {
		t.Error("no circle in list but has")
	}

	last := l.last()
	last.next = l.one()
	r = l.HasCircle()
	if !r {
		t.Error("there has circle in list but no")
	}
}

func TestDeleteFromTail(t *testing.T) {
	l := NewList()

	//data := []int{1, 2, 3, 4}
	l.TailAdd(1)
	l.TailAdd(2)
	l.TailAdd(3)
	l.TailAdd(4)

	l.DeleteFromTail(2)
	if l.String() != "1,2,4" {
		t.Errorf("Delete from tail error, expect 1,2,4 actual %s", l)
	}
}

func TestListCombine(t *testing.T) {
	// normal combine
	one := NewList()
	one.TailAdd(1)
	one.TailAdd(3)
	one.TailAdd(5)

	two := NewList()
	two.TailAdd(2)
	two.TailAdd(4)
	two.TailAdd(6)

	one.Combine(two)
	if one.String() != "1,2,3,4,5,6" {
		t.Errorf("combine error, expect 1,2,3,4,5,6 actual %s", one)
	}

	// empty list combine
	one = NewList()
	two = NewList()
	one.Combine(two)
	if one.String() != "" {
		t.Errorf("combine error, expect ‘“”’ actual %s", one)
	}

	// same list combine
	one = NewList()
	two = NewList()

	one.TailAdd(1)
	one.TailAdd(2)
	two.TailAdd(1)
	two.TailAdd(2)
	one.Combine(two)
	if one.String() != "1,1,2,2" {
		t.Errorf("combine error, expect 1,1,2,2 actual %s", one)
	}
}
