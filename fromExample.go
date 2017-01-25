package main
import "fmt"
func looper(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
func main() {
	looper("direct")
	go looper("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}