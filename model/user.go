package model

import "time"

//Contacts Users
// type Contact struct {
// 	nameContact  string
// 	phoneContact string
// }

//User model
type User struct {
	ID        uint
	Name      string `json:"name"`
	Gmail     string `json:"gmail"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// contacts []Contact
}
