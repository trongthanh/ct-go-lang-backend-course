package main

import (
	"fmt"
	"runtime"
)

func main() {
	printAllocation()
}

func printAllocation() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)

}
