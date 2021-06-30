# First stage: build the executable. Have to use alpine if not doesnt work
FROM golang:1.16.4-alpine3.13 AS source
WORKDIR /build
ADD . .
#install git 
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

COPY ./go.mod ./go.sum ./
RUN go mod download
RUN go build -o gateway-service .
# Move to /dist directory as the place for resulting binary folder
WORKDIR /app

# Copy binary from build to main folder
RUN mv /build/gateway-service .

FROM alpine:latest
# copy the repository form the previous image
COPY --from=source /app /app
# Command to run when starting the container
CMD ["./app/gateway-service"]
