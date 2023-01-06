FROM golang:1.19-alpine as build-base

WORKDIR /app

 COPY go.mod ./
# COPY go.sum ./

RUN go mod download

COPY . .

# RUN CGO_ENABLED=0 go test -v

RUN go build -o ./out/assessment .

# ====================

FROM alpine:3.16.2
COPY --from=build-base /app/out/assessment /app/assessment
EXPOSE 2565
CMD ["/app/assessment"]
