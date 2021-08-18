#!/bin/bash

nohup $KAFKA_HOME/mate/kafka_mate >$KAFKA_HOME/kafka_mate.stdout.log 2>$KAFKA_HOME/kafka_mate.stderr.log &
