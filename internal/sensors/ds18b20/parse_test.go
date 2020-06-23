package ds18b20

import "testing"

func TestPositive(t *testing.T) {
	recv := []byte("Temp C: 123.456\r\n\r\n")

	message, _ := parseBytes(recv)

	if message.Temperature != 123.456 {
		t.Errorf("%f is not 123.456 that I expected", message.Temperature)
	}
}

func TestPartialData(t *testing.T) {
	recv := []byte("C: 123.456\r\n\r\n")
	_, err := parseBytes(recv)

	if err == nil {
		t.Error("partial read should give no data")
	}
}

func TestBadFloat(t *testing.T) {
	recv := []byte("Temp C: .....\r\n\r\n")

	_, err := parseBytes(recv)

	if err == nil {
		t.Errorf("how did you parsed '.....' as a float?..")
	}
}
