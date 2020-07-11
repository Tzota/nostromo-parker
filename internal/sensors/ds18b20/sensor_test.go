package ds18b20

import "testing"

func TestPositiveListenTo(t *testing.T) {
	dp := make(chan []byte)

	h := New()
	go h.ListenTo(dp)
	dp <- []byte("Temp C: 11.22\r\n")
	im := <-h.Messages
	m := im.(Message)
	if m.Temperature != 11.22 {
		t.Errorf("have %f instead of 11.22", m.Temperature)
	}
}
func TestNegativeListenTo(t *testing.T) {
	dp := make(chan []byte)

	h := New()
	go h.ListenTo(dp)
	dp <- []byte(" C: 11.22\r\n")
	if len(h.Messages) != 0 {
		t.Errorf("have %d instead of 11.22", len(h.Messages))
	}
}
