// 7 kyu Round up to the next multiple of 5
// Given an integer as input, can you round it to the next (meaning, "greater than or equal") multiple of 5?

// Examples:

// input:    output:
// 0    ->   0
// 2    ->   5
// 3    ->   5
// 12   ->   15
// 21   ->   25
// 30   ->   30
// -2   ->   0
// -5   ->   -5
// etc.
// Input may be any positive or negative integer (including 0).

// You can assume that all inputs are valid integers.

package kata


func RoundToNext5(n int) int {
	remainder := n % 5

	if remainder == 0 {
		return n
	}

	if n < 0 {
		return n - remainder
	}

	return n + (5 - remainder)
}

// I saw a very clever solution for this that used a while loop:
// for n % 5 != 0 {n++}