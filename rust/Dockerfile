FROM rust:alpineAS builder
WORKDIR /app
COPY . .
RUN cargo build --release

FROM alpine
COPY --from=builder /app/target/release/app /usr/local/bin/app
CMD ["app"
