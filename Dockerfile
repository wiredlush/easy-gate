FROM golang:1.23 AS go-builder
WORKDIR /easy-gate
COPY . .
RUN make

FROM alpine:3.18 AS easy-gate
ENV EASY_GATE_CONFIG_PATH="/etc/easy-gate/easy-gate.json"
WORKDIR /etc/easy-gate
COPY ./assets/easy-gate.json .
WORKDIR /usr/bin
COPY --from=go-builder ./easy-gate/build/easy-gate .
ENTRYPOINT ["/usr/bin/easy-gate"]