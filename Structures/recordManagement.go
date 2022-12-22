package Structures

import (
	"net/http"
	"time"
)


type Record struct{
	Id uint64					 `json:",omitempty"`
	StudentId uint64		     `json:"student_id,omitempty"`
	ScoreId uint64				 `json:"score_id,omitempty"`
	UserAdded uint64   			 `json:"user_added,omitempty"`
	Comment string				 `json:"comment"`

	// This part will be used for GET request
	StdName string				`json:"student_name,omitempty"`
	StdSurname string			`json:"student_surname,omitempty"`
	Action string				`json:"action,omitempty"`
	Points string				`json:"points,omitempty"`
	DateScoreAdded time.Time	`json:"date_added,omitempty"`
}


func(R *Record) Validate(r *http.Request) error{
	var err error
	if R.ScoreId == 0 || R.StudentId == 0{
		err = ClassError{
			Err:"Enter student name and activity type",
		}

	}

	return err
}


type MyRecords struct{

	// This part will be used for GET request
	Id uint64					`json:"id,omitempty"`
	StdId uint64				`json:"student_id,omitempty"`
	StdName string				`json:"student_name,omitempty"`
	StdSurname string			`json:"student_surname,omitempty"`
	StdClass string				`json:"student_class,omitempty"`
	Action string				`json:"action,omitempty"`
	Points int64				`json:"points,omitempty"`
	DateScoreAdded time.Time	`json:"date_added,omitempty"`
	Formatted string          `json:"formatted,omitempty"`
}

type RecordsForStudent struct{
	Id uint64           	`json:"id,omitempty"`
	Action string			`json:"action"`
	Points int64			`json:"points"`
	FromUser uint64			`json:"from_user"`
	Date time.Time			`json:"date"`
	DateFormatted string    `json:"date_formatted,omitempty"`
	TimeFormatted string 	`json:"time_formatted,omitempty"`
}