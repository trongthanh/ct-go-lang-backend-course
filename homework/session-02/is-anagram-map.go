package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	mapS := make(map[string]int)
	mapT := make(map[string]int)

	for _, char := range s {
		key := string(char)
		mapS[key]++
	}

	for _, char := range t {
		key := string(char)
		mapT[key]++
	}

	return mapsEqual(mapS, mapT)
}

// shamelessly copy from chatgpt
// Function to check if two maps are equal
func mapsEqual(map1, map2 map[string]int) bool {
	// Check the length of the maps
	if len(map1) != len(map2) {
		return false
	}

	// Iterate over the keys of map1
	for key, value1 := range map1 {
		// Check if the key exists in map2
		value2, ok := map2[key]
		if !ok {
			return false
		}

		// Check if the values match
		if value1 != value2 {
			return false
		}
	}

	return true
}


func main() {
	fmt.Println("expect true", isAnagram("anagram", "nagaram"))
	fmt.Println("expect true", isAnagram("team", "meat"))
	fmt.Println("expect false", isAnagram("rat", "car"))
	fmt.Println("expect false", isAnagram("she", "his"))
}



