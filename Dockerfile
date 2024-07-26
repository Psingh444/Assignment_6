


# Use the official Golang image as a base
FROM golang:1.20

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod file
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod file is not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o /browser-info

# Expose port 80 to the outside world
EXPOSE 80

# Command to run the executable
CMD ["/browser-info"]
