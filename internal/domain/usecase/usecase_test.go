package usecase

import (
	"testing"
)

func TestResourceNotFoundError(t *testing.T) {
	name := "rank"
	id := "539c3211-c7b7-4aa1-8385-1b60f46cf801"
	e := &ResourceNotFoundError{name, id}
	t.Run("Error", func(t *testing.T) {
		want := "rank not found: 539c3211-c7b7-4aa1-8385-1b60f46cf801"
		if got := e.Error(); got != want {
			t.Errorf("Error() got %v, want %v", got, want)
		}
	})
}
