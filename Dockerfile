FROM golang:1.24 AS build

WORKDIR /build

COPY . .
RUN go build -o bin/pingmon .

FROM scratch

WORKDIR /app

# COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
# COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /build/bin/pingmon /app

CMD [ "/app/pingmon" ]
