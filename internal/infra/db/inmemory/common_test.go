package inmemory

import "testing"

func TestClearDatabase(t *testing.T) {
	ClearDatabase()
	if len(ranks) > 0 || len(attrs) > 0 || len(entries) > 0 {
		t.Error("database was not cleared")
	}
}
