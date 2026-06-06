FROM eclipse-temurin:21-jdk
WORKDIR /app
RUN apt-get install -y curl
COPY srtc ./src
RUN javac src/App.java -d .
EXPOSE 8080
CMD ["java", "App"]
