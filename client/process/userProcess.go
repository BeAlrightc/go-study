package process
import (
	"fmt"
	"net"
	"encoding/json"
	"encoding/binary"
	"go_code/chatroom/common/message"
	"go_code/chatroom/client/utils"
	"os"
)
type UserProcess struct {
	//暂时不需要字段
}

func (this *UserProcess) Register(userId int,userPwd string,userName string)(err error){
	//1.连接到服务器端
	conn, err :=net.Dial("tcp","localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=",err)
		return
	}

	//延时关闭
	defer conn.Close()

	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType
	//3.创建一个RegisterMes 结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd 
	registerMes.User.UserName = userName

	//4.将registerMes序列化
	data, err :=json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Mashal err=",err)
		return
	}

	//5.将data赋给了mes.Data字段
	mes.Data = string(data)

	//6.将mes进行序列化
	data, err =json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Mashal err=",err)
		return
	}
	
	//7.到这个时候，data就是我们要发送的消息
    //7.1先把data的长度发送给服务器
	//先获取data的长度->转成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkgLen) //将该、长度转成了byte类型是数据
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err !=nil {
		fmt.Println("connWrite(buf) fail ",err)
		return
	}

	fmt.Printf("客户端发送数据的消息长度=%d 内容是=%s",len(data),string(data))
	//发送消息本身
	_, err = conn.Write(data)
	if err !=nil {
		fmt.Println("connWrite(data) fail ",err)
		return
	}

	//创建一个Transfer实例
	tf := &utils.Transfer{
		Conn : conn,
	}

	//发送data给服务器端
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误 err=",err)
	}

	mes, err = tf.ReadPkg() //mes 就是RegisterResMes

	if err != nil {
		fmt.Println("readPkg(conn) err=",err)
		return
	}

	//将mes的Data部分反序列化为RegisterResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data),&registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，你重新登录一把")
		os.Exit(0)
	}else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return

}




//给关联一个用户登录的方法
//写一个函数，完成登录操作
func (this *UserProcess) Login(userId int,userPwd string) (err error) {

	//下一个就要开始定协议
	// fmt.Printf("userId = %d userPwd = %s\n",userId,userPwd)
	// return nil

	//1.连接到服务器端
	conn, err :=net.Dial("tcp","localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=",err)
		return
	}

	//延时关闭
	defer conn.Close()

	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType 

	//3.创建一个LoginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd 

	//4.将loginMes序列化
	data, err :=json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Mashal err=",err)
		return
	}
	//5.将data赋给了mes.Data字段
	mes.Data = string(data)
	//6.将mes进行序列化
	data, err =json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Mashal err=",err)
		return
	}
	//7.到这个时候，data就是我们要发送的消息
    //7.1先把data的长度发送给服务器
	//先获取data的长度->转成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkgLen) //将该、长度转成了byte类型是数据
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err !=nil {
		fmt.Println("connWrite(buf) fail ",err)
		return
	}

	fmt.Printf("客户端发送数据的消息长度=%d 内容是=%s",len(data),string(data))
	//发送消息本身
	_, err = conn.Write(data)
	if err !=nil {
		fmt.Println("connWrite(data) fail ",err)
		return
	}
	//休眠20秒
	// time.Sleep(10 * time.Second)
	// fmt.Println("休眠了20秒..")
	//这里还需要处理服务器端返回的消息
	//创建一个Transfer实例
	tf := &utils.Transfer{
		Conn : conn,
	}
	mes, err = tf.ReadPkg() //mes 就是

	if err != nil {
		fmt.Println("readPkg(conn) err=",err)
		return
	}
	
	//将mes的Data部分反序列化为LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	if loginResMes.Code == 200 {
		//fmt.Println("登录成功")

		//现在可以显示当前在线的列表 遍历loginResMes.UsersId
		fmt.Println("当前在线用户列表如下")
		for _, v := range loginResMes.UsersId {
			//如果我们要求不显示自己在线，下面我们增加一个代码
			if v == userId {
				continue
			}
			
			fmt.Println("用户id:\t",v)
		}
		fmt.Println("\n\n")


		//这里我们还需要再客户端启动一个协程
		//该协程保持和服务器端的通讯，如果服务器有数据推送给客户端
		//则可以接受并显示在客户端的终端
		go serverProcessMes(conn)
		//1.显示登录成功后的菜单[循环显示]
		for {
			ShowMenu()
		}
	}else {
		fmt.Println(loginResMes.Error)
	}
	return
}