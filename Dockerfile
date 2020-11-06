FROM golang:latest as builder
LABEL maintainer="Jose Ramon Mañes - github.com/jrmanesdk"
ADD . /app
WORKDIR /app
#RUN go test -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./interfaces/cmd/server
######## Start a new stage from scratch #######
FROM alpine:latest
# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/.env .env
COPY --from=builder /app/main .
# Expose port 8080 to the outside world
EXPOSE 8080
# Command to run the executable
CMD ["./main"]