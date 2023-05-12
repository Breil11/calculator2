#FROM golang:1.19-buster AS build
FROM golang:1.19-buster-latest AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

#ADD . .
COPY . .

RUN go build -o /calculator

#FROM gcr.io/distroless/base-debian10
FROM gcr.io/distroless/base-debian10:v2.1

WORKDIR /

COPY --from=build /calculator /calculator

USER nonroot:nonroot

ENTRYPOINT ["/calculator"]
