package jiekou

// import (
// 	"github.com/toophy/pangu/help"
// 	lua "github.com/toophy/gopher-lua"
// )

// // 场景线程接口
// type IScreenThread interface {
// 	Get_thread_id() int32                         // 获取线程ID
// 	Get_thread_name() string                      // 获取线程名称
// 	GetLuaState() *lua.LState                     // !!!只能获取, 不准许保存指针, 获取LState
// 	PostEvent(a help.IEvent) bool                // 投递定时器事件
// 	PostThreadMsg(tid int32, a help.IEvent) bool // 投递线程间消息
// 	GetEvent(name string) help.IEvent            // 通过别名获取事件
// 	RemoveEvent(e help.IEvent) bool              // 删除事件, 只能操作本线程事件
// 	RemoveEventList(header help.IEvent)          // 删除一整个事件列表
// 	PopTimer(e help.IEvent)                      // 从线程事件中弹出指定事件, 只能操作本线程事件
// 	PopObj(e help.IEvent)                        // 从关联对象中弹出指定事件, 只能操作本线程事件
// 	LogDebug(f string, v ...interface{})          // 线程日志 : 调试[D]级别日志
// 	LogInfo(f string, v ...interface{})           // 线程日志 : 信息[I]级别日志
// 	LogWarn(f string, v ...interface{})           // 线程日志 : 警告[W]级别日志
// 	LogError(f string, v ...interface{})          // 线程日志 : 错误[E]级别日志
// 	LogFatal(f string, v ...interface{})          // 线程日志 : 致命[F]级别日志
// }
