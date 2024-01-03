package table

import (
	"Demo/global"
	"Demo/utils"
	"gorm.io/gorm"
)

type BlankUser struct {
	gorm.Model
	Name      string `json:"name"`
	StudentId string `json:"studentId" gorm:"unique"`
	Grade     string `json:"grade"`
	PassWord  string `json:"passWord"`
}

func (BlankUser) TableName() string {
	return "blank_user"
}

// FindUser method to find all information of one user by studentId
func FindUser(studentId string) (currentUser BlankUser, err error) {
	result := global.GLO_DB.Take(&currentUser).Where("student_id", studentId)
	return currentUser, result.Error
}

// FindUserList method to find all the user in databse
func FindUserList() (users []BlankUser, err error) {
	result := global.GLO_DB.Find(&users)
	return users, result.Error
}

// CreateUser method to create one user account
func CreateUser(currentUser BlankUser) (err error) {
	currentUser.PassWord = utils.EncryptPassWord(currentUser.PassWord)
	return global.GLO_DB.Create(&currentUser).Error
}

// UpdateUser method to update information of one user
func UpdateUser(newUser BlankUser) (err error) {
	currentUser, err := FindUser(newUser.StudentId)
	if err != nil {
		return err
	}

	updateData := make(map[string]interface{})
	if newUser.Name != "" {
		updateData["Name"] = newUser.Name
	}
	if newUser.PassWord != "" {
		updateData["PassWord"] = utils.EncryptPassWord(newUser.PassWord)
	}
	if newUser.StudentId != "" {
		updateData["StudentId"] = newUser.StudentId
	}
	if newUser.Grade != "" {
		updateData["Grade"] = newUser.Grade
	}

	if len(updateData) == 0 {
		return nil
	}
	result := global.GLO_DB.Model(&BlankUser{}).Where("id = ?", currentUser.ID).Updates(updateData)
	return result.Error
}

// DeleteUser method to delete admin account
func DeleteUser(studentId string) (err error) {
	currentUser, err := FindUser(studentId)
	if err != nil {
		return err
	}

	// real delete
	return global.GLO_DB.Unscoped().Delete(&currentUser).Error
}
