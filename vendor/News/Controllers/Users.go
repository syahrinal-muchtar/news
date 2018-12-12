package Controllers

import (
	"News/ApiHelpers"
	"News/Models"
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func getPwd(password string) []byte {
	return []byte(password)
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func LoginUser(c *gin.Context) {
	var user Models.User

	c.Bind(&user)

	pwd := getPwd(user.Password)
	err := Models.LoginUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "User Not Found",
		})
	}
	pwdMatch := comparePasswords(user.Password, pwd)

	if pwdMatch != true {
		c.JSON(401, gin.H{
			"status":  401,
			"message": "wrong password",
		})
		return
	}
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
	fmt.Println("TEST")
}

func ShowUsers(c *gin.Context) {
	var user []Models.User
	err := Models.ShowUsers(&user)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, user)
	} else {
		ApiHelpers.RespondJSON(c, 200, user)
	}
}

func ShowUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.User
	err := Models.ShowUser(&user, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, user)
	} else {
		ApiHelpers.RespondJSON(c, 200, user)
	}
}

func RegisterUser(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)
	err := Models.RegisterUser(&user)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, user)
	} else {
		ApiHelpers.RespondJSON(c, 200, user)
	}
}

func UpdateUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.ShowUser(&user, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, user)
	}
	c.BindJSON(&user)
	err = Models.UpdateUser(&user, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, user)
	} else {
		ApiHelpers.RespondJSON(c, 200, user)
	}
}

func DeleteUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.DeleteUser(&user, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, user)
	} else {
		ApiHelpers.RespondJSON(c, 200, user)
	}
}
