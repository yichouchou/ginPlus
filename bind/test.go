package bind

type ReqTest struct {
	Accesstoken string `json:"Accesstoken"`
	UserName    string `json:"UserName"`
	Password    string `json:"Password"`
	Age         string `json:"Age"`
}
