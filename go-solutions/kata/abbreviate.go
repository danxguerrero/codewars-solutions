// 8kyu Abbreviate a Two Word Name
// Write a function to convert a name into initials. This kata strictly takes two words with one space in between them.
// The output should be two capital letters with a dot separating them.
// It should look like this:
// Sam Harris => S.H
// patrick feeney => P.F

package kata

import (
  "fmt"
  "strings"
)

func AbbrevName(name string) string{
  split_name := strings.Split(name, " ")
  first_initial := strings.ToUpper(string(split_name[0][0]))
  last_initial := strings.ToUpper(string(split_name[1][0]))
  initials := fmt.Sprintf("%s.%s", first_initial, last_initial)
  return initials
}