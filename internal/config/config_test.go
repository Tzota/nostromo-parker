package config

import (
	"reflect"
	"testing"
)

func TestReadFromFilePositive(t *testing.T) {
	cfg, err := ReadFromFile("../../test/testdata/config/config2.json")
	if err != nil {
		t.Error(err)
	}
	if len(cfg.Points) != 2 {
		t.Errorf("Got %d points instead of 2", len(cfg.Points))
	}
}

func TestReadFromFileAbsent(t *testing.T) {
	_, err := ReadFromFile("../../test/testdata/config/THERE_S_NO_SUCH_FILE.json")
	if err == nil {
		t.Error("somehow absent file was read")
	}
}

func TestReadFromFileMalformed(t *testing.T) {
	_, err := ReadFromFile("../../test/testdata/config/config_malformed.json")
	if err == nil {
		t.Error("somehow malformed file was read")
	}
}

func TestJsonTags(t *testing.T) {
	cfg, err := ReadFromFile("../../test/testdata/config/jsonTags.json")
	if err != nil {
		t.Error(err)
	}
	if len(cfg.Points) != 2 {
		t.Errorf("Got %d points instead of 2", len(cfg.Points))
	}

	exists := func(point Point) bool {
		for _, p := range cfg.Points {
			if reflect.DeepEqual(point, p) {
				return true
			}
		}
		return false
	}

	point := Point{Mac: "00:01:02:03:04:05", Kind: "foo"}
	if !exists(point) {
		t.Error("seems like som json tags are broken")
	}
}
