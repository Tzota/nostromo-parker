package counter

import (
	"fmt"
	"strconv"
	"strings"
)

func parseString(message string) (Message, error) {
	m := strings.TrimSpace(message) // don't wanna regex
	val, err := strconv.Atoi(m)
	if err != nil {
		return Message{}, fmt.Errorf("can't find integer in a message '%s'", message)
	}
	return Message{Counter: val}, nil
}
