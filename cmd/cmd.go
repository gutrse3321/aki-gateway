package main

import "flag"

/**
 * @Author: Tomonori
 * @Date: 2020/6/18 17:15
 * @Title:
 * --- --- ---
 * @Desc:
 */

var configFile = flag.String("f", "resource/gateway.yml", "set config file which viper will loading")

func main() {
	flag.Parse()

	app, err := CreateApp(*configFile)
	if err != nil {
		panic(err)
	}

	if err := app.Start(); err != nil {
		panic(err)
	}

	app.AwaitSignal()
}
