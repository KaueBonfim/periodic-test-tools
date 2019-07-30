package main

import (
	"fmt"
	"periodic_service_tools/app"
	"os"
)

func main(){

	runner_app := app.App{}
	runner_app.Init(os.Getenv("TYPE_BANK"))
	runner_app.InitRouters()
	fmt.Print("Iniciando server!")
	runner_app.Run(":5080")
}
