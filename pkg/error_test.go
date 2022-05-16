package goa

import (
	"errors"
	"reflect"
	"testing"
)

type MyError struct {
	name string
}

func (e MyError) Error() string {
	return e.name
}

func TestServiceError_Unwrap(t *testing.T) {

	t.Run("unwrap service error", func(t *testing.T) {
		err := &MyError{name: "wrapped error"}
		se := &ServiceError{
			Name: "service error",
			err:  err,
		}
		if !errors.Is(se, err) {
			t.Errorf("expected errors.Is(ServiceError, err) = true, but false")
		}
		var target *MyError
		if !errors.As(se, &target) {
			t.Fatalf("expected errors.As(ServiceError, MyError) = true, but false")
		}
		if got, want := target.Error(), err.Error(); got != want {
			t.Errorf("expected %v, got %v", want, got)
		}
	})

	t.Run("unwrap merged service error", func(t *testing.T) {
		err1 := errors.New("wrapped error #1")
		se1 := &ServiceError{
			Name: "service error #1",
			err:  err1,
		}
		err2 := &MyError{name: "wrapped error #2"}
		se2 := &ServiceError{
			Name: "service error #2",
			err:  err2,
		}

		se := MergeErrors(se1, se2)

		if !errors.Is(se, err1) {
			t.Errorf("expected errors.Is(ServiceError, err1) = true, but false")
		}
		if !errors.Is(se, err2) {
			t.Errorf("expected errors.Is(ServiceError, err2) = true, but false")
		}

		var target *MyError
		if !errors.As(se, &target) {
			t.Fatalf("expected errors.As(ServiceError, MyError) = true, but false")
		}
		if got, want := target.Error(), err2.Error(); got != want {
			t.Errorf("expected %v, got %v", want, got)
		}
	})

	t.Run("service error is still service error", func(t *testing.T) {
		field := "some field"
		method := func() error {
			return &ServiceError{
				Name:    "some method error",
				ID:      "#123",
				Message: "method error",
				Field:   &field,
				err:     nil,
			}
		}
		err := method()
		se := &ServiceError{Name: "fault"}
		if !errors.As(err, &se) {
			t.Errorf("expected errors.As(err, ServiceError) = true, but false")
		}
		if !reflect.DeepEqual(err, se) {
			t.Errorf("expected err == se, but false, err=%#+v, se=%#+v", err, se)
		}
	})
}
