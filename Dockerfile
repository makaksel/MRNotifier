FROM golang:1.25

WORKDIR /app

# зависимости
RUN go install github.com/air-verse/air@latest

# кэш модулей
COPY go.mod go.sum ./
RUN go mod download

# код (будет перезаписан volume)
COPY . .

CMD ["air"]