package config

import "path/filepath"

// Config 應用配置
type Config struct {
	OutputDir      string
	AudioFormat    string
	AudioQuality   string
	Bitrate        string
	OutputTemplate string
}

// NewConfig 創建默認配置
func NewConfig() *Config {
	outputDir := "output"
	return &Config{
		OutputDir:      outputDir,
		AudioFormat:    "mp3",
		AudioQuality:   "0",
		Bitrate:        "320k",
		OutputTemplate: filepath.Join(outputDir, "%(title)s.%(ext)s"),
	}
}

// WithOutputDir 設置輸出目錄
func (c *Config) WithOutputDir(dir string) *Config {
	c.OutputDir = dir
	c.OutputTemplate = filepath.Join(dir, "%(title)s.%(ext)s")
	return c
}

// WithBitrate 設置比特率
func (c *Config) WithBitrate(bitrate string) *Config {
	c.Bitrate = bitrate
	return c
}
