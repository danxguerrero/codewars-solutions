// 7 kyu Vowel Count
// Return the number (count) of vowels in the given string.
// We will consider a, e, i, o, u as vowels for this Kata (but not y).
// The input string will only consist of lower case letters and/or spaces.

package kata

func GetCount(str string) (count int) {
  vowels := []byte{'a','e','i','o','u'}
  count = 0
  for i:=0; i < len(vowels); i++ {
    for j := 0; j < len(str); j++ {
      if vowels[i]==str[j]{
        count++
      }
    }
  }
  return count
}