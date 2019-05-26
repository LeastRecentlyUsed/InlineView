package model

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var priceKeys map[string]string

func TestMain(m *testing.M) {
	setup()
	res := m.Run()
	os.Exit(res)
}

func setup() {
	priceKeys = map[string]string{
		"xxx": "MK17 9AU",
		"yyy": "MK17 8AQ",
	}
}

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

func TestReadIndexCanReturnEmptyMap(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte("")))

	m, c := readIndex(r)
	if c > 0 {
		t.Error("Invalid number of map entries created from empty reader.  Expected zero but got:", c)
	}
	if len(m) > 0 {
		t.Error("Map should be empty but has a length greater than zero:", len(m))
	}
}

func TestWriteIndexCreatesValidWriterContents(t *testing.T) {
	expectedBuff := "xxx|MK17 9AU\nyyy|MK17 8AQ\n"
	res := writeIndex(priceKeys)
	out := fmt.Sprint(res)

	if out != expectedBuff {
		t.Error("Invalid write buffer created from map", out)
	}
}

func TestFoundIndexReturnsValidBool(t *testing.T) {
	var indexTable = []struct {
		in1 map[string]string
		in2 string
		out bool
	}{
		{priceKeys, "xxx", true},
		{priceKeys, "yyy", true},
		{priceKeys, "zzz", false},
	}

	for _, testEntry := range indexTable {
		res := foundIndex(testEntry.in1, testEntry.in2)
		if res != testEntry.out {
			t.Error("Invalid response from index search.  Expecting", testEntry.out, "but received", res, "for key:", testEntry.in2)
		}
	}
}
