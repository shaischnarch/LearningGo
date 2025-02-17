// Package Learning Go
// Simple packege with add function for practicing writing a public go library
package learninggo

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// Add func for adding two integers
// https://www.mathisfun.com/numbers/addition.html
func Add[T Number](x T, y T) T {
	return x + y
}