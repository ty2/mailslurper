FROM golang:1.22.2-alpine3.19 as builder

LABEL maintainer="erguotou525@gmail.compute"

WORKDIR /app
COPY ./go.mod .
COPY ./go.sum .

RUN go mod download
RUN go install github.com/mjibson/esc@latest # TODO: Consider using native file embedding

COPY . .
WORKDIR cmd/mailslurper

RUN go mod tidy
RUN go generate

RUN apk --no-cache add build-base

RUN CGO_ENABLED=1 go build -o mailslurper

FROM alpine:3.19.1

RUN apk add --no-cache ca-certificates \
 && echo -e '{\n\
  "wwwAddress": "0.0.0.0",\n\
  "wwwPort": 8080,\n\
  "wwwPublicURL": "",\n\
  "serviceAddress": "0.0.0.0",\n\
  "servicePort": 8085,\n\
  "servicePublicURL": "",\n\
  "smtpAddress": "0.0.0.0",\n\
  "smtpPort": 2500,\n\
  "dbEngine": "SQLite",\n\
  "dbHost": "",\n\
  "dbPort": 0,\n\
  "dbDatabase": "./mailslurper.db",\n\
  "dbUserName": "",\n\
  "dbPassword": "",\n\
  "maxWorkers": 1000,\n\
  "autoStartBrowser": false,\n\
  "keyFile": "",\n\
  "certFile": "",\n\
  "adminKeyFile": "",\n\
  "adminCertFile": ""\n\
  }'\
  >> config.json

COPY --from=builder /app/cmd/mailslurper/mailslurper mailslurper

EXPOSE 8080 8085 2500

CMD ["./mailslurper"]
