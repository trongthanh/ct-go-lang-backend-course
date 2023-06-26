// Problem source: https://leetcode.com/problems/valid-parentheses/
package main

import (
	"fmt"
	"strings"
	"thanhtran-s01/stack"
)

var bracketPairs = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
}

func isValid(s string) bool {
	slice := strings.Split(s, "")
	bracketStack := stack.New()

	valid := true
	for _, char := range slice {
		if (char == "(" || char == "[" || char == "{") {
			bracketStack.Push(char)
			// fmt.Println("stack pushed:", bracketStack)
		} else if (char == ")" || char == "]" || char == "}") {
			if (bracketStack.Len() == 0) {
				valid = false
			} else {
				topStack := bracketStack.Pop().(string)
				// fmt.Println("stack pop:", topStack)
				if (bracketPairs[topStack] != char) {
					valid = false
				}
			}
		}
	}

	// there are unclosed bracketStacks
	if (bracketStack.Len() > 0) {
		valid = false
	}

	return valid
}


func main() {
	fmt.Println("expect false; receive: ",isValid(")"))
	// valid
	fmt.Println("expect true; receive: ", isValid("()"))
	// valid
	fmt.Println("expect true; receive: ",isValid("()[]{}"))
	// valid
	fmt.Println("expect true; receive: ",isValid("{[]}"))
	// valid
	fmt.Println("expect true; receive: ",isValid("{[]}"))
	// invalid
	fmt.Println("expect false; receive: ",isValid("(]"))
	// invalid
	fmt.Println("expect false; receive: ",isValid("()[}"))
	// invalid
	fmt.Println("expect false; receive: ",isValid("["))
	// invalid
	// invalid
	fmt.Println("expect false; receive: ",isValid("{}("))
}
