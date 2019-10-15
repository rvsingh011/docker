FROM golang
WORKDIR '/app'
RUN go get -u github.com/go-redis/redis/
COPY . . 
CMD ["go", "run", "main.go"]

