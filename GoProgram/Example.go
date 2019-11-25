package main

import (  
	"fmt"
	"math"
	"os"
	"strings"
)
    //Structure with two fields
	type Rectangle struct{
 
	height float64
	width float64
	}

	// Structure with one field
	type Circle struct{
		radius float64
		}
	// An interface defines a list of methods that a type must implement

	type Shape interface {
		area() float64
		}

	//Function declration Which takes two arguements of type int
	func calcSquares(number int, squareop chan int) {  
	fmt.Println("calcSquares Goroutine Executed")
    sum := 0
    for number != 0 {
        digit := number % 10
        sum += digit * digit
        number /= 10
    }
    squareop <- sum
	}

	//Function declration Which takes two arguements of type int
    func calcCubes(number int, cubeop chan int) {  
	fmt.Println("CalcCubes Goroutine Executed")
    sum := 0 
    for number != 0 {
        digit := number % 10
        sum += digit * digit * digit
        number /= 10
    }
    cubeop <- sum
	} 

	// Method Which has Reciever
	func (rect Rectangle) area() float64{
			return rect.width * rect.height 
	}

	// Method Which has Reciever
	func (c Circle) area() float64 {
		return math.Pi * math.Pow(c.radius, 2)
	} 

	//Function which is taking Interface instance as an arguement
	func getArea(shape Shape) float64{
 
		return shape.area()
		 
	}

	// Variadic function to join strings which takes multiple agruements
	func joinstr(element...string)string{ 
    	return strings.Join(element, "-") 
	} 

	//Function which takes pointer as an arguement
	func change(val *int) {
		*val = 55
	}

	func main() {  

	fmt.Println("------------------Variables and Types---------------------")

	//Variables Declration without initial Value
	var t int
	var u string
	var v float64
	var w bool

	fmt.Printf(" value of t is %v \n", t)
	fmt.Printf(" value of u is %v \n", u)
	fmt.Printf(" value of v is %v \n", v)
	fmt.Printf(" value of w is %v \n", w)
	fmt.Println("")

	//Variables  with initial Value
	t=10
	u="Hello"
	v=21.34
	w=true

	fmt.Printf(" value of t is %v \n", t)
	fmt.Printf(" value of u is %v \n", u)
	fmt.Printf(" value of v is %v \n", v)
	fmt.Printf(" value of w is %v \n", w)
	fmt.Println("")
	
	//Creation of channels
	number := 2
    sqrch := make(chan int)
	cubech := make(chan int)

	fmt.Println("------------------Goroutines and Channnels---------------------")

	//Calling Goroutines for two functions
    go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	
    squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final Square output", squares)
	fmt.Println("Final Cube output", cubes)
	
	rect := Rectangle{20, 50}
	circ := Circle{4}
	
	fmt.Println("Rectangle Area =", getArea(rect))
	fmt.Println("Circle Area =", getArea(circ))


	// Use of  switch Statement
    switch {
    case getArea(rect) <= 500:
        fmt.Println("Small Rectangle")
    case getArea(rect) <= 1000:
        fmt.Println("Medium Rectangle")
    case getArea(rect) > 1000:
        fmt.Println("Big Rectangle")
	}
	fmt.Println("")

	fmt.Println("------------------Arrays and Slices---------------------")

	// an array of 5 integers
	arr := [5]int{11, 12, 31, 4, 1}
    fmt.Println("Initial array is:", arr)
	fmt.Println("")
	
	// Sorting of an Array
    var min int = 0
    var tmp int = 0

    for i := 0; i < len(arr); i++ {
        min = i
        for j := i + 1; j < len(arr); j++ {
            if arr[j] < arr[min] {
                min = j
            }
        }
		
        tmp = arr[i]
        arr[i] = arr[min]
        arr[min] = tmp
    }

	fmt.Println("Sorted array:    ", arr)
	 fmt.Println("")
	
	var s []int = arr[1:4]

	fmt.Println("Slice s = ", s)
	slice1 := arr[1:4]
	slice2 := arr[:3]
	slice3 := arr[2:]
	slice4 := arr[:]

	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)
	fmt.Println("slice3 = ", slice3)
	fmt.Println("slice4 = ", slice4)
	
	fmt.Println("")

	fmt.Println("------------------Variadic function---------------------")


	fmt.Println(joinstr()) 
     
	// Passing multiple arguments to variadic function
	fmt.Println("After passing two arguements two function")
	fmt.Println(joinstr("Data", "Platform")) 
	fmt.Println("")
	fmt.Println("After passing three arguements two function")
	fmt.Println(joinstr("Flexera", "Data", "Platform")) 
	fmt.Println("")
	fmt.Println("After passing twelve arguements two function")
    fmt.Println(joinstr("d", "a", "t", "a", "p","l","a","t","f","o","r","m"))
	fmt.Println("")
	fmt.Println("------------------Pointers---------------------")

		//Creating a pointer
		b := 255
		var a *int = &b
		fmt.Printf("Type of a is %T\n", a)
		fmt.Println("address of b is", a)
	
		//Zero value of a pointer is nil
		d := 25
		var p *int
		if p == nil {
			fmt.Println("\np is", p)
			p = &d
			fmt.Println("p after initialization is", p)
		}
	
		//Dereferencing a pointer
		i := 255
		j := &i
		fmt.Println("\nAddress of i is", j)
		fmt.Println("value of i is", *j)

		//Changing the value pointed using dereference
		*j++
		fmt.Println("new value of i is", i)
	
		//Passing pointer to a function
		k := 58
		fmt.Println("\nvalue of k before function call is", k)
		l := &k
		change(l)
		fmt.Println("value of k after function call is", k)
		fmt.Println("")

		fmt.Println("------------------Error Handling---------------------")

		// Trying to open the text file
		f, err := os.Open("/test.txt")
        if err != nil {
					 fmt.Println(err)
					 fmt.Println("")
                     return
        }
		fmt.Println(f.Name(), "opened successfully")

}
