FROM eclipse-temurin:21-jdk-alpine AS builder

WORKDIR /app

COPY . .

# Used for health check
RUN apk add --no-cache curl

VOLUME [ "/app/index" ]

# Index
# ENTRYPOINT ["./gradlew", "index", "--args=--spring.profiles.active=prod"]

# Search
EXPOSE 3000
ENTRYPOINT ["./gradlew", "bootRun", "--args=--spring.profiles.active=prod"]