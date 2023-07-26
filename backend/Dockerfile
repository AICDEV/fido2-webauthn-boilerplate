FROM golang:1.20.4-alpine3.16 AS builder
RUN mkdir /app
COPY . /app
WORKDIR /app 
RUN export GOOS="linux" \
    && export PATH=$PATH:/go/bin \
    && export GARC="amd64" \
    && go build -ldflags="-s -w" -o fido


FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/fido .
CMD ["./fido"]
