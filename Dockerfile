# 本番用
FROM golang:1.16-alpine as builder
# ワーキングディレクトリの設定
ENV ROOT=/usr/src
WORKDIR ${ROOT}
# アップデートとgitのインストール
RUN apk update && apk add git alpine-sdk
# ホストのファイルをコンテナの作業ディレクトリに移行
# 依存関係
COPY go.mod go.sum ./
RUN go mod download

WORKDIR ${ROOT}/app

# ソースコードとか変更頻度が高いものは後
COPY . .

RUN GOOS=linux go build -o ${ROOT}/bin ./cmd/pokewordle_solver/main.go
# WORKDIR ${ROOT}/bin
# RUN touch .bin

FROM alpine:latest as release

ENV ROOT=/usr/src
ENV GIN_MODE=release
WORKDIR ${ROOT}
COPY --from=builder  ${ROOT}/bin .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY ./db/ ./db/
EXPOSE 8080
CMD ["./bin"]