//go:build js && wasm

package localstorage

import "syscall/js"

var LocalStorage = js.Global().Get("localStorage")

// Get an item.
func Get(key string) string {
	v := LocalStorage.Call("getItem", key)

	if v.IsNull() {
		return ""
	}

	return v.String()
}

// Set an item.
func Set(key string, value string) {
	LocalStorage.Call("setItem", key, value)
}

// Remove an item.
func Remove(key string) {
	LocalStorage.Call("removeItem", key)
}
