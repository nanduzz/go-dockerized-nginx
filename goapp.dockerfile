############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

# All these steps will be cached
RUN mkdir /number
WORKDIR /number
COPY . . 

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o /go/bin/number

############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
#COPY --from=builder /number .
COPY --from=builder /go/bin/number .

#exposes the 80
EXPOSE 80
# Run the hello binary.
ENTRYPOINT ["/number"]

