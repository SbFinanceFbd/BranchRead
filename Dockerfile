FROM golang:1.19.4 AS build
WORKDIR /go/src
COPY src ./src
COPY main.go .
COPY go.mod .
COPY docs ./docs

ENV CGO_ENABLED=0
RUN go env -w GOPRIVATE="github.com/SbFinanceFbd/golib"
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o sbfin .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src ./
COPY --from=build /go/src/sbfin ./
EXPOSE 3081
ENTRYPOINT ["./sbfin"]
