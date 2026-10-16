FROM golang:1.26-alpine AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /corlys .

FROM scratch
COPY --from=build /corlys /corlys
EXPOSE 8888
ENTRYPOINT ["/corlys"]
