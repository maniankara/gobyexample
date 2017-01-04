package main

import "fmt"

func array()  {
	s := make([]string, 4)
	fmt.Println("empty list: ", s)
	// init
	s[0] = "abc"
	s[1] = "123"
	// s[2] = 'c' // this is not allowed
	s[2] = "c"
	s[3] = "xyz"
	fmt.Println("filled list: ", s)
	// add more
	s = append(s, "test")
	fmt.Println("Size of s: ", len(s))
	// slice copy
	c := s
	c1 := []string{}
	copy(c1,s)
	fmt.Println("Content of c: ", c)
	fmt.Println("Content of c1: ", c1)
	// 2D
	twoD := make([][]int, 5) // only row is mentioned here
	for i := 0; i<5; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen) // make every time a row, crappy
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i+j
		}
	}
	fmt.Println("2D array: ", twoD)
}

func maps()  {
	m := make(map[string]int)
	fmt.Println(m)
	// simple assginment
	m["abc"] = 10
	m["xyz"] = 20
	fmt.Println(m)
	fmt.Println(len(m))
	// deletion
	delete(m, "abc")
	fmt.Println(m)
	// create
	n := map[string]int {"qwe" : 20, "poi" : 50} // json values
	fmt.Println(n)

}

// Range returns the list/array/slice/map in a loop
func ranges()  {
	n := map[string]string { "key1":"value1", "key2":"value2"}
	for key, val := range n {
		fmt.Println(key,val)
	}

}

func main() {

	/* Array examples*/
	// array()

	/* Maps examples*/
	// maps()
	
	/* Range examples */
	ranges()
}
