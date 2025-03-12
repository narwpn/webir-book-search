#!/bin/sh

cd /app

./gradlew index --args='--spring.profiles.active=prod'

sleep 5

exec ./gradlew bootRun --args='--spring.profiles.active=prod'