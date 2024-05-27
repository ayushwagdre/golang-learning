package password_learning

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func PasswordStorageInit() {
	password := "mysecretpassword"

	// Hashing the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hashed Password:", string(hashedPassword))

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		fmt.Println("Password is incorrect")
	} else {
		fmt.Println("Password is correct")
	}

}
