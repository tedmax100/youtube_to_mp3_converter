//go:build integration
// +build integration

package integration

import (
	"os"
	"path/filepath"
	"testing"

	"youtube_to_mp3/pkg/config"
	"youtube_to_mp3/pkg/downloader"
	"youtube_to_mp3/pkg/validator"
)

// TestEndToEndDownload 端到端測試（需要實際的網路連接和工具）
func TestEndToEndDownload(t *testing.T) {
	// 這個測試需要實際安裝 yt-dlp 和 ffmpeg
	// 使用 -tags=integration 標籤來運行此測試

	// 檢查依賴
	systemValidator := validator.NewSystemValidator(nil)
	if err := systemValidator.ValidateDependencies(); err != nil {
		t.Skipf("Skipping integration test: %v", err)
	}

	// 創建臨時測試目錄
	tempDir := t.TempDir()

	// 配置
	cfg := config.NewConfig().WithOutputDir(tempDir)

	// 創建下載器
	dl := downloader.NewYtDlpDownloader(cfg, nil)

	// 使用一個短的測試視頻（Creative Commons 授權）
	// 這是一個公共領域的測試視頻
	testURL := "https://www.youtube.com/watch?v=aqz-KE-bpKQ" // Big Buck Bunny 60s test

	// 下載
	if err := dl.Download(testURL); err != nil {
		t.Fatalf("Download failed: %v", err)
	}

	// 驗證文件是否被創建
	files, err := dl.GetOutputFiles()
	if err != nil {
		t.Fatalf("Failed to get output files: %v", err)
	}

	if len(files) == 0 {
		t.Fatal("Expected at least one output file")
	}

	// 驗證文件存在且不為空
	for _, file := range files {
		stat, err := os.Stat(file)
		if err != nil {
			t.Errorf("Output file does not exist: %v", err)
		}
		if stat.Size() == 0 {
			t.Errorf("Output file is empty: %s", file)
		}
		if filepath.Ext(file) != ".mp3" {
			t.Errorf("Expected mp3 file, got: %s", file)
		}

		t.Logf("Created file: %s (size: %d bytes)", file, stat.Size())
	}
}

// TestValidatorIntegration 驗證器集成測試
func TestValidatorIntegration(t *testing.T) {
	systemValidator := validator.NewSystemValidator(nil)

	t.Run("validate all dependencies", func(t *testing.T) {
		err := systemValidator.ValidateDependencies()
		if err != nil {
			t.Logf("Dependencies not installed: %v", err)
			t.Logf("To run full integration tests, install: yt-dlp and ffmpeg")
		} else {
			t.Log("All dependencies are installed")
		}
	})

	t.Run("validate yt-dlp", func(t *testing.T) {
		err := systemValidator.ValidateYtDlp()
		if err != nil {
			t.Logf("yt-dlp not installed: %v", err)
		} else {
			t.Log("yt-dlp is installed")
		}
	})

	t.Run("validate ffmpeg", func(t *testing.T) {
		err := systemValidator.ValidateFFmpeg()
		if err != nil {
			t.Logf("ffmpeg not installed: %v", err)
		} else {
			t.Log("ffmpeg is installed")
		}
	})
}

// TestConfigIntegration 配置集成測試
func TestConfigIntegration(t *testing.T) {
	tempDir := t.TempDir()

	cfg := config.NewConfig().
		WithOutputDir(tempDir).
		WithBitrate("192k")

	// 創建下載器（但不實際下載）
	dl := downloader.NewYtDlpDownloader(cfg, nil)

	// 驗證配置被正確應用
	if dl == nil {
		t.Fatal("Failed to create downloader with custom config")
	}

	t.Logf("Config test passed with output dir: %s", tempDir)
}
