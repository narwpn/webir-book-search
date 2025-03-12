FROM eclipse-temurin:21-jdk-alpine AS builder

WORKDIR /app

COPY . .

RUN ./gradlew build -x test


FROM eclipse-temurin:21-jre-alpine

WORKDIR /app

# Used for health check
RUN apk add --no-cache curl

COPY --from=builder /app/build/libs/*.jar app.jar
COPY --from=builder /app/gradlew /app/gradlew
COPY --from=builder /app/gradle /app/gradle
COPY --from=builder /app/index /app/index

EXPOSE 3000 5000

ENTRYPOINT ["java", "-jar", "app.jar", "--spring.profiles.active=prod"]
