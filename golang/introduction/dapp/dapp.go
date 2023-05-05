package main

import "fmt"

var deeta="loves Jesus"




/* Declare a single variable */
var a int

/* Declare a group of variables */
var (
    b bool
    c float32
    d string
)


func main() {

	customers := GetCustomers()

	for _, customer := range customers {
	  //we can access the "customer" variable in this approach
	  fmt.Println(customer)
	}



	a = 42                  // Assign single value
	b, c = true, 32.0       // Assign multiple values
	d = "string"            // Strings must contain double quotes
	fmt.Println(a, b, c, d) // 42 true 32 string

	// fmt.Println("Hello, world.")

	rocker:=getData(1)
	fmt.Println(rocker)
	// fmt.Println(deeta)

	// get the wrong customer
	rocker=getData(15)
	
	customers1 := getData3()
fmt.Println(customers1)

// for {
// 	//any code in here will run forever!
	
// 	fmt.Println("Infinite Loop 1")
// 	  time.Sleep(time.Second)
  
// 	//unless we break out the loop like this
// 	break
//   }


//loop for x number of loops
// for x := 0; x < 10; x++ {

// 	//any code in here will run 10 times! (unless we break!)
// 	fmt.Println(customers[x])
  
//   }

for _, customer := range customers {

	//we can access the "customer" variable in this approach
	// customer = customers[x]
	// customer = customers[_]
	// fmt.Println(customer)
  
	//OR
	//we can use the supplied customer from the loop
	// and silence the x variable, replace it with a _ character
	fmt.Println(customer)
  }
  



}



func getData(customerId int) (customer string) {

// example 1
	// return " oludayo " + deeta
	// // fmt.Println(deeta)

	// // var firstName = "Dee"
	// // lastName := "Deoye"
		
	// // 	fullName := firstName + " " + lastName
	// // 	return fullName
  

	// //or we can return the computation instead of adding another variable!
	// // return firstName + " " + lastName
  
	// //or we dont even need to declare variables :)
	// // return "Marcel Dempers"


	if customerId == 1 {
		return "Marcel Dempers"
	  } else if customerId == 2 {
		return "Bob Smith"
	  } else {
		return " what did you say again? 🤦 "
	  }
  
  }

func getData2() (customers [2]string) {
	//create 1 record
	customer := "Marcel Dempers"
  
	//assign our customer to the array
	customers[0] = customer
  
	//OR we can assign it like this
	customers[1] = "Bob Smith"
  
	//send it back to the caller
	return customers
	
  }

//   Slices
  func getData3() (customers []string) {
  
	//initialise our slice of type string
	customers = []string{ "Marcel Dempers", "Bob Smith", "John Smith"}
	
	//add more legendary customers dynamically 
	customers = append(customers, "Ben Spain")
	customers = append(customers, "Aleem Janmohamed")
	customers = append(customers, "Jamie le Notre")
	customers = append(customers, "Victor Savkov")
	customers = append(customers, "P The Admin")
	customers = append(customers, "Adrian Oprea")
	customers = append(customers, "Jonathan D")
  
	//send it back to the caller
	return customers
	
  }

  