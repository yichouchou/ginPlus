package bind

type ReqTest struct {
	Access_token string `json:"access_token"`
	UserName     string `json:"user_name" binding:"required"` // With verification mode
	Password     string `json:"password"`
	Age          string `json:"age"`
}
