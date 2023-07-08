FROM golang:1.19 AS Production

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o farmacare
EXPOSE 8000
CMD ["/app/farmacare"]