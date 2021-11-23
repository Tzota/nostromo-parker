package dataprovider

import (
	"errors"
	"strconv"
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

	c, _ := GetChunker(conn)
	data := <-c

	if strings.Compare(string(data), "passed!") != 0 {
		t.Errorf("wrong data came back (%s, %d, %d)", string(data), strings.Compare(string(data), "passed!"), len(data))
	}
}

//---------------------------------------

var current int16

type FirstNSuccessReader struct {
	Count int16 // успешно вернет первые Count ответов
}

func (r FirstNSuccessReader) GetAnswer() (string, error) {
	if current < r.Count {
		current = current + 1
		return strconv.Itoa(int(current)), nil
	}

	return "", errors.New("bar")
}

func (r FirstNSuccessReader) Read(p []byte) (n int, err error) {
	ans, err := r.GetAnswer()
	if err != nil {
		return 0, err
	}
	copy(p, []byte(ans))
	return len(ans), nil
}

func Test2SuccessfullAnswers(t *testing.T) {
	conn := FirstNSuccessReader{Count: 2}

	c, flag := GetChunker(conn)

	gonnaBreak := false
	for {
		if gonnaBreak {
			break
		}

		select {
		case data, ok := <-c:
			if ok && strings.Compare(string(data), "") == 0 {
				t.Errorf("В тесте что-то пошло не так %s %v", string(data), conn)
			}
		case <-flag:
			gonnaBreak = true
		}
	}
	if !gonnaBreak {
		t.Error("Как это мы вышли без сигнального канала?..")
	}
}
