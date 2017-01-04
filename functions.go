package  main

import "fmt"

func passByValue(a,b int)  {
	a = 100;
	b = 200;
}

func passByReference(ptrA, ptrB *int)  {
	*ptrA = 100;
	*ptrB = 200;
}

func multipleReturn() (int, int)  {
	return 1000, 2000
}

func variadicFunctions(nums ...int)  {
	for _,i := range nums {
		print(i)
	}
}

func closure(a int) func () int {

	return func () int {
		a++
		return a
	}
}

func recurssion(n int) int  {
	if n == 0 { // end the recurssion here
		return 1
	}
	return n * recurssion(n-1) // recurse here
}

func main() {
	a := 10;
	b := 20;
	passByValue(a,b)
	fmt.Println(a,b)
	passByReference(&a,&b) // pointers
	fmt.Println(a,b)
	a,b = multipleReturn()
	fmt.Println(a,b)
	variadicFunctions(3,5,7,9)
	fmt.Println(closure(10)) // this returns a function (pointer to the function)
	fmt.Println(closure(10)()) // returns the expected value
	fmt.Println(recurssion(5)) // returns 5!
}
