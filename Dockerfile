FROM golang
#Setting environment variables, which provides running our golang binary file on linux
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
COPY . /snippetboxapp
WORKDIR /snippetboxapp
RUN go mod download
RUN chmod +x /snippetboxapp/cmd/web
RUN go build /snippetboxapp/cmd/web # This will create a binary file named web on /snippetboxapp
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
#getting binary file from already created app image to alpine image to run
COPY --from=0 /snippetboxapp .
ENTRYPOINT ./web
EXPOSE 4000