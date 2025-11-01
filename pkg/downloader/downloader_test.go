package downloader

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"youtube_to_mp3/pkg/config"
)

// MockCommandExecutor 模擬命令執行器
type MockCommandExecutor struct {
	executeFunc func(name string, args []string, stdout, stderr io.Writer) error
	lastCommand string
	lastArgs    []string
}

// Execute 執行命令（模擬實現）
func (m *MockCommandExecutor) Execute(name string, args []string, stdout, stderr io.Writer) error {
	m.lastCommand = name
	m.lastArgs = args

	if m.executeFunc != nil {
		return m.executeFunc(name, args, stdout, stderr)
	}
	return nil
}

func TestNewYtDlpDownloader(t *testing.T) {
	t.Run("with nil executor", func(t *testing.T) {
		cfg := config.NewConfig()
		downloader := NewYtDlpDownloader(cfg, nil)

		if downloader == nil {
			t.Error("Expected downloader to be created, got nil")
		}
		if downloader.executor == nil {
			t.Error("Expected default executor to be set, got nil")
		}
	})

	t.Run("with custom executor", func(t *testing.T) {
		cfg := config.NewConfig()
		mock := &MockCommandExecutor{}
		downloader := NewYtDlpDownloader(cfg, mock)

		if downloader == nil {
			t.Error("Expected downloader to be created, got nil")
		}
		if downloader.executor != mock {
			t.Error("Expected custom executor to be set")
		}
	})
}

func TestBuildArgs(t *testing.T) {
	cfg := config.NewConfig()
	downloader := NewYtDlpDownloader(cfg, nil)

	url := "https://www.youtube.com/watch?v=test123"
	args := downloader.buildArgs(url)

	// 檢查必要的參數是否存在
	expectedParams := []string{
		"--extract-audio",
		"--audio-format",
		"mp3",
		"--audio-quality",
		"0",
		"--postprocessor-args",
		"ffmpeg:-b:a 320k",
		"-o",
		url,
	}

	argsStr := strings.Join(args, " ")
	for _, param := range expectedParams {
		if !strings.Contains(argsStr, param) {
			t.Errorf("Expected args to contain '%s', got: %v", param, args)
		}
	}

	// 檢查最後一個參數是 URL
	if args[len(args)-1] != url {
		t.Errorf("Expected last arg to be URL '%s', got '%s'", url, args[len(args)-1])
	}
}

func TestDownload(t *testing.T) {
	t.Run("successful download", func(t *testing.T) {
		// 使用臨時目錄
		tempDir := t.TempDir()

		cfg := config.NewConfig().WithOutputDir(tempDir)
		mock := &MockCommandExecutor{
			executeFunc: func(name string, args []string, stdout, stderr io.Writer) error {
				// 模擬成功執行
				return nil
			},
		}

		downloader := NewYtDlpDownloader(cfg, mock)
		url := "https://www.youtube.com/watch?v=test123"

		err := downloader.Download(url)

		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		// 驗證執行了正確的命令
		if mock.lastCommand != "yt-dlp" {
			t.Errorf("Expected command to be 'yt-dlp', got '%s'", mock.lastCommand)
		}

		// 驗證輸出目錄被創建
		if _, err := os.Stat(tempDir); os.IsNotExist(err) {
			t.Error("Expected output directory to be created")
		}
	})

	t.Run("download fails", func(t *testing.T) {
		tempDir := t.TempDir()

		cfg := config.NewConfig().WithOutputDir(tempDir)
		mock := &MockCommandExecutor{
			executeFunc: func(name string, args []string, stdout, stderr io.Writer) error {
				return errors.New("download error")
			},
		}

		downloader := NewYtDlpDownloader(cfg, mock)
		url := "https://www.youtube.com/watch?v=test123"

		err := downloader.Download(url)

		if err == nil {
			t.Error("Expected error when download fails")
		}
		if !strings.Contains(err.Error(), "下載失敗") {
			t.Errorf("Expected error message to contain '下載失敗', got: %v", err)
		}
	})

	t.Run("custom bitrate", func(t *testing.T) {
		tempDir := t.TempDir()

		cfg := config.NewConfig().
			WithOutputDir(tempDir).
			WithBitrate("256k")

		mock := &MockCommandExecutor{}
		downloader := NewYtDlpDownloader(cfg, mock)

		url := "https://www.youtube.com/watch?v=test123"
		_ = downloader.Download(url)

		// 檢查比特率參數
		argsStr := strings.Join(mock.lastArgs, " ")
		if !strings.Contains(argsStr, "256k") {
			t.Errorf("Expected args to contain custom bitrate '256k', got: %v", mock.lastArgs)
		}
	})
}

func TestGetOutputFiles(t *testing.T) {
	t.Run("find mp3 files", func(t *testing.T) {
		tempDir := t.TempDir()

		// 創建一些測試文件
		testFiles := []string{"song1.mp3", "song2.mp3", "video.mp4"}
		for _, filename := range testFiles {
			filepath := filepath.Join(tempDir, filename)
			if err := os.WriteFile(filepath, []byte("test"), 0644); err != nil {
				t.Fatalf("Failed to create test file: %v", err)
			}
		}

		cfg := config.NewConfig().WithOutputDir(tempDir)
		downloader := NewYtDlpDownloader(cfg, nil)

		files, err := downloader.GetOutputFiles()

		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		// 應該只找到 2 個 mp3 文件
		if len(files) != 2 {
			t.Errorf("Expected 2 mp3 files, got %d", len(files))
		}

		// 驗證文件名
		for _, file := range files {
			if !strings.HasSuffix(file, ".mp3") {
				t.Errorf("Expected mp3 file, got: %s", file)
			}
		}
	})

	t.Run("no files found", func(t *testing.T) {
		tempDir := t.TempDir()

		cfg := config.NewConfig().WithOutputDir(tempDir)
		downloader := NewYtDlpDownloader(cfg, nil)

		files, err := downloader.GetOutputFiles()

		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if len(files) != 0 {
			t.Errorf("Expected 0 files, got %d", len(files))
		}
	})
}

func TestDefaultCommandExecutor(t *testing.T) {
	executor := &DefaultCommandExecutor{}

	t.Run("execute successful command", func(t *testing.T) {
		err := executor.Execute("echo", []string{"test"}, os.Stdout, os.Stderr)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}
	})

	t.Run("execute failing command", func(t *testing.T) {
		err := executor.Execute("false", []string{}, os.Stdout, os.Stderr)
		if err == nil {
			t.Error("Expected error for failing command")
		}
	})
}
