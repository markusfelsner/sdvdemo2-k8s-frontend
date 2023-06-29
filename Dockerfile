FROM golang:1.17-alpine AS build

WORKDIR /src/
COPY main.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/sdvdemo

FROM scratch
COPY --from=build /bin/sdvdemo /bin/sdvdemo
ENTRYPOINT ["/bin/sdvdemo"]
