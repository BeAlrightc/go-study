package process
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
		onlineUsers : make(map[int]*UserProcess,1024)
	}
}

//完成对onlineUsers的添加
func (this *UserMgr) AddOnlinesUser() {
	
}

