FROM golang:1.21-alpine3.19

RUN mkdir -p /opt/github.com/jbullfrog81/fishing-buddy-service

WORKDIR /opt/github.com/jbullfrog81/fishing-buddy-service

# Custom cache invalidation
#ARG CACHEBUST=1

#RUN ls -la
#RUN ls -la
#RUN sleep 60
#COPY . .
COPY go.mod go.mod
COPY go.sum go.sum
COPY cmd/main.go main.go
COPY internal ./internal/
#COPY ../go.sum ./
RUN go mod vendor

#COPY cmd/main.go ./
#ADD internal ./internal/

RUN go build -o /bin/fishing-buddy -v .

EXPOSE 8080

ENTRYPOINT [ "fishing-buddy" ]