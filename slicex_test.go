package slicex_test

import (
	"testing"

	"github.com/tomsalmo/slicex"
)

func TestFind(t *testing.T) {
	t.Run("find existing int", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		got, err := slicex.Find(slice, 3)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if got != 3 {
			t.Errorf("got %v, want %v", got, 3)
		}
	})

	t.Run("int not found", func(t *testing.T) {
		slice := []int{1, 2, 4, 5}
		got, err := slicex.Find(slice, 3)

		if err != slicex.ErrNotFound {
			t.Errorf("expected ErrNotFound, got %v", err)
		}
		if got != 0 {
			t.Errorf("got %v, want %v", got, 0)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		var slice []string
		got, err := slicex.Find(slice, "test")

		if err != slicex.ErrNotFound {
			t.Errorf("expected ErrNotFound, got %v", err)
		}
		if got != "" {
			t.Errorf("got %q, want empty string", got)
		}
	})

	t.Run("find string", func(t *testing.T) {
		slice := []string{"foo", "bar", "baz"}
		got, err := slicex.Find(slice, "bar")

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if got != "bar" {
			t.Errorf("got %q, want %q", got, "bar")
		}
	})

	t.Run("find struct", func(t *testing.T) {
		type person struct {
			name string
			age  int
		}
		people := []person{
			{name: "Alice", age: 25},
			{name: "Bob", age: 30},
		}
		want := person{name: "Bob", age: 30}
		got, err := slicex.Find(people, want)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestFindFunc(t *testing.T) {
	t.Run("find existing int", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		got, err := slicex.FindFunc(slice, func(x int) bool { return x == 3 })

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if got != 3 {
			t.Errorf("got %v, want %v", got, 3)
		}
	})

	t.Run("int not found", func(t *testing.T) {
		slice := []int{1, 2, 4, 5}
		got, err := slicex.FindFunc(slice, func(x int) bool { return x == 3 })

		if err != slicex.ErrNotFound {
			t.Errorf("expected ErrNotFound, got %v", err)
		}
		if got != 0 {
			t.Errorf("got %v, want %v", got, 0)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		var slice []string
		got, err := slicex.FindFunc(slice, func(s string) bool { return s == "test" })

		if err != slicex.ErrNotFound {
			t.Errorf("expected ErrNotFound, got %v", err)
		}
		if got != "" {
			t.Errorf("got %q, want empty string", got)
		}
	})

	t.Run("find person", func(t *testing.T) {
		type person struct {
			name string
			age  int
		}
		people := []person{
			{name: "Alice", age: 25},
			{name: "Bob", age: 30},
		}

		got, err := slicex.FindFunc(people, func(p person) bool { return p.age == 30 })

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		want := person{name: "Bob", age: 30}
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("person not found", func(t *testing.T) {
		type person struct {
			name string
			age  int
		}
		people := []person{
			{name: "Alice", age: 25},
			{name: "Bob", age: 30},
		}

		got, err := slicex.FindFunc(people, func(p person) bool { return p.age == 40 })

		if err != slicex.ErrNotFound {
			t.Errorf("expected ErrNotFound, got %v", err)
		}
		if got != (person{}) {
			t.Errorf("got %v, want empty person", got)
		}
	})
}
