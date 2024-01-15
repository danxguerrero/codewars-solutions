// 8kyu Is it a palindrome?
// Write a function that checks if a given string (case insensitive) is a palindrome.
// A palindrome is a word, number, phrase, or other sequence of symbols that reads the same backwards as forwards, such as madam or racecar.
package kata

import "strings"

func IsPalindrome(str string) bool {
  lowerStr := strings.ToLower(str)
  
  for i := 0; i < len(lowerStr); i++ {
    if lowerStr[i] != lowerStr[len(lowerStr)-1-i] {
      return false
    }
  }
  return true
}