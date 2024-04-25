# Use an efficient base image 
FROM golang:alpine 

# Working directory for your app
WORKDIR /app

# Copy dependencies
COPY go.mod go.sum ./

RUN go mod download

# Copy source code
COPY . .

# Build the executable
RUN go build -o /app/backend 

# Command to run the container
CMD [ "/app/backend" ]