package usecase

import (
	"context"
	"fmt"
)

type Usecase[I any, O any] interface {
	Execute(context.Context, I) (*O, error)
}

type ResourceNotFoundError struct {
	name string
	id   string
}

func (e *ResourceNotFoundError) Error() string {
	return fmt.Sprintf("%v not found: %v", e.name, e.id)
}
