package main

import (
	"storage/config"
	"storage/pkg"
)

func main() {
	eng := initApp(config.Instance.LocalStorage, config.Instance.Server)
	pkg.Logger.Info("try to start listen on " + config.Instance.Server.Addr)
	err := eng.Run(config.Instance.Server.Addr)
	//pkg.Logger.Error("try to start listen on " + config.Instance.Server.Addr)
	if err != nil {
		return
	}
}
