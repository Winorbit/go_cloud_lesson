FROM golang:1.17.0-alpine3.13
RUN mkdir /app
ADD . /app
WORKDIR /app
ENV GOOGLE_APPLICATION_CREDENTIALS=/app/kyrsypython_gcp_key.json

RUN go mod tidy

RUN go build -o main .
CMD ["/app/main"]