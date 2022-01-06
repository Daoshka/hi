package main

import "fmt"

func main() {

	var x int = 7
	fmt.Println(even(x))

	fmt.Println(max_el(8, 5, 4, 3, 2, 9, 10, 2, 4, 0, 5, 10.5, 12, 16.5, 1, 2.2))

}

func my_w() {

}

//even?
func even(x int) bool {
	if x%2 == 0 {
		return true
	} else {
		return false
	}
}

//MAX
func max_el(arg ...float64) float64 {

	mx := 0.0
	for _, v := range arg {
		if v > mx {
			mx = v
		} else {

		}

	}
	return mx

}
