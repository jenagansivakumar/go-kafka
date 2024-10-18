# Use a base image with Go installed
FROM golang:1.20

# Set working directory inside container
WORKDIR /app

# Copy your app's source code into the container
COPY . .

# Build the Go app
RUN go build -o worker ./worker.go

# Command to run the app when the container starts
CMD ["./worker"]
