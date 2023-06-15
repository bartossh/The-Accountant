
FROM golang:1.20
WORKDIR /app
COPY . .
RUN go mod tidy

ENTRYPOINT [ "make", "start"]