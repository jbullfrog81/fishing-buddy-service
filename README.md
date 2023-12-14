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

## Local Setup
```
docker compose -f ./docker/compose.yaml up
```

## Run Locally
```
# Bring up cache and mysql
docker compose up -d cache
docker compose up -d mysql
# Build the app
docker build -f dockerfile .
# Run the app
docker run -it --network fishing-buddy-service_default -p 8080:8080 165e542b0dcb4ab0b6bfaf04cf9d30abaf3847edf89197137e36a1c708e7dcca
```

