package slicex

import (
	"errors"
	"slices"
)

var (
	ErrNotFound = errors.New("element not found in slice")
)

// Find is like slices.Index but returns the element instead of the index, and an error if it is not found.
func Find[T ~[]E, E comparable](s T, v E) (E, error) {
	idx := slices.Index(s, v)
	if idx == -1 {
		var zero E
		return zero, ErrNotFound
	}

	return s[idx], nil
}

// FindFunc is like slices.IndexFunc but returns the element instead of the index, and an error if it is not found.
func FindFunc[T ~[]E, E any](s T, fn func(E) bool) (E, error) {
	idx := slices.IndexFunc(s, fn)
	if idx == -1 {
		var zero E
		return zero, ErrNotFound
	}

	return s[idx], nil
}

// Filter returns a new slice of the elements that match the callback
func Filter[T ~[]E, E comparable](s T, fn func(E) bool) []E {
	ss := slices.Clone(s)
	return slices.DeleteFunc(ss, func(e E) bool {
		return !fn(e)
	})
}
