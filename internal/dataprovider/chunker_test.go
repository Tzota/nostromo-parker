package dataprovider

import (
	"strings"
	"testing"
)

type SimpleUnixReader struct {
	answer string
}

func (r SimpleUnixReader) Read(p []byte) (n int, err error) {
	copy(p, r.answer)
	return len(r.answer), nil
}

func TestPositive(t *testing.T) {
	conn := SimpleUnixReader{"passed!"}
	c := GetChunker(conn)
	data := <-c

	if strings.Compare(string(data), "passed!") != 0 {
		t.Errorf("wrong data came back (%s, %d, %d)", string(data), strings.Compare(string(data), "passed!"), len(data))
	}
}
