package response

type SignUp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SignIn struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshToken struct {
	AccessToken string `json:"access_token"`
}

type GetProfile struct {
	Username           string `json:"username"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Email              string `json:"email"`
	ValidatedEmail     bool   `json:"validated_email"`
	Cellphone          string `json:"cellphone"`
	ValidatedCellphone bool   `json:"validated_cellphone"`
	CreatedAt          int64  `json:"created_at"`
}
