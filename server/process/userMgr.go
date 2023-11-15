package process2
import (
	"fmt"
)
//因为UserMge实例在服务其中有且只有一个
//因为在很多的地方，都会使用，因此，我们
//将其定义为全局变量
var (
	userMgr *UserMgr
)
type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

//完成对userMge的初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers : make(map[int]*UserProcess,1024),
	}
}

//完成对onlineUsers的添加
func (this *UserMgr) AddOnlinesUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

//删除
func (this *UserMgr) DeleteOnlinesUser(userId int ) {
	delete(this.onlineUsers,userId)
}

//返回当前所有在线的用户
func (this *UserMgr)GetAllUsers() map[int]*UserProcess {
	return this.onlineUsers
}

//根据id返回对应的值
func(this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess,err error){
	//如何从map中取出一个值，待检测的方式
	up, ok := this.onlineUsers[userId]
	if !ok { //说明你要查找的用户，当前不在线
		err = fmt.Errorf("用户id不存在",userId)
		return
	} 
	return
}

