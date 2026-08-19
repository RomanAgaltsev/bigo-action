// Package main is a fixture for the contract workflow: a function with a
// bound bigo can prove, so a successful run produces a real report rather
// than an empty one.
package main

//bigo:max O(n)
func Sum(xs []int) int {
	total := 0
	for _, x := range xs {
		total += x
	}
	return total
}

func main() {}
