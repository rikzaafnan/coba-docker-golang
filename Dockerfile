# ambil image golang dari docker.hub
FROM golang:1.18

# copy web.go ke dalam /app/main.go
COPY web.go /app/main.go

# menjalankan perintahnya dengan perintah CMD dan dijalankan dengan array
# CMD go run main.go
CMD ["go", "run", "/app/main.go"]


