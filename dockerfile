FROM golang:1.23
WORKDIR /app
COPY . .
ARG GO_BUILD_FLAGS=""
RUN go build ${GO_BUILD_FLAGS} -o server .
EXPOSE 8080
CMD ["./server"]