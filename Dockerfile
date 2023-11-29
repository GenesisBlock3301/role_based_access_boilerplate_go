FROM golang:alpine

# This is working directory inside container
WORKDIR /app

# COPY go.mod or go.sum inside container directory.
COPY go.mod go.sum ./

# Now that you have the module files inside the Docker
# image that you are building, you can use the RUN command
# to run the command go mod download there as well.
# This works exactly the same as if you were running go
# locally on your machine, but this time these Go modules
# will be installed into a directory inside the image.
RUN go mod download

#The next thing you need to do is to copy your source code into the image
COPY . .

RUN go build -o role_based_authentication_golang
EXPOSE 8080

# RUN
CMD ["./role_based_authentication_golang"]






