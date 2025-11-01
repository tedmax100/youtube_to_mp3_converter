# YouTube 轉 MP3 工具

[![CI](https://github.com/YOUR_USERNAME/youtube_to_mp3/workflows/CI/badge.svg)](https://github.com/YOUR_USERNAME/youtube_to_mp3/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/YOUR_USERNAME/youtube_to_mp3)](https://goreportcard.com/report/github.com/YOUR_USERNAME/youtube_to_mp3)
[![codecov](https://codecov.io/gh/YOUR_USERNAME/youtube_to_mp3/branch/main/graph/badge.svg)](https://codecov.io/gh/YOUR_USERNAME/youtube_to_mp3)

這是一個用 Go 編寫的工具，可以從 YouTube 下載視頻並轉換為高品質 MP3 格式（320kbps）。

## 功能特點

- 下載 YouTube 視頻並自動轉換為 MP3
- 輸出 320kbps 高品質音頻
- 自動保存到 `output` 目錄
- 使用視頻原始標題作為文件名
- 模塊化設計，易於測試和擴展
- 完整的單元測試和集成測試覆蓋

## 項目結構

```
youtube_to_mp3/
├── main.go                    # 主程序入口
├── main_test.go               # 主程序測試
├── go.mod                     # Go 模塊定義
├── Makefile                   # 構建和測試命令
├── pkg/
│   ├── config/               # 配置管理
│   │   ├── config.go
│   │   └── config_test.go
│   ├── downloader/           # 下載器實現
│   │   ├── downloader.go
│   │   └── downloader_test.go
│   └── validator/            # 依賴驗證器
│       ├── validator.go
│       └── validator_test.go
└── test/
    ├── integration/          # 集成測試
    │   └── integration_test.go
    └── mocks/                # 測試用的 mock 對象
        └── mocks.go
```

## 系統需求

在使用此工具之前，需要安裝以下依賴：

### 1. yt-dlp

**macOS:**
```bash
brew install yt-dlp
```

**Linux (Ubuntu/Debian):**
```bash
sudo apt update
sudo apt install yt-dlp
```

或使用 pip：
```bash
pip install yt-dlp
```

**Windows:**
```bash
pip install yt-dlp
```

### 2. FFmpeg

**macOS:**
```bash
brew install ffmpeg
```

**Linux (Ubuntu/Debian):**
```bash
sudo apt update
sudo apt install ffmpeg
```

**Windows:**
從 [FFmpeg 官網](https://ffmpeg.org/download.html) 下載並安裝

### 3. Go

需要 Go 1.16 或更高版本。從 [Go 官網](https://golang.org/dl/) 下載安裝。

## 使用方法

### 使用 Makefile（推薦）

```bash
# 查看所有可用命令
make help

# 編譯程序
make build

# 運行程序
make run URL="https://www.youtube.com/watch?v=VIDEO_ID"
```

### 方法 1: 直接運行

```bash
go run main.go "https://www.youtube.com/watch?v=VIDEO_ID"
```

### 方法 2: 編譯後運行

```bash
# 編譯
make build
# 或
go build -o youtube_to_mp3

# 運行
./youtube_to_mp3 "https://www.youtube.com/watch?v=VIDEO_ID"
```

## 使用範例

```bash
# 下載單個視頻
go run main.go "https://www.youtube.com/watch?v=dQw4w9WgXcQ"

# 使用 Makefile
make run URL="https://www.youtube.com/watch?v=dQw4w9WgXcQ"

# 下載播放列表中的視頻
go run main.go "https://www.youtube.com/playlist?list=PLAYLIST_ID"
```

## 輸出

所有轉換後的 MP3 文件將保存在 `output` 目錄中，文件名為視頻的原始標題。

## 注意事項

- 請確保您有權下載和轉換視頻內容
- 下載速度取決於您的網路和視頻大小
- 轉換過程可能需要一些時間，特別是對於較長的視頻
- 請遵守 YouTube 的服務條款

## 故障排除

### 錯誤: "未找到 yt-dlp"
請按照上述說明安裝 yt-dlp。

### 錯誤: "未找到 ffmpeg"
請按照上述說明安裝 FFmpeg。

### 下載失敗
- 檢查 URL 是否正確
- 確認網路正常
- 某些視頻可能有地區限制或其他訪問限制

## 測試

本項目包含完整的單元測試和集成測試。

### 運行測試

```bash
# 運行所有單元測試
make test

# 運行單元測試（詳細模式）
make test-unit

# 運行E2E測試（需要 yt-dlp 和 ffmpeg）
make test-integration

# 運行所有測試（包括集成測試）
make test-all

# 生成測試覆蓋率報告
make test-coverage

# 查看覆蓋率報告（在瀏覽器中打開）
make coverage-view
```

### 測試結構

#### 單元測試

- **config 包測試** (`pkg/config/config_test.go`)
  - 配置創建和修改
  - 方法鏈式調用

- **validator 包測試** (`pkg/validator/validator_test.go`)
  - 依賴檢查功能
  - Mock 對象測試
  - 錯誤處理

- **downloader 包測試** (`pkg/downloader/downloader_test.go`)
  - 下載功能
  - 命令參數構建
  - 文件輸出檢查

#### E2E測試

- **端到端測試** (`test/integration/integration_test.go`)
  - 完整的下載流程測試
  - 實際的 yt-dlp 和 ffmpeg 集成
  - 需要網路連接

### 運行E2E測試

集成測試需要：
- 安裝 yt-dlp 和 ffmpeg
- 網路連接

```bash
# 運行E2E測試
go test -v -tags=integration ./test/integration/...

# 或使用 Makefile
make test-integration
```

### 測試覆蓋率

```bash
# 生成覆蓋率報告
make test-coverage

# 這將生成：
# - coverage.txt (覆蓋率數據)
# - coverage.html (HTML 格式報告)
```

## 開發指南

### 添加新功能

1. 在相應的包中添加新代碼
2. 為新功能編寫單元測試
3. 如果需要，添加E2E測試
4. 運行測試確保通過：`make test`
5. 檢查測試覆蓋率：`make test-coverage`

### 代碼結構

- **pkg/config**: 配置管理，支持自定義輸出目錄、比特率等
- **pkg/validator**: 依賴驗證，檢查系統是否安裝必要工具
- **pkg/downloader**: 下載和轉換邏輯，使用接口設計便於測試
- **test/mocks**: 測試用的 mock 對象

### Mock 對象

項目使用interface和DI模式，便於單元測試：

```go
// 例如：測試下載器時使用 mock 命令執行器
mockExecutor := &MockCommandExecutor{
    executeFunc: func(name string, args []string, stdout, stderr io.Writer) error {
        // 模擬命令執行
        return nil
    },
}
downloader := NewYtDlpDownloader(cfg, mockExecutor)
```

## CI/CD

本項目使用 GitHub Actions 進行持續集成和測試。

### GitHub Actions 工作流

項目包含以下 CI 工作流（`.github/workflows/ci.yml`）：

1. **測試（Test）**
   - 在多個 Go 版本（1.20, 1.21, 1.22）上運行
   - 執行所有單元測試
   - 生成測試覆蓋率報告
   - 上傳覆蓋率到 Codecov

2. **代碼檢查（Lint）**
   - 使用 golangci-lint 進行代碼質量檢查
   - 檢查代碼風格、潛在錯誤和最佳實踐

3. **構建（Build）**
   - 編譯程序確保無編譯錯誤
   - 上傳build artifacts

### 觸發條件

CI 會在以下情況自動運行：
- 推送到 `main`、`master` 或 `develop` 分支
- 創建 Pull Request 到這些分支

### 本地運行 CI 檢查

在提交程式碼前，你可以在local運行相同的檢查：

```bash
# 運行測試
make test

# 運行 linter
make lint

# 生成覆蓋率報告
make test-coverage

# 構建程序
make build
```

### 配置 Badge

將 README 中的 `YOUR_USERNAME` 替換為你的 GitHub 用戶名：

```markdown
[![CI](https://github.com/tedmax100/youtube_to_mp3/workflows/CI/badge.svg)](https://github.com/tedmax100/youtube_to_mp3/actions)
```

### Codecov 設置（可選）

如果想使用 Codecov 追蹤測試覆蓋率：

1. 在 [codecov.io](https://codecov.io) 註冊並連接你的 GitHub 倉庫
2. Codecov 會自動接收 GitHub Actions 上傳的覆蓋率報告
3. 無需額外配置 token（對於公開倉庫）

## 技術細節

- 使用 `yt-dlp` 下載 YouTube 視頻
- 使用 `ffmpeg` 轉換音頻為 MP3 格式
- 音頻比特率: 320kbps
- 音頻格式: MP3
- 採用interface設計模式，便於測試和擴展
- 完整的錯誤處理和依賴檢查

## 授權

本項目僅供學習和個人使用。請尊重版權，不要用於商業用途。
