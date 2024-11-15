package main

// import (
// 	"errors"
// 	"fmt"
// 	"math"
// )

// //ERROR handling

// func Sqrt(value float64) (float64, error) {
// 	if value < 0 {
// 		return 0, errors.New("Math, passed negative ")
// 	}

// 	return math.Sqrt(value), nil
// }

// func main() {
// 	result, err := Sqrt(-1)

// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(result)
// 	}

// 	result, err = Sqrt(81)

// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(result)
// 	}
// }

// Interfaces

// /* define an interface */
// type interface_name interface {
// 	method_name1 [return_type]
// 	method_name2 [return_type]
// 	method_name3 [return_type]
// 	...
// 	method_namen [return_type]
//  }

//  /* define a struct */
//  type struct_name struct {
// 	/* variables */
//  }

//  /* implement interface methods*/
//  func (struct_name_variable struct_name) method_name1() [return_type] {
// 	/* method implementation */
//  }
//  ...
//  func (struct_name_variable struct_name) method_namen() [return_type] {
// 	/* method implementation */
//  }

// /* Interfaces example */
// type Shape interface {
// 	area() float64
// }

// type Circle struct {
// 	x, y, radius float64
// }

// type Rectangle struct {
// 	width, legnth float64
// }

// func (circle Circle) area() float64 {
// 	return math.Pi * circle.radius * circle.radius
// }

// func (rectangle Rectangle) area() float64 {
// 	return rectangle.legnth * rectangle.width
// }

// func getArea(shape Shape) float64 {
// 	return shape.area()
// }

// func main() {
// 	circle := Circle{x: 0, y: 0, radius: 4}
// 	rectangle := Rectangle{width: 40, legnth: 2}

// 	fmt.Printf("The are of Rectangle is %f\n", getArea(rectangle))
// 	fmt.Printf("The are of Circle is %f\n", getArea(circle))

// }

// Type conversion

// func main() {
// 	var sum = 17
// 	var m = 5
// 	var men float32

// 	men = float32(sum) / float32(m)
// 	fmt.Printf("Value of m : %f\n ", men)
// }

//Fibonacci recursion

// func fibonacci(i int) (ret int) {
// 	if i == 0 {
// 		return 0
// 	}

// 	if i == 1 {
// 		return 1
// 	}
// 	return fibonacci(i-1) + fibonacci(i-2)
// }

// func main() {
// 	var i int
// 	for i = 0; i < 20; i++ {
// 		fmt.Printf("%d", fibonacci(i))
// 	}
// }

//Factorial recursion

// func factorial(i int) int {
// 	if i <= 1 {
// 		return 1
// 	}
// 	return i * factorial(i-1)
// }
// func main() {
// 	var i int = 5

// 	fmt.Printf("Factorial of %d is %d", i, factorial(i))

// }

// Maps

// func main() {
// 	var cntrCptlMp map[string]string

// 	// create a map
// 	cntrCptlMp = make(map[string]string)

// 	// insert key value pairs in the map
// 	cntrCptlMp["Ru"] = "Mscw"
// 	cntrCptlMp["Italy"] = "Rome"
// 	cntrCptlMp["Tjk"] = "Dshnb"

// 	fmt.Printf(cntrCptlMp["Ru"])

// 	// print map using keys

// 	for cntr := range cntrCptlMp {
// 		fmt.Println("\nCapital of", cntr, " is ", cntrCptlMp[cntr])
// 	}

// 	//test if entry is present in the map or not

// 	capital, ok := cntrCptlMp["US"]

// 	// test, ok condition

// 	if ok {
// 		fmt.Print("The capital of US is ", capital)
// 	} else {
// 		fmt.Print("The capital isnt in the list")
// 	}

// 	//deleting from map
// 	// create one more map
// 	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New Delhi"}
// 	fmt.Println("Original Map")

// 	/* print map */
// 	for country := range countryCapitalMap {
// 		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
// 	}

// 	delete(countryCapitalMap, "Italy")

// 	fmt.Printf("Italy is gone")

// 	fmt.Println("Map is updated ")

// 	//map over the countries again

// 	for country := range countryCapitalMap {
// 		fmt.Println("Capital of ", country, "is ", countryCapitalMap[country])
// 	}

// }

// // Range

// func main() {
// 	/* create a slice */
// 	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

// 	/* print the numbers */
// 	for i := range numbers {
// 		fmt.Println("SLice ", i, "is", numbers[i])
// 	}

// 	/* create a map*/
// 	countryCapitalMap := map[string]string{"France": "Paris", "Tajikistan": "Dushanbe"}

// 	/* print map using keys*/
// 	for i := range countryCapitalMap {
// 		fmt.Println("Capital of", i, "is", countryCapitalMap[i])
// 	}

// 	/* print map using key-values*/
// 	for country, capital := range countryCapitalMap {
// 		fmt.Println("Capital of", country, "is", capital)
// 	}

// }

// Slices

// func main() {
/* create a slice */
// var numbers []int
// printSlice(numbers)

// /* print the original slice */
// fmt.Println("numbers ==", numbers)

// /* print the sub slice starting from index 1(included) to index 4(excluded)*/
// fmt.Println("numbers[1:4] ==", numbers[1:4])

// /* missing lower bound implies 0*/
// fmt.Println("numbers[:3] ==", numbers[:3])

// /* missing upper bound implies len(s)*/
// fmt.Println("numbers[4:] ==", numbers[4:])

// numbers1 := make([]int, 0, 5)
// printSlice(numbers1)

// /* print the sub slice starting from index 0(included) to index 2(excluded) */
// number2 := numbers[:2]
// printSlice(number2)

// /* print the sub slice starting from index 2(included) to index 5(excluded) */
// number3 := numbers[2:5]
// fmt.Println("Use this", number3)
// printSlice(number3)

// appending

// /* append allows nil slice */
// numbers = append(numbers, 0)
// printSlice(numbers)

// /* add one element to slice*/
// numbers = append(numbers, 1)
// printSlice(numbers)

// /* add more than one element at a time*/
// numbers = append(numbers, 2, 3, 4)
// printSlice(numbers)

// /* create a slice numbers1 with double the capacity of earlier slice*/
// numbers1 := make([]int, len(numbers), (cap(numbers))*2)

// /* copy content of numbers to numbers1 */
// copy(numbers1, numbers)
// printSlice(numbers1)
// }

// func main() {
// 	var numbers = make([]int, 3, 5)
// 	printSlice(numbers)

//		if numbers == nil {
//			fmt.Printf("slice is nil")
//		} else {
//			fmt.Printf("This is the value of slice %d", numbers)
//		}
//	}
// func printSlice(x []int) {
// 	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
// }

// Structures

// type Products struct {
// 	title       string
// 	description string
// 	order_id    int
// }

// func main() {
// 	var Product1 Products /*Declare Product1 to Product */
// 	var Product2 Products /*Declare Product2 to Product */

// 	//Product1 spec
// 	Product1.title = "Ftank"
// 	Product1.description = "Sharp and wide"
// 	Product1.order_id = 2974624

// 	//Product2 spec
// 	Product2.title = "Fpump"
// 	Product2.description = "Fast and excessive"
// 	Product2.order_id = 2974623

// 	printProduct(&Product1)

// 	printProduct(&Product2)
// }

// func printProduct(product *Products) {
// 	fmt.Printf("Product title : %s\n ", product.title)
// 	fmt.Printf("Product desc : %s\n ", product.description)
// 	fmt.Printf("Product orderId : %d\n ", product.order_id)
// }

// Pointers

// func main() {
// 	var a int = 200
// 	var b int = 300
// 	var ret int
// 	var ip *int

// 	ret = max(a, b)

// 	ip = &ret

// 	fmt.Printf("Address of a variable: %x\n", &ret)

// 	/* address stored in pointer variable */
// 	fmt.Printf("Address stored in ip variable: %x\n", ip)

// 	/* access the value using the pointer */
// 	fmt.Printf("Value of *ip variable: %d\n", *ip)
// 	// ip = &ret

// 	// fmt.Printf("The address of the max value is %x\n", &ret)
// 	// fmt.Printf("The value of the max value itself is %d", ip)
// }

/* function returning the max between two numbers */
// func max(num1, num2 int) int {
// 	/* local variable declaration */
// 	var result int

// 	if num1 > num2 {
// 		result = num1
// 	} else {
// 		result = num2
// 	}
// 	return result
// }

// Array in use

// func main() {
// 	var n [10]int /* n is an array of 10 integers */
// 	var i, j int

// 	/* initialize elements of array n to 0 */
// 	for i = 0; i < 10; i++ {
// 		n[i] = i + 100
// 	}

// 	/* output each array element's value */
// 	for j = 0; j < 10; j++ {
// 		fmt.Printf("Element[%d] = %d\n", j, n[j])
// 	}
// }

// // Concatination of strings //

// func main() {
// 	greetings := []string{"Hello", "      world!"}
// 	fmt.Println(strings.Join(greetings, " "))
// }

// // Creating strings and string lengths //

// func main() {
// 	var greeting = "Samsung"

// 	fmt.Printf("normal string: ")
// 	fmt.Printf("%s", greeting)
// 	fmt.Printf("\n")
// 	fmt.Printf("hex bytes: ")

// 	for i := 0; i < len(greeting); i++ {
// 		fmt.Printf("%x ", greeting[i])
// 	}

// 	fmt.Printf("\n")
// 	const sampleText = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

// 	/*q flag escapes unprintable characters, with + flag it escapses non-ascii
// 	characters as well to make output unambigous */
// 	fmt.Printf("quoted string: ")
// 	fmt.Printf("%q", sampleText)
// 	fmt.Printf("\n")

// 	fmt.Printf("\n")
// 	fmt.Printf("%d", len(greeting))
// }

/* global variable declaration */
// var a int = 20

// func main() {
// 	/* local variable declaration in main function */
// 	var a int = 10
// 	var b int = 20
// 	var c int = 0

// 	fmt.Printf("value of a in main() = %d\n", a)
// 	c = sum(a, b)
// 	fmt.Printf("value of c in main() = %d\n", c)
// }

// /* function to add two integers */
// func sum(a, b int) int {
// 	fmt.Printf("value of a in sum() = %d\n", a)
// 	fmt.Printf("value of b in sum() = %d\n", b)

// 	return a + b
// }
