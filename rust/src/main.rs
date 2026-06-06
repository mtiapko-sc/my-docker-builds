use std::io::{Read, Write};
use std::net::TcpListener;

fn main() {
    let listener = TcpListener::bind("0.0.0.0:8080").unwrap();
    println!("Server running on port 8080");
    for stream in listener.incoming() {
        let mut stream = stream.unwrap();
        let mut buf = [0u8; 1024];
        stream.read(&mut buf).unwrap();
        let req = String::from_utf8_lossy(&buf);
        let body = if req.contains("GET /healthz") { "ok" } else { "Hello from Rust!" };
        let resp = format!(
            "HTTP/1.1 200 OK\r\nContent-Length: {}\r\nConnection: close\r\n\r\n{}",
            body.len(), body
        );
        stream.write_all(resp.as_bytes()).unwrap();
    }
}
