# tart by building the application.
# Build em apis3produto com distroless
FROM golang:1.12.7 as builder

WORKDIR /go/src/apiw
ENV GO111MODULE=on
COPY . .
#RUN go install -v ./...
RUN GOOS=linux go  build -ldflags="-s -w" -o apiw main.go
RUN cp apiw /go/bin/apiw
#RUN ls -lh

# Now copy it into our base image.
FROM gcr.io/distroless/base
COPY --from=builder /go/bin/apiw /
CMD ["/apiw"]