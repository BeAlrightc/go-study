package process
import (
	"fmt"
	"go_code/chatroom/common/message"
)

//客户端要维护的Map
var onlineUsers map[int]*message.User = make(map[int]*message.User,10)

//在客户端显示当前在线的用户
func outputOnlineUser() {
	//遍历一把onlineUsers
	fmt.Println("当前在线用户列表：")
	for id,_ := range onlineUsers{
		//如果不显示自己
		fmt.Println("用户id:\t\t",id)
	}
}

//编写一个方法，处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	
	//适当的优化
	user,ok :=onlineUsers[notifyUserStatusMes.UserId]
	if !ok { //原来没有
		user = &message.User{
			UserId : notifyUserStatusMes.UserId,
		}	
	}

	user.UserStatus = string(notifyUserStatusMes.Status)
	onlineUsers[notifyUserStatusMes.UserId] = user

	outputOnlineUser()
}