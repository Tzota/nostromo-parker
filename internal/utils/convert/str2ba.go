package convert

import (
	"fmt"
	"strconv"
	"strings"
)

// Str2ba converts MAC address string representation to little-endian byte array
func Str2ba(addr string) ([6]byte, error) {
	var b [6]byte
	if strings.Count(addr, ":") != 5 {
		return b, fmt.Errorf("seems like it's not a MAC address")
	}
	a := strings.Split(addr, ":")

	for i, tmp := range a {
		u, err := strconv.ParseUint(tmp, 16, 8)
		if err != nil {
			return b, err
		}
		b[len(b)-1-i] = byte(u)
	}
	return b, nil
}
