package isogram

import (
	"isogram/bst"
	"strings"
)

// IsIsogram returns wheather a world is an isogram (has no repeating letters except dashes and spaces)
func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	print(word, "\n")
	var tree bst.Tree
	for _, char := range word {
		if char == ' ' || char == '-' {
			continue
		}
		found := tree.Insert(byte(char))
		if found == 0 {
			print(string(char), "\n\n")
			return false
		}
	}
	return true
}
