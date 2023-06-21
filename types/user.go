package types

type UserRegisterReq struct {
	UserName string `form:"user_name" json:"user_name" from:"user_name"`
	Password string `form:"password" json:"password" from:"password"`
}

type UserLoginReq struct {
	UserName string `form:"user_name" json:"user_name" from:"user_name"`
	Password string `form:"password" json:"password" from:"password"`
}

type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

type UserResp struct {
	ID       uint   `json:"id" form:"id" example:"1"`                    // 用户ID
	UserName string `json:"user_name" form:"user_name" example:"FanOne"` // 用户名
	CreateAt int64  `json:"create_at" form:"create_at"`                  // 创建
}
