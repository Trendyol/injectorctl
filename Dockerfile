FROM golang:1.13-buster as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o injectorctl ./*.go

FROM alpine
COPY --from=build /app/injectorctl /usr/local/bin/injectorctl
COPY --from=build app/entrypoint.sh entrypoint.sh
RUN chmod +x entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]