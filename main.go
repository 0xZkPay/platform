package main

import (
	"fmt"

	"github.com/0xZkPay/platform/app"
	"github.com/0xZkPay/platform/config/envconfig"
	"github.com/0xZkPay/platform/util/pkg/logwrapper"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app.Init()
	logwrapper.Log.Info("Starting ZkPay Platform")
	addr := fmt.Sprintf(":%d", envconfig.EnvVars.APP_PORT)
	err := app.GinApp.Run(addr)
	if err != nil {
		logwrapper.Log.Fatal(err)
	}
}