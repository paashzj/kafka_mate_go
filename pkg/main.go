package main

import (
	"go.uber.org/zap"
	"kafka_mate_go/pkg/config"
	"kafka_mate_go/pkg/kafka"
	"kafka_mate_go/pkg/path"
	"kafka_mate_go/pkg/util"
	"os"
	"os/signal"
)

func main() {
	util.Logger().Debug("this is a debug msg")
	util.Logger().Info("this is a info msg")
	util.Logger().Error("this is a error msg")
	err := kafka.Config()
	if err != nil {
		util.Logger().Error("generate config failed ", zap.Error(err))
	}
	if config.RaftEnable {
		err := util.CallScript(path.KfkStartRaftScript)
		if err != nil {
			util.Logger().Error("start kafka server failed ", zap.Error(err))
			os.Exit(1)
		}
	} else {
		if !config.ClusterEnable {
			err := util.CallScript(path.KfkStartStandaloneScript)
			if err != nil {
				util.Logger().Error("start kafka server failed ", zap.Error(err))
				os.Exit(1)
			}
		} else {
			err := util.CallScript(path.KfkStartScript)
			if err != nil {
				util.Logger().Error("start kafka server failed ", zap.Error(err))
				os.Exit(1)
			}
		}
	}
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case <-interrupt:
			return
		}
	}
}
