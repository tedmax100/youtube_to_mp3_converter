.PHONY: help test test-unit test-integration test-coverage test-verbose clean build run install-deps

# 默認目標
help:
	@echo "可用的命令："
	@echo "  make test              - 運行所有單元測試"
	@echo "  make test-unit         - 運行單元測試"
	@echo "  make test-integration  - 運行E2E測試（需要 yt-dlp 和 ffmpeg）"
	@echo "  make test-coverage     - 運行測試並生成覆蓋率報告"
	@echo "  make test-verbose      - 運行詳細模式的測試"
	@echo "  make build             - 編譯程序"
	@echo "  make run URL=<url>     - 運行程序"
	@echo "  make clean             - 清理編譯文件和輸出"
	@echo "  make install-deps      - 安裝依賴（yt-dlp 和 ffmpeg）"

# 運行所有單元測試
test:
	@echo "運行單元測試..."
	go test -v ./pkg/...

# 只運行單元測試
test-unit:
	@echo "運行單元測試..."
	go test -v ./pkg/config ./pkg/validator ./pkg/downloader

# 運行E2E測試
test-integration:
	@echo "運行E2E測試..."
	@echo "注意：此測試需要網路連接和實際的 yt-dlp/ffmpeg 安裝"
	go test -v -tags=integration ./test/integration/...

# 生成測試覆蓋率報告
test-coverage:
	@echo "生成測試覆蓋率報告..."
	go test -v -coverprofile=coverage.txt ./pkg/...
	go tool cover -html=coverage.txt -o coverage.html
	@echo "覆蓋率報告已生成: coverage.html"

# 詳細模式測試
test-verbose:
	@echo "運行詳細模式測試..."
	go test -v -race ./...

# 運行所有測試（包括E2E測試）
test-all:
	@echo "運行所有測試..."
	go test -v ./pkg/...
	@echo "\n運行E2E測試..."
	go test -v -tags=integration ./test/integration/...

# 編譯程序
build:
	@echo "編譯程序..."
	go build -o youtube_to_mp3 .
	@echo "編譯完成: youtube_to_mp3"

# 運行程序
run:
ifndef URL
	@echo "錯誤: 請提供 URL 參數"
	@echo "用法: make run URL=https://www.youtube.com/watch?v=..."
	@exit 1
endif
	go run main.go "$(URL)"

# 清理編譯文件和輸出
clean:
	@echo "清理文件..."
	rm -f youtube_to_mp3
	rm -f coverage.out coverage.txt coverage.html
	rm -rf output/
	@echo "清理完成"

# 安裝依賴（僅作為參考，實際安裝可能需要 sudo）
install-deps:
	@echo "請手動安裝以下依賴："
	@echo ""
	@echo "yt-dlp:"
	@echo "  pip install yt-dlp"
	@echo "  或 brew install yt-dlp (macOS)"
	@echo ""
	@echo "ffmpeg:"
	@echo "  sudo apt install ffmpeg (Ubuntu/Debian)"
	@echo "  或 brew install ffmpeg (macOS)"

# 格式化代碼
fmt:
	@echo "格式化代碼..."
	go fmt ./...

# 運行 linter
lint:
	@echo "運行 linter..."
	@which golangci-lint > /dev/null || (echo "請先安裝 golangci-lint"; exit 1)
	golangci-lint run ./...

# 查看測試覆蓋率
coverage-view: test-coverage
	@echo "在瀏覽器中打開覆蓋率報告..."
	@which xdg-open > /dev/null && xdg-open coverage.html || open coverage.html
