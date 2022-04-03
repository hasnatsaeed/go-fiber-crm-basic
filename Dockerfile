#  base image for Go
FROM golang:latest
# Set the Current Working Directory inside the container

WORKDIR /app

RUN git clone https://github.com/hasnatsaeed/go-fiber-crm-basic.git

WORKDIR /app/go-fiber-crm-basic

EXPOSE 9010

ENTRYPOINT ["go", "run", "/app/go-fiber-crm-basic/cmd/main/main.go"]