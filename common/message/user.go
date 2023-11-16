package message

// User 定义一个用户的结构体
type User struct {
	//确定字段信息
	//为了序列化和反序列化成功
	//用户信息的json字符串与结构体字段对应的Tag名字一致
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
	UserStatus string  `json:"userStatus"` //用户状态
	Sex string `json:"sex"` //性别
}
