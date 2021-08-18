#!/bin/bash

nohup $KAFKA_HOME/bin/zookeeper-server-start.sh $KAFKA_HOME/config/zookeeper.properties >$KAFKA_HOME/zookeeper-normal.log &
nohup $KAFKA_HOME/bin/kafka-server-start.sh $KAFKA_HOME/config/server.properties >$KAFKA_HOME/kafka-normal.log &