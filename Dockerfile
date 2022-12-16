FROM golang:1.19.4-alpine3.17

# ADD . /app
# WORKDIR /app
# RUN go mod download
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
# RUN chmod +x ./kaskus
# EXPOSE 7788
# CMD /app/kaskus

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# working directory
WORKDIR /build

# Copy and download depedencies using go mod
COPY    go.mod .
COPY    go.sum .

# copy folder into container
COPY . .

# Build the application
RUN go build -o main .

EXPOSE 3030

# command to running executable file
CMD ["/build/main"]
