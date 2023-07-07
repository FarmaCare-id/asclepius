FROM golang:1.19 AS Production
WORKDIR /app
COPY go.mod ./
RUN go mod tidy
COPY . .
RUN go build -o farmacare
EXPOSE 8080
CMD /app/farmacare