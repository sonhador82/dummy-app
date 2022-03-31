FROM golang:1.16 as builder

ENV CGO_ENABLED=0

ADD . /app
WORKDIR /app

RUN go build -o service ./main.go


FROM scratch
COPY --from=builder /app/service /service
ENTRYPOINT [ "/service" ]
EXPOSE 8000
