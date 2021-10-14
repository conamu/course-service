FROM golang:latest as build-env

WORKDIR /app
COPY . .
RUN go install && go build -o app

FROM gcr.io/distroless/base
WORKDIR app/
COPY --from=build-env /app/app .
EXPOSE 8080
CMD ["./app"]
