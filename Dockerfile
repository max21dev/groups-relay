ARG GO_VERSION=1.23.1

FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS build

WORKDIR /src
COPY . .

# Download dependencies as a separate step to take advantage of Docker's caching.
RUN go mod download -x

# Build the application.
RUN go build -o /app/groups-relay .

# Expose the port that the application listens on.
EXPOSE 9005

# What the container should run when it is started.
ENTRYPOINT [ "/app/groups-relay" ]
