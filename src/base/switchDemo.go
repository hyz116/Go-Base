package main
import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("write ", i, " as ")
	// basic switch
	switch i {
		case 1 :
			fmt.Println("one")
		case 2 :
			fmt.Println("two")
		case 3 :
			fmt.Println("three")				
	}
	fmt.Println("-----------")

	// You can use commons to separate multiple expressions in the same case statement
	// Use the optional default case in this example as well
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday :
		fmt.Println("It`s the weekend")
	default:
		fmt.Println("It`s a weekday")	
	}

	fmt.Println("----------------")
	/*
		switch without an expression is an alternate way to express if/else logic.
	*/
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It`s before noon")
	default: 
		fmt.Println("It`s after noon")
	}

	fmt.Println("------------------")

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I`m a bool")
		case int:
			fmt.Println("I`m a int")
		default:
			fmt.Println("Don`t know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("Go")
}