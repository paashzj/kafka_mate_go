package kafka

import (
	"fmt"
	"github.com/paashzj/gutil"
	"io/ioutil"
	"kafka_mate_go/pkg/config"
	"kafka_mate_go/pkg/path"
	"kafka_mate_go/pkg/util"
	"os"
	"strconv"
	"strings"
)

func Config() error {
	configProp, err := initFromFile(path.KfkOriginalConfig)
	if err != nil {
		return err
	}
	if !config.ClusterEnable {
		if config.KafkaAdvertiseAddress != "" {
			configProp.Set("advertised.listeners", config.KafkaAdvertiseAddress)
		}
	} else {
		configProp.Set("broker.id", getBrokerId())
		configProp.Set("listeners", "PLAINTEXT://"+os.Getenv("HOSTNAME")+".kafka:9092")
		configProp.Set("zookeeper.connect", config.ZkAddress)
	}
	configProp.SetInt64("socket.send.buffer.bytes", config.KafkaSocketSendBufferBytes)
	configProp.SetInt64("socket.receive.buffer.bytes", config.KafkaSocketReceiveBufferBytes)
	if config.KafkaMessageMaxBytes != -1 {
		configProp.SetInt64("message.max.bytes", config.KafkaMessageMaxBytes)
	}
	if config.KafkaFetchMessageMaxBytes != -1 {
		configProp.SetInt64("fetch.message.max.bytes", config.KafkaFetchMessageMaxBytes)
	}
	if config.ReplicaFetchMaxBytes != -1 {
		configProp.SetInt64("replica.fetch.max.bytes", config.ReplicaFetchMaxBytes)
	}
	return configProp.Write(path.KfkConfig)
}

func initFromFile(file string) (*gutil.ConfigProperties, error) {
	configProp := gutil.ConfigProperties{}
	configProp.Init()
	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	split := strings.Split(string(fileBytes), "\n")
	for _, line := range split {
		if strings.HasPrefix(line, "#") {
			continue
		}
		array := strings.Split(line, "=")
		if len(array) != 2 {
			util.Logger().Error(fmt.Sprintf("line error %s", line))
			continue
		}
		configProp.Set(array[0], array[1])
	}
	return &configProp, nil
}

func getBrokerId() string {
	hostname := os.Getenv("HOSTNAME")
	index := strings.LastIndex(hostname, "-")
	zkIndex := hostname[index+1:]
	index, _ = strconv.Atoi(zkIndex)
	return strconv.Itoa(index + 1)
}
