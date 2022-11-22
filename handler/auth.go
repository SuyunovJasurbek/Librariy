package handler

type LoginServise interface {

	LoginUser(email string, password string) bool

}

type LoginInformation struct{
	email string
	passwod string
}

func StaticLoginServise() LoginServise{
	return &LoginInformation{
		email:   "suyunovjas7053@gmail.com",
		passwod: "assalom",
	}
}

func (info *LoginInformation) LoginUser( email string, password string )bool {
 return info.email==email && info.passwod==password
}