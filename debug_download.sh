#!/bin/bash

# Debug 腳本 - 用於診斷下載和轉換問題

URL="$1"

if [ -z "$URL" ]; then
    echo "用法: ./debug_download.sh <YouTube URL>"
    exit 1
fi

echo "=== 開始調試下載 ==="
echo "URL: $URL"
echo ""

echo "=== 檢查依賴 ==="
which yt-dlp
which ffmpeg
echo ""

echo "=== 顯示 yt-dlp 版本 ==="
yt-dlp --version
echo ""

echo "=== 顯示 ffmpeg 版本 ==="
ffmpeg -version | head -1
echo ""

echo "=== 獲取視頻信息 ==="
yt-dlp --print "%(title)s - %(duration)s 秒 - %(filesize_approx)s bytes" "$URL"
echo ""

echo "=== 開始下載並轉換（帶詳細輸出）==="
set -x  # 顯示執行的命令

yt-dlp \
    --extract-audio \
    --audio-format mp3 \
    --audio-quality 0 \
    --postprocessor-args "ffmpeg:-b:a 320k" \
    --progress \
    --newline \
    --no-playlist \
    --verbose \
    -o "output/%(title)s.%(ext)s" \
    "$URL"

EXIT_CODE=$?
set +x

echo ""
echo "=== 下載完成 ==="
echo "退出碼: $EXIT_CODE"
echo ""

echo "=== 輸出目錄內容 ==="
ls -lh output/
echo ""

if [ $EXIT_CODE -eq 0 ]; then
    echo "✓ 成功！"
else
    echo "✗ 失敗，退出碼: $EXIT_CODE"
fi
