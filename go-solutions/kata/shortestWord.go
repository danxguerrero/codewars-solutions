// 7 kyu Shortest Word
// Simple, given a string of words, return the length of the shortest word(s).

// String will never be empty and you do not need to account for different data types.

package kata

import "strings"

func FindShort(s string) int {
  strings := strings.Split(s, " ")
  shortest := strings[0]
  for i, _ := range strings {
    if len(strings[i]) < len(shortest) {
      shortest = strings[i]
    }
  }
  return len(shortest)
}