package slicex_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestFilter(t *testing.T) {
	t.Run("filter ints with matches", func(t *testing.T) {
		slice := []int{1, 2, 3, 2, 4, 2, 5}
		got := slicex.Filter(slice, func(e int) bool { return e == 2 })
		want := []int{2, 2, 2}

		if len(got) != len(want) {
			t.Errorf("got length %v, want length %v", len(got), len(want))
		}
		for i, v := range got {
			if v != want[i] {
				t.Errorf("at index %d: got %v, want %v", i, v, want[i])
			}
		}
	})

	t.Run("filter ints with no matches", func(t *testing.T) {
		slice := []int{1, 3, 4, 5}
		got := slicex.Filter(slice, func(e int) bool { return e == 2 })

		if len(got) != 0 {
			t.Errorf("got length %v, want length 0", len(got))
		}
	})

	t.Run("filter ints all match", func(t *testing.T) {
		slice := []int{3, 3, 3, 3}
		got := slicex.Filter(slice, func(e int) bool { return e == 3 })
		want := []int{3, 3, 3, 3}

		if len(got) != len(want) {
			t.Errorf("got length %v, want length %v", len(got), len(want))
		}
		for i, v := range got {
			if v != want[i] {
				t.Errorf("at index %d: got %v, want %v", i, v, want[i])
			}
		}
	})

	t.Run("filter empty slice", func(t *testing.T) {
		var slice []int
		got := slicex.Filter(slice, func(e int) bool { return e == 1 })

		if len(got) != 0 {
			t.Errorf("got length %v, want length 0", len(got))
		}
	})

	t.Run("filter single element match", func(t *testing.T) {
		slice := []int{5}
		got := slicex.Filter(slice, func(e int) bool { return e == 5 })
		want := []int{5}

		if len(got) != len(want) {
			t.Errorf("got length %v, want length %v", len(got), len(want))
		}
		if got[0] != want[0] {
			t.Errorf("got %v, want %v", got[0], want[0])
		}
	})

	t.Run("filter single element no match", func(t *testing.T) {
		slice := []int{5}
		got := slicex.Filter(slice, func(e int) bool { return e == 3 })

		if len(got) != 0 {
			t.Errorf("got length %v, want length 0", len(got))
		}
	})

	t.Run("filter strings", func(t *testing.T) {
		slice := []string{"foo", "bar", "baz", "foo", "qux", "foo"}
		got := slicex.Filter(slice, func(e string) bool { return e == "foo" })
		want := []string{"foo", "foo", "foo"}

		if len(got) != len(want) {
			t.Errorf("got length %v, want length %v", len(got), len(want))
		}
		for i, v := range got {
			if v != want[i] {
				t.Errorf("at index %d: got %q, want %q", i, v, want[i])
			}
		}
	})

	t.Run("filter structs", func(t *testing.T) {
		type person struct {
			name string
			age  int
		}
		people := []person{
			{name: "Alice", age: 25},
			{name: "Bob", age: 30},
			{name: "Alice", age: 25},
			{name: "Charlie", age: 35},
		}
		target := person{name: "Alice", age: 25}
		got := slicex.Filter(people, func(e person) bool { return cmp.Equal(e, target, cmp.AllowUnexported(person{})) })
		want := []person{
			{name: "Alice", age: 25},
			{name: "Alice", age: 25},
		}

		if len(got) != len(want) {
			t.Errorf("got length %v, want length %v", len(got), len(want))
		}
		for i, v := range got {
			if v != want[i] {
				t.Errorf("at index %d: got %v, want %v", i, v, want[i])
			}
		}
	})

	t.Run("original slice unchanged", func(t *testing.T) {
		original := []int{1, 2, 3, 2, 4, 2, 5}
		want := []int{1, 2, 3, 2, 4, 2, 5}

		_ = slicex.Filter(original, func(e int) bool { return e == 2 })

		if len(original) != len(want) {
			t.Errorf("original slice length changed: got %v, want %v", len(original), len(want))
		}
		for i, v := range original {
			if v != want[i] {
				t.Errorf("original slice modified at index %d: got %v, want %v", i, v, want[i])
			}
		}
	})
}

func BenchmarkFilter(b *testing.B) {
	b.Run("small slice some matches", func(b *testing.B) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = slicex.Filter(slice, func(e int) bool { return e%2 == 0 })
		}
	})

	b.Run("large slice some matches", func(b *testing.B) {
		slice := make([]int, 10000)
		for i := range slice {
			slice[i] = i
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = slicex.Filter(slice, func(e int) bool { return e%2 == 0 })
		}
	})

	b.Run("small slice no matches", func(b *testing.B) {
		slice := []int{1, 3, 5, 7, 9}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = slicex.Filter(slice, func(e int) bool { return e%2 == 0 })
		}
	})

	b.Run("small slice all matches", func(b *testing.B) {
		slice := []int{2, 4, 6, 8, 10}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = slicex.Filter(slice, func(e int) bool { return e%2 == 0 })
		}
	})
}
