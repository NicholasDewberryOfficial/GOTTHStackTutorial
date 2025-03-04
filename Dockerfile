#https://betterstack.com/community/guides/scaling-go/dockerize-golang/
# Use Go 1.23 bookworm as base image

FROM golang:1.23.5-bookworm

# Move to working directory /build
WORKDIR /build

# Copy the go.mod and go.sum files to the /build directory
COPY go.mod ./

#COPY go.sum ./

# Install dependencies
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the application
RUN go build -o /out/dist .

# Start the application
CMD ["./out/dist"]
# Document the port that may need to be published
EXPOSE 8000
