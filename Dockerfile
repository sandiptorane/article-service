FROM golang:1.19-alpine as builder
# create app dir
RUN mkdir /app
# copy all files to app directory
ADD ./ /app
# set workdir
WORKDIR /app
# build the application
RUN go build -o main ./cmd

# final stage
FROM alpine:latest

# create app dir
RUN mkdir /app

# set workdir to app
WORKDIR /app

# copy prebuilt binary from previous stage to new stage
COPY --from=builder /app/main /app
#copy configs
COPY --from=builder /app/configs /app/configs

# run the executable
CMD ["/app/main"]