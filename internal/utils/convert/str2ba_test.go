package convert

import (
	"bytes"
	"testing"
)

func TestNoString(t *testing.T) {
	_, err := Str2ba("")

	if err == nil {
		t.Error("Empty error should give an error")
	}
}

func TestNonMac(t *testing.T) {
	_, err := Str2ba("asdasd")

	if err == nil {
		t.Error("Non-MAC string should give an error")
	}
}

func TestKindsMac(t *testing.T) {
	_, err := Str2ba("TT:TT:TT:TT:TT:TT")

	if err == nil {
		t.Error("Wrong kinda-MAC string should give an error")
	}
}

func TestPositive(t *testing.T) {
	_, err := Str2ba("01:02:03:04:05:06")

	if err != nil {
		t.Error("Cant parse valid MAC")
	}
}

func TestNeedLittleEndian(t *testing.T) {
	actual, err := Str2ba("01:02:03:04:05:06")

	if err != nil {
		t.Error("Cant parse valid MAC")
	}
	expected := [6]byte{6, 5, 4, 3, 2, 1}
	if 0 != bytes.Compare(actual[:], expected[:]) {
		t.Error("not little endian")
	}
}
