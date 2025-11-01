package validator

import (
	"fmt"
	"os/exec"
)

// CommandChecker 定義檢查命令的接口
type CommandChecker interface {
	CheckCommand(name string) error
}

// SystemValidator 系統依賴驗證器
type SystemValidator struct {
	checker CommandChecker
}

// NewSystemValidator 創建新的系統驗證器
func NewSystemValidator(checker CommandChecker) *SystemValidator {
	if checker == nil {
		checker = &DefaultCommandChecker{}
	}
	return &SystemValidator{
		checker: checker,
	}
}

// ValidateDependencies 驗證所有必需的依賴
func (v *SystemValidator) ValidateDependencies() error {
	// 檢查 yt-dlp
	if err := v.checker.CheckCommand("yt-dlp"); err != nil {
		return fmt.Errorf("未找到 yt-dlp，請先安裝: pip install yt-dlp 或 brew install yt-dlp")
	}

	// 檢查 ffmpeg
	if err := v.checker.CheckCommand("ffmpeg"); err != nil {
		return fmt.Errorf("未找到 ffmpeg，請先安裝: sudo apt install ffmpeg 或 brew install ffmpeg")
	}

	return nil
}

// ValidateYtDlp 單獨驗證 yt-dlp
func (v *SystemValidator) ValidateYtDlp() error {
	if err := v.checker.CheckCommand("yt-dlp"); err != nil {
		return fmt.Errorf("未找到 yt-dlp")
	}
	return nil
}

// ValidateFFmpeg 單獨驗證 ffmpeg
func (v *SystemValidator) ValidateFFmpeg() error {
	if err := v.checker.CheckCommand("ffmpeg"); err != nil {
		return fmt.Errorf("未找到 ffmpeg")
	}
	return nil
}

// DefaultCommandChecker 默認的命令檢查器
type DefaultCommandChecker struct{}

// CheckCommand 檢查命令是否存在
func (c *DefaultCommandChecker) CheckCommand(name string) error {
	_, err := exec.LookPath(name)
	return err
}
