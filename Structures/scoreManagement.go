package Structures

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"net/http"
	"strings"
)

type Score struct{
	Id uint64 		  `json:"id,omitempty"`
	Action string 	  `json:"action"`
	Points int64	  `json:"points"`
	UserAdded uint64  `json:"user_added"`

	NewAction string  `json:"new_action,omitempty"`
	NewPoints int64   `json:"new_points,omitempty"`
}

func (s *Score)Validate(r *http.Request) error{
	var err error
	if r.Method == "POST"{
		lower := strings.ToLower(s.Action)

		// Make slice of strings
		splitStr := strings.Split(lower," ")

		/// Capitalize only first word ////
		a := []cases.Caser{ cases.Title(language.Und)}[0]
		lower= a.String(splitStr[0])

		// Select right side of the whole expression
		right := strings.Join(splitStr[1:]," ")

		// Join in
		s.Action = lower + " "+ right

		if s.Action == "" || s.Action==" "{
			err = ClassError{
				Err: "Score type must not be empty",
			}

		}

		if len(s.Action) > 200 {
			err = ClassError{
				Err: "Score type is too long",
			}

		}

		if s.Points == 0{
			err = ClassError{
				Err: "Score can't have 0 points",
			}
		}

	}

	if r.Method == "PUT"{
		lower := strings.ToLower(s.NewAction)

		// Make slice of strings
		splitStr := strings.Split(lower," ")


		/// Capitalize only first word ////
		a := []cases.Caser{ cases.Title(language.Und)}[0]
		lower= a.String(splitStr[0])

		// Select right side of the whole expression
		right := strings.Join(splitStr[1:]," ")


		// Join in
		s.NewAction = lower + " "+ right


		if s.NewAction == " "|| s.NewAction==""{
			err = ClassError{
				Err: "New score type can't be empty",
			}
		}

		if s.NewPoints == 0{
			err = ClassError{
				Err: "Score can't have 0 points",
			}
		}

		if len(s.NewAction) > 200 {
			err = ClassError{
				Err: "Score type is too long",
			}

		}
	}
	return err
}