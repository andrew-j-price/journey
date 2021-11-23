FROM golang:1.17-alpine AS go_builder
WORKDIR /app
ADD . /app/
# RUN go mod tidy
RUN CGO_ENABLED=0 go build -o drive

# FROM alpine:3.14
# RUN apk update && apk add --no-cache curl
FROM scratch
WORKDIR /app
COPY --from=go_builder /app/drive .
# COPY --from=go_builder /app/static/* static/
# COPY --from=go_builder /app/templates/* templates/
CMD ["./drive", "-api"]
