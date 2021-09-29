package main

import (
	"github.com/paashzj/gutil"
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
		stdout, stderr, err := gutil.CallScript(path.KfkStartRaftScript)
		util.Logger().Info("shell result ", zap.String("stdout", stdout), zap.String("stderr", stderr))
		if err != nil {
			util.Logger().Error("start kafka server failed ", zap.Error(err))
			os.Exit(1)
		}
	} else {
		if !config.ClusterEnable {
			stdout, stderr, err := gutil.CallScript(path.KfkStartStandaloneScript)
			util.Logger().Info("shell result ", zap.String("stdout", stdout), zap.String("stderr", stderr))
			if err != nil {
				util.Logger().Error("start kafka server failed ", zap.Error(err))
				os.Exit(1)
			}
		} else {
			stdout, stderr, err := gutil.CallScript(path.KfkStartScript)
			util.Logger().Info("shell result ", zap.String("stdout", stdout), zap.String("stderr", stderr))
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
