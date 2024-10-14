# Use the official Golang image as the base image
FROM golang:1.22.2

# Set the Current Working Directory inside the container
WORKDIR /app-golang

# Using ARG, ENV and EXPOSE
ARG PORT=8080
ENV PORT=${PORT}
EXPOSE ${PORT}

# Copy go.mod and go.sum files
COPY go.mod ./
COPY go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd/server/main.go

# Run program
CMD ["./main"]