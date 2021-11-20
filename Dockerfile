FROM golang:1.17-alpine AS go_builder
WORKDIR /app
ADD . /app/
# RUN go mod tidy
RUN CGO_ENABLED=0 go build -o drive

FROM alpine:3.15
WORKDIR /app
RUN apk add --no-cache bash
COPY --from=go_builder /app/drive .
# COPY --from=go_builder /app/static/* static/
# COPY --from=go_builder /app/templates/* templates/
EXPOSE 8080
CMD ["./drive"]
