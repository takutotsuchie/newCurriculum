FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
# sqlboilerのダウンロード
RUN go install github.com/volatiletech/sqlboiler/v4@latest
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
RUN go install github.com/cosmtrek/air@v1.29.0


COPY . .
RUN go mod tidy


CMD ["air", "-c", ".air.toml"]