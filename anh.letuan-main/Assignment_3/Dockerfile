FROM golang:1.16-alpine
#ENV GO111MODULE=on

# Set destination for COPY
WORKDIR /Assignment_3/src

## Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . .

# Build
RUN go build -o main

# This is for documentation purposes only.
# To actually open the port, runtime parameters
# must be supplied to the docker command.
EXPOSE 8080

# (Optional) environment variable that our dockerised
# application can make use of. The value of environment
# variables can also be set via parameters supplied
# to the docker command on the command line.
ENV HTTP_PORT=8080

# Runs
# Command to run when starting the container
CMD ["/Assignment_3/src/main"]


