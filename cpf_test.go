package cpf

import "testing"

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
