// 7 kyu Highest and Lowest
// In this little assignment you are given a string of space separated numbers, and have to return the highest and lowest number.

// Examples
// HighAndLow("1 2 3 4 5")  // return "5 1"
// HighAndLow("1 2 -3 4 5") // return "5 -3"
// HighAndLow("1 9 3 4 -5") // return "9 -5"
// Notes
// All numbers are valid Int32, no need to validate them.
// There will always be at least one number in the input string.
// Output string must be two numbers separated by a single space, and highest number is first.

package kata

import (
        "fmt"
        "strings"
        "strconv"
        )

func HighAndLow(in string) string { 
  nums := convertToNums(in)
  lowest := nums[0]
  highest:= nums[0]
  
  for i, _ := range nums {
    if lowest > nums[i] {
      lowest = nums[i]
    }
    if highest < nums[i] {
      highest = nums[i]
    }
  }
  
  return fmt.Sprintf("%d %d", highest, lowest)
}

func convertToNums(in string) []int {
  split := strings.Split(in, " ")
  nums := make([]int, len(split))
  for i, _ := range split {
    if num, err := strconv.Atoi(split[i]); err == nil{
      nums[i] = num
    }
  }
  return nums
}