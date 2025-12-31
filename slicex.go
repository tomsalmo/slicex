package slicex

import (
	"errors"
	"iter"
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

// Filter returns a new slice of the elements of slice s that match the callback
func Filter[T ~[]E, E comparable](s T, match func(E) bool) []E {
	ss := make([]E, 0, len(s))
	return slices.AppendSeq(ss, Filtered(s, match))
}

// Filtered returns an iter.Seq over the values of slice s that evaluate true with the match filter.
func Filtered[Slice ~[]E, E any](s Slice, match func(E) bool) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, v := range s {
			if !match(v) {
				continue
			}

			if !yield(v) {
				return
			}
		}
	}
}
