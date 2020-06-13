FROM golang:1.12-alpine AS build
#Install git
RUN apk add --no-cache git
#Get the hello world package from a GitHub repository
RUN go get github.com/cloudtechlearncode/app-summary
WORKDIR /go/src/github.com/cloudtechlearncode/app-summary
# Build the project and send the output to /bin/HelloWorld 
RUN go build -o /bin/app-summary

FROM golang:1.12-alpine
#Copy the build's output binary from the previous build container
COPY --from=build /bin/app-summary /bin/app-summary
ENTRYPOINT ["/bin/app-summary"]
