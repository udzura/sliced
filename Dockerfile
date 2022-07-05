FROM golang:1.18

RUN mkdir /app
ADD . /app
RUN mkdir /app/build && cd /app && \
    env CGO_ENABLED=0 GOOS=linux  GOARCH=amd64 go build -ldflags '-s -w' -trimpath -o /app/build/sliced-linux-amd64  sliced.go && \
    env CGO_ENABLED=0 GOOS=linux  GOARCH=arm64 go build -ldflags '-s -w' -trimpath -o /app/build/sliced-linux-arm64  sliced.go && \
    env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags '-s -w' -trimpath -o /app/build/sliced-darwin-amd64 sliced.go && \
    env CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags '-s -w' -trimpath -o /app/build/sliced-darwin-arm64 sliced.go && \
    echo finished

CMD ["sleep", "inf"]
