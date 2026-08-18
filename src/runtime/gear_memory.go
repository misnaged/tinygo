//go:build tinygo.wasm && gear

package runtime

func GearMemorySize() uintptr {
	return uintptr(wasm_memory_size(wasmMemoryIndex)) * wasmPageSize
}

func GearHeapStart() uintptr {
	return heapStart
}

func GearHeapEnd() uintptr {
	return heapEnd
}

func GearHeapSize() uintptr {
	return heapEnd - heapStart
}
