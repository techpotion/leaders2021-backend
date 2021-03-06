FROM golang:alpine AS build

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o api .

FROM alpine AS bin
COPY --from=build /go/src/app/api /api
COPY --from=build /go/src/app/static /static
COPY --from=build /go/src/app/config.example.yaml /config.yaml
COPY --from=build /go/src/app/gen/swagger/api.swagger.json /gen/swagger/api.swagger.json
CMD ["/api"]