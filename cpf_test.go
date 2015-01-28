package cpf

import (
	"strconv"
	"testing"
)

func TestValidWithDots(t *testing.T) {
	v := Valid("886.100.254-46")
	if v != true {
		t.Error("Expected true, got", v)
	}
}

func TestValidWithoutDots(t *testing.T) {
	v := Valid("88610025446")
	if v != true {
		t.Error("Expected true, got", v)
	}
}

func TestInValidWithDots(t *testing.T) {
	v := Valid("123.123.123.12")
	if v != false {
		t.Error("Expected false, got", v)
	}
}

func TestInValidWithoutDots(t *testing.T) {
	v := Valid("12312312312")
	if v != false {
		t.Error("Expected false, got", v)
	}
}

func TestInValidSameNumbersWithDots(t *testing.T) {
	var n string
	var v bool

	for i := 1; i <= 9; i++ {
		n = ""
		for x := 1; x <= 11; x++ {
			n += strconv.Itoa(i)
		}
		v = Valid(n)
		if v != false {
			t.Errorf("Expected false, got %v for %v", v, n)
		}
	}
}

func TestInValidSameNumbersWithoutDots(t *testing.T) {
	var n string
	var v bool

	for i := 1; i <= 9; i++ {
		n = ""
		for x := 1; x <= 11; x++ {
			if x == 4 || x == 7 {
				n += "."
			}
			if x == 10 {
				n += "-"
			}
			n += strconv.Itoa(i)
		}
		v = Valid(n)
		if v != false {
			t.Errorf("Expected false, got %v for %v", v, n)
		}
	}
}
