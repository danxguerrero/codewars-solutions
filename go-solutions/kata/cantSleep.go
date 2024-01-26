// 8 kyu If you can't sleep, just count sheep!!
// If you can't sleep, just count sheep!!

// Task:
// Given a non-negative integer, 3 for example, return a string with a murmur: "1 sheep...2 sheep...3 sheep...". Input will always be valid, i.e. no negative integers.

package kata

import "fmt"

func countSheep(num int) string {
  count := ""
  for i := 1; i <= num; i++ {
    count = count + fmt.Sprintf("%d sheep...", i)
  }
  return count
}