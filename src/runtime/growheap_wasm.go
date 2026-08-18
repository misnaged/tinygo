//go:build tinygo.wasm && !gear

package runtime

func growHeap() bool {
	memorySize := wasm_memory_size(wasmMemoryIndex)

	result := wasm_memory_grow(wasmMemoryIndex, memorySize)
	if result == -1 {
		return false
	}

	setHeapEnd(
		uintptr(wasm_memory_size(wasmMemoryIndex) * wasmPageSize),
	)

	return true
}
