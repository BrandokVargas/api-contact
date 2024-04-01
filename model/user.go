package model

//Contacts Users
type Contact struct {
	nameContact  string
	phoneContact string
}

//User model
type User struct {
	name     string
	gmail    string
	contacts []Contact
}
