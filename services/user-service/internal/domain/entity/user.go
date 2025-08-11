package entity

import "time"

type User struct {
	Id              int
	Login           string
	Email           string
	passHashHex     string
	PrivilegesLevel int
	Created         time.Time
}

func (u *User) CheckPass(passHashHex string) bool {
	if u.passHashHex == passHashHex {
		return true
	}
	return false
}
