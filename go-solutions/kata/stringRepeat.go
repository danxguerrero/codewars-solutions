// 8 kyu String repeat

// Write a function that accepts an integer n and a string s as parameters, and returns a string of s repeated exactly n times.

// Examples (input -> output)
// 6, "I"     -> "IIIIII"
// 5, "Hello" -> "HelloHelloHelloHelloHello"

package kata

func RepeatStr(repetitions int, value string) string {
  var repeat string
  for i := 0; i < repetitions; i++ {
    repeat = repeat + value
  }
  return repeat
}