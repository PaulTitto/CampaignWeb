package user

import "time"

type User struct {
	Id               int
	Name             string
	Occupation       string
	Email            string
	PasswordHash     string
	Avatar_file_name string
	Role             string
	Create_at        time.Time
	Update_at        time.Time
}
