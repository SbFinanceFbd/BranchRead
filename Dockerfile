FROM golang:1.19.4 AS build
WORKDIR /go/src
COPY src ./src
COPY BranchRead.go .
COPY go.mod .
COPY docs ./docs

ENV CGO_ENABLED=0
RUN go env -w GOPRIVATE="github.com/SbFinanceFbd/golib"
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o sbfin .

FROM alpine:3.17.2 AS runtime
ENV GIN_MODE=release
RUN mkdir app
RUN apk update && apk add bash && apk --no-cache add tzdata
RUN mkdir -p app/src/main
COPY --from=build /go/src/src/main/resources ./app/src/main/resources
COPY --from=build /go/src/docs ./app/docs
COPY --from=build /go/src/sbfin ./app/
WORKDIR /app
EXPOSE 3081
ENTRYPOINT ["./sbfin"]
