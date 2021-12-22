FROM golang:alpine AS build
# git required for go mod
# RUN apk add --no-cache git
# Working directory will be created if it does not exist
ENV HOME /src
ENV CGO_ENABLED 0
ENV GOOS linux

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .

#ARG VERSION
# Build the executable
RUN go build -ldflags="-s -w -X main.appVersion=${VERSION}" -o app main.go

# STAGE 2: build the container to run
#FROM gcr.io/distroless/static AS final
FROM scratch AS final

WORKDIR /root/

# copy compiled app
COPY --from=build /src/app .

EXPOSE 8080
# run binary
ENTRYPOINT ["./app"]