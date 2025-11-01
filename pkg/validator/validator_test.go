package validator

import (
	"errors"
	"strings"
	"testing"
)

// MockCommandChecker 模擬命令檢查器
type MockCommandChecker struct {
	commands map[string]error
}

// NewMockCommandChecker 創建新的模擬檢查器
func NewMockCommandChecker() *MockCommandChecker {
	return &MockCommandChecker{
		commands: make(map[string]error),
	}
}

// SetCommandResult 設置命令檢查結果
func (m *MockCommandChecker) SetCommandResult(name string, err error) {
	m.commands[name] = err
}

// CheckCommand 檢查命令（模擬實現）
func (m *MockCommandChecker) CheckCommand(name string) error {
	if err, exists := m.commands[name]; exists {
		return err
	}
	return nil
}

func TestNewSystemValidator(t *testing.T) {
	t.Run("with nil checker", func(t *testing.T) {
		validator := NewSystemValidator(nil)
		if validator == nil {
			t.Error("Expected validator to be created, got nil")
		}
		if validator.checker == nil {
			t.Error("Expected default checker to be set, got nil")
		}
	})

	t.Run("with custom checker", func(t *testing.T) {
		mock := NewMockCommandChecker()
		validator := NewSystemValidator(mock)
		if validator == nil {
			t.Error("Expected validator to be created, got nil")
		}
		if validator.checker != mock {
			t.Error("Expected custom checker to be set")
		}
	})
}

func TestValidateDependencies(t *testing.T) {
	t.Run("all dependencies present", func(t *testing.T) {
		mock := NewMockCommandChecker()
		mock.SetCommandResult("yt-dlp", nil)
		mock.SetCommandResult("ffmpeg", nil)

		validator := NewSystemValidator(mock)
		err := validator.ValidateDependencies()

		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}
	})

	t.Run("yt-dlp missing", func(t *testing.T) {
		mock := NewMockCommandChecker()
		mock.SetCommandResult("yt-dlp", errors.New("not found"))
		mock.SetCommandResult("ffmpeg", nil)

		validator := NewSystemValidator(mock)
		err := validator.ValidateDependencies()

		if err == nil {
			t.Error("Expected error when yt-dlp is missing")
		}
		if !strings.Contains(err.Error(), "yt-dlp") {
			t.Errorf("Expected error message to mention yt-dlp, got: %v", err)
		}
	})

	t.Run("ffmpeg missing", func(t *testing.T) {
		mock := NewMockCommandChecker()
		mock.SetCommandResult("yt-dlp", nil)
		mock.SetCommandResult("ffmpeg", errors.New("not found"))

		validator := NewSystemValidator(mock)
		err := validator.ValidateDependencies()

		if err == nil {
			t.Error("Expected error when ffmpeg is missing")
		}
		if !strings.Contains(err.Error(), "ffmpeg") {
			t.Errorf("Expected error message to mention ffmpeg, got: %v", err)
		}
	})

	t.Run("all dependencies missing", func(t *testing.T) {
		mock := NewMockCommandChecker()
		mock.SetCommandResult("yt-dlp", errors.New("not found"))
		mock.SetCommandResult("ffmpeg", errors.New("not found"))

		validator := NewSystemValidator(mock)
		err := validator.ValidateDependencies()

		if err == nil {
			t.Error("Expected error when all dependencies are missing")
		}
	})
}

func TestValidateYtDlp(t *testing.T) {
	t.Run("yt-dlp present", func(t *testing.T) {
		mock := NewMockCommandChecker()
		mock.SetCommandResult("yt-dlp", nil)

		validator := NewSystemValidator(mock)
		err := validator.ValidateYtDlp()

		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}
	})

	t.Run("yt-dlp missing", func(t *testing.T) {
		mock := NewMockCommandChecker()
		mock.SetCommandResult("yt-dlp", errors.New("not found"))

		validator := NewSystemValidator(mock)
		err := validator.ValidateYtDlp()

		if err == nil {
			t.Error("Expected error when yt-dlp is missing")
		}
	})
}

func TestValidateFFmpeg(t *testing.T) {
	t.Run("ffmpeg present", func(t *testing.T) {
		mock := NewMockCommandChecker()
		mock.SetCommandResult("ffmpeg", nil)

		validator := NewSystemValidator(mock)
		err := validator.ValidateFFmpeg()

		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}
	})

	t.Run("ffmpeg missing", func(t *testing.T) {
		mock := NewMockCommandChecker()
		mock.SetCommandResult("ffmpeg", errors.New("not found"))

		validator := NewSystemValidator(mock)
		err := validator.ValidateFFmpeg()

		if err == nil {
			t.Error("Expected error when ffmpeg is missing")
		}
	})
}

func TestDefaultCommandChecker(t *testing.T) {
	checker := &DefaultCommandChecker{}

	t.Run("check existing command", func(t *testing.T) {
		// 'go' should exist since we're running tests with go
		err := checker.CheckCommand("go")
		if err != nil {
			t.Errorf("Expected 'go' command to exist, got error: %v", err)
		}
	})

	t.Run("check non-existing command", func(t *testing.T) {
		err := checker.CheckCommand("this-command-definitely-does-not-exist-12345")
		if err == nil {
			t.Error("Expected error for non-existing command")
		}
	})
}
