FROM golang:1.16-alpine as dev
# ワーキングディレクトリの設定
ENV ROOT=/usr/src
WORKDIR ${ROOT}
# アップデートとgitのインストール
RUN apk update && apk add git && apk add bash && apk add --no-cache coreutils
# ホストのファイルをコンテナの作業ディレクトリに移行
# 依存関係
COPY go.mod go.sum ./
RUN go mod download

WORKDIR ${ROOT}/app

# ソースコードとか変更頻度が高いものは後
COPY . .

CMD ["./wait-for-it.sh", "db:5432", "--", "go", "run", "./cmd/pokewordle_solver/main.go"]

FROM golang:1.16-alpine as builder
# ワーキングディレクトリの設定
ENV ROOT=/usr/src
WORKDIR ${ROOT}
# アップデートとgitのインストール
RUN apk update && apk add git
# ホストのファイルをコンテナの作業ディレクトリに移行
# 依存関係
COPY go.mod go.sum ./
RUN go mod download

WORKDIR ${ROOT}/app

# ソースコードとか変更頻度が高いものは後
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ${ROOT}/bin ./cmd/pokewordle_solver/main.go

FROM alpine:latest as release

RUN apk update && apk add bash && apk add --no-cache coreutils

ENV ROOT=/usr/src
ENV GIN_MODE=release
WORKDIR ${ROOT}
COPY --from=builder  ${ROOT}/bin .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY ./csv/ ./csv/
COPY ./wait-for-it.sh /usr/src/
CMD ["./wait-for-it.sh", "db:5432", "--", "/usr/src/bin"]