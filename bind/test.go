package bind

type ReqTest struct {
	Accesstoken string `json:"Accesstoken"`
	UserName    string `json:"UserName"`
	Password    string `json:"Password"`
	Age         string `json:"Age"`
}

type ReqGet struct {
	Accesstoken string `form:"accesstoken"`
	UserName    string `form:"userName"`
	Password    string `form:"password"`
	Age         string `form:"age"`
}
