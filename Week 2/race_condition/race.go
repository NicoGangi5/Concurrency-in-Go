package main

import "fmt"

func main()  {
	for {
		var x int

		go func(num *int) {
			*num = 1
			*num = *num * 2
		}(&x)

		go func(num int) {
			fmt.Println("Num:", num)
		}(x)

		/*
		Race condition is when multiple threads are trying to access and manipulate the same variable.
		The code above accesses to print and modify the variable.
		The correct working would be to initialize the variable to 1, then multiply it by 2 and finally print it on the screen.
		But due to the uncertainty of the schedule it is not possible to know in which order the instructions will be executed.
		It is possible that the variable is initialized to 1, then print and finally multiplied by 2.
		 */
	}
}