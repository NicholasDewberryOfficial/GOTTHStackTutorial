#####################################
FROM golang:1.23.5-alpine AS build-env

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first if you are using modules (so Docker can cache module downloads)
# If your project doesn't use modules, you can skip these lines.
COPY go.mod ./
RUN go mod download

# Copy the rest of your application code
COPY . .


# Build the Go program
# Replace <YOUR_MAIN_FILE.go> with the actual file that contains your main() function.
# The -o flag names the output binary "app".
RUN go build -o app .

RUN ls -l /app

######################################
# Step 2: Final minimal image
######################################
FROM alpine:3.17

# Set the working directory for our minimal image
WORKDIR /app

# Copy binary from build stage
COPY --from=build-env /app/app /app/

# Copy any assets you need at runtime (e.g., static files), if not baked into templates
#
COPY --from=build-env /app/static /app/static
COPY --from=build-env /app/slices /app/slices

# Make sure the port matches what your Go server listens on
EXPOSE 8080

# Start the Go web server when the container runs
CMD ["./app"]
