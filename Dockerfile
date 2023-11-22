FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application files to the working directory
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port the application runs on
EXPOSE 8089

# Command to run the executable
CMD ["./main"]
