FROM gradle:8.7-jdk21
WORKDIR /app
COPY . .
RUN ./gradlew jar --no-daemon -q
CMD ["java", "-jar", "build/libs/app.jar"]
