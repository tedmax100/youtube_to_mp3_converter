package config

import (
	"path/filepath"
	"testing"
)

func TestNewConfig(t *testing.T) {
	cfg := NewConfig()

	if cfg.OutputDir != "output" {
		t.Errorf("Expected OutputDir to be 'output', got '%s'", cfg.OutputDir)
	}

	if cfg.AudioFormat != "mp3" {
		t.Errorf("Expected AudioFormat to be 'mp3', got '%s'", cfg.AudioFormat)
	}

	if cfg.AudioQuality != "0" {
		t.Errorf("Expected AudioQuality to be '0', got '%s'", cfg.AudioQuality)
	}

	if cfg.Bitrate != "320k" {
		t.Errorf("Expected Bitrate to be '320k', got '%s'", cfg.Bitrate)
	}

	expectedTemplate := filepath.Join("output", "%(title)s.%(ext)s")
	if cfg.OutputTemplate != expectedTemplate {
		t.Errorf("Expected OutputTemplate to be '%s', got '%s'", expectedTemplate, cfg.OutputTemplate)
	}
}

func TestWithOutputDir(t *testing.T) {
	cfg := NewConfig()
	customDir := "custom_output"

	cfg.WithOutputDir(customDir)

	if cfg.OutputDir != customDir {
		t.Errorf("Expected OutputDir to be '%s', got '%s'", customDir, cfg.OutputDir)
	}

	expectedTemplate := filepath.Join(customDir, "%(title)s.%(ext)s")
	if cfg.OutputTemplate != expectedTemplate {
		t.Errorf("Expected OutputTemplate to be '%s', got '%s'", expectedTemplate, cfg.OutputTemplate)
	}
}

func TestWithBitrate(t *testing.T) {
	cfg := NewConfig()
	customBitrate := "192k"

	cfg.WithBitrate(customBitrate)

	if cfg.Bitrate != customBitrate {
		t.Errorf("Expected Bitrate to be '%s', got '%s'", customBitrate, cfg.Bitrate)
	}
}

func TestConfigChaining(t *testing.T) {
	cfg := NewConfig().
		WithOutputDir("downloads").
		WithBitrate("256k")

	if cfg.OutputDir != "downloads" {
		t.Errorf("Expected OutputDir to be 'downloads', got '%s'", cfg.OutputDir)
	}

	if cfg.Bitrate != "256k" {
		t.Errorf("Expected Bitrate to be '256k', got '%s'", cfg.Bitrate)
	}
}
