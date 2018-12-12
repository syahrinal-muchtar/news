package Models

import (
	"News/Config"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(b *User) (err error) {

	if err := Config.DB.Raw("select email,password from users WHERE email = ?", b.Email).Scan(b).Error; err != nil {
		return err
	}
	return nil
}

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

func RegisterUser(b *User) (err error) {
	pwd := getPwd(b.Password)
	b.Password = hashAndSalt(pwd)
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func ShowUsers(b *[]User) (err error) {
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func ShowUser(b *User, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(b *User, id string) (err error) {
	Config.DB.Save(b)
	return nil
}

func DeleteUser(b *User, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(b)
	return nil
}
