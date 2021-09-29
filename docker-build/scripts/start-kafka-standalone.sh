#!/bin/bash

nohup $KAFKA_HOME/bin/zookeeper-server-start.sh $KAFKA_HOME/config/zookeeper.properties >>$KAFKA_HOME/zookeeper.stdout.log 2>>$KAFKA_HOME/zookeeper.stderr.log &
nohup $KAFKA_HOME/bin/kafka-server-start.sh $KAFKA_HOME/config/server.properties >>$KAFKA_HOME/kafka.stdout.log 2>>$KAFKA_HOME/kafka.stderr.log &