package model

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestReadIndexReturnsValidMapAndNumberOfEntries(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte("xxx|MK17 9AU\nyyy|MK17 8AQ")))

	m, c := readIndex(r)
	if c != 2 {
		t.Error("Invalid number of map entries created from reader.  Expected 2 but got:", c)
	}
	if m["xxx"] != "MK17 9AU" || m["yyy"] != "MK17 8AQ" {
		t.Error("Map entry returned does not match reader contents:", m)
	}
}

func TestWriteIndexCreatesValidWriteBuffer(t *testing.T) {
	expectedBuff := "xxx|MK17 9AU\nyyy|MK17 8AQ\n"
	m := map[string]string{
		"xxx": "MK17 9AU",
		"yyy": "MK17 8AQ",
	}

	b := new(bytes.Buffer)
	for key, val := range m {
		fmt.Fprintf(b, "%s|%s\n", key, val)
	}

	if b.String() != expectedBuff {
		t.Error("Invalid write buffer created from map", b.String())
	}
}
