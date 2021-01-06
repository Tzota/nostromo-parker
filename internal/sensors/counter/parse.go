package counter

import (
	"fmt"
	"strconv"
	"strings"
)

func parseString(message string) (counterPayload, error) {
	m := strings.TrimSpace(message) // don't wanna regex
	val, err := strconv.Atoi(m)
	if err != nil {
		return counterPayload{}, fmt.Errorf("can't find integer in a message '%s'", message)
	}
	return counterPayload{Counter: val}, nil
}
