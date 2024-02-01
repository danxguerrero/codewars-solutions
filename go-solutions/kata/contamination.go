// 8 kyu Contamination #1 -String-

// An AI has infected a text with a character!!

// This text is now fully mutated to this character.

// If the text or the character are empty, return an empty string.
// There will never be a case when both are empty as nothing is going on!!

// Note: The character is a string of length 1 or an empty string.

// Example
// text before = "abc"
// character   = "z"
// text after  = "zzz"

package kata


func Contamination(text, char string) string {
  var result string
  if text == "" || char == ""{
    return ""
  }
  
  for i := 0; i < len(text ); i++{
    result = result + char
  }
  return result
}