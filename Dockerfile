FROM golang:latest

WORKDIR /app

# Goモジュールの初期化
RUN go mod init tview-tutorial

RUN go get -u github.com/rivo/tview
