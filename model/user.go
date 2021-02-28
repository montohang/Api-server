package model

import (
	"regexp"
	"strings"
)


type Metadata struct {
	CurrentPage int `json:"currentPage"`
	FirstPage   int `json:"fisrtPage"`
	LastPage    int `json:"lastPage"`
	TotalData   int `json:"total"`
}

type UserList struct {
	Users    []*User  `json:"users"`
	Metadata Metadata `json:"metadata"`
}

type Job struct {
	JobID    string `json:"jobID"`
	JobLabel string `json:"jobLabel"`
}

type Education struct {
	EducationID    string `json:"educationID"`
	EducationLabel string `json:"educationLabel"`
}

type User struct {
	UserID      string            `json:"userID"`
	IDCard      string            `json:"idCard"`
	Username    string            `json:"username"`
	DateOfBirth string            `json:"dataOfBirth"`
	Job         Job               `json:"job"`
	Education   Education         `json:"education"`
	UserStatus  int               `json:"userStatus"`
	CreatedDate string            `json:"createdDate"`
	UpdatedDate string            `json:"updatedDate"`
	Errors      map[string]string `json:"errors"`
}


func (user *User) Validate() bool {
	user.Errors = make(map[string]string)
	reg := regexp.MustCompile("((19|20)\\d\\d)-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])")
	if len(strings.TrimSpace(user.IDCard)) > 16 || strings.TrimSpace(user.IDCard) == "" {
		user.Errors["IDCard"] = "enter your id card"
	}
	if strings.TrimSpace(user.Username) == ""{
		user.Errors["Username"] = "enter your username"
	}
	if strings.TrimSpace(user.DateOfBirth) == "" || !reg.MatchString(user.DateOfBirth) {
		user.Errors["DateOfBirth"] = "enter your date of birth"
	}
	if strings.TrimSpace(user.Job.JobID) == "" {
		user.Errors["JobID"] = "enter your job id"
	}
	if strings.TrimSpace(user.Education.EducationID) == "" {
		user.Errors["EducationID"] = "enter your education id"
	}
	return len(user.Errors) == 0
}