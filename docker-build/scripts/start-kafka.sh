#!/bin/bash

nohup $KAFKA_HOME/bin/kafka-server-start.sh $KAFKA_HOME/config/server.properties >$KAFKA_HOME/kafka.stdout.log 2>$KAFKA_HOME/kafka.std.stderr.log &