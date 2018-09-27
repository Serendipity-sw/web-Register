package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/guotie/config"
	"github.com/guotie/deferinit"
	"github.com/smtc/glog"
	"github.com/swgloomy/gutil"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var (
	rt                  *gin.Engine
	pidStrPath          = "./web-Register.pid"
	configFn            = flag.String("config", "./config.json", "config file path")
	debugFlag           = flag.Bool("d", false, "debug mode")
	serverListeningPort int    //服务监听端口
	logsDir             string //日志存放目录
	forwardingDomain    string //转发域
	rootPrefix          string //服务运行所需要的二级目录名称

)

/**
初始化web工具
创建人:邵炜
创建时间:2017年2月9日13:45:26
输入参数:是否为调试模式(true为调试模式 false为正式运行模式)
*/
func ginInit(debug bool) {
	//设置gin的工作方式
	gin.SetMode(gutil.If(debug, gin.DebugMode, gin.ReleaseMode).(string))
	rt = gin.Default()
	setGinRouter(rt)
	go rt.Run(fmt.Sprintf(":%d", serverListeningPort))
}

/**
服务运行
创建人:邵炜
创建时间:2017年2月8日18:01:18
输入参数:配置文件路径 是否为调试模式(d表示为调试模式,否则为正式运行模式)
*/
func serverRun(cfn string, debug bool) {
	config.ReadCfg(cfn)
	configRead()
	gutil.LogInit(debug, logsDir)

	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("set many cpu successfully!")

	deferinit.InitAll()
	fmt.Println("init all module successfully!")

	deferinit.RunRoutines()
	fmt.Println("init all run successfully!")

	ginInit(debug)
	fmt.Println("ginInit run successfully!")
}

/**
服务停止
创建人:邵炜
创建时间:2017年2月9日14:06:27
*/
func serverExit() {
	deferinit.StopRoutines()
	fmt.Println("stop routine successfully!")

	deferinit.FiniAll()
	fmt.Println("stop all modules successfully!")

	glog.Close()
}

/**
服务构造函数(程序启动主入口)
创建人:邵炜
创建时间:2017年2月9日14:08:21
*/
func main() {
	if gutil.CheckPid(pidStrPath) {
		return
	}
	flag.Parse()
	serverRun(*configFn, *debugFlag)
	c := make(chan os.Signal, 1)
	gutil.WritePid(pidStrPath)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	//信号等待
	<-c
	gutil.RmPidFile(pidStrPath)
	serverExit()
	os.Exit(0)
}
