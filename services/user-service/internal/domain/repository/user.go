package repository

type UserRepository interface {
	Register()
	Login()
	ChangePrivileges()
	GetPrivileges()
}
