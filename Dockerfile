FROM --platform=linux/amd64 golang:1.20-alpine AS go_builder
WORKDIR /journey
COPY . .
RUN go build -o ./bin/api ./cmd/api
RUN go build -o ./bin/boondocks ./cmd/boondocks
RUN go build -o ./bin/hello_colors ./cmd/hello_colors
RUN go build -o ./bin/identity ./cmd/identity
RUN go build -o ./bin/math ./cmd/math
RUN go build -o ./bin/random ./cmd/random


FROM --platform=linux/amd64 alpine:3.18
ENV PATH="${PATH}:/journey/bin"
WORKDIR /journey
COPY --from=go_builder /journey/bin/api /journey/bin/api
COPY --from=go_builder /journey/bin/boondocks /journey/bin/boondocks
COPY --from=go_builder /journey/bin/hello_colors /journey/bin/hello_colors
COPY --from=go_builder /journey/bin/identity /journey/bin/identity
COPY --from=go_builder /journey/bin/math /journey/bin/math
COPY --from=go_builder /journey/bin/random /journey/bin/random
