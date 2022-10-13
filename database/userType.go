package database

import "notcoder_exe__blog/models"

func InsertUserType(usertype *models.UserType) {
	db.Create(usertype)
}
