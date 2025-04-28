FROM golang:1.24

# Set app workdir
WORKDIR /go/src/app

# Copy dependencies list
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy application sources
COPY . .

# Build app
RUN go build -o app github.com/IliaSobolev/Torrseal/cmd

# Run app
CMD ["./app"]
