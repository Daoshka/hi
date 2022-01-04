package main

import "fmt"



func main() {
	//output string\

	/*
	var x string ="Hi friends"
    fmt.Println(x)

	fmt.Println("32132x42452=", 32132*42452)
	*/
	
	//input strin

	/*
	fmt.Print("Enter a number: ")
    var input float64
    fmt.Scanf("%f", &input)

    output := input * 2

    fmt.Println(output)
	*/ 
	
	//F in C
	/*
	var F float64
	var C float64
	fmt.Println("Enter temperature in Fahrenheit: ")
	fmt.Scanf("%f", &F)
	C=(F-32)*5/9
	fmt.Println("Fahrenheit-", F,"\nCelsius=", C)
	*/

	//которая выводит числа от 1 до 100. Но для кратных трём нужно вывести «Fizz» вместо числа, для кратных пяти - «Buzz», а для кратных как трём, так и пяти — «FizzBuzz»
	/*
	for i:=1; i<=100; i++{
		
	if i%3==0&&i%5==0{ fmt.Println("FizzBuzz")
	}else if i%3==0 { fmt.Println("Fizz")
	}else if i%5==0 { fmt.Println("Buzz")
	}else {fmt.Println(i)}
	}
	*/

	//switch
	/*
	for i:=1; i<=100; i++{
	switch i {
		case 0: fmt.Println("Zero")
		case 1: fmt.Println("One")
		case 2: fmt.Println("Two")
		case 3: fmt.Println("Three")
		case 4: fmt.Println("Four")
		case 5: fmt.Println("Five")
		default: fmt.Println("Unknown Number")
		}
	}
	*/
	//среднее массива х
	/*
	x := [5]float64{ 98, 93, 77, 82, 83 }
	var total float64 = 0
for _, value := range x {
    total += value
}
fmt.Println(total / float64(len(x)))
*/

//append
/*
slice1 := []int{1,2,3}
slice2 := append(slice1, 4, 5)
fmt.Println(slice1, slice2)
*/

//maps
/*
    elements := make(map[string]string)
    elements["H"] = "Hydrogen"
    elements["He"] = "Helium"
    elements["Li"] = "Lithium"
    elements["Be"] = "Beryllium"
    elements["B"] = "Boron"
    elements["C"] = "Carbon"
    elements["N"] = "Nitrogen"
    elements["O"] = "Oxygen"
    elements["F"] = "Fluorine"
    elements["Ne"] = "Neon"

    fmt.Println(elements)
	*/

	//min x[]
	x:=[]int {48,96,86,68,57,82,63,70,37,34,83,27,19,97,9,17,
	}
	min:=x[0] 
		for y:=1;y<len(x); y++{
		min1:=x[y]
		if min1<min		{
		min=min1
		} else{}
	}
	fmt.Println(min)

}