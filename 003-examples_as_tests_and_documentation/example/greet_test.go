package example

import "fmt"

// This is how you say hello to Jon.
func ExampleHello() {
	greeting, err := Hello("Jon")
	if err != nil {
		panic(err)
	}
	fmt.Println(greeting)

	// Output:
	// Hello, Jon
}

// This is how you say hello to Mike using a method receiver.
func ExampleDemo_Hello() {
	g := Demo{}

	greeting, err := g.Hello("Mike")
	if err != nil {
		panic(err)
	}
	fmt.Println(greeting)

	// Output:
	// Hello, Mike
}

// This is how you say hello to Jane.
func ExampleHello_jane() {
	greeting, err := Hello("Jane")
	if err != nil {
		panic(err)
	}
	fmt.Println(greeting)

	// Output:
	// Hello, Jane
}

// This is how you use the Page() method to show who hasn't checked in yet.
func ExamplePage() {
	checkIns := map[string]bool{
		"Bob":   true,
		"Alice": false,
		"Eve":   false,
		"Janet": true,
		"Susan": true,
		"John":  false,
	}
	Page(checkIns)

	// Unordered output:
	// Paging Alice; please see the front desk to check in.
	// Paging Eve; please see the front desk to check in.
	// Paging John; please see the front desk to check in.
}
