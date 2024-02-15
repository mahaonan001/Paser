package dto

import "PaSer/model"

type AdminDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func AdminInfo(Admin model.Admin) AdminDto {
	return AdminDto{
		Name:  Admin.Name,
		Email: Admin.Email,
	}
}

type UserDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func UserInfo(User model.User) UserDto {
	return UserDto{
		Name:  User.Name,
		Email: User.Email,
	}
}
