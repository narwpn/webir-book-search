FROM eclipse-temurin:21-jdk-alpine AS builder

WORKDIR /app

COPY . .

# Used for health check
RUN apk add --no-cache curl

RUN chmod +x entrypoint.sh

EXPOSE 3000 5000

ENTRYPOINT ["./entrypoint.sh"]