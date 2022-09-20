package Structures

import (
	"net/http"
	"strings"
)




type ClassError struct {
	Err string
}


func (C ClassError) Error() string{
	return C.Err
}



type Class struct{
	Id uint64         `json:"id,omitempty"`
	Class string      `json:"class"`
	UserAdded uint64  `json:"user_added,omitempty"`
	NewName string     `json:"newName,omitempty"`
}

// Validate function checks if the DB already has
// this class name stored. If the returned value is 0,
// then proceed to save the new class. Otherwise, result is
// 1 , hence return error message from API
func (c *Class) Validate(r *http.Request) error{
	var err error

	if r.Method == "POST" {
		upperClass := strings.ToUpper(c.Class)
		c.Class = upperClass

		if c.Class == "" {
			err = ClassError{
				Err: "Class can't be empty",
			}
			return err
		}

		if len(c.Class) > 10 {
			err = ClassError{
				Err: "Class must be not greater than 10 characters",
			}

		}

	}

	if r.Method == "PUT"{
		upperNewName := strings.ToUpper(c.NewName)
		c.NewName = upperNewName

		if c.NewName == ""{
			err = ClassError{
				Err: "Class can't be empty",
			}
			return err
		}

		if len(c.NewName) > 10 {
			err = ClassError{
				Err: "Class must be not greater than 10 characters",
			}

		}
	}
	return err

}