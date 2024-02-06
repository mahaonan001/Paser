package dto

import "PaSer/model"

type AdminDto struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func AdminInfo(Admin model.Admin) AdminDto {
	return AdminDto{
		Name:  Admin.Name,
		Phone: Admin.Phone,
		Email: Admin.Email,
	}
}
