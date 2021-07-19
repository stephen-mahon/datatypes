package main

import (
	"fmt"

	"github.com/stephen-mahon/datatypes/organization"
)

func main() {
	p := organization.NewPerson("Steve", "Mahon", organization.NewEuroUnionIdentifier("my real id", "Ireland"))
	err := p.SetTwitterHandler("@steve")
	fmt.Printf("%T\n", organization.TwitterHandler("Test"))
	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s\n", err)
	}
}
