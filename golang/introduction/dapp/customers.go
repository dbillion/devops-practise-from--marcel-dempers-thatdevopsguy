package main

type(Customer struct {
	FirstName string
	LastName string
	FullName string
  }
)
  func GetCustomers()(customers []Customer) {

	//we can declare customers like this:
	// marcel := 	Customer{ FirstName: "Marcel", LastName: "Dempers" }
  
	  customers = append(customers,
		  Customer{ FirstName: "Marcel", LastName: "Dempers" },
		  Customer{ FirstName: "Ben", LastName: "Spain" },
		  Customer{ FirstName: "Aleem", LastName: "Janmohamed" },
		  Customer{ FirstName: "Jamie", LastName: "le Notre" },
		  Customer{ FirstName: "Victor", LastName: "Savkov" },
		  Customer{ FirstName: "P", LastName: "The Admin" },
		  Customer{ FirstName: "Adrian", LastName: "Oprea" },
		  Customer{ FirstName: "Jonathan", LastName: "D" },
	  )
  
	  return customers
  
  }


