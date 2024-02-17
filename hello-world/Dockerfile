FROM golang AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o hello-world -ldflags '-extldflags "-static"' .

FROM scratch
COPY --from=builder /app /
CMD ["/hello-world"]
