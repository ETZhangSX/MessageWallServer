package main

import (
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/TarsCloud/TarsGo/tars/util/rogger"

	"LifeService"
)

var comm *tars.Communicator
//SLOG 日志对象
var SLOG = rogger.GetLogger("ServerLog")

func main() { //Init servant
	comm = tars.NewCommunicator()
	imp := new(MessageWallImp)                                    //New Imp
	imp.init()
	app := new(LifeService.MessageWall)                                 //New init the A Tars
	cfg := tars.GetServerConfig()                               //Get Config File Object
	app.AddServant(imp, cfg.App+"."+cfg.Server+".MessageWallObj") //Register Servant
	tars.Run()
}
