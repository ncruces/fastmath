package fastmath_test

import (
	"fmt"

	"github.com/ncruces/fastmath"
)

func ExampleSqrt() {
	fmt.Println(fastmath.Sqrt(2))
	// Output: 1.4633539
}

func ExampleCbrt() {
	fmt.Println(fastmath.Cbrt(2))
	// Output: 1.2996774
}

func ExampleRcpr() {
	fmt.Println(fastmath.Rcpr(2))
	// Output: 0.47474486
}

func ExampleRSqrt() {
	fmt.Println(fastmath.RSqrt(2))
	// Output: 0.71637243
}

func ExampleRCbrt() {
	fmt.Println(fastmath.RCbrt(2))
	// Output: 0.80025184
}
