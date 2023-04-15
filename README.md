# fishing-buddy-service
This is the backend service for the fishing buddy application

## Build Image
```
docker build -f docker/dockerfile .
```

## Run Image
```
docker run -d -p 127.0.0.1:8080:8080/tcp <IMAGE ID>
```