package process2
import (
	"fmt"
	"net"
	"encoding/json"
	"go_code/chatroom/common/message"
	"go_code/chatroom/client/utils"
)
type SmsProcess struct {
	//
}
//写方法转发消息
func (this *SmsProcess) SendGroupMes(mes *message.Message) {

	//遍历服务器端的onlineUsers map[int]*UserProgress
	//将消息转发出去
	//取出mes的内容SmsMes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data),&smsMes)
	if err != nil {
		fmt.Println(" json.Unmarshal err=",err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println(" json.Marshal err=",err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		//这里还需要过滤自己，即不要发给自己
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte,conn net.Conn) {

	//创建一个Transfer实例，发送data
	tf := &utils.Transfer{
		Conn : conn, //
	}
	err :=tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败 err=",err)
	}
}