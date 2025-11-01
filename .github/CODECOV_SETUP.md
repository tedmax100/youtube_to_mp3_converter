# Codecov 設置指南

本文檔說明如何為此項目設置 Codecov 測試覆蓋率追蹤。

## 步驟 1: 註冊 Codecov

1. 訪問 [codecov.io](https://codecov.io)
2. 使用 GitHub 帳號登錄
3. 授權 Codecov 訪問你的 GitHub 倉庫

## 步驟 2: 添加倉庫

1. 在 Codecov 控制台中，點擊 "Add new repository"
2. 找到並選擇 `youtube_to_mp3` 倉庫
3. Codecov 會生成一個上傳 token

## 步驟 3: 添加 GitHub Secret

1. 前往 GitHub 倉庫頁面
2. 點擊 **Settings** > **Secrets and variables** > **Actions**
3. 點擊 **New repository secret**
4. 設置：
   - Name: `CODECOV_TOKEN`
   - Secret: 粘貼從 Codecov 獲取的 token
5. 點擊 **Add secret**

## 步驟 4: 驗證設置

1. 推送代碼到 GitHub
2. GitHub Actions 會自動運行
3. 測試完成後，覆蓋率報告會自動上傳到 Codecov
4. 在 Codecov 控制台查看覆蓋率報告

## Badge 配置

在 README.md 中已經包含了 Codecov badge：

```markdown
[![codecov](https://codecov.io/gh/YOUR_USERNAME/youtube_to_mp3/branch/main/graph/badge.svg)](https://codecov.io/gh/YOUR_USERNAME/youtube_to_mp3)
```

記得將 `YOUR_USERNAME` 替換為你的 GitHub 用戶名。

## 注意事項

- 對於公開倉庫，Codecov token 是可選的（但推薦使用）
- 對於私有倉庫，必須設置 CODECOV_TOKEN
- CI 配置中已設置 `fail_ci_if_error: false`，所以即使 Codecov 上傳失敗，CI 也不會失敗

## 本地測試覆蓋率

你也可以在本地生成覆蓋率報告：

```bash
# 生成覆蓋率報告
make test-coverage

# 在瀏覽器中查看
make coverage-view
```

這將生成：
- `coverage.txt` - 覆蓋率數據文件
- `coverage.html` - HTML 格式的可視化報告
