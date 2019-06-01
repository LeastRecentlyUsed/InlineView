package entities

import (
	"os"
	"testing"
)

var testStore PriceStore

func TestMain(m *testing.M) {
	setup()
	res := m.Run()
	os.Exit(res)
}

func setup() {
	testData := []PriceRecord{
		{"xxx", "MK17 9AU", "data string one"},
		{"yyy", "MK17 8AQ", "data string two"},
	}
	testStore.Store = testData
}

func Test_SyncIndexWithStore(t *testing.T) {
	logicTest := testStore
	logicTest.SyncIndexWithStore()

	if len(logicTest.Index) != 2 {
		t.Error("Expected 2 index entries, got", len(logicTest.Index))
	}
	if !logicTest.ModIndex {
		t.Error("Expected ModIndex to be true but got", logicTest.ModIndex)
	}
}

func Benchmark_SyncIndexWithStore(b *testing.B) {
	speedTest := testStore
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		speedTest.SyncIndexWithStore()
	}
}
