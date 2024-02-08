//8 kyu Rock Paper Scissors!

// Let's play! You have to return which player won! In case of a draw return Draw!.

// Examples(Input1, Input2 --> Output):

// "scissors", "paper" --> "Player 1 won!"
// "scissors", "rock" --> "Player 2 won!"
// "paper", "paper" --> "Draw!"

package kata

import "fmt"

func Rps(p1, p2 string) string {
  if p1 == "rock" && p2 == "paper" {
    return fmt.Sprintf("Player 2 won!")
  } else if p1 == "paper" && p2 == "rock" {
    return fmt.Sprintf("Player 1 won!")
  } else if p1 == "scissors" && p2 == "rock" {
    return fmt.Sprintf("Player 2 won!")
  } else if p1 == "rock" && p2 == "scissors" {
    return fmt.Sprintf("Player 1 won!")
  } else if p1 == "scissors" && p2 == "paper" {
    return fmt.Sprintf("Player 1 won!")
  } else if p1 == "paper" && p2 == "scissors" {
    return fmt.Sprintf("Player 2 won!")
  }
  return fmt.Sprintf("Draw!")
}