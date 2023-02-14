package main

import (
	"runtime/debug"
	"storage/config"
	"storage/pkg"
)

func main() {
	eng := initApp(config.Instance.LocalStorage, config.Instance.Server)
	pkg.Logger.Info("try to start listen on " + config.Instance.Server.Addr)
	defer func() {
		err := recover()
		if err != nil {
			pkg.Logger.Info(string(debug.Stack()))
		}
	}()
	err := eng.Run(config.Instance.Server.Addr)
	if err != nil {
		return
	}
}
