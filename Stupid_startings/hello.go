package main

import (
	"fmt"
)

//"math"

const s string = "constant"

func main() {
	//--------------------------------------------------------------
	//---------------------Values-----------------------------------
	//--------------------------------------------------------------
	// fmt.Println("go" + "lang")
	// fmt.Println("1+1 =", 1+1)
	// fmt.Println("7.0/3.0 =", 7.0/3.0)

	// fmt.Println(true && false)
	// fmt.Println(true || false)
	// fmt.Println(!true)
	//--------------------------------------------------------------
	//----------------------Variables-------------------------------
	//--------------------------------------------------------------
	// var a = "init"
	// fmt.Println(a)

	// var b, c int = 2, 3
	// fmt.Println(b, c)

	// var d = true
	// fmt.Println(d == (5*4 == 20))

	// var e int
	// fmt.Println(e)

	// var f int = 344
	// fmt.Println(f)

	// hello := 2324
	// fmt.Println(hello)
	//--------------------------------------------------------------
	//----------------------Constants-------------------------------
	//--------------------------------------------------------------
	// fmt.Println(s)

	// const n = 500000000

	// const d = 3e20 / n
	// fmt.Println(d)

	// fmt.Println(int64(d))

	// fmt.Println(math.Cos(n))
	//--------------------------------------------------------------
	//-----------------------For------------------------------------
	//--------------------------------------------------------------
	// i := 1
	// for i <= 2 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// for k := 3; k <= 9; k++ {
	// 	fmt.Println(k)
	// }

	// for {
	// 	fmt.Println("loop")
	// 	break
	// }

	// for n := 0; n <= 5; n++ {
	// 	if n%2 == 1 {
	// 		continue
	// 	}
	// 	fmt.Println(n)
	// }
	//--------------------------------------------------------------
	//-----------------------If/Else--------------------------------
	//--------------------------------------------------------------
	// if 7%2 == 0 {
	// 	fmt.Println("7 is even")
	// } else {
	// 	fmt.Println("7 is odd")
	// }

	// if 8%4 == 0 {
	// 	fmt.Println("8 is divisible by 4")
	// }

	// if num := 9; num < 0 {
	// 	fmt.Println(num, "is negative")
	// } else if num < 10 {
	// 	fmt.Println(num, "has 1 digit")
	// } else {
	// 	fmt.Println(num, "has multiple digits")
	// }
	//--------------------------------------------------------------
	//-------------------------Switch-------------------------------
	//--------------------------------------------------------------
	// i := 2
	// fmt.Print("Write ", i, " as ")
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("twpoooo")
	// }

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("Weekend")
	// default:
	// 	fmt.Println("Weekend, NOOOOT")
	// }

	// t := time.Now()
	// switch {
	// case t.Hour() < 12:
	// 	fmt.Println("its before noon")
	// default:
	// 	fmt.Println("Its afternoooooon")
	// }

	// whatAmI := func(i interface{}) {
	// 	switch t := i.(type) {
	// 	// case bool:
	// 	// 	fmt.Println("Im a bool")
	// 	case int:
	// 		fmt.Println("Im an int")
	// 	default:
	// 		fmt.Printf("Dont know type %T\n", t)
	// 	}
	// }
	// whatAmI(true)
	// whatAmI(1)
	// whatAmI("hey")

	//--------------------------------------------------------------
	//-------------------------Arrays-------------------------------
	//--------------------------------------------------------------
	// var a [5]int
	// fmt.Println("emp:", a)

	// a[4] = 100
	// fmt.Println("set:", a)
	// fmt.Println("get:", a[4])
	// fmt.Println("len:", len(a))

	// b := [5]int{1, 2, 3, 4, 5}
	// fmt.Println("dcl:", b)

	// var twoD [2][3]int
	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		twoD[i][j] = i + j
	// 	}
	// }
	// fmt.Println("2d: ", twoD

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("set:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "d", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl2:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl2:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
