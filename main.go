package main

import (
	"oss-helper/grpc_service"
	"oss-helper/internal/config"
)

func main() {
	cfg := config.LoadCfg()

	//db := dal.InitDb(cfg)
	//defer db.Close()

	grpc_service.Run(cfg)

	println("Service Run On ---→ " + cfg.Host)

}
