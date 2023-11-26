# ビルドされる実行ファイルの名前
BINARY_NAME=main

build:
	go build -o $(BINARY_NAME) .

run:
	./$(BINARY_NAME)

clean:
	go clean
	rm $(BINARY_NAME)