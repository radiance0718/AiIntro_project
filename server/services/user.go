package services

//
//import (
//	"Demo/table"
//	"Demo/utils"
//	"github.com/gin-gonic/gin"
//)
//
//// Login
//// @Summary Admin Login
//// @Description Verify a user account
//// @Tags User
//// @Accept application/x-www-form-urlencoded
//// @Produce json
//// @Param name formData string true "User's ID"
//// @Param password formData string true "User's password"
//// @Success 200 {object} LoginResponse  "Login Success"
//// @Failure 200 {object} Response "Error"
//// @Router /puser/login [post]
//func Login(c *gin.Context) {
//	currentUserId := c.PostForm("studentId")
//	currentUser, err := table.FindUser(currentUserId)
//	if err != nil {
//		c.JSON(200, gin.H{
//			"status":  "Error",
//			"message": err,
//		})
//		return
//	}
//
//	userPassWord := c.PostForm("password")
//	if currentUser.PassWord != utils.EncryptPassWord(userPassWord) {
//		c.JSON(200, gin.H{
//			"status":  "Error",
//			"message": "User password not correct",
//		})
//		return
//	}
//
//	c.JSON(200, gin.H{
//		"status":  "Success",
//		"message": "Login Success",
//	})
//	return
//}
//
//// UserCreateUser
//// @Summary Create a user account
//// @Description Create a user account
//// @Tags User
//// @Accept application/x-www-form-urlencoded
//// @Produce json
//// @Param name formData string true "User's name"
//// @Param password formData string true "User's password"
//// @Param studentId formData string true "User's student ID"
//// @Param grade formData string true "User's grade"
//// @Success 200 {object} LoginResponse  "Create Success"
//// @Failure 200 {object} Response "Error"
//// @Router /user/create [post]
//func UserCreateUser(c *gin.Context) {
//	newUser := table.BlankUser{
//		Name:      c.PostForm("name"),
//		PassWord:  c.PostForm("password"),
//		StudentId: c.PostForm("studentid"),
//		Grade:     c.PostForm("grade"),
//	}
//
//	err := table.CreateUser(newUser)
//	if err != nil {
//		c.JSON(200, gin.H{
//			"status":  "Error",
//			"message": err,
//		})
//		return
//	}
//
//	c.JSON(200, gin.H{
//		"status":  "Success",
//		"message": "Create Success",
//	})
//	return
//}
//
//// UserUpdateUser
//// @Summary Update a user account
//// @Description Update a user account
//// @Tags User
//// @Accept application/x-www-form-urlencoded
//// @Produce json
//// @Param name formData string true "User's name"
//// @Param password formData string "User's password"
//// @Param studentId formData string "User's student ID"
//// @Param grade formData string "User's grade"
//// @Param token formData string true "User's current token"
//// @Success 200 {object} LoginResponse  "Update Success"
//// @Failure 200 {object} Response "Error"
//// @Router /user/update [post]
//func UserUpdateUser(c *gin.Context) {
//	StudentId := c.PostForm("studentId")
//	if StudentId == "" {
//		c.JSON(200, gin.H{
//			"status":  "Error",
//			"message": "Can not found user ID",
//		})
//		return
//	}
//
//	newUser := table.BlankUser{
//		Name:      c.PostForm("name"),
//		PassWord:  c.PostForm("password"),
//		StudentId: StudentId,
//		Grade:     c.PostForm("grade"),
//	}
//
//	err := table.UpdateUser(newUser)
//	if err != nil {
//		c.JSON(200, gin.H{
//			"status":  "Error",
//			"message": err,
//		})
//		return
//	}
//
//	c.JSON(200, gin.H{
//		"status":  "Success",
//		"message": "Modify Success",
//	})
//	return
//}
//
//// UserDeleteUser
//// @Summary Delete a user account
//// @Description Delete a user account
//// @Tags User
//// @Accept application/x-www-form-urlencoded
//// @Produce json
//// @Param studentId formData string true "User's student ID"
//// @Param token formData string true "User's current token"
//// @Success 200 {object} LoginResponse  "Update Success"
//// @Failure 200 {object} Response "Error"
//// @Router /user/delete [post]
//func UserDeleteUser(c *gin.Context) {
//	StudentId := c.PostForm("studentId")
//	if StudentId == "" {
//		c.JSON(200, gin.H{
//			"status":  "Error",
//			"message": "Can not found user ID",
//		})
//		return
//	}
//
//	err := table.DeleteUser(StudentId)
//	if err != nil {
//		c.JSON(200, gin.H{
//			"status":  "Error",
//			"message": err,
//		})
//		return
//	}
//
//	c.JSON(200, gin.H{
//		"status":  "Success",
//		"message": "Delete Success",
//	})
//	return
//}
//
//// UserFindUser
//// @Summary Delete a user account
//// @Description Delete a user account
//// @Tags User
//// @Accept application/x-www-form-urlencoded
//// @Produce json
//// @Param studentId formData string true "User's student ID"
//// @Param token formData string true "User's current token"
//// @Success 200 {object} LoginResponse  "Update Success"
//// @Failure 200 {object} Response "Error"
//// @Router /user/find [get]
//func UserFindUser(c *gin.Context) {
//	StudentId := c.Query("studentId")
//	if StudentId == "" {
//		c.JSON(200, gin.H{
//			"status":  "Error",
//			"message": "Can not found user ID",
//		})
//		return
//	}
//
//	foundUser, err := table.FindUser(StudentId)
//	if err != nil {
//		c.JSON(200, gin.H{
//			"status":  "Error",
//			"message": err,
//		})
//		return
//	}
//
//	c.JSON(200, gin.H{
//		"status":        "Success",
//		"message":       "Success in finding User",
//		"name":          foundUser.Name,
//		"password hash": foundUser.PassWord,
//		"student id":    foundUser.StudentId,
//		"grade":         foundUser.Grade,
//	})
//	return
//}
