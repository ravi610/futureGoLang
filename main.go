package main

import (
	"fmt"
	"time"
)

func main() {

	// future get method use case
	future1 := newFuture(func() Result {
		res := "Hello"
		time.Sleep(time.Second)
		return Result{resultValue: res}
	})

	res := future1.get()
	fmt.Printf("Result : %v \n", res)

	//future getWithTimeout Method usecase, failure case
	future2 := newFuture(func() Result {
		x := 4
		y := 8
		res := 3*x + 4*y
		time.Sleep(2 * time.Second)
		return Result{resultValue: res}
	})

	res2 := future2.getWithTimeout(time.Second)
	fmt.Printf("Result : %v \n", res2)

	//future getWithTimeout Method usecase, success case
	future3 := newFuture(func() Result {
		x := 4
		y := 8
		res := 3*x + 4*y
		time.Sleep(2 * time.Second)
		return Result{resultValue: res}
	})

	res3 := future3.getWithTimeout(3 * time.Second)
	fmt.Printf("Result : %v \n", res3)

	//future cancel method usecase
	future4 := newFuture(func() Result {
		x := 4
		y := 8
		res := 3*x + 4*y
		time.Sleep(4 * time.Second)
		return Result{resultValue: res}
	})

	future4.cancel()
	res4 := future4.get()
	fmt.Printf("Result : %v \n", res4)
}
