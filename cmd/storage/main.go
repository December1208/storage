package main

import "storage/config"

func main() {
	eng := initApp(config.Instance.LocalStorage)
	println(config.Instance.LocalStorage.Domain)
	err := eng.Run()
	if err != nil {
		return
	}
}
