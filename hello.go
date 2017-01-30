package main

import "math"
import "fmt"
import "./fib"

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println(math.Log2(3))
	fmt.Println(math.Log10(3))
  	fn := fib.FibFunc()
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
}
