package dataprovider

import (
	"testing"
)

type SimpleUnixReader struct {
	answer string
}

func (r SimpleUnixReader) Read(fd int, p []byte) (n int, err error) {
	copy(p, r.answer)
	return len(r.answer), nil
}

func TestPositive(t *testing.T) {
	t.Skip()
	c := GetChunker(SimpleUnixReader{"passed!"}, 1)
	data := <-c

	if string(data) != "passed!" {
		t.Error("no data came back")
	}
}
