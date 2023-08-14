package request

import (
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/derror"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/derror/message"
)

type SignUp struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Cellphone string `json:"cellphone"`
}

func (req SignUp) Validate() error {
	if len(req.Username) == 0 {
		return derror.NewBadRequestError(message.InvalidRequest)
	}

	//if len(req.Username) <= 3 || len(req.Username) >= 30 {
	//	return derror.NewBadRequestError(derror.InvalidRequest)
	//}

	if len(req.Password) == 0 {
		return derror.NewBadRequestError(message.InvalidRequest)
	}
	//
	//if len(req.NationalID) == 0 {
	//	return derror.NewBadRequestError(derror.InvalidRequest)
	//}
	//
	//if len(req.PassportNo) == 0 {
	//	return derror.NewBadRequestError(derror.InvalidRequest)
	//}
	//
	//// validation password
	//{
	//	var (
	//		upp, low, num, sym bool
	//		tot                uint8
	//	)
	//
	//	for _, char := range req.Password {
	//		switch {
	//		case unicode.IsUpper(char):
	//			upp = true
	//			tot++
	//		case unicode.IsLower(char):
	//			low = true
	//			tot++
	//		case unicode.IsNumber(char):
	//			num = true
	//			tot++
	//		case unicode.IsPunct(char) || unicode.IsSymbol(char):
	//			sym = true
	//			tot++
	//		default:
	//			return derror.NewBadRequestError(derror.InvalidRequest)
	//		}
	//	}
	//
	//	if !upp || !low || !num || !sym || tot < 8 {
	//		return derror.NewBadRequestError(derror.InvalidRequest)
	//	}
	//}

	return nil
}

type SignIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req SignIn) Validate() error {
	if len(req.Username) == 0 {
		return derror.NewBadRequestError(message.InvalidRequest)
	}

	if len(req.Password) == 0 {
		return derror.NewBadRequestError(message.InvalidRequest)
	}

	return nil
}

type EditProfile struct {
	UserID    int    `json:"-"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Cellphone string `json:"cellphone"`
	Birthday  int64  `json:"birthday"`
}

func (req EditProfile) Validate() error {
	// todo
	return nil
}
