package mocks

import (
	"errors"
	"io"
)

// CommandChecker 模擬命令檢查器
type CommandChecker struct {
	commands map[string]error
}

// NewCommandChecker 創建新的模擬檢查器
func NewCommandChecker() *CommandChecker {
	return &CommandChecker{
		commands: make(map[string]error),
	}
}

// SetCommandResult 設置命令檢查結果
func (m *CommandChecker) SetCommandResult(name string, err error) {
	m.commands[name] = err
}

// CheckCommand 檢查命令（模擬實現）
func (m *CommandChecker) CheckCommand(name string) error {
	if err, exists := m.commands[name]; exists {
		return err
	}
	return nil
}

// CommandExecutor 模擬命令執行器
type CommandExecutor struct {
	ExecuteFunc func(name string, args []string, stdout, stderr io.Writer) error
	LastCommand string
	LastArgs    []string
	CallCount   int
}

// Execute 執行命令（模擬實現）
func (m *CommandExecutor) Execute(name string, args []string, stdout, stderr io.Writer) error {
	m.LastCommand = name
	m.LastArgs = args
	m.CallCount++

	if m.ExecuteFunc != nil {
		return m.ExecuteFunc(name, args, stdout, stderr)
	}
	return nil
}

// Downloader 模擬下載器
type Downloader struct {
	DownloadFunc     func(url string) error
	GetOutputFunc    func() ([]string, error)
	DownloadedURLs   []string
	ShouldFailOnURL  string
}

// Download 模擬下載
func (m *Downloader) Download(url string) error {
	m.DownloadedURLs = append(m.DownloadedURLs, url)

	if m.ShouldFailOnURL != "" && url == m.ShouldFailOnURL {
		return errors.New("mock download error")
	}

	if m.DownloadFunc != nil {
		return m.DownloadFunc(url)
	}
	return nil
}

// GetOutputFiles 模擬獲取輸出文件
func (m *Downloader) GetOutputFiles() ([]string, error) {
	if m.GetOutputFunc != nil {
		return m.GetOutputFunc()
	}
	return []string{}, nil
}
