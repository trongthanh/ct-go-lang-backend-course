package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(s string) string {
	sorted := strings.Split(s, "")
	sort.Strings(sorted)
	// fmt.Println(s, "->", sorted);
	return strings.Join(sorted, "")
}

func isAnagram(s string, t string) bool {
	return sortString(s) == sortString(t)
}


func main() {
	fmt.Println("expect true", isAnagram("anagram", "nagaram"))
	fmt.Println("expect true", isAnagram("team", "meat"))
	fmt.Println("expect false", isAnagram("rat", "car"))
	fmt.Println("expect false", isAnagram("she", "his"))
}



