package data

func UserByEmail(email string)(u User,err error){
	err = Db.Model(User{Email: email}).First(&u).Error
	return
}
func Encrypt(password string) string{
	return password
}

