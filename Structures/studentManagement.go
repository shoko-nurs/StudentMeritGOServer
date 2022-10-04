package Structures

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"net/http"
	"strings"
)

type Student struct {
	Id uint64 			`json:"id,omitempty"`
	Name string			`json:"name,omitempty"`
	Surname string		`json:"surname,omitempty"`
	ClassId uint64		`json:"class_id,omitempty"`
	ClassName string	`json:"class_name,omitempty"`
	CurrentScore int64	`json:"current_score"`
	UserAdded uint64    `json:"user_added"`

	//NewName string		`json:"new_name,omitempty"`
	//NewSurname string   `json:"new_surname,omitempty"`
	//NewClass uint64     `json:"new_class,omitempty"`

}


func (s *Student) Validate(r *http.Request) error{
	var err error

	if r.Method=="POST" || r.Method=="PUT" {
		// Lower name and surname, remove all spaces
		name := strings.ToLower(s.Name)
		name = strings.ReplaceAll(name, " ", "")

		surname := strings.ToLower(s.Surname)
		surname = strings.ReplaceAll(surname, " ", "")

		//Capitalize Name
		a := []cases.Caser{cases.Title(language.Und)}[0]
		name = a.String(name)
		s.Name = name

		//Capitalize Surname
		a = []cases.Caser{cases.Title(language.Und)}[0]
		surname = a.String(surname)
		s.Surname = surname

		if s.Name == "" || s.Surname == "" {
			err = ClassError{
				Err: "Name and Surname can't be empty",
			}

		}

		if s.ClassId == 0 {

			err = ClassError{
				Err: "Class can't be empty",
			}
		}


	}

	return err
}
