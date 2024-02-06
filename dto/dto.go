package dto

import "PaSer/model"

type UserDto struct {
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Identity   string `json:"identity"`
	Grade      string `json:"grade"`
	Gander     string `json:"gander"`
	Zone       string `json:"zone"`
	School     string `json:"school"`
	ParentName string `json:"parentname"`
	ParentRole string `json:"parentrole"`
}

func UserInfo(user model.User) UserDto {
	return UserDto{
		Name:       user.Name,
		Phone:      user.Phone,
		Email:      user.Email,
		Identity:   user.Identity,
		Grade:      user.Grade,
		Gander:     user.Gander,
		Zone:       user.Zone,
		School:     user.School,
		ParentName: user.ParentName,
		ParentRole: user.ParentRole,
	}
}
