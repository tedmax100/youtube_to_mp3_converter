package downloader

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"youtube_to_mp3/pkg/config"
)

// Downloader 定義下載器接口
type Downloader interface {
	Download(url string) error
	GetOutputFiles() ([]string, error)
}

// CommandExecutor 定義命令執行器接口
type CommandExecutor interface {
	Execute(name string, args []string, stdout, stderr io.Writer) error
}

// YtDlpDownloader YouTube 下載器實現
type YtDlpDownloader struct {
	config   *config.Config
	executor CommandExecutor
}

// NewYtDlpDownloader 創建新的 YtDlp 下載器
func NewYtDlpDownloader(cfg *config.Config, executor CommandExecutor) *YtDlpDownloader {
	if executor == nil {
		executor = &DefaultCommandExecutor{}
	}
	return &YtDlpDownloader{
		config:   cfg,
		executor: executor,
	}
}

// Download 下載並轉換視頻為 MP3
func (d *YtDlpDownloader) Download(url string) error {
	// 創建輸出目錄
	if err := os.MkdirAll(d.config.OutputDir, 0755); err != nil {
		return fmt.Errorf("創建輸出目錄失敗: %v", err)
	}

	// 構建 yt-dlp 命令參數
	args := d.buildArgs(url)

	// 執行命令
	if err := d.executor.Execute("yt-dlp", args, os.Stdout, os.Stderr); err != nil {
		return fmt.Errorf("下載失敗: %v", err)
	}

	return nil
}

// buildArgs 構建 yt-dlp 命令參數
func (d *YtDlpDownloader) buildArgs(url string) []string {
	return []string{
		"--extract-audio",
		"--audio-format", d.config.AudioFormat,
		"--audio-quality", d.config.AudioQuality,
		"--postprocessor-args", fmt.Sprintf("ffmpeg:-b:a %s", d.config.Bitrate),
		"-o", d.config.OutputTemplate,
		url,
	}
}

// GetOutputFiles 獲取輸出文件列表
func (d *YtDlpDownloader) GetOutputFiles() ([]string, error) {
	pattern := filepath.Join(d.config.OutputDir, fmt.Sprintf("*.%s", d.config.AudioFormat))
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, fmt.Errorf("查找輸出文件失敗: %v", err)
	}
	return files, nil
}

// DefaultCommandExecutor 默認命令執行器
type DefaultCommandExecutor struct{}

// Execute 執行系統命令
func (e *DefaultCommandExecutor) Execute(name string, args []string, stdout, stderr io.Writer) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	return cmd.Run()
}
