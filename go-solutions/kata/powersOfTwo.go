// 8 kyu Powers of 2
// Complete the function that takes a non-negative integer n as input, and returns a list of all the powers of 2 with the exponent ranging from 0 to n ( inclusive ).

// Examples
// n = 0  ==> [1]        # [2^0]
// n = 1  ==> [1, 2]     # [2^0, 2^1]
// n = 2  ==> [1, 2, 4]  # [2^0, 2^1, 2^2]

package kata

func PowersOfTwo(n int) []uint64 {
  powers := make([]uint64, n+1)
  powers[0] = 1
  for i := 1; i <= n; i++ {
    powers[i] = powers[i-1] * 2
  }
  return powers
}