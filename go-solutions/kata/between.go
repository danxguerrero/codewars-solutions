//8 kyu What is between?
// Complete the function that takes two integers (a, b, where a < b) and return an array of all integers between the input parameters, including them.
// For example:
// a = 1
// b = 4
// --> [1, 2, 3, 4]

package kata

func Between(a, b int) []int {
  count := make([]int, b - a + 1)
  for i := range count {
    count[i] = a + i
  }
  return count
}