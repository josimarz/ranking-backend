package inmemory

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	result := m.Run()
	ClearDatabase()
	os.Exit(result)
}
