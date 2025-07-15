# This is still a work in progress

FROM golang AS build

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY main.go .
COPY internal .

RUN go mod tidy -v
RUN go build -ldflags "-s -w" -o bin/pingmon .

FROM scratch

WORKDIR /app

COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /build/bin/pingmon /app

CMD [ "/app/pingmon" ]
