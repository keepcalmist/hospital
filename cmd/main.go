package main

import (
	"github.com/keepcalmist/hospital/pkg/config"
	app "github.com/keepcalmist/hospital/pkg/server"
)

func main() {
	cfg := config.InitConfig()
	aplication := app.InitApp(cfg)
	aplication.ListenAndServe(cfg)
}
