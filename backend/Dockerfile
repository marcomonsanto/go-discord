# Start from the latest golang base image
FROM golang:1.21-bullseye

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

ENV APP_ENV=production

# Command to run the executable
CMD ["./main"]