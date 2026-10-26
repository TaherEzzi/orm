package val_test

import (
	"fmt"

	"github.com/hanzoai/orm/val"
)

type User struct {
	Email    string
	Name     string
	Password string
}

// ExampleValidator demonstrates struct field validation using the val package.
func ExampleValidator() {
	user := User{
		Email:    "bad",
		Name:     "",
		Password: "hi",
	}

	v := val.New()
	v.Check("Email").Exists().IsEmail()
	v.Check("Name").Exists().MinLength(2)
	v.Check("Password").MinLength(6)

	errs := v.Exec(&user)
	for _, err := range errs {
		if fe, ok := err.(*val.FieldError); ok {
			fmt.Printf("%s: %s\n", fe.Field, fe.Error())
		}
	}
	// Output:
	// Email: Field must be an email.
	// Name: Field cannot be blank.
	// Password: Field must be atleast 6 characters long.
}
