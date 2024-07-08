package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `Password@123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	// loginPassword := `Password@123`
	loginPassword := `Password@1234`
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPassword))
	if err != nil {
		fmt.Println("YOU CAN'T LOGIN", "The password is ", loginPassword, "You want this", s, "password to login")
		return
	}
	fmt.Println("You're logged in")

}
