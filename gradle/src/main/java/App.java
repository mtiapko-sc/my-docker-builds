public class App {
    public static void main(String[] args) throws Exception {
        var server = com.sun.net.httpserver.HttpServer.create(
            new java.net.InetSocketAddress(8080), 0);
        server.createContext("/healthz", e -> {
            byte[] body = "ok".getBytes();
            e.sendResponseHeaders(200, body.length);
            e.getResponseBody().write(body);
            e.getResponseBody().close();
        });
        server.createContext("/", e -> {
            byte[] body = "Hello from Gradle/Java!".getBytes();
            e.sendResponseHeaders(200, body.length);
            e.getResponseBody().write(body);
            e.getResponseBody().close();
        });
        server.start();
        System.out.println("Server running on port 8080");
    }
}
