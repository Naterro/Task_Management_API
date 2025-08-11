package entity

import "time"

type User struct {
	Id       int
	Login    string
	Email    string
	PassHash string
	Created  time.Time
}

func (u *User) CheckPass(passHash string) bool {
	if u.PassHash == passHash {
		return true
	}
	return false
}
