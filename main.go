package main

import (
	"fmt"

	"github.com/stephen-mahon/datatypes/organization"
)

func main() {
	p := organization.NewPerson("Steve", "Mahon")
	err := p.SetTwitterHandler("@steve")
	fmt.Printf("%T\n", organization.TwitterHandler("Test"))
	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s\n", err)
	}
	fmt.Println(p.TwitterHandler())
	fmt.Println(p.TwitterHandler().RedirectURL())
	fmt.Println(p.ID())
	fmt.Println(p.FullName())
}
