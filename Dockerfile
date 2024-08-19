# Build stage
# FROM golang:alpine AS builder

# ENV GO111MODULE=on \
#     CGO_ENABLED=0 \
#     GOOS=linux \
#     GOARCH=amd64

# WORKDIR /app

# COPY go.mod go.sum main.go 
# COPY ./ ./
# RUN go mod download
# RUN go build -o main
# RUN find 

#RUN cp /app/main .

# Final stage
FROM scratch

WORKDIR /app

COPY main /app/main
COPY *.env /app/
# COPY --from=builder /app/main .
RUN chmod 644 /app/*.env

ENTRYPOINT ["./main"]
