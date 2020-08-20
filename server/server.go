package server

import (
	"fmt"
	"gin-admin/config"
	"gin-admin/router"
	"gin-admin/utils"
	"os"
)

func Start() {
	cfg := config.NewConfig()
	ginServer := router.InitRouter()
	//ginpprof.Wrapper(ginServer)
	utils.Logger.Info("Current ENV: ", os.Getenv("env"))
	utils.Logger.Info("Start gin-admin on Port: ", cfg.HTTP.Port)
	utils.Logger.Infof("Log path: %s, Log name: %s", cfg.Log.LogPath, cfg.Log.LogName)
	//err := ginServer.RunTLS(":"+ string(cfg.HTTP.Port), cfg.HTTP.Certificate, cfg.HTTP.CertificateKey)
	err := ginServer.Run(cfg.HTTP.Host + ":" + fmt.Sprintf("%d", cfg.HTTP.Port))

	if err != nil {
		utils.Logger.Fatalf("Gin  Start err: %s", err.Error())
	}
}
