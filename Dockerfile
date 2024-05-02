FROM golang:1.22-bullseye

RUN go install github.com/beego/bee/v2@latest

ENV GO111MODULE=on

ENV APP_HOME /go/src/app
RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"

COPY . .

EXPOSE 8000

CMD ["bee", "run"]