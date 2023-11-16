//go:build js && wasm

package localstorage

import "testing"

func TestBasic(t *testing.T) {
	k := "hello"
	v := "world"

	t.Run("set", func(t *testing.T) {
		Set(k, v)
	})

	t.Run("get", func(t *testing.T) {
		rv := Get(k)

		if rv != v {
			t.Fatalf("expected a value of %s but got %s", rv, v)
		}
	})

	t.Run("remove", func(t *testing.T) {
		Remove(k)

		rv := Get(k)

		if rv != "" {
			t.Fatal("expected value to be removed but it remained")
		}
	})
}
