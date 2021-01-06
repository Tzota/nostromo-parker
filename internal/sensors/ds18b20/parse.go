package ds18b20

import (
	"fmt"
	"regexp"
	"strconv"
)

var findRe *regexp.Regexp

func init() {
	findRe = regexp.MustCompile(`Temp C:\s*([\-\d\.]*)\s*`)
}

// ParseBytes tries to find a temperature info in serial data
func parseBytes(bytes []byte) (ds18b20Payload, error) {
	message := string(bytes)

	return parseString(message)
}

// ParseString tries to find a temperature info in serial data
func parseString(message string) (ds18b20Payload, error) {
	matches := findRe.FindStringSubmatch(message)

	if len(matches) < 2 {
		return ds18b20Payload{}, fmt.Errorf("no data (partial read?) `%s`", message)
	}

	temperature, err := strconv.ParseFloat(matches[1], 32)
	if err != nil {
		return ds18b20Payload{}, fmt.Errorf("can't find a temperature in the message")
	}

	return ds18b20Payload{float32(temperature)}, nil
}
