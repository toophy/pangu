package thread

/*import (
	"errors"
	lua "github.com/toophy/gopher-lua"
	"github.com/toophy/pangu/help"
	"net"
)

// // 消息函数类型
// type MsgFunc func(*ScsKernel, *ClientConn)

// 网络线程
type NetThread struct {
	Thread

	OpenTcpSvr bool                // 打开Tcp服务器端口
	Address    string              // 地址
	Listener   *net.TCPListener    // 本地侦听端口
	Conns      map[int]*ClientConn // 连接池
	ConnLast   int                 // 最后连接Id
	// MsgProc    map[int]MsgFunc     // 消息处理函数注册表
}

// 新建网络线程
func New_net_thread(id int32, name string, heart_time int64, lay1_time uint64) (*NetThread, error) {
	a := new(NetThread)
	err := a.Init_net_thread(id, name, heart_time, lay1_time)
	if err == nil {
		return a, nil
	}
	return nil, err
}

// 初始化网络线程
func (this *NetThread) Init_net_thread(id int32, name string, heart_time int64, lay1_time uint64) error {
	if id < Tid_net_1 || id > Tid_net_3 {
		return errors.New("线程ID超出范围 [Tid_net_1,Tid_net_3]")
	}
	err := this.Init_thread(this, id, name, heart_time, lay1_time)
	if err == nil {

		// 打开本地TCP侦听
		serverAddr, err := net.ResolveTCPAddr("tcp", this.Address)

		if err != nil {
			return errors.New("Net Listen : '%s' %s", this.Address, err.Error())
		}

		listener, err := net.ListenTCP("tcp", serverAddr)
		if err != nil {
			return errors.New("TcpSerer ListenTCP: %s", err.Error())
		}

		this.LogInfo("ScsKernel Listening to: " + listener.Addr().String())

		this.OpenTcpSvr = true
		this.Listener = listener

		return nil
	}
	return err
}

// 响应线程首次运行
func (this *NetThread) on_first_run() {

	// 	this.set_close()

	// 	this.Conns = make(map[int]*ClientConn, 20)
	// 	this.RegMsgProc()

	if t.OpenTcpSvr {
		for {
			conn, err := t.Listener.AcceptTCP()

			if err != nil {
				GetLog().Error("Accept: %s", err.Error())
				GetLog().Flush()
				panic("ERROR: " + "Accept: " + " " + err.Error()) // terminate program
			}

			new_client := new(ClientConn)
			new_client.InitClient(t, t.ConnLast, conn)
			t.AddConn(new_client)
			go t.ConnProc(new_client)
		}
	}

}

// 响应线程最先运行
func (this *NetThread) on_pre_run() {
}

// 响应线程运行
func (this *NetThread) on_run() {
}

// 响应线程退出
func (this *NetThread) on_end() {
}
*/
