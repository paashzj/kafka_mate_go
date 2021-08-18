#!/bin/bash

uuid=`bash $KAFKA_HOME/bin/kafka-storage.sh random-uuid`
bash $KAFKA_HOME/bin/kafka-storage.sh format -t $uuid $KAFKA_HOME/config/kraft/server.properties

nohup $KAFKA_HOME/bin/kafka-server-start.sh $KAFKA_HOME/config/kraft/server.properties >$KAFKA_HOME/kafka.stdout.log 2>$KAFKA_HOME/kafka.std.stderr.log