
# # STEP 1 build executable binary
# FROM golang:alpine as builder
# # install git to get all the dependencies via go get
# RUN apk update && apk add --no-cache git
# # Set workdir to the app source code path
# WORKDIR $GOPATH/src/github.com/chandanghosh/goexercises/ex-mongo-mvc/
# COPY . .
# # RUN git init
# # # fetch dependencies
# # RUN go get -u -d -v
# # build the binary
# RUN go build -ldflags "-s -w" -o /go/bin/mvc-mongo-goapp

# build a small image now
FROM scratch
WORKDIR /app
COPY ./ex-mongo-mvc .

#RUN chmod +x ./ex-mongo-mvc
ENTRYPOINT ["./ex-mongo-mvc" ]