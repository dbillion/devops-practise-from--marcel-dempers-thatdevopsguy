package main // Required for a standalone executable.

// import "fmt" // fmt implements formatted I/O.

/* When this program is executed the first function that runs is main.main() */
import (
    "fmt"
    "reflect"
)


/* Declare a single variable */
var a int

/* Declare a group of variables */
var (
    b bool
    c float32
    d string
)


/* Define a stack type using a struct */
type stack struct {
	index int
	data  [5]int
}

/* Define push and pop methods */
func (s *stack) push(k int) {
	s.data[s.index] = k
	s.index++
}

/* Notice the stack pointer s passed as an argument */
func (s *stack) pop() int {
	s.index--
	return s.data[s.index]
}


func main() {
/* Create a pointer to the new stack and push 2 values */
s := new(stack)
s.push(23)
s.push(14)
fmt.Printf("stack: %v\n", *s) // stack: {2 [23 14 0 0 0]}


	fmt.Println("Hello, world") // Call Println() from the fmt package.

	// a = 42                  // Assign single value
	// b, c = true, 32.0       // Assign multiple values
	// d = "string"            // Strings must contain double quotes
	// fmt.Println(a, b, c, d) // 42 true 32 string

	// a := 42            // Initialize and assign to a single variable
	// b, c := true, 32.0 // Initialize and assign to multiple variables
	// d := "string"
	// fmt.Println(a, b, c, d) // 42 true 32 string


	        // /* User specified types */
			// const a int32 = 12         // 32-bit integer
			// const b float32 = 20.5      // 32-bit float
			// var c complex128 = 1 + 4i  // 128-bit complex number
			// var d uint16 = 14          // 16-bit unsigned integer
	
			// /* Default types */
			// n := 42              // int
			// pi := 3.14           // float64
			// x, y := true, false  // bool
			// z := "Go is awesome" // string
	
			// fmt.Printf("user-specified types:\n %T %T %T %T\n", a, b, c, d)
			// fmt.Printf("default types:\n %T %T %T %T %T\n", n, pi, x, y, z)

			/* Define an array of size 4 that stores deployment options */
	// var DeploymentOptions = [4]string{"R-pi", "AWS", "GCP", "Azure"}

	// /* Loop through the deployment options array */
	// for i := 0; i < len(DeploymentOptions); i++ {
	// 	option := DeploymentOptions[i]
	// 	fmt.Println(i, option)
	// }


	// /* Loop through the deployment options array */
	// for index, option := range DeploymentOptions {
	// 	fmt.Println(index, option)
	// }


		/* Define an array containing programming languages */
	 	languages := [9]string{
			"C", "Lisp", "C++", "Java", "Python",
			"JavaScript", "Ruby", "Go", "Rust", // Must include the trailing comma
		}

	// 	/* Define slices */
	// classics := languages[0:3]  // alternatively languages[:3]
	// modern := make([]string, 4) // len(modern) = 4
	// modern = languages[3:7]     // include 3 exclude 7
	// new := languages[7:9]       // alternatively languages[7:]

	// fmt.Printf("classic languagues: %v\n", classics) // classic languagues: [C Lisp C++]
	// fmt.Printf("modern languages: %v\n", modern)     // modern languages: [Java Python JavaScript Ruby]
	// fmt.Printf("new languages: %v\n", new)           // new languages: [Go Rust]
	

	  // -- snip -- //
	  allLangs := languages[:]                      // copy of the array
	  fmt.Println(reflect.TypeOf(allLangs).Kind())   // slice

	  /* Create a slice containing web frameworks */
	  frameworks := []string{
		  "React", "Vue", "Angular", "Svelte",
		  "Laravel", "Django", "Flask", "Fiber",
	  }

	  jsFrameworks := frameworks[0:4:4]          // length 4 capacity 4
	  frameworks = append(frameworks, "Meteor")  // not possible with arrays

	  fmt.Printf("all frameworks: %v\n", frameworks)
	  fmt.Printf("js frameworks: %v\n", jsFrameworks)


	  /* Define a map containing the release year of several languages */
	firstReleases := map[string]int{
		"C": 1972, "C++": 1985, "Java": 1996,
		"Python": 1991, "JavaScript": 1996, "Go": 2012,
	}

	/* Loop through each entry and output the name and release year */
	for k, v := range firstReleases {
		fmt.Printf("%s was first released in %d\n", k, v)
	}

	x := []float64{2.15, 3.14, 42.0, 29.5}
	fmt.Println(average(x))   // 19.197499999999998


	y := []float64{2.15, 3.14, 42.0, 29.5}
	fmt.Println(average2(y)) // 19.197499999999998



	count := 1
	for count < 5 {
		count += count
	}
	fmt.Println(count) // 8


	var address *int  // declare an int pointer
	number := 42      // int
	address = &number // address stores the memory address of number
	value := *address // dereferencing the value 

	fmt.Printf("address: %v\n", address) // address: 0xc0000ae008
	fmt.Printf("value: %v\n", value)     // value: 42

// concurren🔮 :https://fireship.io/lessons/learn-go-in-100-lines/

c := make(chan int) // Create a channel to pass ints
	for i := 0; i < 5; i++ {
		go cookingGopher(i, c) // Start a goroutine
	}

	for i := 0; i < 5; i++ {
		gopherID := <-c // Receive a value from a channel
		fmt.Println("gopher", gopherID, "finished the dish")
	} // All goroutines are finished at this point
}

/* Notice the channel as an argument */
func cookingGopher(id int, c chan int) {
	fmt.Println("gopher", id, "started cooking")
	c <- id // Send a value back to main
	
}

/* Define a function to find the average of the floats contained in a slice */
func average(x []float64) (avg float64) {
	total := 0.0
	if len(x) == 0 {
		avg = 0
	} else {
		for _, v := range x {
			total += v
		}
		avg = total / float64(len(x))
	}
	return
}


func average2(x []float64) (avg float64) {
	total := 0.0
	switch len(x) {
	case 0:
		avg = 0
	default:
		for _, v := range x {
			total += v
		}
		avg = total / float64(len(x))
	}
	return
}
