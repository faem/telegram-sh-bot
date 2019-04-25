FROM golang

COPY . /telegram-bot

WORKDIR /telegram-bot

RUN go mod tidy
RUN go build

ENTRYPOINT ["/telegram-bot/tbot"]